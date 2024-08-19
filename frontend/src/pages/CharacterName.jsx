import React, { useState } from 'react';

const CharacterName = () => {
  const [characterName, setCharacterName] = useState('');

  const handleCharacterNameChange = (event) => {
    setCharacterName(event.target.value);
  };

  return (
    <div>
      <h2>Enter your new character name</h2>
      <form>
        <label htmlFor="characterName">Character Name:</label>
        <input type="text" id="characterName" value={characterName} onChange={handleCharacterNameChange} />
      </form>
    </div>
  );
};

export default CharacterName;