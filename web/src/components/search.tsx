import { Search } from 'lucide-react';
import { useMemo, useState } from 'react';

import { Button } from '@/components/ui/button';
import { Skeleton } from '@/components/ui/skeleton';
import {
  Command,
  CommandDialog,
  CommandEmpty,
  CommandGroup,
  CommandInput,
  CommandItem,
  CommandList,
} from '@/components/ui/command';
import { useBondWithRatings, usePickedBonds } from '#/entities/bonds/hooks';
import { getBondWithRating } from '#/entities/bonds/model';

import { Dialog, DialogContent } from './ui/dialog';
import { BondCard } from './bond-card';
import { AddChosenForm } from './add-chosen-form';

type BondSearchProps = {
  isUserLogedIn: boolean
}

const MAX_RESULTS = 50;

function isIsinQuery(search: string) {
  const s = search.trim().toUpperCase();

  return /^[A-Z0-9]{2,12}$/.test(s);
}

function filter(value: string, search: string) {
  const [name, isin] = value.split('|');

  if (!name || !isin) {
    return 1;
  }

  const s = search.toLowerCase();
  const n = name.toLowerCase();
  const i = isin.toLowerCase();

  if (!s) return 1;

  if (isIsinQuery(search)) {
    return i.includes(s) ? 1 : 0;
  }

  return n.includes(s) ? 1 : 0;
}

export function BondSearch({ isUserLogedIn }: BondSearchProps) {
  const [open, setOpen] = useState(false);
  const [bondOpen, setBondOpen] = useState(false);
  const [value, setValue] = useState('');
  const [search, setSearch] = useState('');
  const { data: bonds, isLoading } = useBondWithRatings();
  const { refetch } = usePickedBonds(isUserLogedIn);

  const allBonds = bonds?.bonds || [];

  const shownBonds = useMemo(() => {
    if (!search.trim()) {
      return [];
    }

    return (bonds?.bonds || [])
      .filter((bond) => filter(`${bond.name}|${bond.isin ?? ''}`, search) === 1)
      .slice(0, MAX_RESULTS);
  }, [bonds, search]);

  const bond = allBonds.find(bond =>bond.name === value);
  const bondWithRatings = getBondWithRating( bond === undefined ? '' : bond.id, allBonds, bonds?.companies || []);

  return (
    <div className='flex flex-col gap-4 w-full'>
      <Dialog open={ bondOpen } onOpenChange={ setBondOpen }>

        <DialogContent showCloseButton={ false } className='bg-transparent! border-0! ring-0! gap-4'>
          {
            bondWithRatings 
     && <>
            <BondCard bond={ bondWithRatings }/>
            { isUserLogedIn && <AddChosenForm bond={ bondWithRatings } refetch={ refetch }/> }
          </>
          }
        </DialogContent>
      </Dialog>

      <Button onClick={ () => setOpen(true) } variant='ghost' className='border border-input'>
        <Search/>
        Поиск облигаций
      </Button>
      <CommandDialog open={ open } onOpenChange={ (open) => {
        setOpen(open);
        setValue('');
        setSearch('');
      } }>
        <Command
          filter={ filter }    
        >
          <CommandInput
            value={ search }
            onValueChange={ setSearch }
            placeholder='Введите имя или ISIN облигации'
          />
          <CommandList>
            {isLoading
              ? (
                <div className='flex flex-col gap-2 p-2'>
                  {Array.from({ length: 6 }).map((_, index) => (
                    <Skeleton key={ index } className='h-5 w-full' />
                  ))}
                </div>
              )
              : (
                <>
                  {
                    search.trim()
                  && <CommandEmpty>
                    Ничего не найдено.
                  </CommandEmpty>
                  }
                  <CommandGroup heading='Совпадения'>
                    {shownBonds.map(bond => (
                      <CommandItem
                        key={ bond.id }
                        value={ `${bond.name}|${bond.isin ?? ''}` }
                        onSelect={ (value) => {
                          setOpen(false);
                          setBondOpen(true);
                          setValue(value.split('|')[0]);
                        } }
                      >
                        {bond.name}
                      </CommandItem>
                    ))}
                  </CommandGroup>
                </>
              )}
          </CommandList>
        </Command>
      </CommandDialog>
    </div>
  );
}
