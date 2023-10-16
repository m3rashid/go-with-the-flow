import { Routes, Route } from 'react-router-dom';
import NotFound from '../pages/404';
import routes from './routes';

const AppRoutes = () => {
  return (
    <Routes>
      {routes.map(({ path, Component }) => (
        <Route {...{ key: path, path, Component }} />
      ))}

      <Route path='*' element={<NotFound />} />
    </Routes>
  );
};

export default AppRoutes;
