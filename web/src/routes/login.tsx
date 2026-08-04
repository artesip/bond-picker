import { createFileRoute } from '@tanstack/react-router';

import { LoginForm } from '#/pages/login';
import { useLoginForm } from '#/entities/auth/schema';

export const Route = createFileRoute('/login')({
  component: RouteComponent,
});

function RouteComponent() {
  const { rhf, onSubmit, isSubmiting } = useLoginForm();

  return (
    <div className='flex h-screen w-full items-center justify-center'>
      <LoginForm 
        className='w-100'
        rhf={ rhf }
        isSubmiting={ isSubmiting }
        onSubmit={ (e) => {
          e.preventDefault();
          onSubmit(e);
        } }
      />
    </div>
  );
}
