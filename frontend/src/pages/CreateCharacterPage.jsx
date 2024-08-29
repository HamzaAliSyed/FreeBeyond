import React, { useReducer, useState } from "react";
import CCTabBackground from "./pagespecificcomponents/cctabbackground";
import CCTabClass from "./pagespecificcomponents/cctabclass";
import CCTabRace from "./pagespecificcomponents/cctabrace";
import CCTabNameStats from "./pagespecificcomponents/cctabnamestats";

const CreateCharacterPage = () => {
  const [newCharacterData, setNewCharacterData] = useState({
    namestats: { name: "", characterPortrait: undefined },
    background: {},
    race: {},
    class: {},
  });

  const initialStats = {
    strength: 0,
    dexterity: 0,
    constitution: 0,
    intelligence: 0,
    wisdom: 0,
    charisma: 0,
  };

  const characterStatsReducer = (state = initialStats, action) => {
    switch (action.type) {
      case "UPDATE_STAT":
        return {
          ...state,
          [action.payload.stat]: action.payload.value,
        };
      default:
        return state;
    }
  };

  const updateStat = (stat, value) => {
    return {
      type: "UPDATE_STAT",
      payload: { stat, value },
    };
  };

  const [characterStatsState, CharacterStatsDispatcher] = useReducer(
    characterStatsReducer,
    initialStats
  );

  const tabs = [
    { id: "namestats", label: "Name and Stats", component: CCTabNameStats },
    { id: "background", label: "Background", component: CCTabBackground },
    { id: "race", label: "Race", component: CCTabRace },
    { id: "class", label: "Class", component: CCTabClass },
  ];
  const [activeTab, setActiveTab] = useState("namestats");

  const handleTabDataChange = (tabId, newData) => {
    if (tabId === "namestats") {
      Object.entries(newData).forEach(([key, value]) => {
        if (initialStats.hasOwnProperty(key)) {
          CharacterStatsDispatcher(updateStat(key, value));
        }
      });
    }
    setNewCharacterData((prevData) => ({
      ...prevData,
      [tabId]: { ...prevData[tabId], ...newData },
    }));
  };

  const handleSubmit = () => {
    console.log("Submitting character data:", newCharacterData);
  };

  return (
    <div className="max-w-4xl mx-auto mt-8">
      <div className="border-b border-gray-200">
        <nav className="-mb-px flex">
          {tabs.map((tab) => (
            <button
              key={tab.id}
              onClick={() => setActiveTab(tab.id)}
              className={`py-2 px-4 text-sm font-medium ${
                activeTab === tab.id
                  ? "border-b-2 border-blue-500 text-blue-600"
                  : "text-gray-500 hover:text-gray-700"
              }`}
            >
              {tab.label}
            </button>
          ))}
        </nav>
      </div>
      <div className="mt-4">
        {tabs.map((tab) => (
          <div key={tab.id} className={activeTab === tab.id ? "" : "hidden"}>
            <tab.component
              data={newCharacterData[tab.id]}
              onDataChange={(newData) => handleTabDataChange(tab.id, newData)}
            />
          </div>
        ))}
      </div>
      <div className="mt-6 text-right">
        <button
          onClick={handleSubmit}
          className="px-4 py-2 bg-gray-500 text-white rounded hover:bg-gray-600"
        >
          Submit Character
        </button>
      </div>
    </div>
  );
};

export default CreateCharacterPage;
