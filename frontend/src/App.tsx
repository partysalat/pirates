import React from 'react';
import './App.css';
import { Accounting } from './components/accounting';
import { QueryClient, QueryClientProvider } from 'react-query';
import { ThemeProvider } from '@mui/material';
import { accountingTheme, mainTheme } from './Theme';
// import { Game } from './game';
import { MainLayout } from './layouts/MainLayout';
import { GameLayout } from './layouts/GameLayout';
import { createBrowserRouter, RouterProvider } from 'react-router-dom';

const queryClient = new QueryClient({
  defaultOptions: {
    queries: {
      staleTime: Infinity,
    },
  },
});
const router = createBrowserRouter([
  {
    path: '/',
    element: (
      <ThemeProvider theme={mainTheme}>
        <MainLayout />
      </ThemeProvider>
    ),
  },
  {
    path: '/game',
    element: (
      <ThemeProvider theme={mainTheme}>
        <GameLayout />
      </ThemeProvider>
    ),
  },
  {
    path: '/abrechnung',
    element: (
      <ThemeProvider theme={accountingTheme}>
        <Accounting />
      </ThemeProvider>
    ),
  },
]);

function App() {
  return (
    <div className="App">
      <QueryClientProvider client={queryClient}>
        <RouterProvider router={router} />
      </QueryClientProvider>
    </div>
  );
}

export default App;
