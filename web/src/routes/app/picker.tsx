import { createFileRoute } from '@tanstack/react-router';
import z from 'zod';

import { PickerPage } from '#/pages/picker';
import { GlobalSkeleton } from '#/components/global-skeleton';

export const Route = createFileRoute('/app/picker')({
  component     : RouteComponent,
  validateSearch: z.object({
    id: z.string().optional(),
  }),
  pendingComponent: () => <GlobalSkeleton/>,
});

function RouteComponent() {
  return (
    <div className='h-full w-full flex flex-col'>
      <PickerPage />
    </div>
  );
}
