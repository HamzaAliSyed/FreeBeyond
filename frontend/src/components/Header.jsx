import { Navigate, useNavigate } from 'react-router-dom';
import React, { useState, useEffect } from 'react';

function Header() {
    const [isSignedIn, setIsSignedIn] = useState(false)
    const [username, setUsername] = useState("")
    const [profilePicture, setProfilePicture] = useState('')
    const accountCreateNavigate = useNavigate()
    const signInNavigate = useNavigate()
    const gotoHome = useNavigate()

    useEffect( 
        () => {
            const storedUsername = localStorage.getItem("Username")
            if (storedUsername) {
                setUsername(storedUsername)
                setIsSignedIn(true)
            }
        }, []
    )

    const handleCreateAccountClick = async () => {
        try {
            const response = await fetch('http://localhost:2712/api/accounts/createrequest', {
                method: 'POST',
            });

            if (response.ok) {
                // If the response is OK, display the CreateAccountForm component in a popup
                accountCreateNavigate("/create-account")
                // Logic to display the CreateAccountForm component goes here
            } else {
                // If the response is not 200, show a server down message
                alert("Server is down. Please try again later.");
            }
        } catch (error) {
            console.error('Error fetching data:', error);
            alert("An error occurred while communicating with the server.");
        }
    };

    const handleSignInClick = async () => {
        try {
            const response =await fetch('http://localhost:2712/api/accounts/signinrequest', {
                method: 'POST',
            })

            if (response.ok) {
                signInNavigate("/sign-in")
            } else {
                // If the response is not 200, show a server down message
                alert("Server is down. Please try again later.");
            }
        } catch (error) {
            console.error('Error fetching data:', error);
            alert("An error occurred while communicating with the server.");
        }
    }

    const handleSignOutClick = () => {
        localStorage.removeItem('token')
        localStorage.removeItem('Username')
        setIsSignedIn(false)
        setUsername("")
        gotoHome("/")
    }

    return (
        <nav className='flex w-full h-20 justify-between items-center'>
            <div className="p-4">logo</div>
            <div className='p-4'>search</div>
            <div className="p-4">
                <ul>
                    {isSignedIn ? (
                        <>
                            <li>
                            <span>Welcome, {username}</span>
                            </li>
                            <li>
                                <a href="#" onClick={handleSignOutClick}>
                                    Sign Out
                                </a>
                            </li>
                        </>
                    ) : (
                        <>
                            <li>
                                <a href='#' onClick={handleSignInClick}>
                                    Sign In
                                </a>
                            </li>
                            <li>
                                <a href='#' onClick={handleCreateAccountClick}>
                                    Create Account
                                </a>
                            </li>
                        </>
                    )}
                </ul>
            </div>
        </nav>
    )
}

export default Header