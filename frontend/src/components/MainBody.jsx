import React, { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";

const MainBody = () => {
    const navigate = useNavigate()
    const [hasUsername, setHasUsername] = useState(false)
    useEffect(
        () => {
            const storedUsername = localStorage.getItem('Username')
            setHasUsername(storedUsername !== null)
        }
    )
    const HandleCreateCharacter = async () => {
        if (hasUsername) {
            const storedUsername = localStorage.getItem('Username');
      const response = await fetch('http://localhost:2712/api/accounts/createacharacter', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({ username:   
 storedUsername })
      });

      if (response.ok) {
        navigate('/create-character');
      } else {
        alert('Internal Server Error');
      }

        } else {
            navigate('/sign-in')
        }
    }
return (
    <div>
        <button onClick={HandleCreateCharacter}>Create a Character</button>
    </div>
)
}

export default MainBody