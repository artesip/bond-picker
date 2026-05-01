import { useRef, useState } from 'react';
import { Plus, Minus } from 'lucide-react';
import { Controller } from 'react-hook-form';

import { useFavoriteAddForm } from '#/entities/bonds/shemas';

import { Button } from './ui/button';
import { Input } from './ui/input';
import { Field, FieldError } from './ui/field';
import { Spinner } from './ui/spinner';

import type { Bond } from '#/entities/bonds/model';

type AddChosenFormProps = {
  bond: Bond
  refetch: () => void
}

export function AddChosenForm({ bond, refetch }: AddChosenFormProps) {
  const [isSubmiting, setIsSubmiting] = useState(false);
  const [isLoading, setIsLoading] = useState(false);
  const { rhf, onSubmit } = useFavoriteAddForm(bond.id);
  const timeoutRef = useRef<NodeJS.Timeout | null>(null);

  function onFavoriteClick() {
    if (isSubmiting) return;

    setIsSubmiting(true);

    startCloserTime();
  }

  function startCloserTime() {
    timeoutRef.current = setTimeout(() => {      
      setIsSubmiting(false);
      timeoutRef.current = null;
    }, 5000);
  }

  const number = rhf.watch('number');

  return (
    <form onSubmit={ async (e) => {
      e.preventDefault();
      setIsLoading(true);

      try {
        await onSubmit(e);
        refetch();
      } finally {
        setIsLoading(false);
        setIsSubmiting(false);
      }
    } }
    >
      {!isSubmiting && <Button onClick={ onFavoriteClick } className='w-full' variant={ 'secondary' }>В избранное</Button>}

      {
        isSubmiting
        && <div className='flex gap-2'>
          <Button
            variant={ 'ghost' }
            type='button'
            onClick={ () => {
              rhf.setValue('number', number + 1 );

              if (timeoutRef.current) {
                clearTimeout(timeoutRef.current);
                startCloserTime();
              }
            } }>
            <Plus />
          </Button>
          
          <Controller
            name='number'
            control={ rhf.control }
            render={ ({ field, fieldState }) => (
              <Field data-invalid={ fieldState.invalid }>
                <Input
                  { ...field }
                  id={ field.name }
                  aria-invalid={ fieldState.invalid }
                  onChange={ (e) => {
                    if (timeoutRef.current) {
                      clearTimeout(timeoutRef.current);
                      startCloserTime();
                    }

                    field.onChange(e);
                  } }
                  placeholder='Кол-во'
                />
                {fieldState.invalid && <FieldError errors={ [fieldState.error] } />}
              </Field>
            ) }
          />

          
          <Button
            variant={ 'ghost' }
            type='button'
            onClick={ () => {
              rhf.setValue('number', number - 1 );
              if (timeoutRef.current) {
                clearTimeout(timeoutRef.current);
                startCloserTime();
              }
            } }>
            <Minus />
          </Button>
          
          <Button
            variant={ 'secondary' }
            className='w-1/2'
            type='submit'
            disabled={ isLoading }
          >
            {isLoading ? <Spinner /> : 'Добавить'}
          </Button>
        </div>
      }
    </form>
  );
}