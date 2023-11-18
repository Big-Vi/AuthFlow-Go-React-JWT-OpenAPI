import { Link } from "react-router-dom";

export const RenderHeader = () => {
     return (
          <div className="header">
               <div className="menuItem">
                    <Link to={'/'}>
                         Auth flow
                    </Link>
               </div>
          </div>
     )
}
