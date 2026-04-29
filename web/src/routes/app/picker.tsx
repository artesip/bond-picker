import { createFileRoute } from '@tanstack/react-router';
import z from 'zod';

import { PickerPage } from '#/pages/picker';

export const Route = createFileRoute('/app/picker')({
  component     : RouteComponent,
  validateSearch: z.object({
    id: z.string().optional(),
  }),
});

function RouteComponent() {
  return (
    <div className='h-full w-full flex flex-col'>
      <PickerPage />
    </div>
  );
}
