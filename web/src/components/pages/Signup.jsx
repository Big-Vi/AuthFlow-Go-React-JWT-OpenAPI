import { useState } from "react";
import { Formik, Form, useField } from "formik";
import * as Yup from "yup";
import { AuthData } from "../../auth/AuthWrapper"

export const Signup = () => {

     const { signup } = AuthData();
     const [ errorMessage, setErrorMessage ] = useState(null)

     const TextInput = ({ label, ...props }) => {
          const [field, meta] = useField(props);
          return (
            <>
              <label className="mb-2 mt-6 block text-black" htmlFor={props.id || props.name}>{label}</label>
              <input className="w-96 rounded-lg border-[1.5px] border-stroke bg-transparent py-3 px-5 font-medium outline-none transition focus:border-primary active:border-primary disabled:cursor-default disabled:bg-whiter dark:border-form-strokedark dark:bg-form-input dark:focus:border-primary" {...field} {...props} />
              {meta.touched && meta.error ? (
                <p className="text-red w-96 mt-4">{meta.error}</p>
              ) : null}
            </>
          );
     };

     return (
          <div className="flex flex-col items-center">
               <div className="w-96 mt-8 flex justify-between">
                    <h2 className="text-2xl font-bold text-black">Sign Up</h2>
                    <div className="flex"><p>Already have an account?</p><a className="pl-2 inline-block text-blue" href="/login">Log In</a></div>
               </div>
               {errorMessage ?
                    <p className="text-red w-96 mt-8">{errorMessage}</p>
               : null }
               <Formik
                    initialValues={{
                         userName: "",
                         email: "",
                         password: ""
                    }}
                    validationSchema={Yup.object({
                         userName: Yup.string()
                           .max(15, "Must be 15 characters or less")
                           .required("Required"),
                         email: Yup.string()
                           .email("Invalid email addresss`")
                           .required("Required"),
                         password: Yup.string()
                           .min(6, "Must be 6 characters or more")
                           .required("Required"),
                    })}
                    onSubmit={async (values, { setSubmitting }) => {
                         try {
                              await signup(values.userName, values.email, values.password)
                         } catch (error) {
                              setErrorMessage(error)
                         }
                         setSubmitting(false);
                    }}
               >
                    {({ isSubmitting }) => (
                         <Form>
                              <TextInput
                                   label="Username*"
                                   name="userName"
                                   type="text"
                                   placeholder="Enter username"
                                   autoComplete="off" 
                              />
                              <TextInput
                                   label="Email*"
                                   name="email"
                                   type="email"
                                   placeholder="email@gmail.com"
                                   autoComplete="whatever" //Some browsers attempt to autosuggest even it's value set to off.
                              />
                              <TextInput
                                   label="Password*"
                                   name="password"
                                   type="password"
                                   placeholder="Password"
                                   autoComplete="off" 
                              />
                              <div className="button">
                                   <button type="submit" disabled={isSubmitting} className="flex w-96 justify-center rounded bg-primary p-3 mt-8 font-medium text-gray">
                                        {isSubmitting ? 'Submitting...' : 'Submit'}
                                   </button>
                              </div>
                         </Form>
                    )}
               </Formik>
          </div>
     )
}