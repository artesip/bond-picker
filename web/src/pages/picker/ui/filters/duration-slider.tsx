import  { Controller } from 'react-hook-form';


import  { Field, FieldLabel } from '#/components/ui/field';
import  { InputGroup, InputGroupInput } from '#/components/ui/input-group';
import { Slider } from '#/components/ui/slider';

import type { UseFormReturn } from 'react-hook-form';
import type { FilterInput } from '#/entities/bonds/shemas';

type DurationSliderProps = {
    rhf: UseFormReturn<FilterInput>
    max: number
}

export function DurationSlider({ rhf, max }: DurationSliderProps) {
  return (
    <Field>
      <FieldLabel>Доходность</FieldLabel>
            
      <Controller
        name='durationFrom'
        control={ rhf.control }
        render={ ({ field: fromField }) => (
          <Controller
            name='durationTo'
            control={ rhf.control }
            render={ ({ field: toField }) => (
              <div className='flex gap-4'>
          
                <InputGroup className='w-25'>
                  <InputGroupInput
                    value={ fromField.value }
                    onChange={ (e) => {
                      const value = Number(e.target.value);
                      fromField.onChange(value);
                    } }
                  />
                </InputGroup>

                <Slider
                  value={ [fromField.value, toField.value] }
                  max={ max }
                  min={ 0 }
                  step={ 0.1 }
                  onValueChange={ ([from, to]) => {
                    fromField.onChange(from);
                    toField.onChange(to);
                  } }
                  className='mx-auto w-full'
                />

                <InputGroup className='w-25'>
                  <InputGroupInput
                    value={ toField.value }
                    onChange={ (e) => {
                      const value = Number(e.target.value);
                      toField.onChange(value);
                    } }
                  />
                </InputGroup>

              </div>
            ) }
          />
        ) }
      />
    </Field>
  );
}