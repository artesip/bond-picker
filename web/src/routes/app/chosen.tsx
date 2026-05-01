import { createFileRoute } from '@tanstack/react-router';

import { ChosenPage } from '#/pages/chosen';
import { GlobalSkeleton } from '#/components/global-skeleton';

export const Route = createFileRoute('/app/chosen')({
  component       : Page,
  pendingComponent: () => <GlobalSkeleton/>,
});

function Page() {
  return (
    <ChosenPage/>
  );
}
