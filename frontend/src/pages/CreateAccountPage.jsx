import { useState } from "react";
import healerimage from "../assets/healer.jpg";
import { useNavigate } from 'react-router-dom';


function CreateAccountPage() {
    const [firstName, setFirstName] = useState("")
    const [lastName, setLastName] = useState("")
    const [userName, setUserName] = useState("")
    const [password, setPassword] = useState("")
    const [email, setEmail] = useState("")
    const [validPassword, setValidPassword] = useState("")
    const [touchedpassword, settouchedPassword] = useState(false)
    const [confirmPassword, setConfirmedPassword] = useState("")
    const [touchedConfirmPassword, setTouchedConfirmPassword] = useState(false)
    const [emailValid, setEmailValid] = useState("")
    const [touchedEmail, setTouchedEmail] = useState(false)
    const [isUserNameValid, setUsernameValid] = useState("")
    const [isUserNameTouched, setUserNameTouched] = useState(false)
    const navigate = useNavigate();


    const validatePassword = (password) => {
        const minLength = password.length >= 10
        const hasUpperCase = /[A-Z]/.test(password)
        const hasLowerCase = /[a-z]/.test(password)
        const hasNumber = /[0-9]/.test(password)
        const hasSpecialChar = /[!@#$%^&*(),.?":{}|<>]/.test(password)

        return minLength && hasUpperCase && hasLowerCase && hasNumber && hasSpecialChar
    }

    const renderpasswordValidation = () => {
        if (touchedpassword == false) {
            return null
        } else {
            if (validPassword == false ) {
                return (
                    <div className="text-red-500 mt-2 text-sm flex items-center">
                        <span className="mr-2">❌</span>
                        Password must contain one capital letter, one lowercase letter, one number, and one special character.
                    </div>
                );
            } else {
                return (
                    <div className="text-green-500 mt-2 text-sm flex items-center">
                        <span className="mr-2">✔️</span>
                        Password is valid.
                    </div>
                )
            }
        }
    }

    const renderConfirmPasswordValidation = () => {
        if (touchedConfirmPassword == false) {
            return null
        } else {
            if (confirmPassword !== password) {
                return (
                    <div className="text-red-500 mt-2 text-sm flex items-center">
                        <span className="mr-2">❌</span>
                        Confirm password doesn't match.
                    </div>
                )
            } else {
                return (
                    <div className="text-green-500 mt-2 text-sm flex items-center">
                        <span className="mr-2">✔️</span>
                        Confirm password matches.
                    </div>
                )
            }
        }
    }

    const renderEmailValidation = () => {
        if (touchedEmail == false) {
            return null
        } else {
            if (emailValid == false) {
                return (
                    <div className="text-red-500 mt-2 text-sm flex items-center">
                        <span className="mr-2">❌</span>
                        Email must be longer than 3 characters and contain "@".
                    </div>
                )
            } else {
                return (
                    <div className="text-green-500 mt-2 text-sm flex items-center">
                        <span className="mr-2">✔️</span>
                        Email is valid.
                    </div>
                )
            }
        }
    }

    const renderUserNameMatch = () => {
        if (isUserNameTouched == false) {
            return null
        } else {
            if(isUserNameValid == false) {
                return (
                    <div className="text-red-500 mt-2 text-sm flex items-center">
                        <span className="mr-2">❌</span>
                        Please select a new username. This one is already taken.
                    </div>
                )
            }
        }
    }

    const handlePasswordChange = (e) => {
        const newPassword = e.target.value
        setPassword(newPassword)

        settouchedPassword(true)

        if (validatePassword(newPassword) != true) {
            setValidPassword(false)
        }
        else {
            setValidPassword(true)
        }

    }

    const handleConfirmPassword = (e) => {
        const newConfirmPassword = e.target.value
        setConfirmedPassword(newConfirmPassword)
        setTouchedConfirmPassword(true)
    }

    const handleEmailChange = (e) => {
        const newEmail = e.target.value
        setEmail(newEmail)

        setTouchedEmail(true)
        if (newEmail.length >= 3 && newEmail.includes('@')) {
            setEmailValid(true)
        } else {
            setEmailValid(false)
        }
    }

    const handleUserNameChange = async (e) => {
        const newUsername = e.target.value
        setUserNameTouched(true)
        setUserName(newUsername)
        try {
            const response = await fetch (
                'http://127.0.0.1:2712/api/accounts/usernamematchrequest',
                {
                    method: 'POST',
                    headers : {
                        'Content-Type': 'application/json',
                    },body:JSON.stringify({ username: newUsername }),
                }
            )

            if (response.status === 409 ) {
                setUsernameValid(false)
            } else {
                setUsernameValid(true)
            }
        }
        catch (error) {
            setUsernameValid(false)
        }
    }

    const handleSubmit = async (e) => {
        e.preventDefault()
        const userData = {
            firstName,
            lastName,
            userName,
            email,
            password,
            role : "player"
        }
        try {
            const response = await fetch (
                "http://127.0.0.1:2712/api/accounts/createaccount",
                {
                    method : "POST",
                    headers : {
                        'Content-Type': 'application/json',
                    },
                    body: JSON.stringify(userData),
                }
            )

            if (response.status === 201 ) {
                alert("Account created successfully!");
                navigate('/')
            } else {
                const errorText = await response.text()
                alert(`Error: ${errorText}`)
            }
        } catch (error){
            console.error('Error submitting form:', error);
        }
    }

    return (
        <div className="grid grid-cols-12 h-screen">
            <div className="col-span-3">
                <img src={healerimage} className="w-full h-full object-cover" />
            </div>
            <div className="col-span-9 flex items-center justify-center">
                <form className="w-2/3 bg-white p-8 shadow-lg rounded-lg" onSubmit={handleSubmit}>
                    <h2 className="text-2xl font-bold mb-6 text-center">Create Your Account</h2>

                    {/* First Name */}
                    <div className="mb-4">
                        <label className="block text-gray-700 text-sm font-bold mb-2" htmlFor="firstName">
                            First Name
                        </label>
                        <input className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline" id="firstName" type="text" placeholder="First Name" value={firstName}  onChange={(e) => setFirstName(e.target.value)} required/>
                    </div>

                    {/* Last Name */}
                    <div className="mb-4">
                        <label className="block text-gray-700 text-sm font-bold mb-2" htmlFor="lastName">
                            Last Name
                        </label>
                        <input className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline" id="lastName" type="text" placeholder="Last Name"  value={lastName} onChange={(e) => setLastName(e.target.value)} required/>
                    </div>

                    {/* Username */}
                    <div className="mb-4">
                        <label className="block text-gray-700 text-sm font-bold mb-2" htmlFor="username">
                            Username
                        </label>
                        <input className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline" id="username" type="text" placeholder="Username" value={userName} onChange={handleUserNameChange} required/>
                    </div>
                    {renderUserNameMatch()}

                    {/* Email */}
                    <div className="mb-4">
                        <label className="block text-gray-700 text-sm font-bold mb-2" htmlFor="email">
                            Email
                        </label>
                        <input className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline" id="email" type="email" placeholder="Email"  value={email} onChange={handleEmailChange} required/>
                    </div>
                    { renderEmailValidation()}

                    
                    {/* Password */}
                    <div className="mb-4">
                        <label className="block text-gray-700 text-sm font-bold mb-2" htmlFor="password">
                            Password
                        </label>
                        <input className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 mb-3 leading-tight focus:outline-none focus:shadow-outline" id="password" type="password" placeholder="Password"  value={password} onChange={handlePasswordChange}/>
                    </div>
                    {
                       renderpasswordValidation()
                    }

                    {/* Confirm Password */}
                    <div className="mb-6">
                        <label className="block text-gray-700 text-sm font-bold mb-2" htmlFor="confirmPassword">
                            Confirm Password
                        </label>
                        <input className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline" id="confirmPassword" type="password" placeholder="Confirm Password"  value={confirmPassword} onChange={handleConfirmPassword} required/>
                    </div>
                    {
                        renderConfirmPasswordValidation()
                    }
                    {/* Sign Up Button */}
                    <div className="flex items-center justify-between">
                        <button className="bg-[#BC0F0F] hover:bg-red-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline" type="submit">
                            Sign Up
                        </button>
                    </div>
                </form>
            </div>
        </div>
    );
}

export default CreateAccountPage;
