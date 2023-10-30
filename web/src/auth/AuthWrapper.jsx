import { createContext, useContext, useState, useEffect } from "react";
import { RenderHeader } from "../components/structure/Header";
import {
  RenderMenu,
  RenderRoutes,
} from "../components/structure/RenderNavigation";
import { useNavigate } from "react-router-dom";

const AuthContext = createContext();
export const AuthData = () => useContext(AuthContext);

export const AuthWrapper = () => {
  const navigate = useNavigate();

  const [user, setUser] = useState({ name: "", isAuthenticated: false });

  const checkAuthenticationStatus = async () => {
    try {
      const response = await fetch("http://localhost:8000/api/user/auth-status", {
        method: "GET",
        credentials: "include",
        headers: {
          "Content-Type": "application/json",
        },
      });

      if (response.ok) {
        const userData = await response.json();
        setUser({ name: userData.userName, isAuthenticated: true });
      } else {
        setUser({ name: "", isAuthenticated: false });
      }
    } catch (error) {
      console.error("Network error:", error);
    }
  };

  // Call checkAuthenticationStatus when the component is mounted
  useEffect(() => {
    checkAuthenticationStatus();
  }, []);

  const login = async (email, password) => {
    try {
      const response = await fetch("http://localhost:8000/api/user/login", {
        method: "POST",
        credentials: "include",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ email, password }),
      });

      if (response.ok) {
        const userData = await response.json();
        setUser({ name: userData.userName, isAuthenticated: true });
      } else {
        // Handle login failure here, e.g., show an error message
        console.error("Login failed");
      }
    } catch (error) {
      // Handle network error, e.g., display a message to the user
      console.error("Network error:", error);
    }
  };

  const logout = async () => {
    await fetch("http://localhost:8000/api/user/logout", {
      method: "POST",
      credentials: "include",
      headers: {
        "Content-Type": "application/json",
      },
    });

    setUser({ ...user, isAuthenticated: false });
    navigate("/login");
  };

  return (
    <AuthContext.Provider value={{ user, login, logout }}>
      <>
        <RenderHeader />
        <RenderMenu />
        <RenderRoutes />
      </>
    </AuthContext.Provider>
  );
};
