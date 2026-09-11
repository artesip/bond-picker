import { memo, useEffect, useMemo, useRef } from 'react';
import { useTheme } from 'next-themes';
import { useNavigate } from '@tanstack/react-router';
import { ColorType, CrosshairMode, LineSeries, createOptionsChart } from 'lightweight-charts';

import { Skeleton } from '#/components/ui/skeleton';
import { useIsMobile } from '#/hooks/use-mobile';
import { useIsUserLoggedIn } from '#/stores/auth';
import { useSelectedBondStore } from '#/stores/selected-bond';
import { cn } from '#/lib/utils';

import { ScatterSeries, buildScatterLayout, toScatterPoints } from './scatter-series';

import type { IChartApiBase, ISeriesApi, MouseEventParams } from 'lightweight-charts';
import type { Bond } from '#/entities/bonds/model';
import type { ScatterPoint } from './scatter-series';

const ALL_COLOR = '#2962FF';
const PICKED_COLOR = '#FF6D00';

type ScatterChartProps = {
  data: Bond[]
  picked: Bond[]
  isLoading: boolean
  className?: string
};

function renderTooltip(point: ScatterPoint): string {
  return `
    <div class="mb-1.5 max-w-48 font-medium">${point.name}</div>
    <div class="flex items-center gap-2">
      <span class="text-muted-foreground">YTM</span>
      <span class="ml-auto pl-3 font-medium tabular-nums">${point.ytm.toFixed(2)}%</span>
    </div>
    <div class="flex items-center gap-2">
      <span class="text-muted-foreground">Дюрация</span>
      <span class="ml-auto pl-3 font-medium tabular-nums">${point.duration.toFixed(2)}</span>
    </div>
  `;
}

export const ScatterChart = memo(function ScatterChart({ data, picked, isLoading, className }: ScatterChartProps) {
  const { theme } = useTheme();
  const isMobile = useIsMobile();
  const isUserLogedIn = useIsUserLoggedIn();
  const navigate = useNavigate({ from: isUserLogedIn ? '/app/picker' : '/app/watch' });
  const setSelectedId = useSelectedBondStore((state) => state.setSelectedId);

  const containerRef = useRef<HTMLDivElement>(null);
  const tooltipRef = useRef<HTMLDivElement>(null);
  const chartRef = useRef<IChartApiBase<number> | null>(null);
  const gridSeriesRef = useRef<ISeriesApi<'Line', number> | null>(null);
  const allSeriesRef = useRef<ISeriesApi<'Custom', number> | null>(null);
  const pickedSeriesRef = useRef<ISeriesApi<'Custom', number> | null>(null);
  const pointsRef = useRef<Map<string, ScatterPoint>>(new Map());
  const navigateRef = useRef(navigate);
  const isMobileRef = useRef(isMobile);
  const fitKeyRef = useRef<string | null>(null);

  const pickedIds = useMemo(() => new Set(picked.map(bond => bond.id)), [picked]);

  useEffect(() => {
    navigateRef.current = navigate;
  }, [navigate]);

  useEffect(() => {
    isMobileRef.current = isMobile;
  }, [isMobile]);

  useEffect(() => {
    const container = containerRef.current;
    if (!container) return;

    const foregroundColor = theme === 'light' ? 'black' : 'white';
    const gridLineColor = theme === 'light' ? 'oklch(0.922 0 0)' : 'oklch(1 0 0 / 10%)';

    const chart = createOptionsChart(container, {
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
      crosshair: {
        mode: CrosshairMode.Normal,
      },
      rightPriceScale: {
        borderVisible: false,
      },
      timeScale: {
        borderVisible: false,
      },
      localization: {
        precision: 2,
      },
    });

    const gridSeries = chart.addSeries(LineSeries, {
      color                 : 'transparent',
      lineVisible           : false,
      pointMarkersVisible   : false,
      crosshairMarkerVisible: false,
      lastValueVisible      : false,
      priceLineVisible      : false,
      autoscaleInfoProvider : () => null,
    });

    const sharedOptions = {
      priceLineVisible: false,
      lastValueVisible: false,
      hoverStrokeColor: foregroundColor,
      priceFormat     : {
        type     : 'custom' as const,
        formatter: (price: number) => `${price.toFixed(2)}%`,
        minMove  : 0.01,
      },
    };

    const allSeries = chart.addCustomSeries(new ScatterSeries(), {
      ...sharedOptions,
      color : ALL_COLOR,
      radius: 4,
    });
    const pickedSeries = chart.addCustomSeries(new ScatterSeries(), {
      ...sharedOptions,
      color      : PICKED_COLOR,
      radius     : 5,
      hoverRadius: 7,
    });

    chartRef.current = chart;
    gridSeriesRef.current = gridSeries;
    allSeriesRef.current = allSeries;
    pickedSeriesRef.current = pickedSeries;

    const tooltip = tooltipRef.current;

    const hideTooltip = () => {
      if (tooltip) tooltip.style.display = 'none';
    };

    const handleCrosshairMove = (param: MouseEventParams<number>) => {
      const info = param.hoveredInfo;

      if (!tooltip || isMobileRef.current || !param.point || info?.objectKind !== 'custom-object' || typeof info.objectId !== 'string') {
        hideTooltip();

        return;
      }

      const point = pointsRef.current.get(info.objectId);
      if (!point) {
        hideTooltip();

        return;
      }

      tooltip.innerHTML = renderTooltip(point);
      tooltip.style.display = 'block';

      const tooltipWidth = tooltip.offsetWidth;
      const left = param.point.x + 12 + tooltipWidth > container.clientWidth
        ? param.point.x - tooltipWidth - 12
        : param.point.x + 12;

      tooltip.style.left = `${Math.max(0, left)}px`;
      tooltip.style.top = `${param.point.y + 12}px`;
    };

    const handleClick = (param: MouseEventParams<number>) => {
      const info = param.hoveredInfo;

      if (info?.objectKind !== 'custom-object' || typeof info.objectId !== 'string') return;

      setSelectedId(info.objectId);

      navigateRef.current({
        search: (prev) => ({
          ...prev,
          id: info.objectId as string,
        }),
        replace    : false,
        resetScroll: false,
      });
    };

    chart.subscribeCrosshairMove(handleCrosshairMove);
    chart.subscribeClick(handleClick);

    return () => {
      chart.unsubscribeCrosshairMove(handleCrosshairMove);
      chart.unsubscribeClick(handleClick);
      chart.remove();
      chartRef.current = null;
      gridSeriesRef.current = null;
      allSeriesRef.current = null;
      pickedSeriesRef.current = null;
      hideTooltip();
    };
  }, [theme, setSelectedId]);

  // Должен перезапускаться вместе с эффектом создания чарта (зависимость от `theme`),
  // чтобы после пересоздания серий заново заполнить их данными.
  useEffect(() => {
    const width = containerRef.current?.clientWidth ?? 1000;
    const layout = buildScatterLayout(data, width);
    const points = toScatterPoints(data, layout);
    const all = points.filter(point => !pickedIds.has(point.id));
    const favorite = points.filter(point => pickedIds.has(point.id));

    pointsRef.current = new Map(points.map(point => [point.id, point]));

    gridSeriesRef.current?.setData(layout.grid.map(time => ({ time, value: 0 })));
    allSeriesRef.current?.setData(all);
    pickedSeriesRef.current?.setData(favorite);

    // fitContent() сбрасывает зум/пан — вызываем только когда реально изменился
    // набор облигаций или тема, а не при каждом ререндере (например, после клика).
    const fitKey = `${theme}:${data.map(bond => bond.id).join(',')}`;
    if (fitKey !== fitKeyRef.current) {
      fitKeyRef.current = fitKey;
      chartRef.current?.timeScale().fitContent();
    }
  }, [data, pickedIds, theme]);

  return (
    <div className={ cn('relative h-full w-full', className) } ref={ containerRef }>
      <div
        ref={ tooltipRef }
        style={ { display: 'none' } }
        className='pointer-events-none absolute z-10 rounded-lg border border-border bg-popover px-3 py-2 text-xs text-popover-foreground shadow-md'
      />
      {
        isLoading && <Skeleton className='absolute inset-0 z-20'/>
      }
    </div>
  );
});
