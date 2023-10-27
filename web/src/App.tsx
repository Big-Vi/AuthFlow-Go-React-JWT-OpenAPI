import {
  RouterProvider,
  createBrowserRouter,
  createRoutesFromElements,
  Route,
} from "react-router-dom"
import Dashboard , { loader as dashboardLoader } from "./pages/Dashboard"
import NotFound from "./pages/NotFound"
import Login, { loader as loginLoader, action as loginAction } from "./pages/Login"
import Layout from "./components/Layout"
import Home from "./pages/Home"
// import Error from "./components/Error"
// import { requireAuth } from "./utils"


const router = createBrowserRouter(createRoutesFromElements(
  <Route path="/" element={<Layout />}>
     <Route path="/" element={<Home />} />
    <Route
      path="login"
      element={<Login />}
      loader={loginLoader}
      action={loginAction}
    />

    <Route path="dashboard">
      <Route
        index
        element={<Dashboard />}
        loader={dashboardLoader}
      />
    </Route>
    <Route path="*" element={<NotFound />} />
  </Route>
))

export default function App() {
  return (
    <RouterProvider router={router} />
  );
}
