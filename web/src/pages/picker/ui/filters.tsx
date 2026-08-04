import { useWatch  } from 'react-hook-form';

import { useBonds, useRatings } from '#/entities/bonds/hooks';
import { compareRatings } from '#/entities/bonds/model';

import { FilterToggles } from './filters/toggles';
import { RatingSelect } from './filters/rating-select';
import { CurrencySelect } from './filters/currency-select';
import { OfferToggle } from './filters/offer-toggle';
import { YtmSlider } from './filters/ytm-slider';
import { DurationSlider } from './filters/duration-slider';

import type { UseFormReturn } from 'react-hook-form';
import type { FilterInput } from '#/entities/bonds/shemas';

type FilterBlockProps = {
    rhf: UseFormReturn<FilterInput>
}

export function FilterBlock({ rhf }: FilterBlockProps) {
  const { data: ratings, isLoading: ratingLoading } = useRatings();
  const { data: bonds, isLoading: bondsLoading } = useBonds();

  const filters = useWatch({
    control: rhf.control,
    name   : [
      'ratingEnabled',
      'ytmEnabled',
      'durationEnabled',
      'offerEnabled',
      'currencyEnabled',
    ],
  });
  const [
    ratingEnabled,
    ytmEnabled,
    durationEnabled,
    offerEnabled,
    currencyEnabled
  ] = filters;

  const ratingValues = Array.from(
    new Set((ratings ?? []).map(el => el.ratingValue))
  ).sort(compareRatings).filter(r => r !== '');

  const currencyValues = Array.from(
    new Set((bonds ?? []).map(el => el.currencyID))
  );

  const ytmMax = Math.max(...bonds?.map(bond => bond.ytm) || []) > 150 ? Math.max(...bonds?.map(bond => bond.ytm) || []) : 150;
  const duraionMax = Math.max(...bonds?.map(bond => bond.duration) || []) > 15 ? Math.max(...bonds?.map(bond => bond.duration) || []) : 15;

  return (
    <div className='flex flex-col gap-4'>
      <FilterToggles rhf={ rhf }/>
    
      {
        ratingEnabled
         && <RatingSelect rhf={ rhf } ratingLoading={ ratingLoading } ratingValues={ ratingValues }/>
      }

      {
        ytmEnabled
        && <YtmSlider rhf={ rhf } max={ ytmMax }/>
      }

      {
        durationEnabled
        && <DurationSlider rhf={ rhf } max={ duraionMax }/>
      }
      
      {
        currencyEnabled
        && <CurrencySelect rhf={ rhf } bondsLoading={ bondsLoading } currencyValues={ currencyValues }/>
      }
      
      {
        offerEnabled
         && <OfferToggle rhf={ rhf }/>
      }
    </div>

  );
}