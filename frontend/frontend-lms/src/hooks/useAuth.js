import { useState, useEffect } from 'react';
import authService from '../services/authService.js';

export const useAuth = () => {
  const [user, setUser] = useState(null);
  const [isAuthenticated, setIsAuthenticated] = useState(false);
  const [isLoading, setIsLoading] = useState(true);

  useEffect(() => {
    const initAuth = () => {
      // Clean up any invalid localStorage data first
      const userStr = localStorage.getItem('user');
      if (userStr === 'undefined' || userStr === 'null') {
        localStorage.removeItem('user');
        localStorage.removeItem('token'); // Also remove token if user data is invalid
      }
      
      const currentUser = authService.getCurrentUser();
      const authenticated = authService.isAuthenticated();
      
      setUser(currentUser);
      setIsAuthenticated(authenticated);
      setIsLoading(false);
    };

    initAuth();
  }, []);

  const login = async (email, password) => {
    const response = await authService.login(email, password);
    setUser(response.user);
    setIsAuthenticated(true);
    return response;
  };

  const logout = () => {
    authService.logout();
    setUser(null);
    setIsAuthenticated(false);
  };

    const isAdmin = () => {
    return authService.isAdmin();
  };

  const isTeacher = () => {
    return authService.isTeacher();
  };

  return {
    user,
    isAuthenticated,
    isLoading,
    login,
    logout,
    isAdmin,
    isTeacher,
  };
};

export default useAuth;
