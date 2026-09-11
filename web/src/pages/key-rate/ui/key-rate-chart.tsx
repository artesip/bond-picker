import { useEffect, useRef } from 'react';
import { useTheme } from 'next-themes';
import { ColorType, LineSeries, createChart } from 'lightweight-charts';

import { cn } from '#/lib/utils';

import { KEY_RATE_COLOR, RUONIA_COLOR, toLineData } from '../lib';

import { renderTooltip } from './tooltip';

import type { IChartApi, ISeriesApi, LineData, MouseEventParams, Time } from 'lightweight-charts';
import type { KeyRate } from '#/entities/analitics/model';
import type { TooltipRow } from './tooltip';

type KeyRateChartProps = {
  keyRates?: KeyRate[]
  ruonia?: KeyRate[]
  className?: string
};

export function KeyRateChart({ keyRates, ruonia, className }: KeyRateChartProps) {
  const { theme } = useTheme();

  const containerRef = useRef<HTMLDivElement>(null);
  const tooltipRef = useRef<HTMLDivElement>(null);
  const chartRef = useRef<IChartApi | null>(null);
  const keyRateSeriesRef = useRef<ISeriesApi<'Line'> | null>(null);
  const ruoniaSeriesRef = useRef<ISeriesApi<'Line'> | null>(null);

  useEffect(() => {
    const container = containerRef.current;
    if (!container) return;

    const foregroundColor = theme === 'light' ? 'black' : 'white';
    const gridLineColor = theme === 'light' ? 'oklch(0.922 0 0)' : 'oklch(1 0 0 / 10%)';

    const chart = createChart(container, {
      autoSize: true,
      layout  : {
        background     : { type: ColorType.Solid, color: 'transparent' },
        textColor      : foregroundColor,
        attributionLogo: false,
      },
      grid: {
        vertLines: { color: gridLineColor },
        horzLines: { color: gridLineColor },
      },
      rightPriceScale: {
        borderVisible: false,
      },
      timeScale: {
        borderVisible: false,
      },
    });

    const keyRateSeries = chart.addSeries(LineSeries, {
      color           : KEY_RATE_COLOR,
      lineWidth       : 2,
      priceLineVisible: false,
      lastValueVisible: false,
    });
    const ruoniaSeries = chart.addSeries(LineSeries, {
      color           : RUONIA_COLOR,
      lineWidth       : 2,
      priceLineVisible: false,
      lastValueVisible: false,
    });

    chartRef.current = chart;
    keyRateSeriesRef.current = keyRateSeries;
    ruoniaSeriesRef.current = ruoniaSeries;

    const tooltip = tooltipRef.current;
    if (tooltip) tooltip.style.display = 'none';

    const handleCrosshairMove = (param: MouseEventParams) => {
      if (!tooltip) return;

      const keyRateData = param.seriesData.get(keyRateSeries) as LineData<Time> | undefined;
      const ruoniaData = param.seriesData.get(ruoniaSeries) as LineData<Time> | undefined;

      if (!param.point || param.time === undefined || (keyRateData?.value == null && ruoniaData?.value == null)) {
        tooltip.style.display = 'none';

        return;
      }

      const rows: TooltipRow[] = [];
      if (keyRateData?.value != null) {
        rows.push({ color: KEY_RATE_COLOR, name: 'Ключевая ставка', value: keyRateData.value });
      }
      if (ruoniaData?.value != null) {
        rows.push({ color: RUONIA_COLOR, name: 'RUONIA', value: ruoniaData.value });
      }

      tooltip.innerHTML = renderTooltip(param.time, rows);
      tooltip.style.display = 'block';

      const tooltipWidth = tooltip.offsetWidth;
      const left = param.point.x + 12 + tooltipWidth > container.clientWidth
        ? param.point.x - tooltipWidth - 12
        : param.point.x + 12;

      tooltip.style.left = `${Math.max(0, left)}px`;
      tooltip.style.top = `${param.point.y + 12}px`;
    };

    chart.subscribeCrosshairMove(handleCrosshairMove);

    return () => {
      chart.unsubscribeCrosshairMove(handleCrosshairMove);
      chart.remove();
      chartRef.current = null;
      keyRateSeriesRef.current = null;
      ruoniaSeriesRef.current = null;
      if (tooltip) tooltip.style.display = 'none';
    };
  }, [theme]);

  useEffect(() => {
    const series = keyRateSeriesRef.current;
    if (!series) return;

    series.setData(toLineData(keyRates ?? []));
    chartRef.current?.timeScale().fitContent();
  }, [keyRates, theme]);

  useEffect(() => {
    const series = ruoniaSeriesRef.current;
    if (!series) return;

    series.setData(toLineData(ruonia ?? []));
    chartRef.current?.timeScale().fitContent();
  }, [ruonia, theme]);

  return (
    <div className={ cn('relative', className) } ref={ containerRef }>
      <div
        ref={ tooltipRef }
        className='pointer-events-none absolute z-10 rounded-lg border border-border bg-popover px-3 py-2 text-xs text-popover-foreground shadow-md'
      />
    </div>
  );
}
