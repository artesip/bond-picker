import {
  AllCommunityModule,
  ModuleRegistry,
} from 'ag-charts-community';
import { AgCharts } from 'ag-charts-react';
import { useTheme } from 'next-themes';
import { useMemo, useRef } from 'react';
import { useNavigate } from '@tanstack/react-router';

import { Skeleton } from '#/components/ui/skeleton';

import type { AgChartOptions, AgChartTheme, AgChartInstance } from 'ag-charts-community';
import type { Bond, BondWithRatings } from '#/entities/bonds/model';


ModuleRegistry.registerModules([
  AllCommunityModule
]);

function getTheme(theme?: string): AgChartTheme {
  const foregroundColor = theme === 'light' ? 'black' : 'white';
  const gridLineColor = theme === 'light' ? 'oklch(0.922 0 0)' : 'oklch(1 0 0 / 10%)';

  return {
    params: {
      backgroundColor: 'transparent',
      foregroundColor: foregroundColor,
      gridLineColor  : gridLineColor,
    }
  };
}

type BondChartProps = {
  data: BondWithRatings[]
  picked: Bond[]
  isLoading: boolean
}

export const BondChart = ({ data, isLoading, picked }: BondChartProps) => {
  const { theme } = useTheme();
  const navigate = useNavigate({ from: '/app/picker' });
  const chartRef = useRef<AgChartInstance>(null);

  const pickedMap = useMemo(() => new Map(picked.map(bond => [bond.id, bond])), [picked]);
   
  const options: AgChartOptions = useMemo(() => ({
    theme    : getTheme(theme),
    padding  : 1,
    animation: { enabled: false },
    series   : [
      {
        nodeClickRange: 'nearest',
        type          : 'scatter',
        title         : 'All',
        xKey          : 'duration',
        data          : data.filter(bond => !pickedMap.has(bond.id)),
        xName         : 'Дюрация',
        yKey          : 'ytm',
        yName         : 'Доходность(YTM)',
        tooltip       : {
          renderer: (params) => {
            const { datum } = params;
            
            return {
              title: datum.name
            };
          }
        },
        listeners: {
          seriesNodeClick: (event) => {
            const id = event.datum?.id;
            if (!id) return;

            navigate({
              search: (prev) => ({
                ...prev,
                id: String(id),
              }),
              replace: false,
            });
          }
        },
      },
      {
        nodeClickRange: 'nearest',
        type          : 'scatter',
        title         : 'Избранное',
        xKey          : 'duration',
        data          : data.filter(bond => pickedMap.has(bond.id)),
        xName         : 'Дюрация',
        yKey          : 'ytm',
        yName         : 'Доходность(YTM)',
        tooltip       : {
          renderer: (params) => {
            const { datum } = params;
            
            return {
              title: datum.name
            };
          }
        },
        listeners: {
          seriesNodeClick: (event) => {
            const id = event.datum?.id;
            if (!id) return;

            navigate({
              search: (prev) => ({
                ...prev,
                id: String(id),
              }),
              replace: false,
            });
          }
        },
      },
      {
        nodeClickRange: 'nearest',
        type          : 'scatter',
        title         : 'AI',
        xKey          : 'duration',
        data          : data.filter((_, i) => i % 3 === 2),
        xName         : 'Дюрация',
        yKey          : 'ytm',
        yName         : 'Доходность(YTM)',
        tooltip       : {
          renderer: (params) => {
            const { datum } = params;
            
            return {
              title: datum.name
            };
          }
        },
        listeners: {
          seriesNodeClick: (event) => {
            const id = event.datum?.id;
            if (!id) return;

            navigate({
              search: (prev) => ({
                ...prev,
                id: String(id),
              }),
              replace: false,
            });
          }
        },
      }
    ],
    interaction: {
      mode: 'nearest',
    },
    axes: {
      x: {
        type : 'number',
        title: {
          text: 'Дюрация',
        },
      },
      y: {
        type : 'number',
        title: {
          text: 'YTM',
        },
        label: {
          formatter: (params) => {
            return params.value + '%';
          },
        },
      },
    },
  }), [theme, data, navigate, pickedMap]);

  return (
    <>
      {
        isLoading && <Skeleton className='aspect-video w-full h-full'/>
      }
      {
        !isLoading && <AgCharts options={ options } style={ { height: '100%' } } ref={ chartRef }/>
      }
    </>
  );
};
