import { useNavigate } from '@tanstack/react-router';
import { useEffect, useState } from 'react';

import { BondCard } from '#/components/bond-card';
import { useIsMobile } from '#/hooks/use-mobile';
import { Drawer, DrawerContent } from '#/components/ui/drawer';
import { useBondWithRatings, usePickedBonds } from '#/entities/bonds/hook';
import { getBondWithRating } from '#/entities/bonds/model';
import { useIsUserLoggedIn } from '#/stores/auth';
import { useSelectedBondId } from '#/stores/selected-bond';


export function ChosenBond() {
  const isUserLoggedIn = useIsUserLoggedIn();
  const id = useSelectedBondId();
  const { data: pickedBonds, isLoading: pickedBondsLoading, refetch } = usePickedBonds(isUserLoggedIn);

  const isMobile = useIsMobile();
  const navigate = useNavigate({ from: isUserLoggedIn ? '/app/picker' : '/app/watch' });

  const { data: bonds } = useBondWithRatings();

  const selectedBond = getBondWithRating(id || '', bonds?.bonds || [], bonds?.companies || []);
  const isPicked = (pickedBonds || []).some(bond => bond.id === id);

  const [open, setOpen] = useState(false);
  useEffect(() => {
    if (isMobile && id) {
      // eslint-disable-next-line react-hooks/set-state-in-effect
      setOpen(true);
    }
  }, [id, isMobile]);

  if (!id || !selectedBond || pickedBondsLoading) {
    return <div></div>;
  }

  if (isMobile) {
    return (
      <Drawer open={ open } onOpenChange={ (e) => {
        setOpen(e);
        navigate({
          search: (prev) => {
            const { ...rest } = prev;
            return rest;
          },
          resetScroll   : false,
          viewTransition: true,
        }); 
      } }>
        <DrawerContent className='gap-4 mb-4 px-2 items-center justify-center min-w-90'>
          <BondCard
            key={ selectedBond.id }
            bond={ selectedBond }
            className='bg-transparent! border-0! ring-0 shadow-none mt-0'
            isPicked={isPicked}
            refetch={refetch}
          />
        </DrawerContent>
      </Drawer>
    );
  }

  return (
    <div className='flex flex-col gap-2'>
      <BondCard
        key={ selectedBond.id }
        bond={ selectedBond }
        isPicked={isPicked}
        refetch={refetch}
      />
    </div>
  );
}