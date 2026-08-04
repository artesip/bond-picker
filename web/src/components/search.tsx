import { Search } from 'lucide-react';
import { useState } from 'react';

import { Button } from '@/components/ui/button';
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

import { Dialog, DialogContent } from './ui/dialog';
import { BondCard } from './bond-card';
import { AddChosenForm } from './add-chosen-form';

type BondSearchProps = {
  isUserLogedIn: boolean
}

export function BondSearch({isUserLogedIn}: BondSearchProps) {
  const [open, setOpen] = useState(false);
  const [bondOpen, setBondOpen] = useState(false);
  const [value, setValue] = useState('');
  const { data: bonds, isLoading } = useBondWithRatings();
  const { refetch } = usePickedBonds(isUserLogedIn);

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

  return (
    <div className='flex flex-col gap-4 w-full'>
      <Dialog open={ bondOpen } onOpenChange={ setBondOpen }>

        <DialogContent showCloseButton={ false } className='bg-transparent! border-0! ring-0! gap-4'>
          {bonds && <>
            <BondCard bond={ bonds.filter(bond => bond.name === value)[0] }/>
            <AddChosenForm bond={ bonds.filter(bond => bond.name === value)[0] } refetch={ refetch }/>
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
      } }>
        <Command
          filter={ filter }    
        >
          <CommandInput placeholder='Введите имя или isin облигации' />
          <CommandList>
            <CommandEmpty>No results found.</CommandEmpty>
            <CommandGroup heading='Совпадения'>
              {!isLoading && bonds.map(bond => (
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
              ))
              }
            </CommandGroup>
          </CommandList>
        </Command>
      </CommandDialog>
    </div>
  );
}
