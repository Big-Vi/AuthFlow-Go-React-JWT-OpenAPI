import { Account } from "../pages/Account"
import { Home } from "../pages/Home"
import { Signup } from "../pages/Signup"
import { Login } from "../pages/Login"

export const nav = [
     { path:     "/",         name: "Home",        element: <Home />,       isMenu: false,     isPrivate: false  },
     { path:     "/signup",   name: "Signup",      element: <Signup />,     isMenu: false,     isPrivate: false  },
     { path:     "/login",    name: "Login",       element: <Login />,      isMenu: false,    isPrivate: false  },
     { path:     "/account",  name: "Account",     element: <Account />,    isMenu: true,     isPrivate: true  },
]