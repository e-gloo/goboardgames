import { SOCKET_URL } from '@/constants';
import io, { Socket } from 'socket.io-client';
import { ref } from 'vue';
import { joinRoom, randomizeFleet } from './Api';
import { Board } from './Board';
import { EnemyBoard } from './EnemyBoard';
import { GamePhase } from './enums/GamePhase';

export class Game {
  socket: Socket;
  playerNb: number;
  phase = ref<GamePhase>(null);
  yourTurn = ref<boolean>(false);
  generatedCode = ref<string>(null);
  board = ref<Board>(null);
  enemyBoard = ref<EnemyBoard>(null);
  mapWidth = null;
  mapHeight = null;

  constructor() {
    this.socket = io(`${SOCKET_URL}/battleship`, {
      reconnection: false,
    });

    this.socket.connect();

    this.socket.on('connect', () => {
      console.log('connected');
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
    this.board.value = new Board(this.mapWidth, this.mapHeight);
    this.enemyBoard.value = new EnemyBoard(this.mapWidth, this.mapHeight);
  }

  async getBoardFleet() {
    this.board.value.fleet = await randomizeFleet(this.socket);
    return this.board.value.fleet;
  }

  ready() {
    this.socket.emit('ready');
  }
}
