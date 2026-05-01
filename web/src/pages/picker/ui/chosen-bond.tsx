import { useNavigate, useSearch } from '@tanstack/react-router';
import { useEffect, useState } from 'react';

import { AddChosenForm } from '#/components/add-chosen-form';
import { BondCard } from '#/components/bond-card';
import { useIsMobile } from '#/hooks/use-mobile';
import { Drawer, DrawerContent } from '#/components/ui/drawer';

import type { BondWithRatings } from '#/entities/bonds/model';

type ChosenBondProps = {
    data: BondWithRatings[]
    refetch: () => void
}

export function ChosenBond({ data, refetch }: ChosenBondProps) {
  const { id } = useSearch({ from: '/app/picker' });
  const isMobile = useIsMobile();
  const selectedBond = data.find(bond => bond.id === id);
  const navigate = useNavigate({ from: '/app/picker' });

  const [open, setOpen] = useState(false);
  useEffect(() => {
    if (isMobile && id) {
      // eslint-disable-next-line react-hooks/set-state-in-effect
      setOpen(true);
    }
  }, [id, isMobile]);

  if (!id || !selectedBond) {
    return <div></div>;
  }

  if (isMobile) {
    return (
      <Drawer open={ open } onOpenChange={ (e) => {
        setOpen(e);
        navigate({
          search: (prev) => {
            const { id, ...rest } = prev;
            return rest;
          },
          resetScroll   : false,
          viewTransition: true,
        }); 
      } }>
        <DrawerContent className='gap-4 mb-4 px-2'>
          <BondCard bond={ selectedBond } className='bg-transparent! border-0! ring-0 shadow-none mt-0'/> 
          <AddChosenForm bond={ selectedBond } refetch={ refetch }/>
        </DrawerContent>
      </Drawer>
    );
  }

  return (
    <div className='flex flex-col gap-2'>
      <BondCard bond={ data.filter(bond => bond.id === id)[0] }/> 
      <AddChosenForm bond={ data.filter(bond => bond.id === id)[0] } refetch={ refetch }/>
    </div>
  );
}