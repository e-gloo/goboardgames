import { Fleet } from "@/battleship/types/Fleet";
import createWrapper from "event-wrapper";
import { Socket } from "socket.io-client";

class Api {
  socket: Socket;

  constructor(socket: Socket) {
    this.socket = socket;
  }

  randomizeFleet() {
    return new Promise<Fleet>((resolve, reject) => {
      const wrap = createWrapper(this.socket, (result) => {
        result ? resolve(result) : reject();
      });

      wrap("randomizeFleetError", () => wrap.done());
      wrap("NewFleet", (fleet) => {
        const ships = fleet.map((s) => {
          const cells = Array.from(atob(s.Cells)).map((v) => v.charCodeAt(0));
          return { ...s, Cells: cells };
        });
        wrap.done(ships);
      });
      this.socket.emit("randomizeFleet");
    });
  }
}

export default Api;
