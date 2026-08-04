import { createFileRoute, Link, Outlet, useLocation } from '@tanstack/react-router';
import { Star, MousePointerClick } from 'lucide-react';
import { useQuery } from '@tanstack/react-query';

import { PortfolioSwitcher } from '#/components/portfolio-switcher';
import { SidebarProvider, SidebarHeader, SidebarContent, SidebarGroup, SidebarFooter, SidebarTrigger, Sidebar, SidebarMenu, SidebarMenuItem, SidebarMenuButton } from '#/components/ui/sidebar';
import { NavUser } from '#/components/nav-user';
import { Me } from '#/entities/auth/api';
import { Separator } from '#/components/ui/separator';
import { Breadcrumb, BreadcrumbList, BreadcrumbItem, BreadcrumbLink } from '#/components/ui/breadcrumb';
import { Alert, AlertDescription } from '#/components/ui/alert';
import { Tooltip, TooltipContent, TooltipTrigger } from '#/components/ui/tooltip';
import { BondSearch } from '#/components/search';
import { useKeyRate } from '#/entities/bonds/hooks';
import { Button } from '#/components/ui/button';

export const Route = createFileRoute('/app')({
  component: AppLayout,
});

const buttons = [
  {
    icon : <Star />,
    url  : '/app/chosen',
    title: 'Избранное',
  },
  {
    icon : <MousePointerClick/>,
    url  : '/app/picker',
    title: 'Выбор облигаций',
  },
];

function AppLayout() {
  const { pathname } = useLocation();
  const isUserLogedIn = pathname !== '/app/watch';

  const portfolios = ['default'];
  const { data: user } = useQuery({ queryKey: ['me'], queryFn: Me, enabled: isUserLogedIn });
  const { data: keyRate, isLoading: keyRateLoading } = useKeyRate();

  const currentBreadLink = buttons.filter((button) => button.url === pathname)[0];

  return (
    <SidebarProvider className='h-screen'>
      {isUserLogedIn && <Sidebar variant='floating'>
        <SidebarHeader > 
          <PortfolioSwitcher defaultPortfolio='default' portfolios={ portfolios }/>

        </SidebarHeader>
        <SidebarContent>
          <SidebarGroup>
            <SidebarMenu>
              {
                buttons.map(item => 
                  <SidebarMenuItem key={ item.title }>
                    <SidebarMenuButton asChild className='text-[16px]'>
                      <Link
                        to={ item.url }
                        activeProps={ { className: 'bg-sidebar-accent text-sidebar-accent-foreground' } }
                      >
                        {item.icon}
                        {item.title}
                      </Link>
                    </SidebarMenuButton>
                  </SidebarMenuItem>
                )
              }
            </SidebarMenu>
          </SidebarGroup>
          
        </SidebarContent>
        <SidebarFooter className='items-center'>
          <BondSearch isUserLogedIn={ isUserLogedIn }/>
          
          <NavUser
            user={ { username: user?.username || '', avatar: 'https://img.daisyui.com/images/profile/demo/yellingcat@192.webp' } }
          />
          <span className='text-gray-600 text-[14px]'>v1.0.0</span>
        </SidebarFooter>
      </Sidebar>
      }
      
      <div className='flex flex-col p-2 h-full w-full'>
        <header className='flex h-12 shrink-0 items-center gap-2 px-2'>
          { isUserLogedIn 
          && <>
            <SidebarTrigger className='-ml-1'/>
            <div>
              <Separator
                orientation='vertical'
                className='mr-2 data-[orientation=vertical]:h-5'
              />
            </div>
          </>
          }


          {
            currentBreadLink && <Breadcrumb>
              <BreadcrumbList>
                <BreadcrumbItem className='block text-[16px]'>
                  <BreadcrumbLink href={ currentBreadLink.url }>{currentBreadLink.title}</BreadcrumbLink>
                </BreadcrumbItem>
              </BreadcrumbList>
            </Breadcrumb>
          }
        
          {
            currentBreadLink && currentBreadLink.url === '/app/picker'
              && <div className='ml-auto'>
                <Tooltip>
                  <TooltipTrigger asChild>
                    <Alert className='p-1 lg:py-2 lg:px-3'>
                      <AlertDescription className='truncate!'>ИИ может ошибаться. Не является индивидуальной инвестиционной рекомендацией</AlertDescription>
                    </Alert>
                  </TooltipTrigger>
                  <TooltipContent className='items-center'>
                    <p>ИИ может ошибаться. Не является индивидуальной инвестиционной рекомендацией</p>
                  </TooltipContent>
                </Tooltip>
              </div>
          }

          {
            !keyRateLoading && currentBreadLink && currentBreadLink.url === '/app/chosen'
              && <div className='ml-auto'>
                <Tooltip>
                  <TooltipTrigger asChild>
                    <Alert className='p-1 lg:py-2 lg:px-3'>
                      <AlertDescription className='truncate!'>{keyRate}% ─ Ключевая ставка ЦБ РФ</AlertDescription>
                    </Alert>
                  </TooltipTrigger>
                  <TooltipContent className='items-center'>
                    <p>{keyRate}% ─ Ключевая ставка ЦБ РФ</p>
                  </TooltipContent>
                </Tooltip>
              </div>
          }

          {
            !isUserLogedIn
            && <a href='/login' className='ml-auto w-23 cursor-pointer'> 
              <Button variant={ 'secondary' } className='w-full cursor-pointer'>
                Выйти
              </Button>
            </a>
          }
          
        </header>

        <main className='mt-2 ml-2 mr-2 flex-1'>
          <Outlet />
        </main>
      </div>
      
    </SidebarProvider>
  );
}