export const Home = () => {

     return (
          <div className="flex flex-col items-center">
               <div className="text-start mt-8">
                    <h1 className="text-2xl font-bold text-black mb-4 uppercase font-bold">Auth flow demo</h1>
                    <ul className="list-disc list-inside font-semibold">
                         <li>Golang backend(Echo framework) - JWT Auth</li>
                         <li>Postgres</li>
                         <li>OpenAPI Specification endpoint</li>
                         <li>React SPA</li>
                         <li>Access token stored as secure cookie</li>
                         <li>Tailwind CSS</li>
                         <li>Vite.js</li>
                         <li>Formik & Yup</li>
                    </ul>
                    <div className="mt-4">
                         <p className="text-xl font-bold text-black uppercase font-bold">TODO:</p>
                         <ul className="list-disc list-inside font-semibold">
                              <li>Email setup</li>
                              <li>Confirm email</li>
                              <li>Forgot password</li>
                         </ul>
                    </div>
               </div>
          </div>
     )
}