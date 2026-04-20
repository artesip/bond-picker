import { createFileRoute } from '@tanstack/react-router';

export const Route = createFileRoute('/app/chosen')({ component: App });

function App() {
  return (
    <div>123</div>
  );
}
