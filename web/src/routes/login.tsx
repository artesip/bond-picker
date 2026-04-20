import { createFileRoute } from '@tanstack/react-router';

import { LoginForm } from '#/components/login-form';
import { useLoginForm } from '#/entities/auth/schema';

export const Route = createFileRoute('/login')({
  component: RouteComponent,
});

function RouteComponent() {
  const { rhf, onSubmit } = useLoginForm();

  return (
    <div className='flex h-screen w-full items-center justify-center'>
      <LoginForm 
        className='w-100'
        rhf={ rhf }
        onSubmit={ (e) => {
          e.preventDefault();
          onSubmit(e);
        } }
      />
    </div>
  );
}
