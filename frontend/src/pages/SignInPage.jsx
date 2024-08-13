import Archer from "../assets/ranger.jpg"
import { useNavigate } from "react-router-dom"
import { useState } from "react"


function SignInPage() {

const [userName, setUserName] = useState("")
const [password, setPassword] = useState("")
const navigate = useNavigate()

const handleSubmit = async (e) => {
    e.preventDefault()

    const credentials = {
        Username: userName,
        Password: password,
    }

    try {
        const response = await fetch("http://localhost:2712/api/accounts/signin",{
            method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify(credentials),
        })

        if (response.ok) {
            const data = await response.json()
            const token = data.token

            localStorage.setItem("token", token)
            alert("Signed In")
            navigate("/");
        } else {
            const errorText = await response.text();
            alert(`Sign in failed: ${errorText}`);
        }
    } catch (error) {
        console.error("Error during sign-in:", error);
        alert("An error occurred during sign-in. Please try again.");
    }
}

return (<div className="grid grid-cols-12 h-screen"> 
<div className="col-span-3">
    <img src={Archer} className="w-full h-full object-cover" />
</div>
<div className="col-span-9 flex items-center justify-center">
    <form className="w-2/3 bg-white p-8 shadow-lg rounded-lg" onSubmit={handleSubmit}>
    <div className="mb-4">
                        <label className="block text-gray-700 text-sm font-bold mb-2" htmlFor="userName">
                            Username
                        </label>
                        <input className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline" id="userName" type="text" placeholder="UserName" value={userName} onChange={(e) => setUserName(e.target.value)} required/>
                    </div>
                    <div className="mb-4">
                        <label className="block text-gray-700 text-sm font-bold mb-2" htmlFor="password">
                            Password
                        </label>
                        <input className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 mb-3 leading-tight focus:outline-none focus:shadow-outline" id="password" type="password" placeholder="Password" value={password} onChange={(e) => setPassword(e.target.value)} required/>
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