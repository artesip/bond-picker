import { SquareArrowOutUpRight, Trash2 } from 'lucide-react';
import { useNavigate } from '@tanstack/react-router';

import { DeletePicked } from '#/entities/bonds/api';

import { DialogTrigger, DialogContent, DialogHeader, DialogTitle, DialogDescription, Dialog, DialogFooter, DialogClose } from '../ui/dialog';
import { Button } from '../ui/button';

import type { ICellRendererParams } from 'ag-grid-community';
import type { Bond } from '#/entities/bonds/model';

export function DeleteIcon(props: ICellRendererParams) {

  async function handleDeleteClick() {
    if (!props.api) {
      console.error('AG Grid API not available');
      return;
    }
    
    try {
      const selectedNodes = props.api.getSelectedNodes();

      props.api.applyTransaction({
        remove: selectedNodes.map(n => n.data)
      });

      Promise.all(
        selectedNodes.map(el => {
          const bond = el.data as Bond;
          return DeletePicked(bond.id);
        })
      );

    } catch (e) {
      console.log(e);
    }
  }

  return (
    <Dialog>
      <DialogTrigger className='flex items-center justify-center shrink-0 w-full h-full cursor-pointer hover:text-destructive'>
        <Trash2 />
      </DialogTrigger>
      <DialogContent showCloseButton={ false }>
        <DialogHeader>
          <DialogTitle className='text-[18px]'>Удаление из избранного</DialogTitle>
          <DialogDescription>
            Вы уверены что хотите удалить выбранные облигации из избранного ?
          </DialogDescription>
        </DialogHeader>

        <DialogFooter>
          <DialogClose asChild>
            <Button variant={ 'ghost' } onClick={ handleDeleteClick }>Да, удалить</Button>
          </DialogClose>
          
          <DialogClose asChild>
            <Button variant='ghost'>Нет</Button>
          </DialogClose>
        </DialogFooter>
      </DialogContent>
    </Dialog>
  );
}

export function RatingIcon(props: ICellRendererParams) {
  const navigate = useNavigate({ from: '/app/chosen' });

  function onLinkClick() {
    const id = props.data.id;

    navigate({
      search: (prev) => ({
        ...prev,
        id: String(id),
      }),
      replace: false,
    });
  }

  return (
    <button
      className='flex shrink-0 items-center cursor-pointer justify-center mt-2 text-gray-300 opacity-30 hover:opacity-100'
      onClick={ onLinkClick }
    >
      <SquareArrowOutUpRight />
    </button>
  );
}