import { Fleet } from './types/Fleet';
import { Board } from './Board';

export class MyBoard extends Board {
  fleet: Fleet;

  constructor(w: number, h: number) {
    super(w, h);
  }

  isBoatHere(pos: number) {
    return this.fleet?.filter(f => f.Cells.includes(pos)).length
  }
}
