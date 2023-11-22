import { useState } from "react";
import { Formik, Form, useField } from "formik";
import * as Yup from "yup";
import { AuthData } from "../../auth/AuthWrapper"

export const Login = () => {

     const { login } = AuthData();
     const [ errorMessage, setErrorMessage ] = useState(null)

     const TextInput = ({ label, ...props }) => {
          const [field, meta] = useField(props);
          return (
            <>
              <label className="mb-2 mt-6 block text-black" htmlFor={props.id || props.name}>{label}</label>
              <input autoComplete="off"  className="w-96 rounded-lg border-[1.5px] border-stroke bg-transparent py-3 px-5 font-medium outline-none transition focus:border-primary active:border-primary disabled:cursor-default disabled:bg-whiter dark:border-form-strokedark dark:bg-form-input dark:focus:border-primary" {...field} {...props} />
              {meta.touched && meta.error ? (
                <p className="text-red w-96 mt-4">{meta.error}</p>
              ) : null}
            </>
          );
     };

     return (
          <div className="flex flex-col items-center">
               <div className="w-96 mt-8 flex justify-between">
                    <h2 className="text-2xl font-bold text-black">Login</h2>
                    <div className="flex"><p>Don't have an account?</p><a className="pl-2 inline-block text-blue" href="/signup">Sign Up</a></div>
               </div>
               {errorMessage ?
                    <p className="text-red w-96 mt-8">{errorMessage}</p>
               : null }
               <Formik
                    initialValues={{
                         email: "",
                         password: ""
                    }}
                    validationSchema={Yup.object({
                         email: Yup.string()
                           .required("Required"),
                         password: Yup.string()
                           .required("Required"),
                    })}
                    onSubmit={async (values, { setSubmitting }) => {
                         try {
                              await login(values.email, values.password)
                         } catch (error) {
                              setErrorMessage(error)
                         }
                         setSubmitting(false);
                    }}
               >
                    {({ isSubmitting }) => (
                         <Form>
                              <TextInput
                                   label="Email*"
                                   name="email"
                                   type="email"
                                   placeholder="Enter email"
                                   autoComplete="whatever"
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