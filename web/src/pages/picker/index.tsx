import { useMemo } from 'react';
import { useWatch } from 'react-hook-form';

import { Card } from '#/components/ui/card';
import { useBondWithRatings, usePickedBonds } from '#/entities/bonds/hook';
import { getLastNotRevokedRatings, inRange  } from '#/entities/bonds/model';
import { useFilterForm } from '#/entities/bonds/shemas';
import { useIsUserLoggedIn } from '#/stores/auth';

import { ScatterChart } from './ui/scatter-chart';
import { FilterBlock } from './ui/filters';
import { ChosenBond } from './ui/chosen-bond';
import { SelectedBondSync } from './ui/selected-bond-sync';

export function PickerPage() {
  const isUserLoggedIn = useIsUserLoggedIn();
  const { rhf } = useFilterForm();

  const { data: bonds, isLoading: bondsLoading } = useBondWithRatings();

  const { data: pickedBonds, isLoading: pickedBondsLoading } = usePickedBonds(isUserLoggedIn);

  const [
    ratingEnabled,
    ytmEnabled,
    durationEnabled,
    offerEnabled,
    currencyEnabled,
    ratingFrom,
    ratingTo,
    ytmFrom,
    ytmTo,
    durationFrom,
    durationTo,
    currency,
    offer,
  ] = useWatch({
    control: rhf.control,
    name   : [
      'ratingEnabled',
      'ytmEnabled',
      'durationEnabled',
      'offerEnabled',
      'currencyEnabled',
      'ratingFrom',
      'ratingTo',
      'ytmFrom',
      'ytmTo',
      'durationFrom',
      'durationTo',
      'currency',
      'offer',
    ],
  });

  const mapOfRatings = useMemo(
    () => getLastNotRevokedRatings(bonds?.companies || []),
    [bonds]
  );

  const filtered = useMemo(() => (bonds?.bonds || [])
    .filter(bond => bond.type === 'fix')
    .filter(bond => !ytmEnabled || (bond.ytm >= ytmFrom && bond.ytm <= ytmTo))
    .filter(bond => !durationEnabled ||  (bond.duration >= durationFrom && bond.duration <= durationTo))
    .filter(bond => !currencyEnabled || bond.currencyID === currency)
    .filter(bond => !offerEnabled || offer === 'all' || (offer === 'no' && (bond.callOption === null && bond.putOption === null))
                      || (offer === 'put' && (bond.putOption !== null)) || (offer === 'call' && (bond.callOption !== null)))
    .filter(bond => !ratingEnabled || (inRange(ratingFrom, ratingTo, mapOfRatings.get(bond.companyID) ?? ''))),
  [
    bonds,
    mapOfRatings,
    ratingEnabled,
    ytmEnabled,
    durationEnabled,
    offerEnabled,
    currencyEnabled,
    ratingFrom,
    ratingTo,
    ytmFrom,
    ytmTo,
    durationFrom,
    durationTo,
    currency,
    offer,
  ]);

  const isLoading = bondsLoading || pickedBondsLoading;

  return (
    <div className='grid grid-cols-1 lg:grid-cols-10 h-full gap-4'>
      <SelectedBondSync from={ isUserLoggedIn ? '/app/picker' : '/app/watch' }/>

      <Card className='grid col-span-1 lg:col-span-7 p-0 not-lg:order-2 min-h-150'>
        <ScatterChart isLoading={ isLoading } data={ filtered || [] } picked={ pickedBonds || [] }/>
      </Card>

      <div className='grid content-start col-span-1 lg:col-span-3 not-lg:order-1 w-full'>
        <FilterBlock rhf={ rhf }/>
        <span className='text-muted-foreground text-[14px] mt-2' >Показано {filtered.length}. Всего {bonds?.bonds.length}</span>

        {
          !isLoading && bonds?.bonds
            && <ChosenBond/>
        }
      </div>
    </div>
  );
}