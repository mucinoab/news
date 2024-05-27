import React from 'react';
import ReactDOM from 'react-dom/client';
import {
  createBrowserRouter,
  RouterProvider,
} from "react-router-dom";

import Newsletters from './components/Newsletters.tsx';
import NewForm from './components/NewForm.tsx';


const router = createBrowserRouter([
  {
    path: "/",
    element: <Newsletters />,
  },
  {
    path: "/new",
    element: <NewForm />,
  }
]);

ReactDOM.createRoot(document.getElementById('root')!).render(
  <React.StrictMode>
    <RouterProvider router={router} />
  </React.StrictMode>,
);
