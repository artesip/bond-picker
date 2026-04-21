import { useSearch } from '@tanstack/react-router';

import { Card, CardContent, CardHeader, CardTitle } from '#/components/ui/card';
import { useBondWithRatings, usePickedBonds } from '#/entities/bonds/hooks';
import { inRange  } from '#/entities/bonds/model';
import { useFilterForm } from '#/entities/bonds/shemas';
import { Badge } from '#/components/ui/badge';
import { CopyButton } from '#/components/copy-button';
import { AddChosenForm } from '#/components/add-chosen-form';

import { BondChart } from './ui/chart';
import { FilterBlock } from './ui/filters';

import type { Bond } from '#/entities/bonds/model';

type BondCardProps = {
  bond: Bond
}

const formatDate = (date: Date | null) => {
  if (!date) return '—';
  return new Intl.DateTimeFormat('ru').format(new Date(date));
};

const formatNumber = (value: number) => {
  return new Intl.NumberFormat('ru').format(value);
};

export const BondCard = ({ bond }: BondCardProps) => {
  return (
    <Card className='w-full max-w-xl shadow-md rounded-2xl mt-6'>
      <CardHeader>
        <div>
          <CardTitle className='flex text-lg font-semibold gap-2 items-center'>
            {bond.name}

            {bond.callOption && (
              <Badge variant='secondary' className='text-[14px]'>Call</Badge>
            )}
            {bond.putOption && (
              <Badge variant='secondary' className='text-[14px]'>Put</Badge>
            )}

            <div className='ml-auto'>
              <CopyButton value={ bond.isin }/>
            </div>
          </CardTitle>
          <div className='text-sm text-muted-foreground'>
            ISIN: {bond.isin}
          </div>
        </div>
      </CardHeader>

      <CardContent className='grid grid-cols-2 gap-4 text-sm'>
        <div>
          <span className='text-muted-foreground'>Цена:</span>
          <div>{bond.price.toFixed(2)}</div>
        </div>

        <div>
          <span className='text-muted-foreground'>YTM:</span>
          <div>{bond.ytm.toFixed(2)}%</div>
        </div>

        <div>
          <span className='text-muted-foreground'>Дюрация:</span>
          <div>{bond.duration.toFixed(2)}</div>
        </div>

        <div>
          <span className='text-muted-foreground'>Купон / Частота:</span>
          <div>{bond.couponPercent}% / {bond.couponPeriod}</div>
        </div>

        <div>
          <span className='text-muted-foreground'>Номинал:</span>
          <div>{formatNumber(bond.faceValue)}</div>
        </div>

        <div>
          <span className='text-muted-foreground'>НКД:</span>
          <div>{bond.acruedint.toFixed(2)}</div>
        </div>

        <div>
          <span className='text-muted-foreground'>Следующий купон:</span>
          <div>{formatDate(bond.nextCoupon)}</div>
        </div>

        <div>
          <span className='text-muted-foreground'>Погашение:</span>
          <div>{formatDate(bond.matDate)}</div>
        </div>

        <div>
          <span className='text-muted-foreground'>Размер лота:</span>
          <div>{bond.lotSize}</div>
        </div>

        <div>
          <span className='text-muted-foreground'>Размер выпуска:</span>
          <div>{formatNumber(bond.issueSize)} шт.</div>
        </div>
      </CardContent>
    </Card>
  );
};

export function PickerPage() {
  const { rhf } = useFilterForm();
  const { data, isLoading: bondsLoading } = useBondWithRatings();
  const { data: pickedBonds, isLoading: pickedBondsLoading, refetch } = usePickedBonds();

  const filters = rhf.watch();
  const { id } = useSearch({ from: '/app/picker' });

  const filtered = data
    .filter(bond => !filters.ytmEnabled || (bond.ytm >= filters.ytmFrom && bond.ytm <= filters.ytmTo))
    .filter(bond => !filters.durationEnabled ||  (bond.duration >= filters.durationFrom && bond.duration <= filters.durationTo))
    .filter(bond => !filters.currencyEnabled || bond.currencyID === filters.currency)
    .filter(bond => !filters.offerEnabled || filters.offer === 'all' || (filters.offer === 'no' && (bond.callOption === null && bond.putOption === null)) 
                      || (filters.offer === 'put' && (bond.putOption !== null)) || (filters.offer === 'call' && (bond.callOption !== null)))
    .filter(bond => !filters.ratingEnabled || (inRange(filters.ratingFrom, filters.ratingTo, bond.ratings?.[0]?.ratingValue ?? '')));

  const isLoading = bondsLoading || pickedBondsLoading;

  return (
    <div className='flex h-full flex-row gap-4'>
      <Card className='flex flex-1 w-9/12 p-0'>
        <BondChart isLoading={ isLoading } data={ filtered } picked={ pickedBonds || [] }/>
      </Card>

      <div className='w-3/12'>
        <FilterBlock rhf={ rhf }/>
        <span className='text-muted-foreground text-[14px]' >Облигаций {filtered.length}. Всего {data.length}</span>

        {
          !isLoading && id && data 
            && <div className='flex flex-col gap-2'>
              <BondCard bond={ data.filter(bond => bond.id === id)[0] }/> 
              <AddChosenForm bond={ data.filter(bond => bond.id === id)[0] } refetch={ refetch }/>
            </div>
        } 
      </div>
    </div>
  );
}