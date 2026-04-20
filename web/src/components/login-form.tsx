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

import type { UseFormReturn } from 'react-hook-form';
import type { LoginInput } from '#/entities/auth/schema';

type LoginFormProps = React.ComponentProps<'div'> & {
  rhf: UseFormReturn<LoginInput>
}

export function LoginForm({
  className,
  rhf,
  ...props
}: LoginFormProps) {
  
  return (
    <div className={ cn('flex flex-col gap-6', className) } { ...props }>
      <Card>
        <CardHeader>
          <CardTitle>Вход в аккаунт</CardTitle>
          <CardDescription>
            Введите данные ниже для входа в аккаунт
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
                <Button type='submit'>Вход</Button>
                <FieldDescription className='text-center'>
                  Нет аккаунта? <a href='registration'>Зарегестрироваться</a>
                </FieldDescription>
              </Field>
            </FieldGroup>
          </form>
        </CardContent>
      </Card>
    </div>
  );
}
