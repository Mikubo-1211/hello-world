import React, { useState } from 'react';
import { BrowserRouter as Router, Routes, Route, Navigate, Outlet } from 'react-router-dom';
import LoginForm from './LoginForm';
import SignUpForm from './SignUpForm';
import HomePage from './HomePage';

const App: React.FC = () => {
  const [isLoggedIn, setIsLoggedIn] = useState(false);

  const handleLogin = () => {
    setIsLoggedIn(true);
  };

  const handleLogout = () => {
    setIsLoggedIn(false);
  };

  return (
    <Router>
      <Routes>
        <Route path="/">
          <Route path="login" element={isLoggedIn ? <Navigate to="/home" /> : <LoginForm handleLogin={handleLogin} />} />
          <Route path="signup" element={isLoggedIn ? <Navigate to="/home" /> : <SignUpForm handleLogin={handleLogin} />} />
          <Route path="home" element={isLoggedIn ? <HomePage handleLogout={handleLogout} /> : <Navigate to="/login" />} />
          <Route path="/" element={<Navigate to="/login" />} />
        </Route>
      </Routes>
    </Router>
  );
};

export default App;

