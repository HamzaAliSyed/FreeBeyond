import { useState, useReducer } from "react";

function ReducerCharacterStats(state, action) {}

const CCTabNameStats = () => {
  const [ccNameTabSelected, setCCNameTabSelected] = useState(true);
  const [characterName, setCharacterName] = useState("");
  const [characterImage, setCharacterImage] = useState(false);
  const [statsState, statsDispatch] = useReducer(ReducerCharacterStats, {
    strength: 0,
    dexterity: 0,
    constitution: 0,
    intelligence: 0,
    wisdom: 0,
    charisma: 0,
  });

  const [rollResults, setRollResults] = useState([]);
  const [timesrolled, setTimesRolled] = useState(0);

  function RollForAStat() {
    if (timesrolled < 7) {
      const rolls = [];
      for (let i = 0; i < 4; i++) {
        rolls.push(Math.floor(Math.random() * 6) + 1);
      }

      rolls.sort((a, b) => a - b);
      const sum = rolls.slice(1).reduce((acc, roll) => acc + roll, 0);
      setRollResults([...rollResults, sum]);
      setTimesRolled(timesrolled + 1);
    } else {
      alert("You have rolled for alloted 7 times");
    }
  }
  return (
    <div>
      <div>
        <label>Name:</label>
        <input
          type="text"
          value={characterName}
          onChange={(e) => setCharacterName(e.target.value)}
        />
      </div>
      <div>
        <label>Image:</label>
        <input
          type="text"
          value={characterImage}
          onChange={(e) => setCharacterImage(e.target.value)}
        />
      </div>
      <div>
        <button onClick={RollForAStat}>Roll for stats</button>
      </div>
      <div>
        <label>Strength:</label>
        <input type="text" value={statsState.strength} readOnly />
      </div>
      <div>
        <label>Dexterity:</label>
        <input type="text" value={statsState.dexterity} readOnly />
      </div>
      <div>
        <label>Constitution:</label>
        <input type="text" value={statsState.constitution} readOnly />
      </div>
      <div>
        <label>Intelligence:</label>
        <input type="text" value={statsState.intelligence} readOnly />
      </div>
      <div>
        <label>Wisdom:</label>
        <input type="text" value={statsState.wisdom} readOnly />
      </div>
      <div>
        <label>Charisma:</label>
        <input type="text" value={statsState.charisma} readOnly />
      </div>
      <div>
        <h3>Roll Results:</h3>
        <ul>
          {rollResults.map((result, index) => (
            <li key={index}>
              Roll {index + 1}: {result}
            </li>
          ))}
        </ul>
      </div>
    </div>
  );
};
export default CCTabNameStats;
