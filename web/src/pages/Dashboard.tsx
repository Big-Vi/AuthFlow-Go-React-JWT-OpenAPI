// import { useLoaderData } from "react-router-dom"
import { requireAuth } from "../utils"

export async function loader({ request }) {
    await requireAuth(request)
    return "value"
}

export default function Dashboard() {
    // const loaderData = useLoaderData()

    return (
        <>
            <section className="dashboard">
              Dashboard                
            </section>
        </>
    )
}