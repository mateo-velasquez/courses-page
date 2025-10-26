import { RouterProvider } from 'react-router-dom';
import { router } from './router.jsx';
import './styles/variables.css';
import './styles/utilities.css';

function App() {
  return <RouterProvider router={router} />;
}

export default App;
