import { Controller  } from 'react-hook-form';

import { cn } from '#/lib/utils';
import { Button } from '#/components/ui/button';
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from '#/components/ui/card';
import {
  Field,
  FieldDescription,
  FieldError,
  FieldGroup,
  FieldLabel,
} from '#/components/ui/field';
import { Input } from '#/components/ui/input';

import { Spinner } from './ui/spinner';

import type { UseFormReturn } from 'react-hook-form';
import type { RegistrationInput } from '#/entities/auth/schema';

type RegistrationFormProps = React.ComponentProps<'div'> & {
 rhf: UseFormReturn<RegistrationInput>
 isSubmiting: boolean,
}

export function RegistrationForm({
  className,
  rhf,
  isSubmiting,
  ...props
}: RegistrationFormProps) {
  return (
    <div className={ cn('flex flex-col gap-6', className) } { ...props }>
      <Card>
        <CardHeader>
          <CardTitle>Регистрация в системе</CardTitle>
          <CardDescription>
            Введите данные ниже для регистрации
          </CardDescription>
        </CardHeader>
        <CardContent>
          <form>
            <FieldGroup>
              <Controller
                name='username'
                control={ rhf.control }
                render={ ({ field, fieldState }) => (
                  <Field data-invalid={ fieldState.invalid }>
                    <FieldLabel htmlFor={ field.name }>Логин</FieldLabel>
                    <Input
                      { ...field }
                      id={ field.name }
                      aria-invalid={ fieldState.invalid }
                      placeholder='m@example.com'
                    />
                    {fieldState.invalid && <FieldError errors={ [fieldState.error] } />}
                  </Field>
                ) }
              />

              <Controller
                name='password'
                control={ rhf.control }
                render={ ({ field, fieldState }) => (
                  <Field data-invalid={ fieldState.invalid }>
                    <FieldLabel htmlFor='password'>Пароль</FieldLabel>

                    <Input
                      { ...field }
                      id={ field.name }
                      aria-invalid={ fieldState.invalid }
                      type='password'
                    />
                    {fieldState.invalid && <FieldError errors={ [fieldState.error] } />}
                  </Field>
                ) }
              />

              <Field>
                <Button type='submit' disabled={ isSubmiting }>
                  {!isSubmiting && 'Регистрация'}
                  {isSubmiting && <Spinner className='size-4'/>}
                </Button>
                <FieldDescription className='text-center'>
                  Уже есть аккаунт? <a href='login'>Вход</a>
                </FieldDescription>
              </Field>
            </FieldGroup>
          </form>
        </CardContent>
      </Card>
    </div>
  );
}
