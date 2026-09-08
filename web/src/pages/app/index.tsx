import { useQuery } from '@tanstack/react-query';
import { useLocation, Outlet, Link } from '@tanstack/react-router';


import { NavUser } from '#/pages/app/ui/nav-user';
import { PortfolioSwitcher } from '#/pages/app/ui/portfolio-switcher';
import { BondSearch } from '#/components/search';
import { SidebarProvider, SidebarHeader, SidebarContent, SidebarGroup, SidebarMenu, SidebarMenuItem, SidebarMenuButton, SidebarFooter, Sidebar } from '#/components/ui/sidebar';
import { Me } from '#/entities/auth/api';
import { useAuthStore } from '#/stores/auth';

import { buttons } from './buttons';
import { Header } from './ui/header';

export function App() {
  const { pathname } = useLocation();
  const isUserLogedIn = pathname !== '/app/watch';
  const setIsUserLoggedIn = useAuthStore((s) => s.setIsUserLoggedIn);

  if (useAuthStore().isUserLoggedIn !== isUserLogedIn) {
    setIsUserLoggedIn(isUserLogedIn);
  }

  const portfolios = ['default'];
  const { data: user } = useQuery({ queryKey: ['me'], queryFn: Me, enabled: isUserLogedIn });

  

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
                        preload={'intent'}
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
          <BondSearch/>
          
          <NavUser
            user={ { username: user?.username || '', avatar: 'https://img.daisyui.com/images/profile/demo/yellingcat@192.webp' } }
          />
          <span className='text-gray-600 text-[14px]'>v1.0.0</span>
        </SidebarFooter>
      </Sidebar>
      }
      
      <div className='flex flex-col p-2 h-full w-full'>
        <Header/>

        <main className='mt-2 ml-2 mr-2 flex-1'>
          <Outlet />
        </main>
      </div>
      
    </SidebarProvider>
  );
}