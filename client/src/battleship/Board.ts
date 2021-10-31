import { Fleet } from './types/Fleet';

export class Board {
  fleet: Fleet;
  hits: number[];
  w: number;
  h: number;

  constructor(w: number, h: number) {
    this.hits = [];
    this.w = w;
    this.h = h;
  }

  isBoatHere(pos: number) {
    return this.fleet?.filter(f => f.Cells.includes(pos)).length
  }
}
