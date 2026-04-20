import { zodResolver } from '@hookform/resolvers/zod';
import { useForm } from 'react-hook-form';
import z from 'zod';

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