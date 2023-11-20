import { useReducer, useState } from "react";
import { useNavigate } from "react-router-dom";
import { AuthData } from "../../auth/AuthWrapper"

export const Login = () => {

     const navigate = useNavigate();
     const { login } = AuthData();
     const [ formData, setFormData ] = useReducer((formData, newItem) => { return ( {...formData, ...newItem} )}, {email: "", password: ""})
     const [ errorMessage, setErrorMessage ] = useState(null)
     
     const doLogin = async () => {
          try {
               await login(formData.email, formData.password)
          } catch (error) {
               setErrorMessage(error)
          }
     }

     return (
          <div className="flex flex-col items-center">
               <div className="w-96 mt-8 flex justify-between">
                    <h2 className="text-2xl font-bold text-black">Login</h2>
                    <div className="flex"><p>Don't have an account?</p><a className="pl-2 inline-block text-blue" href="/signup">Sign Up</a></div>
               </div>
               {errorMessage ?
                    <p className="text-red w-96 mt-8">{errorMessage}</p>
               : null }
               <div>
                    <label className="mb-2 mt-6 block text-black">
                         Email
                    </label>
                    <input
                    type="email"
                    placeholder="Enter email"
                    value={formData.email} 
                    onChange={(e) => setFormData({email: e.target.value}) }
                    className="w-96 rounded-lg border-[1.5px] border-stroke bg-transparent py-3 px-5 font-medium outline-none transition focus:border-primary active:border-primary disabled:cursor-default disabled:bg-whiter dark:border-form-strokedark dark:bg-form-input dark:focus:border-primary"
                    />
               </div>
               <div>
                    <label className="mb-2 mt-6 block text-black">
                         Password
                    </label>
                    <input
                    type="password"
                    placeholder="Password"
                    value={formData.password}
                    onChange={(e) => setFormData({password: e.target.value}) }
                    className="w-96 rounded-lg border-[1.5px] border-stroke bg-transparent py-3 px-5 font-medium outline-none transition focus:border-primary active:border-primary disabled:cursor-default disabled:bg-whiter dark:border-form-strokedark dark:bg-form-input dark:focus:border-primary"
                    />
               </div>
               
               <div className="button">
                    <button onClick={doLogin} className="flex w-96 justify-center rounded bg-primary p-3 mt-8 font-medium text-gray">
                         Log in
                    </button>
               </div>
          </div>
     )
}