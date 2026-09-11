import { Link, useLocation } from '@tanstack/react-router';

import { Alert, AlertDescription } from '#/components/ui/alert';
import { Breadcrumb, BreadcrumbList, BreadcrumbItem, BreadcrumbLink } from '#/components/ui/breadcrumb';
import { Button } from '#/components/ui/button';
import { SidebarTrigger } from '#/components/ui/sidebar';
import { Separator } from '#/components/ui/separator';
import { useKeyRate } from '#/entities/analitics/hook';
import { WithTooltip } from '#/components/with-tooltip';

import { buttons } from '../buttons';

export function Header() {
  const { pathname } = useLocation();
  const isUserLogedIn = pathname !== '/app/watch';

  const { data: keyRate, isLoading: keyRateLoading } = useKeyRate();
  const union = [...buttons.base, ...buttons.analitics];
  const currentBreadLink = union.find((button) => button.url === pathname);

  return (  
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
        !keyRateLoading && currentBreadLink && currentBreadLink.url === '/app/chosen'
          && <div className='ml-auto'>
            <WithTooltip text={`${keyRate}% ─ Ключевая ставка ЦБ РФ`}>
              <Link to={'/app/key-rate'}>
                <Alert className='p-1 lg:py-2 lg:px-3'>
                  <AlertDescription className='truncate!'>{keyRate}% ─ Ключевая ставка ЦБ РФ</AlertDescription>
                </Alert>
              </Link>
            </WithTooltip>
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
  );
}