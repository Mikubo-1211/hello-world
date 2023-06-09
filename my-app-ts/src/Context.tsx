import React, { createContext, useContext, useState, ReactNode } from 'react';

interface ContextProps {
  email: string;
  setEmail: (email: string) => void;
}

const Context = createContext<ContextProps>({
  email: '',
  setEmail: () => {},
});

export const useEmailContext = (): ContextProps => useContext(Context);

export const EmailProvider: React.FC<{ children: ReactNode }> = ({ children }) => {
  const [email, setEmail] = useState('');

  return (
    <Context.Provider value={{ email, setEmail }}>
      {children}
    </Context.Provider>
  );
};
