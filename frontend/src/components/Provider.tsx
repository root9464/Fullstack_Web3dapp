import { QueryClient, QueryClientProvider } from '@tanstack/react-query';
import { ReactQueryDevtools } from '@tanstack/react-query-devtools';
import { createRouter, RouterProvider } from '@tanstack/react-router';
import { TonConnectUIProvider } from '@tonconnect/ui-react';
import { routeTree } from '../routeTree.gen';

declare module '@tanstack/react-router' {
  interface Register {
    router: typeof router;
  }
}

const queryClient = new QueryClient();
const router = createRouter({ routeTree });
export const Provider = () => {
  return (
    <QueryClientProvider client={queryClient}>
      <TonConnectUIProvider manifestUrl='https://taiga-labs.github.io/dexlot.json'>
        <RouterProvider router={router} />
        <ReactQueryDevtools initialIsOpen={false} />
      </TonConnectUIProvider>
    </QueryClientProvider>
  );
};
