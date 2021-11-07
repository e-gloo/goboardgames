import { Fleet } from '@/battleship/types/Fleet';
import createWrapper from 'event-wrapper';
import { Socket } from 'socket.io-client';

export function randomizeFleet(socket: Socket) {
  return new Promise<Fleet>((resolve, reject) => {
    const wrap = createWrapper(socket, (result) => {
      result ? resolve(result) : reject();
    });

    wrap('randomizeFleetError', () => wrap.done());
    wrap('NewFleet', (fleet) => {
      const ships = fleet.map((s) => {
        const cells = Array.from(atob(s.Cells)).map((v) => v.charCodeAt(0));
        return { ...s, Cells: cells };
      });
      wrap.done(ships);
    });
    socket.emit('randomizeFleet');
  });
}

export function joinRoom(socket: Socket, code: string) {
  return new Promise<boolean>((resolve) => {
    const wrap = createWrapper(socket, resolve);

    wrap('joinRoomError', () => wrap.done(false));
    wrap('joinRoomOk', () => wrap.done(true));

    socket.emit('joinRoom', code);
  });
}

export function ready(socket: Socket, cb: (playerNb) => void) {
  return new Promise<number>((resolve) => {
    const wrap = createWrapper(socket, resolve);

    wrap('PlayerReady', (playerNb) => cb(playerNb));
    wrap('PlayersTurn', wrap.done);
  });
}


export function attackResult(socket: Socket, handleResult: (AttackResult) => void): any {
  const wrap = createWrapper(socket);
  wrap('AttackResult', handleResult)
  return wrap;
}

export function gameOver(socket: Socket, attackWrap: any): any {
  return new Promise<boolean>((resolve) => {
    const wrap = createWrapper(socket, resolve);

    wrap('GameOver', () => {
      attackWrap.done();
      wrap.done();
    });
  });
}
