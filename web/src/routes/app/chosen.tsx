import { createFileRoute } from '@tanstack/react-router';
import z from 'zod';

import { ChosenPage } from '#/pages/chosen';
import { GlobalSkeleton } from '#/components/global-skeleton';

export const Route = createFileRoute('/app/chosen')({
  component       : Page,
  pendingComponent: () => <GlobalSkeleton/>,
  validateSearch  : z.object({
    id: z.string().optional(),
  }),
});

function Page() {
  return (
    <ChosenPage/>
  );
}
