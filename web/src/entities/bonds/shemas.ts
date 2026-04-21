import { zodResolver } from '@hookform/resolvers/zod';
import { useForm } from 'react-hook-form';
import { toast } from 'sonner';
import z from 'zod';

import { PickBond } from './api';

export const FilterSchema = z.object({
  ratingEnabled  : z.boolean(),
  ytmEnabled     : z.boolean(),
  durationEnabled: z.boolean(),
  offerEnabled   : z.boolean(),
  currencyEnabled: z.boolean(),

  ratingFrom: z.string(),
  ratingTo  : z.string(),

  ytmFrom: z.number(),
  ytmTo  : z.number(),

  durationFrom: z.number(),
  durationTo  : z.number(),

  currency: z.string(),
  offer   : z.string(),
});

export type FilterInput = z.infer<typeof FilterSchema>;


export function useFilterForm() {
  const rhf = useForm<FilterInput>({
    resolver     : zodResolver(FilterSchema),
    defaultValues: {
      ratingEnabled  : true,
      ytmEnabled     : true,
      durationEnabled: true,
      offerEnabled   : false,
      currencyEnabled: false,

      ratingFrom: 'BBB+',
      ratingTo  : 'AA',

      ytmFrom: 10,
      ytmTo  : 150,

      durationFrom: 0,
      durationTo  : 8,

      currency: '',
      offer   : 'all',
    }
  });

  return {
    rhf
  };
}


const FavoriteAddSchema = z.object({
  number: z.number().min(1, 'Минимальное кол-во 1'),
});

export type FavoriteAddInput = z.infer<typeof FavoriteAddSchema>;

export function useFavoriteAddForm(id: string) {
  const rhf = useForm<FavoriteAddInput>({
    resolver     : zodResolver(FavoriteAddSchema),
    defaultValues: {
      number: 1,
    }
  });

  const onSubmit = rhf.handleSubmit(async (data) => {
    try {
      await PickBond(id, data.number);
    } catch (e) {
      if (e instanceof Error) {
        toast.error(e.message);
      }
    }
  });

  return {
    rhf,
    onSubmit,
  }; 
}