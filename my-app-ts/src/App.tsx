import React, { useState } from 'react';
import { BrowserRouter as Router, Switch, Route, Redirect } from 'react-router-dom';
import LoginForm from './LoginForm';
import SignUpForm from './SignUpForm';
import HomePage from './HomePage';
import { UserProvider } from './UserContext';

const App: React.FC = () => {
  const [isLoggedIn, setIsLoggedIn] = useState(false);

  const handleLogin = () => {
    setIsLoggedIn(true);
  };

  const handleLogout = () => {
    setIsLoggedIn(false);
  };

  return (
    <UserProvider>
      <Router>
      <Switch>
        <Route exact path="/login">
          {isLoggedIn ? <Redirect to="/home" /> : <LoginForm handleLogin={handleLogin} />}
        </Route>
        <Route exact path="/signup">
          {isLoggedIn ? <Redirect to="/home" /> : <SignUpForm handleLogin={handleLogin} />}
        </Route>
        <Route exact path="/home">
          {isLoggedIn ? <HomePage handleLogout={handleLogout} /> : <Redirect to="/login" />}
        </Route>
        <Route path="/">
          <Redirect to="/login" />
        </Route>
      </Switch>
    </Router>
  </UserProvider>
    
  );
};

export default App;
