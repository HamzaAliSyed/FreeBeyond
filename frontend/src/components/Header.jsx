import { useNavigate } from 'react-router-dom';
function Header() {
    const accountCreateNavigate = useNavigate()
    const signInNavigate = useNavigate()
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
            const response =await fetch('http://localhost:2712/api/accounts/signinrequest', {})

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

    return (
        <nav className="bg-black">
            <ul className='flex justify-end py-3 px-6 border-b'>
                <li><a 
                className="bg-[#BC0F0F]
                        my-3
                        mx-2
                        box-border 
                        text-white 
                        cursor-pointer 
                        block 
                        font-roboto 
                        text-base 
                        font-bold 
                        outline-none 
                        p-[9px_20px] 
                        no-underline 
                        tap-highlight-transparent"
                        onClick={handleCreateAccountClick}
                >Create Account</a></li>
                <li><a className="
                bg-[#BC0F0F]
                        my-3
                        mx-2
                        box-border 
                        text-white 
                        cursor-pointer 
                        block 
                        font-roboto 
                        text-base 
                        font-bold 
                        outline-none 
                        p-[9px_20px] 
                        no-underline 
                        tap-highlight-transparent
                " onClick={handleSignInClick}
                >Sign In</a></li>
            </ul>
        </nav>
    )
}

export default Header