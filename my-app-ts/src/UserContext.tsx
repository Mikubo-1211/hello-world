import React, { createContext, useState } from 'react';

interface UserContextType {
  userEmail: string;
  setUserEmail: (email: string) => void;
}

export const UserContext = createContext<UserContextType>({
  userEmail: '',
  setUserEmail: () => {},
});

export const UserProvider: React.FC<{ children: React.ReactNode }> = ({ children }) => {
  const [userEmail, setUserEmail] = useState('example@example.com');

  return (
    <UserContext.Provider value={{ userEmail, setUserEmail }}>
      {children}
    </UserContext.Provider>
  );
};


