import { AttackResultEnum } from "../enums/AttackResult";

export type AttackResult = {
  Position: number,
  Result: AttackResultEnum,
  PlayerNb: number
}
