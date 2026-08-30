import { keepPreviousData, useQuery } from '@tanstack/react-query';

import { GetBonds, GetBondsFull, GetKeyRate, GetPicked, GetRatings } from './api';

export function useBonds() {
  return useQuery({
    queryKey       : ['bonds'],
    queryFn        : GetBonds,
    staleTime      : 1000 * 60,
    placeholderData: keepPreviousData,
  });
}

export function usePickedBonds(enabled: boolean = true) {
  return useQuery({
    queryKey       : ['picked-bonds'],
    queryFn        : GetPicked,
    staleTime      : 1000 * 60,
    placeholderData: keepPreviousData,
    enabled        : enabled,
  });
}

export function useRatings() {
  return useQuery({
    queryKey       : ['ratings'],
    queryFn        : GetRatings,
    staleTime      : 1000 * 60,
    placeholderData: keepPreviousData,
  });
}

export function useKeyRate() {
  return useQuery({
    queryKey       : ['key-rate'],
    queryFn        : GetKeyRate,
    staleTime      : 1000 * 60,
    placeholderData: keepPreviousData,
  });
}

export function useBondWithRatings() {
  return useQuery({
    queryKey       : ['bonds-full'],
    queryFn        : GetBondsFull,
    staleTime      : 1000 * 60,
    placeholderData: keepPreviousData,
  });
}

