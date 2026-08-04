import { cn } from '#/lib/utils';

import { CopyButton } from './copy-button';
import { Card, CardHeader, CardTitle, CardContent } from './ui/card';
import { Badge } from './ui/badge';

import type { BondWithRatings } from '#/entities/bonds/model';

type BondCardProps = {
  className?: string
  bond: BondWithRatings
}

const formatDate = (date: Date | null) => {
  if (!date) return '—';
  return new Intl.DateTimeFormat('ru').format(new Date(date));
};

const formatNumber = (value: number) => {
  return new Intl.NumberFormat('ru').format(value);
};

export const BondCard = ({ bond, className }: BondCardProps) => {
  let ratingValue = null;
  let isRevoked = false;
  for (const rating of bond.ratings) {
    if (rating.ratingValue !== '') {
      ratingValue = rating.ratingValue;
      isRevoked = rating.isRevoked;
    }
  }

  return (
    <Card className={ cn('w-full max-w-xl shadow-md rounded-2xl mt-6', className) }>
      <CardHeader>
        <div>
          <CardTitle className='flex text-lg font-semibold gap-2 items-center'>
            {bond.name}

            {ratingValue && <Badge variant='secondary' className='text-[14px]'>{ratingValue}</Badge>}
            {isRevoked && <Badge variant='destructive' className='text-[14px]'>Отозван</Badge>}
            

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