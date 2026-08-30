import {
  AllCommunityModule,
  ModuleRegistry,
} from 'ag-charts-community';
import { AgCharts } from 'ag-charts-react';
import { useTheme } from 'next-themes';
import { useMemo, useRef } from 'react';
import { useNavigate } from '@tanstack/react-router';

import { Skeleton } from '#/components/ui/skeleton';
import { useIsMobile } from '#/hooks/use-mobile';

import type { AgChartOptions, AgChartTheme, AgChartInstance } from 'ag-charts-community';
import type { Bond } from '#/entities/bonds/model';


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
  data: Bond[]
  picked: Bond[]
  isLoading: boolean
  isUserLogedIn: boolean
}

export const BondChart = ({ data, isLoading, picked, isUserLogedIn }: BondChartProps) => {
  const { theme } = useTheme();
  const navigate = useNavigate({ from: isUserLogedIn ? '/app/picker' : '/app/watch' });
  const chartRef = useRef<AgChartInstance>(null);

  const pickedMap = useMemo(() => new Map(picked.map(bond => [bond.id, bond])), [picked]);
  const isMobile = useIsMobile();
   
  const options: AgChartOptions = useMemo(() => ({
    theme  : getTheme(theme),
    padding: {
      top   : isMobile ? 20 : 20,
      left  : isMobile ? 0 : 20,
      right : isMobile ? 5 : 20,
      bottom: isMobile ? 10 : 20,
    },
    animation: { enabled: false },
    series   : [
      {
        nodeClickRange: 'nearest',
        type          : 'scatter',
        title         : 'Все',
        xKey          : 'duration',
        data          : data.filter(bond => !pickedMap.has(bond.id)),
        xName         : 'Дюрация',
        yKey          : 'ytm',
        yName         : 'Доходность(YTM)',
        tooltip       : {
          enabled : !isMobile,
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
              replace       : false,
              resetScroll   : false,
              viewTransition: true,
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
          enabled : !isMobile,
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
              replace       : false,
              resetScroll   : false,
              viewTransition: true,
            });
          }
        },
      },
    ],
    interaction: {
      mode: 'nearest',
    },
    axes: {
      x: {
        type : 'number',
        title: {
          text: isMobile ? '' : 'Дюрация',
        },
      },
      y: {
        type : 'number',
        title: {
          text: isMobile ? '' : 'YTM',
        },
        label: {
          formatter: (params) => {
            return params.value + '%';
          },
        },
      },
    },
  }), [theme, isMobile, data, pickedMap, navigate]);

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
