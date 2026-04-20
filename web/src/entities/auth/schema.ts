import * as z from 'zod';
import { useForm } from 'react-hook-form';
import { zodResolver } from '@hookform/resolvers/zod';
import { useNavigate } from '@tanstack/react-router';
import { toast } from 'sonner';

import { Login, Registration } from '#/entities/auth/api';

const LoginSchema = z.object({
  username: z.string().min(3, 'Минимальная длина 3 символа'),
  password: z.string().min(3, 'Минимальная длина 3 символов')
});

export type LoginInput = z.infer<typeof LoginSchema>;

export function useLoginForm() {
  const rhf = useForm<LoginInput>({
    resolver     : zodResolver(LoginSchema),
    defaultValues: {
      username: '',
      password: '',
    }
  });

  const router = useNavigate();

  const onSubmit = rhf.handleSubmit(async (data) => {
    try {
      await Login(data);

      router({ to: '/app/chosen' });
    } catch (e) {
      if (e instanceof Error) {
        toast.error(e.message);
      }
    }
  });

  return {
    rhf,
    onSubmit,
  }; 
}




const RegistrationSchema = z.object({
  username: z.string().min(3, 'Минимальная длина 3 символа'),
  password: z.string().min(3, 'Минимальная длина 3 символов')
});

export type RegistrationInput = z.infer<typeof RegistrationSchema>;

export function useRegistrationForm() {
  const rhf = useForm<RegistrationInput>({
    resolver     : zodResolver(RegistrationSchema),
    defaultValues: {
      username: '',
      password: '',
    }
  });

  const router = useNavigate();

  const onSubmit = rhf.handleSubmit(async (data) => {
    try {
      await Registration(data);

      router({ to: '/app/chosen' });
    } catch (e) {
      if (e instanceof Error) {
        toast.error(e.message);
      }
    }
  });

  return {
    rhf,
    onSubmit,
  }; 
}