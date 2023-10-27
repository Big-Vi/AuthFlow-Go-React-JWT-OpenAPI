import React, { createContext, useContext, useEffect, useState, ReactNode } from 'react';

const AuthContext = createContext<AuthContextType | undefined>(undefined);

export type AuthContextType = {
  isAuthenticated: boolean;
  getAuthStatus: () => Promise<void>;
};

type AuthContextProviderProps = {
  children: ReactNode;
};

export const AuthContextProvider: React.FC<AuthContextProviderProps> = ({ children }) => {
  const [isAuthenticated, setIsAuthenticated] = useState(false);

  async function getAuthStatus() {
    // Implement your authentication status fetching logic here
    const status = await fetchAuthStatus(); // Replace with your actual logic
    setIsAuthenticated(status);
  }

  useEffect(() => {
    getAuthStatus();
  }, []);

  return (
    <AuthContext.Provider value={{ isAuthenticated, getAuthStatus }}>
      {children}
    </AuthContext.Provider>
  );
};

export const useAuth = (): AuthContextType => {
  const context = useContext(AuthContext);
  if (context === undefined) {
    throw new Error('useAuth must be used within an AuthContextProvider');
  }
  return context;
};

async function fetchAuthStatus(): Promise<boolean> {
  try {
    const response = await fetch('http://localhost:8000/api/user/auth-status', {
      method: 'GET',
      credentials: 'include',
      headers: {
        'Content-Type': 'application/json',
      },
    });
    const data = await response.json();
    return data.isAuthenticated;
  } catch (error) {
    console.error(error);
    return false;
  }
}
