import  { Controller  } from 'react-hook-form';


import  { Field, FieldLabel } from '#/components/ui/field';
import  { SelectTrigger, SelectValue, SelectContent, SelectGroup, SelectLabel, SelectItem, Select } from '#/components/ui/select';
import { Separator } from '#/components/ui/separator';

import type { UseFormReturn } from 'react-hook-form';
import type { FilterInput } from '#/entities/bonds/shemas';

type RatingSelectProps = {
    rhf: UseFormReturn<FilterInput>
    ratingValues: string[]
    ratingLoading: boolean
}

export function RatingSelect({ rhf, ratingLoading, ratingValues }: RatingSelectProps) {
  return (
    <Field>
      <FieldLabel>Рейтинг</FieldLabel>
      <div className='flex items-center gap-4 w-full'>
        <Controller
          name='ratingFrom'
          control={ rhf.control }
          render={ ({ field }) => (
            <Select
              disabled={ ratingLoading }
              value={ field.value }
              onValueChange={ field.onChange }
            >
              <SelectTrigger className='w-full max-w-45'>
                <SelectValue placeholder='От' />
              </SelectTrigger>
              <SelectContent>
                <SelectGroup>
                  <SelectLabel>Рейтинг</SelectLabel>
                  {
                    ratingValues.map(rating => (
                      <SelectItem value={ rating } key={ rating }>
                        {rating}
                      </SelectItem>
                    ))
                  }
                </SelectGroup>
              </SelectContent>
            </Select>
          ) }
        />
        
        <Separator className='w-10!'/>

        <Controller
          name='ratingTo'
          control={ rhf.control }
          render={ ({ field }) => (
            <Select
              disabled={ ratingLoading }
              value={ field.value }
              onValueChange={ field.onChange }
            >
              <SelectTrigger className='w-full max-w-45'>
                <SelectValue placeholder='От' />
              </SelectTrigger>
              <SelectContent>
                <SelectGroup>
                  <SelectLabel>Рейтинг</SelectLabel>
                  {
                    ratingValues.map(rating => (
                      <SelectItem value={ rating } key={ rating }>
                        {rating}
                      </SelectItem>
                    ))
                  }
                </SelectGroup>
              </SelectContent>
            </Select>
          ) }
        />

      </div>  
    </Field>
  );
}