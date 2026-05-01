import { createFileRoute } from '@tanstack/react-router';

import { RegistrationForm } from '#/components/registration-form';
import { useRegistrationForm } from '#/entities/auth/schema';

export const Route = createFileRoute('/registration')({
  component: RouteComponent,
});

function RouteComponent() {
  const { rhf, onSubmit, isSubmiting } = useRegistrationForm();

  return (
    <div className='flex h-screen w-full items-center justify-center'>
      <RegistrationForm 
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