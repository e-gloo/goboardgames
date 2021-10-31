import { inject } from 'vue';
import { BATTLESHIP_KEY } from '../components/constants';
import { Game } from '../Game';

export function useBattleship(): Game {
  return inject(BATTLESHIP_KEY);
}
