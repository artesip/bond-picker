import { Controller  } from 'react-hook-form';


import  { Field, FieldLabel } from '#/components/ui/field';
import { SelectTrigger, SelectValue, SelectContent, SelectGroup, SelectLabel, SelectItem, Select } from '#/components/ui/select';

import type { UseFormReturn } from 'react-hook-form';
import type { FilterInput } from '#/entities/bonds/shemas';

type CurrencySelectProps = {
    rhf: UseFormReturn<FilterInput>
    bondsLoading: boolean
    currencyValues: string[]
}

export function CurrencySelect({ rhf, bondsLoading, currencyValues }: CurrencySelectProps) {
  return (
    <Field>
      <FieldLabel>Валюта</FieldLabel>
      <Controller
        name='currency'
        control={ rhf.control }
        render={ ({ field }) => (
          <Select
            disabled={ bondsLoading }
            value={ field.value }
            onValueChange={ field.onChange }
          >
            <SelectTrigger className='w-full'>
              <SelectValue placeholder='Выберите валюту'/>
            </SelectTrigger>
            <SelectContent>
              <SelectGroup>
                <SelectLabel>Валюта</SelectLabel>
                {
                  currencyValues.map(currency => (
                    <SelectItem value={ currency } key={ currency }>{currency}</SelectItem>
                  ))
                }
              </SelectGroup>
            </SelectContent>
          </Select>
        ) }
      />
    </Field>
  );
}