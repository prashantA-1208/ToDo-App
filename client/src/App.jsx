import { BrowserRouter, Routes, Route, Navigate } from 'react-router-dom';
import LoginPage from './pages/LoginPage';
import SignupPage from './pages/SignupPage';
import Dashboard from './pages/Dashboard';
import { isLoggedIn } from './utils/auth';
import { useState, useEffect } from 'react';

export default function App() {

  const [authenticated, setAuthenticated] = useState(isLoggedIn());

  useEffect(() => {
    const handleStorageChange = () => {
      setAuthenticated(isLoggedIn());
    };

    window.addEventListener('storage', handleStorageChange);
    return () => window.removeEventListener('storage', handleStorageChange);
  }, []);

  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<Navigate to={authenticated ? '/dashboard' : '/login'} />} />
        <Route path="/login" element={<LoginPage setAuthenticated={setAuthenticated} />} />
        <Route path="/signup" element={<SignupPage />} />
        <Route path="/dashboard" element={authenticated ? <Dashboard /> : <Navigate to="/login" />} />
      </Routes>
    </BrowserRouter>
  );
}