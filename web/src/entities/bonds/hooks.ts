import { keepPreviousData, useQuery } from '@tanstack/react-query';
import { useMemo } from 'react';

import { GetBonds, GetPicked, GetRatings } from './api';

import type { Rating } from './model';

export function useBonds() {
  return useQuery({
    queryKey       : ['bonds'],
    queryFn        : GetBonds,
    staleTime      : 1000 * 60,
    placeholderData: keepPreviousData,
  });
}

export function usePickedBonds() {
  return useQuery({
    queryKey       : ['picked-bonds'],
    queryFn        : GetPicked,
    staleTime      : 1000 * 60,
    placeholderData: keepPreviousData,
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

export function useBondWithRatings() {
  const { data: bonds, isLoading: bondLoading } = useBonds();
  const { data: ratings, isLoading: ratingLoading } = useRatings();

  const isLoading = bondLoading || ratingLoading;

  const data = useMemo(() => {
    if (isLoading) return [];

    const ratingsByCompanyID = new Map<string, Rating[]>();

    if (ratings?.length) {
      for (const rating of ratings) {
        const key = rating.companyID;

        const list = ratingsByCompanyID.get(key);
        if (list) {
          list.push(rating);
        } else {
          ratingsByCompanyID.set(key, [rating]);
        }
      }
    }

    return (bonds ?? []).map((bond) => ({
      ...bond,
      ratings: ratingsByCompanyID.get(bond.companyID) ?? []
    }));
  }, [bonds, ratings, isLoading]);

  return {
    data,
    isLoading
  };
}

