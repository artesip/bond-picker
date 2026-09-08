import { Link } from '@tanstack/react-router';
import { ArrowLeft } from 'lucide-react';

import { Button } from './ui/button';
import { Empty, EmptyContent, EmptyDescription, EmptyHeader, EmptyTitle } from './ui/empty';


export default function NotFound() {
 return (
  <div className='h-screen flex justify-center items-center'>
    <Empty>
      <EmptyHeader>
        <EmptyTitle className='text-xl'>404 - Ничего не нашлось</EmptyTitle>
        <EmptyDescription className='text-[16px]'>
          Похоже, такой страницы не существует.
        </EmptyDescription>
      </EmptyHeader>
      <EmptyContent>
        <Button variant={'outline'} className='w-60' asChild>
          <Link to='/app/chosen' preload={'render'}>
            <ArrowLeft className='h-4'/>
            <span>На главную</span>
          </Link>
        </Button>
      </EmptyContent>
    </Empty>
    </div>
  );
}