import * as z from 'zod';
import { useForm } from 'react-hook-form';
import { zodResolver } from '@hookform/resolvers/zod';
import { useNavigate } from '@tanstack/react-router';
import { toast } from 'sonner';
import { useState } from 'react';

import { Login, Registration } from '#/entities/auth/api';

const LoginSchema = z.object({
  username: z.string().trim().min(4, 'Минимальная длина 4 символа'),
  password: z.string().trim().min(8, 'Минимальная длина 8 символов')
});

export type LoginInput = z.infer<typeof LoginSchema>;

export function useLoginForm() {
  const [isSubmiting, setIsSubmiting] = useState<boolean>(false);

  const rhf = useForm<LoginInput>({
    resolver     : zodResolver(LoginSchema),
    defaultValues: {
      username: '',
      password: '',
    }
  });

  const router = useNavigate();

  const onSubmit = rhf.handleSubmit(async (data) => {
    if (isSubmiting) {
      return;
    }

    setIsSubmiting(true);

    try {
      await Login(data);
      router({ to: '/app/chosen' });
    } catch (e) {
      if (e instanceof Error) {
        toast.error(e.message);
      }
    } finally {
      setIsSubmiting(false);
    }
  });

  return {
    rhf,
    onSubmit,
    isSubmiting,
  }; 
}




const RegistrationSchema = z.object({
  username: z.string().trim().min(4, 'Минимальная длина 4 символа'),
  password: z.string().trim().min(8, 'Минимальная длина 8 символов')
});

export type RegistrationInput = z.infer<typeof RegistrationSchema>;

export function useRegistrationForm() {
  const [isSubmiting, setIsSubmiting] = useState<boolean>(false);
  const rhf = useForm<RegistrationInput>({
    resolver     : zodResolver(RegistrationSchema),
    defaultValues: {
      username: '',
      password: '',
    }
  });

  const router = useNavigate();

  const onSubmit = rhf.handleSubmit(async (data) => {
    if (isSubmiting) {
      return;
    }

    setIsSubmiting(true);

    try {
      await Registration(data);

      router({ to: '/app/chosen' });
    } catch (e) {
      if (e instanceof Error) {
        toast.error(e.message);
      }
    } finally {
      setIsSubmiting(false);
    }
  });

  return {
    rhf,
    onSubmit,
    isSubmiting,
  }; 
}