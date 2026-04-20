import { Controller } from 'react-hook-form';

import { Field, FieldLabel } from '#/components/ui/field';
import { Toggle } from '#/components/ui/toggle';

import type { UseFormReturn } from 'react-hook-form';
import type { FilterInput } from '#/entities/bonds/shemas';

const offerts = [
  {
    key : 'no',
    text: 'Нет'
  },
  {
    key : 'all',
    text: 'Все'
  },
  {
    key : 'put',
    text: 'Put'
  },
  {
    key : 'call',
    text: 'Call'
  },
];


type OfferToggleProps = {
    rhf: UseFormReturn<FilterInput>
}

export function OfferToggle({ rhf }: OfferToggleProps) {
  return (
    <Field>
      <FieldLabel>Офферта</FieldLabel>
      <div className='flex flex-wrap items-center gap-2'>
        <Controller
          name='offer'
          control={ rhf.control }
          render={ ({ field }) => (
            <div className='flex flex-wrap items-center gap-2'>
              {offerts.map(el => (
                <Toggle
                  key={ el.key }
                  variant='outline'
                  size='sm'
                  pressed={ field.value === el.key }
                  onPressedChange={ (pressed) => {
                    if (pressed) {
                      field.onChange(el.key);
                    } else {
                      field.onChange('');
                    }
                  } }
                >
                  {el.text}
                </Toggle>
              ))}
            </div>
          ) }
        />
      </div>
    </Field>
  );
}