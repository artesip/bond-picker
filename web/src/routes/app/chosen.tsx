import { createFileRoute } from '@tanstack/react-router';

import { ChosenPage } from '#/pages/chosen';

export const Route = createFileRoute('/app/chosen')({ component: Page });

function Page() {
  return (
    <ChosenPage/>
  );
}
