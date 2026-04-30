import { useNavigate, useSearch } from '@tanstack/react-router';
import { useEffect, useState } from 'react';

import { useRatings } from '#/entities/bonds/hooks';
import { Drawer, DrawerContent, DrawerHeader } from '#/components/ui/drawer';

import { RatingItem } from './rating-item';

import type { Bond } from '#/entities/bonds/model';


type RatingDrawerProps = {
    bonds: Bond[]
}

export function RatingDrawer({ bonds }: RatingDrawerProps) {
  const { id } = useSearch({ from: '/app/chosen' });
  const { data: ratings } = useRatings();
  const navigate = useNavigate({ from: '/app/chosen' });
  
  const selectedBond = bonds.find(bond => bond.id === id);
  const bondRatings = ratings?.filter(rating => rating.companyID === selectedBond?.companyID) || []; 

  const [open, setOpen] = useState(false);
  useEffect(() => {
    if (id) {
      // eslint-disable-next-line react-hooks/set-state-in-effect
      setOpen(true);
    }
  }, [id]);

  return (
    <Drawer direction='right' open={ open } onOpenChange={ (e) => {
      setOpen(e);
      navigate({
        search: (prev) => {
          const { id, ...rest } = prev;
          return rest;
        },
      });
    } }>
      <DrawerContent className='max-w-md!'>
        <DrawerHeader className='text-[18px]'>
          История рейтингов
        </DrawerHeader>

        <ul className='flex flex-col gap-2 p-2'>
          {bondRatings.map(rating => 
            <RatingItem rating={ rating } key={ rating.id }/>
          )}
        </ul>
      </DrawerContent>
    </Drawer>
  );
}