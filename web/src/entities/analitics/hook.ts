import { keepPreviousData, useQuery } from '@tanstack/react-query';

import { GetKeyRate, GetKeyRates, GetRuonia } from './api';

export function useKeyRate() {
  return useQuery({
    queryKey       : ['key-rate'],
    queryFn        : GetKeyRate,
    staleTime      : 1000 * 60,
    placeholderData: keepPreviousData,
  });
}

export function useKeyRates() {
  return useQuery({
    queryKey       : ['key-rates'],
    queryFn        : GetKeyRates,
    staleTime      : 1000 * 60,
    placeholderData: keepPreviousData,
  });
}

export function useRuonia() {
  return useQuery({
    queryKey       : ['ruonia'],
    queryFn        : GetRuonia,
    staleTime      : 1000 * 60,
    placeholderData: keepPreviousData,
  });
}