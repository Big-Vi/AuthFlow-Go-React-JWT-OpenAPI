import { redirect } from "react-router-dom"

async function getAuthStatus() {
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

export async function requireAuth(request) {
    const pathname = new URL(request.url).pathname
    const isAuthenticated = await getAuthStatus();

    if (!isAuthenticated) {
        throw redirect(
            `/login?message=You must log in first.&redirectTo=${pathname}`
        )
    }
}
