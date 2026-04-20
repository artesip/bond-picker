import { createFileRoute, Link, Outlet, useLocation } from '@tanstack/react-router';
import { Star, InfoIcon, MousePointerClick } from 'lucide-react';
import { useQuery } from '@tanstack/react-query';

import { PortfolioSwitcher } from '#/components/portfolio-switcher';
import { SidebarProvider, SidebarHeader, SidebarContent, SidebarGroup, SidebarFooter, SidebarTrigger, Sidebar, SidebarMenu, SidebarMenuItem, SidebarMenuButton } from '#/components/ui/sidebar';
import { NavUser } from '#/components/nav-user';
import { Me } from '#/entities/auth/api';
import { Separator } from '#/components/ui/separator';
import { Breadcrumb, BreadcrumbList, BreadcrumbItem, BreadcrumbLink } from '#/components/ui/breadcrumb';
import { Alert, AlertDescription } from '#/components/ui/alert';
import { Tooltip, TooltipContent, TooltipTrigger } from '#/components/ui/tooltip';
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
  const portfolios = ['default'];
  const { data: user } = useQuery({ queryKey: ['me'], queryFn: Me });

  const { pathname } = useLocation();
  const currentBreadLink = buttons.filter((button) => button.url === pathname)[0];

  return (
    <SidebarProvider className='h-screen'>
      <Sidebar variant='floating'>
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
                        // activeOptions={ { exact: true } }
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
          <NavUser
            user={ { username: user?.username || '', avatar: 'https://img.daisyui.com/images/profile/demo/yellingcat@192.webp' } }
          />
          <span className='text-gray-600 text-[14px]'>v1.0.0</span>
        </SidebarFooter>
      </Sidebar>
      
      <div className='flex flex-col p-2 h-full w-full'>
        <header className='flex h-12 shrink-0 items-center gap-2 px-2'>
          <SidebarTrigger className='-ml-1'/>
          
          <div>
            <Separator
              orientation='vertical'
              className='mr-2 data-[orientation=vertical]:h-5'
            />
          </div>


          <Breadcrumb>
            <BreadcrumbList>
              <BreadcrumbItem className='block text-[16px]'>
                <BreadcrumbLink href={ currentBreadLink.url }>{currentBreadLink.title}</BreadcrumbLink>
              </BreadcrumbItem>
            </BreadcrumbList>
          </Breadcrumb>
        
          {
            currentBreadLink.url === '/app/picker'
              && <div className='ml-auto'>
                <Tooltip>
                  <TooltipTrigger asChild>
                    <Alert className='p-1 lg:py-2 lg:px-3'>
                      <InfoIcon />
                      <AlertDescription className='truncate!'>ИИ может ошибаться. Не является индивидуальной инвестиционной рекомендацией</AlertDescription>
                    </Alert>
                  </TooltipTrigger>
                  <TooltipContent className='items-center'>
                    <p>ИИ может ошибаться. Не является индивидуальной инвестиционной рекомендацией</p>
                  </TooltipContent>
                </Tooltip>
              </div>
          }
          
        </header>

        <main className='mt-2 ml-2 mr-2 flex-1'>
          <Outlet />
        </main>
      </div>
      
    </SidebarProvider>
  );
}