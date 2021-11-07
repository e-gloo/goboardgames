import { SOCKET_URL } from '@/constants';
import io, { Socket } from 'socket.io-client';
import { ref, Ref } from 'vue';
import { joinRoom, randomizeFleet } from './Api';
import { MyBoard } from './MyBoard';
import { EnemyBoard } from './EnemyBoard';
import { GamePhase } from './enums/GamePhase';

export class Game {
  socket: Socket;
  playerNb: number;
  phase = ref<GamePhase>(null);
  yourTurn = ref<boolean>(false);
  generatedCode = ref<string>(null);
  board = ref<MyBoard>(null);
  enemyBoards: Record<number, Ref<EnemyBoard>> = {};
  mapWidth = null;
  mapHeight = null;
  nbPlayers = 2;

  constructor(code?: string) {
    this.generatedCode.value = code;
    this.socket = io(`${SOCKET_URL}/battleship`, {
      reconnection: false,
    });

    this.socket.connect();

    this.socket.on('connect', () => {
      console.log('connected');
      if (this.generatedCode.value?.length) {
        this.joinRoom(this.generatedCode.value)
      }
    });
    this.socket.on('disconnect', () => {
      console.log('disconnected');
    });

    this.socket.on('newCode', (data: string) => {
      this.generatedCode.value = data;
    });
    this.socket.on(
      'gameOptions',
      (data: { MapWidth: number; MapHeight: number }) => {
        this.mapWidth = data.MapWidth;
        this.mapHeight = data.MapHeight;
        console.log(this);
      }
    );
    this.socket.on('gamePhaseUpdated', async (data: number) => {
      this.phase.value = data;
    });
    this.socket.on('PlayerNb', (data: number) => {
      this.playerNb = data;
    });
    this.socket.on('PlayersTurn', (playerNb: number) => {
      this.yourTurn.value = this.playerNb === playerNb;
    })
  }

  disconnect() {
    this.socket.disconnect();
  }

  async joinRoom(code: string) {
    if (!code?.length) {
      return;
    }

    this.generatedCode.value = code;
    const success = await joinRoom(this.socket, code);
    if (!success) {
      // TODO: display error
      this.generatedCode.value = null;
    }
  }

  playWithFriend() {
    this.phase.value = GamePhase.WAITING;
    this.socket.emit('playWithFriend');
  }

  createBoard() {
    this.board.value = new MyBoard(this.mapWidth, this.mapHeight);
    // Styve
    // Array(this.nbPlayers).fill(0).forEach(k => {
    //   if (k + 1 !== this.playerNb) {
    //       this.enemyBoards[k + 1] = ref(new EnemyBoard(this.mapWidth, this.mapHeight));
    //   }
    // });

    // VS Dylan
    for (let i = 1; i <= this.nbPlayers; ++i) {
      if (i != this.playerNb) {
        this.enemyBoards[i] = ref(new EnemyBoard(this.mapWidth, this.mapHeight));
      }
    }

    console.log(this.enemyBoards);

    // FIGHT!
  }

  async getBoardFleet() {
    this.board.value.fleet = await randomizeFleet(this.socket);
    return this.board.value.fleet;
  }

  ready() {
    this.socket.emit('ready');
  }

  attack(pos: number) {
    console.log(`attacking in ${pos}`)
    this.socket.emit('attack', pos, this.playerNb == 1 ? 2 : 1);
  }
}
