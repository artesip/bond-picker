import { MoveLeft, Home, AlertCircle } from 'lucide-react';
import { Link } from '@tanstack/react-router';

import { Button } from '@/components/ui/button';

export default function NotFound() {
  return (
    <div className='relative flex min-h-screen flex-col items-center justify-center overflow-hidden bg-background px-6 py-24 sm:py-32 lg:px-8'>
      {/* Фоновый градиент для мягкого свечения */}
      <div className='absolute inset-0 -z-10 bg-[radial-gradient(45rem_50rem_at_top,theme(colors.primary.100),transparent)] opacity-20' />
      
      <div className='text-center'>
        {/* Иконка или бейдж */}
        <div className='flex justify-center'>
          <div className='rounded-full bg-primary/10 p-4 ring-1 ring-primary/20'>
            <AlertCircle className='h-12 w-12 text-primary' />
          </div>
        </div>

        <p className='mt-6 text-base font-semibold leading-8 text-primary'>
          Ошибка 404
        </p>
        
        <h1 className='mt-4 text-3xl font-bold tracking-tight text-foreground sm:text-5xl'>
          Страница не найдена
        </h1>
        
        <p className='mt-6 text-base leading-7 text-muted-foreground max-w-md mx-auto'>
          Извините, мы не смогли найти страницу, которую вы ищете. Возможно, она была перемещена или удалена.
        </p>

        <div className='mt-10 flex items-center justify-center gap-x-4'>
          {/* Кнопка "Назад" с использованием Shadcn Button */}
          <Button 
            variant='outline' 
            onClick={ () => window.history.back() }
            className='gap-2'
          >
            <MoveLeft className='h-4 w-4' />
            
            Назад
          </Button>

          {/* Кнопка "На главную" */}
          <Button asChild className='gap-2'>
            <Link to='/login'>
              <Home className='h-4 w-4' />
              На главную
            </Link>
          </Button>
        </div>
      </div>
    </div>
  );
}