import { createFileRoute } from '@tanstack/react-router';

import { GlobalSkeleton } from '#/components/global-skeleton';
import { KeyRateGraph } from '#/pages/key-rate';

export const Route = createFileRoute('/app/key-rate')({
  component       : RouteComponent,
  pendingComponent: () => <GlobalSkeleton/>,
});

function RouteComponent() {
  return (
    <div className='h-full w-full flex flex-col'>
      <KeyRateGraph/>
    </div>
  );
}
