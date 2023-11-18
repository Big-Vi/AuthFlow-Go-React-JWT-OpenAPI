import { useReducer, useState } from "react";
import { useNavigate } from "react-router-dom";
import { AuthData } from "../../auth/AuthWrapper"

export const Signup = () => {

     const navigate = useNavigate();
     const { signup } = AuthData();
     const [ formData, setFormData ] = useReducer((formData, newItem) => { return ( {...formData, ...newItem} )}, {userName: "", email: "", assword: ""})
     const [ errorMessage, setErrorMessage ] = useState(null)
     
     const doLogin = async () => {
          try {
               await signup(formData.userName, formData.email, formData.password)
          } catch (error) {
               setErrorMessage(error)
          }
     }

     return (
          <div className="flex flex-col items-center">
               <div className="w-96 mt-8 flex justify-between">
                    <h2 className="text-2xl font-bold text-black">Sign Up</h2>
                    <div className="flex"><p>Already have an account?</p><a className="pl-2 inline-block text-blue" href="/login">Log In</a></div>
               </div>
               <div>
                    <label className="mb-2 mt-6 block text-black">
                         Username
                    </label>
                    <input
                    type="text"
                    placeholder="Enter username"
                    value={formData.userName} 
                    onChange={(e) => setFormData({userName: e.target.value}) }
                    className="w-96 rounded-lg border-[1.5px] border-stroke bg-transparent py-3 px-5 font-medium outline-none transition focus:border-primary active:border-primary disabled:cursor-default disabled:bg-whiter dark:border-form-strokedark dark:bg-form-input dark:focus:border-primary"
                    />
               </div>
               <div>
                    <label className="mb-2 mt-6 block text-black">
                         Username
                    </label>
                    <input
                    type="email"
                    placeholder="email@gmail.com"
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
                    type="text"
                    placeholder="Password"
                    value={formData.password}
                    onChange={(e) => setFormData({password: e.target.value}) }
                    className="w-96 rounded-lg border-[1.5px] border-stroke bg-transparent py-3 px-5 font-medium outline-none transition focus:border-primary active:border-primary disabled:cursor-default disabled:bg-whiter dark:border-form-strokedark dark:bg-form-input dark:focus:border-primary"
                    />
               </div>
               
               <div className="button">
                    <button onClick={doLogin} className="flex w-96 justify-center rounded bg-primary p-3 mt-8 font-medium text-gray">
                         Sign Up
                    </button>
               </div>
               {errorMessage ?
               <div className="error">{errorMessage}</div>
               : null }
          </div>
     )
}