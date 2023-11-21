import { createContext, useContext, useState, useEffect } from "react";
import { RenderHeader } from "../components/structure/Header";
import {
  RenderMenu,
  RenderRoutes,
} from "../components/structure/RenderNavigation";
import { useNavigate } from "react-router-dom";
import { ApiEndpoints } from '../const/apiEndpoints';

const AuthContext = createContext();
export const AuthData = () => useContext(AuthContext);

export const AuthWrapper = () => {
  const navigate = useNavigate();

  const [user, setUser] = useState({ name: "", isAuthenticated: false });

  const checkAuthenticationStatus = async () => {
    try {
      const response = await fetch(ApiEndpoints.AUTHSTATUS, {
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
      const response = await fetch(ApiEndpoints.LOGIN, {
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
        navigate("/account")
      } else {
        // Handle login failure here, e.g., show an error message
        const errorData = await response.json();
        throw errorData;
      }
    } catch (error) {
      // Handle network error, e.g., display a message to the user
      console.error("Login error:", error);
      throw error;
    }
  };

  const signup = async (username, email, password) => {
    try {
      const response = await fetch(ApiEndpoints.SIGNUP, {
        method: "POST",
        credentials: "include",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ username, email, password }),
      });

      if (response.ok) {
        console.log("Signup successful.");
        navigate("/login")
      } else {
        const errorData = await response.json();
        throw errorData;
      }
    } catch (error) {
      // Handle network error, e.g., display a message to the user
      console.error("Network error:", error);
      throw error;
    }
  };

  const logout = async () => {
    await fetch(ApiEndpoints.LOGOUT, {
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
    <AuthContext.Provider value={{ user, signup, login, logout }}>
      <>
        <div className="flex w-1/2">
          <RenderHeader />
          <RenderMenu />
        </div>
        <RenderRoutes />
      </>
    </AuthContext.Provider>
  );
};
