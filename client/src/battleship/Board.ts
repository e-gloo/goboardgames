import { AttackResultEnum } from "./enums/AttackResultEnum";

export class Board {
    hits: Record<number, AttackResultEnum> = {};
    w: number;
    h: number;

    constructor(w: number, h: number) {
      this.hits = {};
      this.w = w;
      this.h = h;
    }

    canAttack(pos: number) {
        return !this.hits[pos]
      }

}
