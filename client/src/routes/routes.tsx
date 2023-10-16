import React from 'react';

const Home = React.lazy(() => import('../pages'));

type Route = {
  path: string;
  auth: boolean;
  Component: React.FC;
};

const routes: Route[] = [
  {
    path: '/',
    auth: false,
    Component: Home,
  },
];

export default routes;
