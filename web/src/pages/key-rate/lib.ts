import type { Time } from 'lightweight-charts';
import type { KeyRate } from '#/entities/analitics/model';

export const KEY_RATE_COLOR = '#2962FF';
export const RUONIA_COLOR = '#FF6D00';

export function toTime(date: Date): Time {
  const year = date.getFullYear();
  const month = String(date.getMonth() + 1).padStart(2, '0');
  const day = String(date.getDate()).padStart(2, '0');

  return `${year}-${month}-${day}` as Time;
}

export function toLineData(data: KeyRate[]) {
  return [...data]
    .sort((a, b) => a.date.getTime() - b.date.getTime())
    .map(point => ({
      time : toTime(point.date),
      value: point.rate,
    }));
}
