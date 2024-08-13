import Archer from "../assets/ranger.jpg"
import { Navigate } from "react-router-dom"


function SignInPage() {


return (<div className="grid grid-cols-12 h-screen"> 
<div className="col-span-3">
    <img src={Archer} className="w-full h-full object-cover" />
</div>
<div className="col-span-9 flex items-center justify-center">
    <form className="w-2/3 bg-white p-8 shadow-lg rounded-lg">
    <div className="mb-4">
                        <label className="block text-gray-700 text-sm font-bold mb-2" htmlFor="userName">
                            Username
                        </label>
                        <input className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline" id="userName" type="text" placeholder="UserName" required/>
                    </div>
                    <div className="mb-4">
                        <label className="block text-gray-700 text-sm font-bold mb-2" htmlFor="password">
                            Password
                        </label>
                        <input className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 mb-3 leading-tight focus:outline-none focus:shadow-outline" id="password" type="password" placeholder="Password" required/>
                    </div>
                    <div className="flex items-center justify-between">
                        <button className="bg-[#BC0F0F] hover:bg-red-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline" type="submit">
                            Sign In
                        </button>
                    </div>
    </form>
</div>
</div>)
}

export default SignInPage