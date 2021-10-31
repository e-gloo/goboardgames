import { Fleet } from './types/Fleet';
import { AttackResult } from './enums/AttackResult'

export class EnemyBoard {
  hits: Record<number, AttackResult> = {};
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
