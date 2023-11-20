import { AuthData } from "../../auth/AuthWrapper"

export const Account = () => {

     const { user } = AuthData();

     return (
          <div className="flex flex-col items-center">
               <div className="text-start mt-8">
               <h1 className="text-2xl font-bold text-black mb-4 uppercase font-bold">Your Account</h1>
                    <p className="font-semibold">Username: {user.name}</p>
               </div>
          </div>
     )
}
