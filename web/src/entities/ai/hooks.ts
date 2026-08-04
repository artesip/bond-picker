import { keepPreviousData, useQuery } from '@tanstack/react-query';

import { GetSuspicuous } from './api';

import type { BondWithRatings } from '../bonds/model';

export const AI_ENABLED = import.meta.env.VITE_ENABLE_AI === 'true';

export function useSuspicious(bonds: BondWithRatings[], kr: number, enabled: boolean) {
  return useQuery({
    queryKey       : ['suspicious-bonds', bonds, kr],
    queryFn        : () => GetSuspicuous(bonds, kr),
    staleTime      : 1000 * 60,
    placeholderData: keepPreviousData,
    enabled
  });
}
