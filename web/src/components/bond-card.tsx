import { useState } from 'react';
import { Check, Copy, Heart } from 'lucide-react';
import { toast } from 'sonner';

import { cn } from '#/lib/utils';
import { DeletePicked, PickBond } from '#/entities/bonds/api.ts';
import { useIsUserLoggedIn } from '#/stores/auth';

import { Card, CardHeader, CardTitle, CardContent } from './ui/card';
import { Badge } from './ui/badge';

import type { BondWithRatings } from '#/entities/bonds/model';

type BondCardProps = {
  className?: string
  bond: BondWithRatings
  isPicked: boolean
  refetch: () => void
}

const formatDate = (date: Date | null) => {
  if (!date) return '—';
  return new Intl.DateTimeFormat('ru').format(new Date(date));
};

const formatNumber = (value: number) => {
  return new Intl.NumberFormat('ru').format(value);
};

const determineCompanyLogoPath = (bond: BondWithRatings) => {
  let companyLogoPath;

  if (bond.name.includes('ОФЗ')) {
    companyLogoPath = 'ofz';
  } else if (bond.companyID === '') {
    // RZD Capital P.L.C. has no inn
    companyLogoPath = '7708503727';
  } else {
    companyLogoPath = bond.companyID;
  }

  return companyLogoPath;
};

export const BondCard = ({ bond, className, isPicked, refetch }: BondCardProps) => {
  const isUserLoggedIn = useIsUserLoggedIn();
  const [copied, setCopied] = useState(false);
  const [isBondPicked, setIsBondPicked] = useState(isPicked);

  const handleCopyIsin = async () => {
    await navigator.clipboard.writeText(bond.isin);
    setCopied(true);

    setTimeout(() => setCopied(false), 1500);
  };

  let ratingValue = null;
  let isRevoked = false;
  for (const rating of bond.ratings) {
    if (rating.ratingValue !== '') {
      ratingValue = rating.ratingValue;
      isRevoked = rating.isRevoked;
    }
  }

  const companyLogoPath = determineCompanyLogoPath(bond);

  const pickBond = async () => {
      try {
        await PickBond(bond.id);
        refetch();
        setIsBondPicked(true);
      } catch (e) {
        if (e instanceof Error) {
          toast.error(e.message);
        }
      }
  };

  const unpickBond = async () => {
    try {
      await DeletePicked(bond.id);
      refetch();
      setIsBondPicked(false);
    } catch (e) {
      if (e instanceof Error) {
        toast.error(e.message);
      }
    }
  };

  return (
    <Card className={ cn('w-full max-w-xl shadow-md rounded-2xl mt-6', className) }>
      <CardHeader>
      <div className={'flex flex-row items-center w-full'}>
        <img src={`/logos/${companyLogoPath}.png`} alt={bond.companyID} className={'h-12 w-12 mr-2'}></img>
        <div className={'w-full'}>
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

            {
                  isUserLoggedIn && <div
                  className='group ml-auto cursor-pointer p-2 rounded-lg hover:bg-muted transition-colors'
                  onClick={isBondPicked ? unpickBond : pickBond}
                >
                  <Heart className={cn('h-5', isBondPicked ? 'text-red-500 fill-red-500 opacity-100' : 'opacity-50 group-hover:text-foreground')}/>
                </div>
            }

          </CardTitle>
          <button
            onClick={ handleCopyIsin }
            title='Скопировать ISIN'
            className='group flex items-center gap-1 text-sm text-muted-foreground cursor-pointer transition-colors hover:text-foreground'
          >
            <span>ISIN: {bond.isin}</span>
            {copied
              ? <Check className='h-3 w-3' />
              : <Copy className='h-3 w-3 opacity-0 transition-opacity group-hover:opacity-100' />}
          </button>
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