import { createFileRoute, redirect } from '@tanstack/react-router';

// Only for redirect
export const Route = createFileRoute('/app/')({
  component : RouteComponent,
  beforeLoad: () => {
    throw redirect(
      {
        to: '/app/chosen'
      }
    );
  }
});


function RouteComponent() {
  return <div></div>;
}
