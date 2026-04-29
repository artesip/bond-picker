'use client';

import {
  CalendarClock,
  ChevronsUpDown,
  LogOut,
} from 'lucide-react';
import { toast } from 'sonner';
import { useNavigate } from '@tanstack/react-router';

import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuTrigger,
} from '@/components/ui/dropdown-menu';
import {
  SidebarMenu,
  SidebarMenuButton,
  SidebarMenuItem,
  useSidebar,
} from '@/components/ui/sidebar';
import {
  Avatar,
  AvatarFallback,
  AvatarImage,
} from '@/components/ui/avatar';
import { Logout } from '#/entities/auth/api';

import { Dialog, DialogClose, DialogContent, DialogDescription, DialogFooter, DialogHeader, DialogTitle, DialogTrigger } from './ui/dialog';
import { Button } from './ui/button';

export function NavUser({
  user,
}: {
  user: {
    username: string
    avatar: string
  }
}) {
  const { isMobile } = useSidebar();
  const router = useNavigate();
  
  async function onUserExit() {
    try {
      await Logout();
      router({ to: '/login' });  
    } catch (e) {
      if (e instanceof Error) {
        toast.error(e.message);
      }
    }
  }

  return (
    <Dialog>
      <SidebarMenu>
        <SidebarMenuItem>
          <DropdownMenu>
            <DropdownMenuTrigger asChild>
              <SidebarMenuButton
                size='lg'
                className='data-[state=open]:bg-sidebar-accent data-[state=open]:text-sidebar-accent-foreground'
              >
                <Avatar className='h-8 w-8 rounded-lg'>
                  <AvatarImage src={ user.avatar } alt={ user.username } />
                  <AvatarFallback className='rounded-lg'>BP</AvatarFallback>
                </Avatar>
                <div className='grid flex-1 text-left text-sm leading-tight'>
                  <span className='truncate font-medium'>{user.username}</span>
                </div>
                <ChevronsUpDown className='ml-auto size-4' />
              </SidebarMenuButton>
            </DropdownMenuTrigger>
            <DropdownMenuContent
              className='w-(--radix-dropdown-menu-trigger-width) min-w-56 rounded-lg'
              side={ isMobile ? 'bottom' : 'top' }
              align='end'
              sideOffset={ 4 }
            >
              <DropdownMenuItem disabled>
                <CalendarClock />
                Обновление данных
              </DropdownMenuItem>
              <DialogTrigger asChild>
                <DropdownMenuItem>
                  <LogOut />
                  Выход
                </DropdownMenuItem>
              </DialogTrigger>
            </DropdownMenuContent>
          </DropdownMenu>
        </SidebarMenuItem>
      </SidebarMenu>

      <DialogContent showCloseButton={ false }>
        <DialogHeader>
          <DialogTitle className='text-[18px]'>Выход из аккаунта</DialogTitle>
          <DialogDescription>
            Вы уверены, что хотите выйти из аккаунта ?
          </DialogDescription>
        </DialogHeader>
        <DialogFooter>
          <Button variant={ 'ghost' } onClick={ onUserExit }>Да, выйти</Button>
          <DialogClose asChild>
            <Button variant={ 'ghost' }>Нет</Button>
          </DialogClose>
        </DialogFooter>
      </DialogContent>
      
    </Dialog>
  );
}
