import { SOCKET_URL } from '@/constants';
import io, { Socket } from 'socket.io-client';
import { ref } from 'vue';
import { randomizeFleet } from './Api';
import { Board } from './Board';
import { GamePhase } from './enums/GamePhase';

export class Game {
  socket: Socket;
  phase = ref<GamePhase>(null);
  yourTurn = ref<boolean>(false);
  generatedCode = ref<string>(null);
  board = ref<Board>(null);

  constructor() {
    this.socket = io(`${SOCKET_URL}/battleship`, {
      reconnection: false,
    });

    this.socket.connect();

    this.socket.on('connect', () => {
      console.log('connected');
      if (this.generatedCode.value) {
        this.socket.emit('joinRoom', this.generatedCode.value);
      }
    });
    this.socket.on('disconnect', () => {
      console.log('disconnected');
    });

    this.socket.on('newCode', (data: string) => {
      this.generatedCode.value = data;
    });
    this.socket.on('gamePhaseUpdated', async (data: number) => {
      this.phase.value = data;
    });
  }

  disconnect() {
    this.socket.disconnect();
  }

  playWithFriend() {
    this.phase.value = GamePhase.WAITING;
    this.socket.emit('playWithFriend');
  }

  createBoard() {
    this.board.value = new Board();
  }

  async getBoardFleet() {
    this.board.value.fleet = await randomizeFleet(this.socket);
    return this.board.value.fleet;
  }
}
