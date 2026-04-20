import { GalleryVerticalEnd, ChevronsUpDown, Check } from 'lucide-react';
import { useState } from 'react';

import {
  SidebarMenu,
  SidebarMenuButton,
  SidebarMenuItem,
} from '@/components/ui/sidebar';


import { DropdownMenuTrigger, DropdownMenuContent, DropdownMenuItem, DropdownMenu } from './ui/dropdown-menu';

type PortfolioSwitcherProps = {
  portfolios: string[]
  defaultPortfolio: string
}

export function PortfolioSwitcher({
  portfolios,
  defaultPortfolio,
}: PortfolioSwitcherProps) {
  const DEFAULT_PORTFOLIO = 'default';

  const [selectedPortfolio, setSelectedPortfolio] = useState(defaultPortfolio);

  return (
    <SidebarMenu>
      <SidebarMenuItem>
        <DropdownMenu>
          <DropdownMenuTrigger asChild>
            <SidebarMenuButton
              size='lg'
              className='data-[state=open]:bg-sidebar-accent data-[state=open]:text-sidebar-accent-foreground'
            >
              <div className='flex aspect-square size-8 items-center justify-center rounded-lg bg-sidebar-primary text-sidebar-primary-foreground'>
                <GalleryVerticalEnd className='size-4' />
              </div>
              <div className='flex flex-col gap-0.5 leading-none'>
                <span className='font-medium'>Портфель</span>
                <span className=''>{selectedPortfolio === DEFAULT_PORTFOLIO ? 'По умолчанию' : selectedPortfolio}</span>
              </div>
              <ChevronsUpDown className='ml-auto' />
            </SidebarMenuButton>
          </DropdownMenuTrigger>
          <DropdownMenuContent
            className='w-(--radix-dropdown-menu-trigger-width)'
            align='start'
          >
            {portfolios.map((version) => (
              <DropdownMenuItem
                key={ version }
                onSelect={ () => setSelectedPortfolio(version) }
              >
                {version === DEFAULT_PORTFOLIO ? 'По умолчанию' : selectedPortfolio}{' '}
                {version === selectedPortfolio && <Check className='ml-auto' />}
              </DropdownMenuItem>
            ))}
          </DropdownMenuContent>
        </DropdownMenu>
      </SidebarMenuItem>
    </SidebarMenu>
  );
}