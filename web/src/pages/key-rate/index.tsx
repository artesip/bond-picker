import { useKeyRates, useRuonia } from '#/entities/analitics/hook';
import { Card } from '#/components/ui/card';
import { Skeleton } from '#/components/ui/skeleton';

import { KeyRateChart } from './ui/key-rate-chart';

export function KeyRateGraph() {
  const { data: keyRates, isLoading: keyRatesLoading } = useKeyRates();
  const { data: ruonia, isLoading: ruoniaLoading } = useRuonia();

  const isLoading = keyRatesLoading || ruoniaLoading;

  return (
    <Card className='relative flex min-h-150 h-full'>
      <p className='px-(--card-spacing) font-heading text-base font-medium'>Ключевая ставка и RUONIA</p>
      {
        !isLoading && <KeyRateChart
          className='min-h-0 flex-1'
          keyRates={ keyRates }
          ruonia={ ruonia }
        />
      }
      {
        isLoading && <Skeleton className='absolute inset-0 rounded-xl'/>
      }
    </Card>
  );
}
