import { useNavigate, useSearch } from '@tanstack/react-router';
import { useEffect, useState } from 'react';

import { AddChosenForm } from '#/components/add-chosen-form';
import { BondCard } from '#/components/bond-card';
import { useIsMobile } from '#/hooks/use-mobile';
import { Drawer, DrawerContent } from '#/components/ui/drawer';
import { useBondWithRatings } from '#/entities/bonds/hooks';
import { getBondWithRating } from '#/entities/bonds/model';


type ChosenBondProps = {
    isUserLogedIn: boolean
    refetch: () => void
}

export function ChosenBond({ refetch, isUserLogedIn }: ChosenBondProps) {
  const { id } = useSearch({ from: isUserLogedIn ? '/app/picker' : '/app/watch' });
  const isMobile = useIsMobile();
  const navigate = useNavigate({ from: isUserLogedIn ? '/app/picker' : '/app/watch' });

  const { data: bonds } = useBondWithRatings();

  const selectedBond = getBondWithRating(id || '', bonds?.bonds || [], bonds?.companies || []);

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
          { isUserLogedIn && <AddChosenForm bond={ selectedBond } refetch={ refetch }/> }
        </DrawerContent>
      </Drawer>
    );
  }

  return (
    <div className='flex flex-col gap-2'>
      <BondCard bond={ selectedBond }/> 
      { isUserLogedIn && <AddChosenForm bond={ selectedBond } refetch={ refetch }/> }
    </div>
  );
}