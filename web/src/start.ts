import { redirect } from '@tanstack/react-router';
import { createMiddleware, createStart } from '@tanstack/react-start';

export const authMiddleware = createMiddleware()
  .server(async ({ next, request }) => {
    const rawCookie = request.headers.get('cookie') || '';

    const hasToken = rawCookie.split(';').some(c => c.trim().startsWith('bond-picker-auth='));
    const url = new URL(request.url);

    if (url.pathname === '/') {
      throw redirect({
        to: hasToken ? '/app' : '/login',
      });
    }

    if (hasToken && (url.pathname === '/login' || url.pathname === '/registration')) {
      throw redirect({ to: '/app/chosen' });
    } else if (!(url.pathname === '/login' || url.pathname === '/registration') && !hasToken && url.pathname !== '/app/watch') {
      throw redirect({ to: '/login' });
    }

    return next();
  });

export const startInstance = createStart(() => {
  return {
    requestMiddleware: [authMiddleware],
  };
});

export default startInstance;