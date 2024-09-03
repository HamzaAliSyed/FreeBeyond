import LevelUpImage from "../assets/leveluppage.jpg";
import { useState, useEffect } from "react";
import { useNavigate } from "react-router-dom";

const CreateLevelUp = () => {
  const navigate = useNavigate();
  const [classes, setClasses] = useState([]);
  const [selectedClass, setSelectedClass] = useState("");

  const [level, setLevel] = useState("");

  const [proficiencyBonus, setProficiencyBonus] = useState("");

  const [textbasedAbilities, setTextbasedAbilities] = useState([
    {
      title: "",
      flavourText: "",
      availability: "",
      ruleType: "",
    },
  ]);

  const addTextbasedAbility = () => {
    setTextbasedAbilities([
      ...textbasedAbilities,
      {
        title: "",
        flavourText: "",
        availability: "",
        ruleType: "",
      },
    ]);
  };

  const [chargedAbilities, setChargedAbilities] = useState([
    {
      title: "",
      flavorText: "",
      totalCharge: "",
      numberOfChargesRemaining: "",
      availability: "",
      ruleType: "",
      recovery: "",
    },
  ]);

  const [modifierAbilities, setModifierAbilities] = useState([
    {
      title: "",
      flavorText: "",
      ruleType: "",
      availability: "",
      modifierStatFamily: "",
      modifierStat: "",
      modifierValue: "",
    },
  ]);

  const removeTextbasedAbility = (index) => {
    setTextbasedAbilities(textbasedAbilities.filter((_, i) => i !== index));
  };

  const handleTextbasedAbilityChange = (index, field, value) => {
    const newAbilities = [...textbasedAbilities];
    newAbilities[index][field] = value;
    setTextbasedAbilities(newAbilities);
  };

  const fetchClasses = async () => {
    try {
      const response = await fetch(
        "http://localhost:2712/api/components/getallclasses"
      );
      if (!response.ok) throw new Error("Failed to fetch classes");
      const data = await response.json();
      setClasses(data);
    } catch (error) {
      console.error("Error fetching classes:", error);
    }
  };

  const addModifierAbility = () => {
    setModifierAbilities([
      ...modifierAbilities,
      {
        title: "",
        flavorText: "",
        ruleType: "",
        availability: "",
        modifierStatFamily: "",
        modifierStat: "",
        modifierValue: "",
      },
    ]);
  };

  const removeModifierAbility = (index) => {
    setModifierAbilities(modifierAbilities.filter((_, i) => i !== index));
  };

  const handleModifierAbilityChange = (index, field, value) => {
    const newAbilities = [...modifierAbilities];
    newAbilities[index][field] = value;
    setModifierAbilities(newAbilities);
  };

  const modifierStatFamilies = {
    "Ability Scores": [
      "Strength",
      "Dexterity",
      "Constitution",
      "Intelligence",
      "Wisdom",
      "Charisma",
    ],
    Skills: [
      "Acrobatics",
      "Animal Handling",
      "Arcana",
      "Athletics",
      "Deception",
      "History",
      "Insight",
      "Intimidation",
      "Investigation",
      "Medicine",
      "Nature",
      "Perception",
      "Performance",
      "Persuasion",
      "Religion",
      "Sleight of Hand",
      "Stealth",
      "Survival",
    ],
    "Saving Throws": [
      "Strength",
      "Dexterity",
      "Constitution",
      "Intelligence",
      "Wisdom",
      "Charisma",
    ],
    Other: ["AC", "Initiative", "Speed"],
  };

  const handleSubmit = async (e) => {
    e.preventDefault();

    const levelUpData = {
      class: selectedClass,
      level: parseInt(level, 10),
      proficiencybonus: parseInt(proficiencyBonus, 10),
      flavourabilities: textbasedAbilities.map((ability) => ({
        title: ability.title,
        flavourtext: ability.flavourText,
        availability: ability.availability.toLowerCase(),
        ruletype: ability.ruleType.toLowerCase(),
      })),
      chargebasedability: chargedAbilities.map((ability) => ({
        title: ability.title,
        flavourtext: ability.flavorText,
        totalcharges: parseInt(ability.totalCharge, 10),
        numberofchargesremaining: parseInt(
          ability.numberOfChargesRemaining,
          10
        ),
        availability: ability.availability.toLowerCase(),
        ruletype: ability.ruleType.toLowerCase(),
        recovery: ability.recovery,
      })),
      modifierability: modifierAbilities.map((ability) => ({
        title: ability.title,
        flavourtext: ability.flavorText,
        ruletype: ability.ruleType.toLowerCase(),
        availability: ability.availability.toLowerCase(),
        modifierstatfamily: ability.modifierStatFamily,
        modifierstat: ability.modifierStat,
        modifiervalue: ability.modifierValue,
      })),
      spellcasting: {}, // Add appropriate spellcasting data if needed
    };

    console.log("Sending data:", JSON.stringify(levelUpData, null, 2));

    try {
      const response = await fetch(
        "http://localhost:2712/api/components/addlevel",
        {
          method: "PUT",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify(levelUpData),
        }
      );

      if (!response.ok) {
        alert("Couldnt add level");
        throw new Error(`HTTP error! status: ${response.status}`);
      }

      const result = await response.json();
      console.log("Level Up added successfully:", result);
      alert("Added the level to the Class");
      navigate("/");
    } catch (error) {
      console.error("Error adding Level Up:", error);
    }
  };
  useEffect(() => {
    fetchClasses();
  }, []);

  const addChargedAbility = () => {
    setChargedAbilities([
      ...chargedAbilities,
      {
        title: "",
        flavorText: "",
        totalCharge: "",
        numberOfChargesRemaining: "",
        availability: "",
        ruleType: "",
        recovery: "",
      },
    ]);
  };

  const removeChargedAbility = (index) => {
    setChargedAbilities(chargedAbilities.filter((_, i) => i !== index));
  };

  const handleChargedAbilityChange = (index, field, value) => {
    const newAbilities = [...chargedAbilities];
    newAbilities[index][field] = value;
    setChargedAbilities(newAbilities);
  };

  return (
    <div className="create-level-up-container grid grid-cols-12 h-screen">
      <div className="col-span-3 h-full">
        <img
          src={LevelUpImage}
          className="w-full object-cover h-full"
          alt="Level Up"
        />
      </div>
      <div className="col-span-9 p-8 flex flex-col items-center justify-start">
        <div className="w-full max-w-2xl text-center mb-8">
          <h1 className="text-4xl font-bold">Add a level to class</h1>
        </div>
        <form className="w-full max-w-2xl" onSubmit={handleSubmit}>
          <div className="mb-6">
            <label
              htmlFor="class"
              className="block text-xl font-semibold text-gray-700 mb-2"
            >
              Class
            </label>
            <select
              id="class"
              value={selectedClass}
              onChange={(e) => setSelectedClass(e.target.value)}
              className="w-full text-lg p-2 border-2 border-gray-300 rounded-lg"
            >
              <option value="">Select a class</option>
              {classes.map((cls, index) => (
                <option key={index} value={cls}>
                  {cls}
                </option>
              ))}
            </select>
          </div>
          <div className="mb-6">
            <label
              htmlFor="level"
              className="block text-xl font-semibold text-gray-700 mb-2"
            >
              Level
            </label>
            <select
              id="level"
              value={level}
              onChange={(e) => setLevel(e.target.value)}
              className="w-full text-lg p-2 border-2 border-gray-300 rounded-lg"
            >
              <option value="">Select a level</option>
              {[...Array(20)].map((_, i) => (
                <option key={i} value={i + 1}>
                  {i + 1}
                </option>
              ))}
            </select>
          </div>
          <div className="mb-6">
            <label
              htmlFor="proficiencyBonus"
              className="block text-xl font-semibold text-gray-700 mb-2"
            >
              Proficiency Bonus
            </label>
            <select
              id="proficiencyBonus"
              value={proficiencyBonus}
              onChange={(e) => setProficiencyBonus(e.target.value)}
              className="w-full text-lg p-2 border-2 border-gray-300 rounded-lg"
            >
              <option value="">Select proficiency bonus</option>
              {[...Array(9)].map((_, i) => (
                <option key={i} value={i + 2}>
                  {i + 2}
                </option>
              ))}
            </select>
          </div>
          <div className="mb-6">
            <label className="block text-xl font-semibold text-gray-700 mb-2">
              Textbased Abilities
            </label>
            {textbasedAbilities.map((ability, index) => (
              <div
                key={index}
                className="mb-4 p-4 border border-gray-300 rounded-lg relative"
              >
                <input
                  type="text"
                  value={ability.title}
                  onChange={(e) =>
                    handleTextbasedAbilityChange(index, "title", e.target.value)
                  }
                  className="w-full text-lg p-2 mb-2 border-2 border-gray-300 rounded-lg"
                  placeholder="Title"
                />
                <textarea
                  value={ability.flavourText}
                  onChange={(e) =>
                    handleTextbasedAbilityChange(
                      index,
                      "flavourText",
                      e.target.value
                    )
                  }
                  className="w-full text-lg p-2 mb-2 border-2 border-gray-300 rounded-lg"
                  placeholder="Flavour Text"
                />
                <select
                  value={ability.availability}
                  onChange={(e) =>
                    handleTextbasedAbilityChange(
                      index,
                      "availability",
                      e.target.value
                    )
                  }
                  className="w-full text-lg p-2 mb-2 border-2 border-gray-300 rounded-lg"
                >
                  <option value="">Select Availability</option>
                  {["Passive", "Free", "Action", "Bonus", "Reaction"].map(
                    (option) => (
                      <option key={option} value={option}>
                        {option}
                      </option>
                    )
                  )}
                </select>
                <select
                  value={ability.ruleType}
                  onChange={(e) =>
                    handleTextbasedAbilityChange(
                      index,
                      "ruleType",
                      e.target.value
                    )
                  }
                  className="w-full text-lg p-2 mb-2 border-2 border-gray-300 rounded-lg"
                >
                  <option value="">Select Rule Type</option>
                  {["Core Rule", "Variant", "Optional"].map((option) => (
                    <option key={option} value={option}>
                      {option}
                    </option>
                  ))}
                </select>
                <button
                  type="button"
                  onClick={() => removeTextbasedAbility(index)}
                  className="mt-2 px-4 py-2 bg-red-600 text-white rounded-lg hover:bg-red-700"
                >
                  Remove
                </button>
              </div>
            ))}
            <button
              type="button"
              onClick={addTextbasedAbility}
              className="mt-2 px-4 py-2 bg-green-600 text-white rounded-lg hover:bg-green-700"
            >
              + Add Textbased Ability
            </button>
          </div>
          <div className="mb-6">
            <label className="block text-xl font-semibold text-gray-700 mb-2">
              Charged Based Abilities
            </label>
            {chargedAbilities.map((ability, index) => (
              <div
                key={index}
                className="mb-4 p-4 border border-gray-300 rounded-lg relative"
              >
                <input
                  type="text"
                  value={ability.title}
                  onChange={(e) =>
                    handleChargedAbilityChange(index, "title", e.target.value)
                  }
                  className="w-full text-lg p-2 mb-2 border-2 border-gray-300 rounded-lg"
                  placeholder="Title"
                />
                <textarea
                  value={ability.flavorText}
                  onChange={(e) =>
                    handleChargedAbilityChange(
                      index,
                      "flavorText",
                      e.target.value
                    )
                  }
                  className="w-full text-lg p-2 mb-2 border-2 border-gray-300 rounded-lg"
                  placeholder="Flavor Text"
                />
                <input
                  type="number"
                  value={ability.totalCharge}
                  onChange={(e) =>
                    handleChargedAbilityChange(
                      index,
                      "totalCharge",
                      e.target.value
                    )
                  }
                  className="w-full text-lg p-2 mb-2 border-2 border-gray-300 rounded-lg"
                  placeholder="Total Charge"
                />
                <input
                  type="number"
                  value={ability.numberOfChargesRemaining}
                  onChange={(e) =>
                    handleChargedAbilityChange(
                      index,
                      "numberOfChargesRemaining",
                      e.target.value
                    )
                  }
                  className="w-full text-lg p-2 mb-2 border-2 border-gray-300 rounded-lg"
                  placeholder="Number of Charges Remaining"
                />
                <select
                  value={ability.availability}
                  onChange={(e) =>
                    handleChargedAbilityChange(
                      index,
                      "availability",
                      e.target.value
                    )
                  }
                  className="w-full text-lg p-2 mb-2 border-2 border-gray-300 rounded-lg"
                >
                  <option value="">Select Availability</option>
                  {["Passive", "Free", "Action", "Bonus", "Reaction"].map(
                    (option) => (
                      <option key={option} value={option}>
                        {option}
                      </option>
                    )
                  )}
                </select>
                <select
                  value={ability.ruleType}
                  onChange={(e) =>
                    handleChargedAbilityChange(
                      index,
                      "ruleType",
                      e.target.value
                    )
                  }
                  className="w-full text-lg p-2 mb-2 border-2 border-gray-300 rounded-lg"
                >
                  <option value="">Select Rule Type</option>
                  {["Core Rule", "Variant", "Optional"].map((option) => (
                    <option key={option} value={option}>
                      {option}
                    </option>
                  ))}
                </select>
                <select
                  value={ability.recovery}
                  onChange={(e) =>
                    handleChargedAbilityChange(
                      index,
                      "recovery",
                      e.target.value
                    )
                  }
                  className="w-full text-lg p-2 mb-2 border-2 border-gray-300 rounded-lg"
                >
                  <option value="">Select Recovery</option>
                  <option value="Short Rest">Short Rest</option>
                  <option value="Long Rest">Long Rest</option>
                </select>
                <button
                  type="button"
                  onClick={() => removeChargedAbility(index)}
                  className="mt-2 px-4 py-2 bg-red-600 text-white rounded-lg hover:bg-red-700"
                >
                  Remove
                </button>
              </div>
            ))}
            <button
              type="button"
              onClick={addChargedAbility}
              className="mt-2 px-4 py-2 bg-green-600 text-white rounded-lg hover:bg-green-700"
            >
              + Add Charged Based Ability
            </button>
          </div>
          <div className="mb-6">
            <label className="block text-xl font-semibold text-gray-700 mb-2">
              Modifier Abilities
            </label>
            {modifierAbilities.map((ability, index) => (
              <div
                key={index}
                className="mb-4 p-4 border border-gray-300 rounded-lg relative"
              >
                <input
                  type="text"
                  value={ability.title}
                  onChange={(e) =>
                    handleModifierAbilityChange(index, "title", e.target.value)
                  }
                  className="w-full text-lg p-2 mb-2 border-2 border-gray-300 rounded-lg"
                  placeholder="Title"
                />
                <textarea
                  value={ability.flavorText}
                  onChange={(e) =>
                    handleModifierAbilityChange(
                      index,
                      "flavorText",
                      e.target.value
                    )
                  }
                  className="w-full text-lg p-2 mb-2 border-2 border-gray-300 rounded-lg"
                  placeholder="Flavor Text"
                />
                <select
                  value={ability.availability}
                  onChange={(e) =>
                    handleModifierAbilityChange(
                      index,
                      "availability",
                      e.target.value
                    )
                  }
                  className="w-full text-lg p-2 mb-2 border-2 border-gray-300 rounded-lg"
                >
                  <option value="">Select Availability</option>
                  {["Passive", "Free", "Action", "Bonus", "Reaction"].map(
                    (option) => (
                      <option key={option} value={option}>
                        {option}
                      </option>
                    )
                  )}
                </select>
                <select
                  value={ability.ruleType}
                  onChange={(e) =>
                    handleModifierAbilityChange(
                      index,
                      "ruleType",
                      e.target.value
                    )
                  }
                  className="w-full text-lg p-2 mb-2 border-2 border-gray-300 rounded-lg"
                >
                  <option value="">Select Rule Type</option>
                  {["Core Rule", "Variant", "Optional"].map((option) => (
                    <option key={option} value={option}>
                      {option}
                    </option>
                  ))}
                </select>
                <select
                  value={ability.modifierStatFamily}
                  onChange={(e) => {
                    handleModifierAbilityChange(
                      index,
                      "modifierStatFamily",
                      e.target.value
                    );
                    handleModifierAbilityChange(index, "modifierStat", "");
                  }}
                  className="w-full text-lg p-2 mb-2 border-2 border-gray-300 rounded-lg"
                >
                  <option value="">Select Modifier Stat Family</option>
                  {Object.keys(modifierStatFamilies).map((family) => (
                    <option key={family} value={family}>
                      {family}
                    </option>
                  ))}
                </select>
                <select
                  value={ability.modifierStat}
                  onChange={(e) =>
                    handleModifierAbilityChange(
                      index,
                      "modifierStat",
                      e.target.value
                    )
                  }
                  className="w-full text-lg p-2 mb-2 border-2 border-gray-300 rounded-lg"
                  disabled={!ability.modifierStatFamily}
                >
                  <option value="">Select Modifier Stat</option>
                  {ability.modifierStatFamily &&
                    modifierStatFamilies[ability.modifierStatFamily].map(
                      (stat) => (
                        <option key={stat} value={stat}>
                          {stat}
                        </option>
                      )
                    )}
                </select>
                <input
                  type="text"
                  value={ability.modifierValue}
                  onChange={(e) =>
                    handleModifierAbilityChange(
                      index,
                      "modifierValue",
                      e.target.value
                    )
                  }
                  className="w-full text-lg p-2 mb-2 border-2 border-gray-300 rounded-lg"
                  placeholder="Modifier Value"
                />
                <button
                  type="button"
                  onClick={() => removeModifierAbility(index)}
                  className="mt-2 px-4 py-2 bg-red-600 text-white rounded-lg hover:bg-red-700"
                >
                  Remove
                </button>
              </div>
            ))}
            <button
              type="button"
              onClick={addModifierAbility}
              className="mt-2 px-4 py-2 bg-green-600 text-white rounded-lg hover:bg-green-700"
            >
              + Add Modifier Ability
            </button>
          </div>
          <button
            type="submit"
            className="mt-6 px-6 py-3 bg-blue-600 text-white rounded-lg hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-opacity-50"
          >
            Submit Level Up
          </button>
        </form>
      </div>
    </div>
  );
};

export default CreateLevelUp;
