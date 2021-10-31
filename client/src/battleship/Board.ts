import { Fleet } from './types/Fleet';

export class Board {
  fleet: Fleet;
  hits: number[];

  constructor() {
    this.hits = [];
  }
}
