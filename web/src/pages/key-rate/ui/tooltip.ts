import type { Time } from 'lightweight-charts';

export type TooltipRow = {
  color: string
  name: string
  value: number
};

export function formatTime(time: Time): string {
  if (typeof time === 'string') {
    const [year, month, day] = time.split('-');

    return `${day}.${month}.${year}`;
  }

  if (typeof time === 'number') {
    return new Date(time * 1000).toLocaleDateString('ru-RU');
  }

  return `${String(time.day).padStart(2, '0')}.${String(time.month).padStart(2, '0')}.${time.year}`;
}

export function formatValue(value: number): string {
  return `${value.toLocaleString('ru-RU', { maximumFractionDigits: 2 })}%`;
}

function renderRow(row: TooltipRow): string {
  return `
    <div class="flex items-center gap-2">
      <span class="size-2 shrink-0 rounded-full" style="background:${row.color}"></span>
      <span class="text-muted-foreground">${row.name}</span>
      <span class="ml-auto pl-3 font-medium tabular-nums">${formatValue(row.value)}</span>
    </div>
  `;
}

export function renderTooltip(time: Time, rows: TooltipRow[]): string {
  return `
    <div class="mb-1.5 font-medium">${formatTime(time)}</div>
    <div class="flex flex-col gap-1">${rows.map(renderRow).join('')}</div>
  `;
}
