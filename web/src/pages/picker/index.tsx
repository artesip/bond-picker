import { Card } from '#/components/ui/card';
import { useBondWithRatings, usePickedBonds } from '#/entities/bonds/hooks';
import { inRange  } from '#/entities/bonds/model';
import { useFilterForm } from '#/entities/bonds/shemas';

import { BondChart } from './ui/chart';
import { FilterBlock } from './ui/filters';
import { ChosenBond } from './ui/chosen-bond';


export function PickerPage() {
  const { rhf } = useFilterForm();
  const { data, isLoading: bondsLoading } = useBondWithRatings();
  const { data: pickedBonds, isLoading: pickedBondsLoading, refetch } = usePickedBonds();

  const filters = rhf.watch();
  

  const filtered = data
    .filter(bond => !filters.ytmEnabled || (bond.ytm >= filters.ytmFrom && bond.ytm <= filters.ytmTo))
    .filter(bond => !filters.durationEnabled ||  (bond.duration >= filters.durationFrom && bond.duration <= filters.durationTo))
    .filter(bond => !filters.currencyEnabled || bond.currencyID === filters.currency)
    .filter(bond => !filters.offerEnabled || filters.offer === 'all' || (filters.offer === 'no' && (bond.callOption === null && bond.putOption === null)) 
                      || (filters.offer === 'put' && (bond.putOption !== null)) || (filters.offer === 'call' && (bond.callOption !== null)))
    .filter(bond => !filters.ratingEnabled || (inRange(filters.ratingFrom, filters.ratingTo, bond.ratings?.[0]?.ratingValue ?? '')));

  const isLoading = bondsLoading || pickedBondsLoading;

  return (
    <div className='grid grid-cols-1 lg:grid-cols-10 h-full gap-4'>
      <Card className='grid col-span-1 lg:col-span-7 p-0 not-lg:order-2 min-h-150'>
        <BondChart isLoading={ isLoading } data={ filtered } picked={ pickedBonds || [] }/>
      </Card>

      <div className='grid content-start col-span-1 lg:col-span-3 not-lg:order-1 w-full'>
        <FilterBlock rhf={ rhf }/>
        <span className='text-muted-foreground text-[14px] mt-2' >Облигаций {filtered.length}. Всего {data.length}</span>

        {
          !isLoading && data 
            && <ChosenBond data={ data } refetch={ refetch }/>
        } 
      </div>
    </div>
  );
}