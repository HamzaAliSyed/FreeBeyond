import React, { useState } from "react";

const CCTabNameStats = ({
  data,
  onDataChange,
  characterStatsState,
  characterStatsDispatcher,
}) => {
  const stats = [
    "strength",
    "dexterity",
    "constitution",
    "intelligence",
    "wisdom",
    "charisma",
  ];
  const [rolledStats, setRolledStats] = useState([]);
  const [abilityModifiers, setAbilityModifiers] = useState({});

  const handleStatDrop = async (e, stat) => {
    e.preventDefault();
    const statId = e.dataTransfer.getData("text");
    const statObject = rolledStats.find((s) => s.id === statId);
    if (statObject) {
      try {
        // Update the stat in the parent component's state
        characterStatsDispatcher({
          type: "UPDATE_STAT",
          payload: { stat, value: statObject.value },
        });

        // Update local state
        onDataChange({ [stat]: statObject.value });

        // Fetch the ability modifier
        const response = await fetch(
          "http://localhost:2712/api/components/getabilitymodifier",
          {
            method: "POST",
            headers: {
              "Content-Type": "application/json",
            },
            body: JSON.stringify({ abilityscore: statObject.value }),
          }
        );

        if (!response.ok) {
          throw new Error("Something wrong at the server");
        }

        const data = await response.json();

        // Update the ability modifier
        setAbilityModifiers((prevModifiers) => ({
          ...prevModifiers,
          [stat]: data.abilityscoremodifier,
        }));

        // Remove the used stat from rolledStats
        setRolledStats((prevStats) => prevStats.filter((s) => s.id !== statId));

        console.log(
          `Updated ${stat} to ${statObject.value} with modifier ${data.abilityModifier}`
        );
      } catch (error) {
        console.error(
          "Error updating stat and fetching ability modifier:",
          error
        );
      }
    }
  };
  const handleCharacterPortrait = (event) => {
    const file = event.target.files[0];
    if (file) {
      const reader = new FileReader();
      reader.onload = (event2) =>
        onDataChange({ characterPortrait: event2.target.result });
      reader.readAsDataURL(file);
    }
  };

  const handleNameChange = (event) => {
    onDataChange({ name: event.target.value });
  };

  const rollADie = () => {
    let roll = Math.random() * 6 + 1;
    roll = Math.floor(roll);
    return roll;
  };

  const rollAStat = () => {
    let statroll = Array(4);
    statroll = statroll.fill();
    statroll = statroll.map(rollADie);
    statroll = statroll.sort((a, b) => b - a);
    statroll = statroll.slice(0, 3);
    let sum = statroll.reduce((a, b) => a + b, 0);
    return sum;
  };

  const handleRollForStats = () => {
    const newStats = Array(7)
      .fill()
      .map(() => ({
        id: Math.random().toString(36).substr(2, 9),
        value: rollAStat(),
      }));
    setRolledStats(newStats);
  };

  console.log("Current ability modifiers:", abilityModifiers);

  return (
    <div className="w-full h-[calc(100vh-48px)] bg-gray-200 rounded-md p-3">
      <h1 className="text-gray-800 text-center text-4xl mb-3">Name & Stats</h1>
      <div className="grid grid-cols-2 gap-4 h-[calc(100%-4rem)]">
        <div className="px-5 py-3">
          <input
            type="text"
            value={data.name || ""}
            onChange={handleNameChange}
            placeholder="Character Name"
            className="w-full p-2 border rounded"
          />
          <div className="flex flex-col space-y-2">
            {stats.map((stat) => (
              <div key={stat} className="flex items-center mb-2">
                <label className="w-24 capitalize">{stat}:</label>
                <div
                  className="w-16 h-16 border-2 my-4 border-gray-400 rounded flex items-center justify-center"
                  onDragOver={(e) => e.preventDefault()}
                  onDrop={(e) => handleStatDrop(e, stat)}
                >
                  {characterStatsState[stat] || ""}
                </div>
                <div className="ml-2 w-8 h-8 border border-gray-400 rounded flex items-center justify-center">
                  {abilityModifiers[stat] !== undefined ? (
                    <>
                      {abilityModifiers[stat] >= 0 ? "+" : ""}
                      {abilityModifiers[stat]}
                    </>
                  ) : (
                    ""
                  )}
                </div>
              </div>
            ))}
          </div>
          <div>
            <button
              onClick={handleRollForStats}
              className="bg-gray-500 text-white py-2 px-4 rounded hover:bg-gray-600"
            >
              Roll for Stats
            </button>
            <div className="mt-4 flex flex-wrap justify-center">
              {rolledStats.map((stat) => (
                <div
                  key={stat.id}
                  draggable
                  onDragStart={(e) => e.dataTransfer.setData("text", stat.id)}
                  className="bg-gray-200 p-2 rounded shadow cursor-move flex items-center justify-center w-16 h-16 m-1"
                >
                  {stat.value}
                </div>
              ))}
            </div>
          </div>
        </div>
        <div className="flex flex-col h-full">
          <h2 className="text-xl font-semibold mb-2">Character Portrait</h2>
          <label
            className="flex-grow bg-gray-300 flex items-center justify-center cursor-pointer rounded-md overflow-hidden"
            htmlFor="image-upload"
          >
            {data.characterPortrait ? (
              <img
                src={data.characterPortrait}
                alt="Character Portrait"
                className="w-full h-full object-cover"
              />
            ) : (
              <span className="text-black text-center p-4">
                Click here to upload character portrait
              </span>
            )}
          </label>
          <input
            id="image-upload"
            type="file"
            accept="image/*"
            className="hidden"
            onChange={handleCharacterPortrait}
          />
        </div>
      </div>
    </div>
  );
};

export default CCTabNameStats;
