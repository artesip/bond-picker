import { Controller  } from 'react-hook-form';

import { Toggle } from '#/components/ui/toggle';

import type { UseFormReturn } from 'react-hook-form';
import type { FilterSchema, FilterInput } from '#/entities/bonds/shemas';
import type z from 'zod';
type FilterTogle = {
    key: keyof z.infer<typeof FilterSchema>
    text: string
}

const toggles: FilterTogle[] = [
  {
    key : 'ratingEnabled',
    text: 'Рейтинг'
  },
  {
    key : 'ytmEnabled',
    text: 'YTM'
  },
  {
    key : 'durationEnabled',
    text: 'Дюрация'
  },
  {
    key : 'currencyEnabled',
    text: 'Валюта'
  },
  {
    key : 'offerEnabled',
    text: 'Оферта'
  },
];

type FilterTogglesProps = {
    rhf: UseFormReturn<FilterInput>
}

export function FilterToggles({ rhf }: FilterTogglesProps) {
  return (
    <div className='flex flex-wrap items-center gap-2 mb-2'>
      {
        toggles.map(el => (
          <Controller
            key={ el.key }
            name={ el.key }
            control={ rhf.control }
            render={ ({ field }) => (
              <Toggle
                variant='outline'
                size='sm'
                pressed={ field.value as boolean }
                onPressedChange={ field.onChange }
              >
                {el.text}
              </Toggle>
            ) }
          />
        ))
      }
    </div>
  );
}