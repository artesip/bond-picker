import { customSeriesDefaultOptions } from 'lightweight-charts';

import type {
  Coordinate,
  CustomData,
  CustomSeriesHitTestResult,
  CustomSeriesOptions,
  CustomSeriesWhitespaceData,
  ICustomSeriesPaneRenderer,
  ICustomSeriesPaneView,
  PaneRendererCustomData,
  PriceToCoordinateConverter,
} from 'lightweight-charts';
import type { Bond } from '#/entities/bonds/model';

type RenderTarget = Parameters<ICustomSeriesPaneRenderer['draw']>[0];
type DrawHitTestData = Parameters<ICustomSeriesPaneRenderer['draw']>[3];

export type ScatterPoint = CustomData<number> & {
  value: number
  id: string
  name: string
  duration: number
  ytm: number
};

export type ScatterSeriesOptions = CustomSeriesOptions & {
  radius: number
  hoverRadius: number
  hitRadius: number
  hoverStrokeColor: string
};

const MIN_GRID_POINTS = 400;
const MAX_GRID_POINTS = 6000;
const DUPLICATE_EPSILON = 1e-4;

export type ScatterLayout = {
  grid: number[]
  step: number
  min: number
};

export function buildScatterLayout(bonds: Bond[], width: number): ScatterLayout {
  const durations = bonds.map(bond => bond.duration);
  const min = durations.length > 0 ? Math.min(...durations) : 0;
  const max = durations.length > 0 ? Math.max(...durations) : 0;

  if (!Number.isFinite(min) || !Number.isFinite(max) || max <= min) {
    return { grid: [min], step: 1, min };
  }

  const points = Math.min(MAX_GRID_POINTS, Math.max(MIN_GRID_POINTS, Math.round(width)));
  const step = (max - min) / (points - 1);
  const grid: number[] = [];

  for (let index = 0; index < points; index++) {
    grid.push(min + index * step);
  }

  return { grid, step, min };
}

export function toScatterPoints(bonds: Bond[], layout: ScatterLayout): ScatterPoint[] {
  const sorted = [...bonds].sort((a, b) => a.duration - b.duration);
  const counters = new Map<number, number>();

  return sorted.map((bond) => {
    const cell = Math.round((bond.duration - layout.min) / layout.step);
    const count = counters.get(cell) ?? 0;
    counters.set(cell, count + 1);

    return {
      time    : layout.min + (cell + count * DUPLICATE_EPSILON) * layout.step,
      value   : bond.ytm,
      id      : bond.id,
      name    : bond.name,
      duration: bond.duration,
      ytm     : bond.ytm,
    };
  });
}

class ScatterSeriesRenderer implements ICustomSeriesPaneRenderer {
  private _data: PaneRendererCustomData<number, ScatterPoint> | null = null;
  private _options: ScatterSeriesOptions | null = null;

  update(data: PaneRendererCustomData<number, ScatterPoint>, seriesOptions: ScatterSeriesOptions) {
    this._data = data;
    this._options = seriesOptions;
  }

  draw(target: RenderTarget, priceConverter: PriceToCoordinateConverter, _isHovered: boolean, hitTestData?: DrawHitTestData) {
    const data = this._data;
    const options = this._options;
    if (!data || !options) return;

    const hoveredId = (hitTestData as { objectId?: string } | undefined)?.objectId;
    const range = data.visibleRange;
    const from = range ? range.from : 0;
    const to = range ? range.to : data.bars.length;

    target.useMediaCoordinateSpace((scope) => {
      const ctx = scope.context;

      for (let index = from; index < to; index++) {
        const bar = data.bars[index];
        if (!bar || !Number.isFinite(bar.x)) continue;

        const y = priceConverter(bar.originalData.value);
        if (y === null) continue;

        const isHovered = hoveredId !== undefined && hoveredId === bar.originalData.id;
        const radius = isHovered ? options.hoverRadius : options.radius;

        ctx.beginPath();
        ctx.arc(bar.x, y, radius, 0, Math.PI * 2);
        ctx.fillStyle = options.color;
        ctx.fill();

        if (isHovered) {
          ctx.lineWidth = 2;
          ctx.strokeStyle = options.hoverStrokeColor;
          ctx.stroke();
        }
      }
    });
  }

  hitTest(x: Coordinate, y: Coordinate, priceConverter: PriceToCoordinateConverter): CustomSeriesHitTestResult | null {
    const data = this._data;
    const options = this._options;
    if (!data || !options) return null;

    let nearest: ScatterPoint | null = null;
    let nearestDistance = Infinity;

    const range = data.visibleRange;
    const from = range ? range.from : 0;
    const to = range ? range.to : data.bars.length;

    for (let index = from; index < to; index++) {
      const bar = data.bars[index];
      if (!bar || !Number.isFinite(bar.x)) continue;

      const pointY = priceConverter(bar.originalData.value);
      if (pointY === null) continue;

      const distance = Math.hypot(x - bar.x, y - pointY);
      if (distance < nearestDistance) {
        nearestDistance = distance;
        nearest = bar.originalData;
      }
    }

    if (!nearest || nearestDistance > options.hitRadius) return null;

    return {
      distance   : nearestDistance,
      objectId   : nearest.id,
      type       : 'point',
      cursorStyle: 'pointer',
      hitTestData: { objectId: nearest.id },
    };
  }
}

export class ScatterSeries implements ICustomSeriesPaneView<number, ScatterPoint, ScatterSeriesOptions> {
  private _renderer = new ScatterSeriesRenderer();

  renderer() {
    return this._renderer;
  }

  update(data: PaneRendererCustomData<number, ScatterPoint>, seriesOptions: ScatterSeriesOptions) {
    this._renderer.update(data, seriesOptions);
  }

  priceValueBuilder(plotRow: ScatterPoint) {
    return [plotRow.value];
  }

  isWhitespace(data: ScatterPoint | CustomSeriesWhitespaceData<number>): data is CustomSeriesWhitespaceData<number> {
    return (data as ScatterPoint).value === undefined;
  }

  defaultOptions(): ScatterSeriesOptions {
    return {
      ...customSeriesDefaultOptions,
      radius          : 4,
      hoverRadius     : 6,
      hitRadius       : 12,
      hoverStrokeColor: '#FFFFFF',
    };
  }
}
