import { NavLink, useNavigate } from "react-router-dom"

export default function Header() {
    const activeStyles = {
        fontWeight: "bold",
        textDecoration: "underline",
        color: "#161616"
    }
    const navigate = useNavigate();
    
    async function logoutUser() {
        await fetch('http://localhost:8000/api/user/logout', {
          method: 'POST',
          credentials: 'include',
          headers: {
            'Content-Type': 'application/json'
          },
        })
        
  navigate('/login');
      }
    
    return (
        <header>
            <nav>
                <NavLink 
                    to="/"
                    style={({isActive}) => isActive ? activeStyles : null}
                >
                    Home
                </NavLink>
                <NavLink 
                    to="dashboard"
                    style={({isActive}) => isActive ? activeStyles : null}
                >
                    Dashboard
                </NavLink>
                <button onClick={logoutUser}>Logout</button>
            </nav>
        </header>
    )
}
