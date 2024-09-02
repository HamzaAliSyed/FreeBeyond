import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import CreateRace from "../assets/createracepage.jpg";

const Race = () => {
  const navigate = useNavigate();

  const [raceName, setRaceName] = useState("");

  const [size, setSize] = useState("");

  const [creatureType, setCreatureType] = useState("");

  const [speed, setSpeed] = useState({
    LandSpeed: 0,
    SwimSpeed: 0,
    FlySpeed: 0,
    ClimbSpeed: 0,
    BurrowSpeed: 0,
  });

  const [abilityScores, setAbilityScores] = useState({
    Strength: 0,
    Dexterity: 0,
    Constitution: 0,
    Intelligence: 0,
    Wisdom: 0,
    Charisma: 0,
  });

  const [languages, setLanguages] = useState([""]);

  const [ageRange, setAgeRange] = useState({ min: "", max: "" });

  const [flavourText, setFlavourText] = useState([
    { heading: "", description: "" },
  ]);

  const [otherBoosts, setOtherBoosts] = useState([{ key: "", value: "" }]);

  const [spells, setSpells] = useState([{ name: "", level: "" }]);

  const [attacks, setAttacks] = useState([
    {
      name: "",
      attribute: "",
      damageType: "",
      damageDice: "",
      dieSize: "",
      range: "",
    },
  ]);

  const abilityNames = [
    "Strength",
    "Dexterity",
    "Constitution",
    "Intelligence",
    "Wisdom",
    "Charisma",
  ];
  const damageTypes = [
    "Slashing",
    "Piercing",
    "Bludgeoning",
    "Fire",
    "Cold",
    "Lightning",
    "Acid",
    "Poison",
    "Necrotic",
    "Radiant",
    "Force",
    "Psychic",
    "Thunder",
  ];
  const dieSizes = ["d4", "d6", "d8", "d10", "d12", "d20"];

  const [sources, setSources] = useState([]);
  const [selectedSource, setSelectedSource] = useState("");

  const [image, setImage] = useState({ file: undefined, preview: undefined });

  const handleRaceName = (event) => {
    setRaceName(event.target.value);
  };

  const handleSizeChange = (event) => {
    setSize(event.target.value);
  };

  const handleCreatureTypeChange = (event) => {
    setCreatureType(event.target.value);
  };

  const handleSpeedChange = (event) => {
    const { name, value } = event.target;
    setSpeed((prevSpeed) => ({
      ...prevSpeed,
      [name]: parseInt(value) || 0,
    }));
  };

  const handleAbilityScoreChange = (event) => {
    const { name, value } = event.target;
    setAbilityScores((prevScores) => ({
      ...prevScores,
      [name]: parseInt(value) || 0,
    }));
  };

  const addLanguageField = () => {
    setLanguages([...languages, ""]);
  };

  const removeLanguageField = (index) => {
    const newLanguages = languages.filter((_, i) => i !== index);
    setLanguages(newLanguages);
  };

  const handleLanguageChange = (index, value) => {
    const newLanguages = [...languages];
    newLanguages[index] = value;
    setLanguages(newLanguages);
  };

  const handleAgeRangeChange = (event) => {
    const { name, value } = event.target;
    setAgeRange((prevRange) => ({
      ...prevRange,
      [name]: value,
    }));
  };

  const addFlavourText = () => {
    setFlavourText([...flavourText, { heading: "", description: "" }]);
  };

  const removeFlavourText = (index) => {
    const newFlavourText = flavourText.filter((_, i) => i !== index);
    setFlavourText(newFlavourText);
  };

  const handleFlavourTextChange = (index, field, value) => {
    const newFlavourText = [...flavourText];
    newFlavourText[index][field === "heading" ? "heading" : "description"] =
      value;
    setFlavourText(newFlavourText);
  };

  const addOtherBoost = () => {
    setOtherBoosts([...otherBoosts, { key: "", value: "" }]);
  };

  const removeOtherBoost = (index) => {
    const newOtherBoosts = otherBoosts.filter((_, i) => i !== index);
    setOtherBoosts(newOtherBoosts);
  };

  const handleOtherBoostChange = (index, field, value) => {
    const newOtherBoosts = [...otherBoosts];
    newOtherBoosts[index][field] = value;
    setOtherBoosts(newOtherBoosts);
  };

  const addSpell = () => {
    setSpells([...spells, { name: "", level: "" }]);
  };

  const removeSpell = (index) => {
    const newSpells = spells.filter((_, i) => i !== index);
    setSpells(newSpells);
  };

  const handleSpellChange = (index, field, value) => {
    const newSpells = [...spells];
    newSpells[index][field] = value;
    setSpells(newSpells);
  };

  const addAttack = () => {
    setAttacks([
      ...attacks,
      {
        name: "",
        attribute: "",
        damageType: "",
        damageDice: "",
        dieSize: "",
        range: "",
      },
    ]);
  };

  const removeAttack = (index) => {
    const newAttacks = attacks.filter((_, i) => i !== index);
    setAttacks(newAttacks);
  };

  const handleAttackChange = (index, field, value) => {
    const newAttacks = [...attacks];
    newAttacks[index][field] = field === "range" ? parseInt(value, 10) : value;
    setAttacks(newAttacks);
  };

  const handleSourceChange = (event) => {
    setSelectedSource(event.target.value);
  };

  const fetchAvailableSources = async () => {
    try {
      const response = await fetch(
        "http://localhost:2712/api/components/getallsourcesnames"
      );
      if (response.ok) {
        const sources = await response.json();
        setSources(sources);
      } else {
        throw new Error("Failed to fetch sources");
      }
    } catch (error) {
      console.error("Error fetching sources:", error);
    }
  };

  useEffect(() => {
    fetchAvailableSources();
  }, []);

  const handleImageUpload = (event) => {
    const file = event.target.files[0];
    if (file && file.type.substr(0, 6) === "image/") {
      const reader = new FileReader();
      reader.onloadend = () => {
        setImage({
          file: file,
          preview: reader.result,
        });
      };
      reader.readAsDataURL(file);
    } else {
      alert("Please select an image file");
      event.target.value = null; // Reset the input
    }
  };

  const handleSubmit = async (event) => {
    event.preventDefault();

    const formattedAttacks = attacks.map((attack) => ({
      attackname: attack.name,
      attribute: attack.attribute,
      damage: {
        [attack.damageType]: `${attack.damageDice}${attack.dieSize}`,
      },
      range: parseInt(attack.range, 10),
    }));

    const formattedFlavourText = flavourText.map((item) => ({
      title: item.heading,
      flavourtext: item.description,
    }));

    const raceData = {
      name: raceName,
      abilityscores: abilityScores,
      size: size,
      speed: speed,
      creaturetype: creatureType,
      flavourtext: formattedFlavourText,
      spells: Object.fromEntries(
        spells.map((spell) => [spell.name, parseInt(spell.level)])
      ),
      attacks: formattedAttacks,
      otherboost: Object.fromEntries(
        otherBoosts.map((boost) => [boost.key, boost.value])
      ),
      agerange: [parseInt(ageRange.min), parseInt(ageRange.max)],
      languages: languages,
      image: image.preview,
      source: selectedSource,
    };

    try {
      const response = await fetch(
        "http://localhost:2712/api/components/createrace",
        {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify(raceData),
        }
      );

      console.log("Sending data:", JSON.stringify(raceData, null, 2));

      if (!response.ok) {
        const errorText = await response.text();
        throw new Error(
          `HTTP error! status: ${response.status}, message: ${errorText}`
        );
      }

      const result = await response.json();
      console.log("Success:", result);
      alert("Race Created successfully");
      navigate("/");
    } catch (error) {
      console.error("Error:", error);
      alert(error);
    }
  };

  return (
    <div className="create-race-container grid grid-cols-12 h-screen">
      <div className="col-span-3 h-full">
        <img src={CreateRace} className="w-full object-cover h-full" />
      </div>
      <div className="col-span-9 p-8 flex items-center justify-center">
        <form className="w-full max-w-2xl text-center" onSubmit={handleSubmit}>
          <div className="mb-8">
            <label
              htmlFor="raceName"
              className="block text-3xl font-semibold text-gray-800 mb-4"
            >
              Race Name
            </label>
            <input
              type="text"
              id="raceName"
              value={raceName}
              onChange={handleRaceName}
              className="w-full text-2xl p-3 border-2 border-gray-300 rounded-lg focus:border-indigo-500 focus:ring focus:ring-indigo-200 focus:ring-opacity-50"
            />
          </div>
          <div className="mb-8">
            <label
              htmlFor="size"
              className="block text-3xl font-semibold text-gray-800 mb-4"
            >
              Size
            </label>
            <select
              id="size"
              value={size}
              onChange={handleSizeChange}
              className="w-full text-2xl p-3 border-2 border-gray-300 rounded-lg focus:border-indigo-500 focus:ring focus:ring-indigo-200 focus:ring-opacity-50"
            >
              <option value="">Select size</option>
              <option value="tiny">Tiny</option>
              <option value="small">Small</option>
              <option value="medium">Medium</option>
              <option value="huge">Huge</option>
              <option value="gargantuan">Gargantuan</option>
            </select>
          </div>
          <div className="mb-8">
            <label
              htmlFor="creatureType"
              className="block text-3xl font-semibold text-gray-800 mb-4"
            >
              Creature Type
            </label>
            <select
              id="creatureType"
              value={creatureType}
              onChange={handleCreatureTypeChange}
              className="w-full text-2xl p-3 border-2 border-gray-300 rounded-lg focus:border-indigo-500 focus:ring focus:ring-indigo-200 focus:ring-opacity-50"
            >
              <option value="">Select creature type</option>
              <option value="Organic">Organic</option>
              <option value="Construct">Construct</option>
            </select>
          </div>
          <div className="mb-8">
            <label className="block text-3xl font-semibold text-gray-800 mb-4">
              Speed
            </label>
            {Object.keys(speed).map((speedType) => (
              <div key={speedType} className="mb-2 flex items-center">
                <label
                  htmlFor={speedType}
                  className="w-1/3 text-xl text-right mr-2"
                >
                  {speedType}:
                </label>
                <input
                  type="number"
                  id={speedType}
                  name={speedType}
                  value={speed[speedType]}
                  onChange={handleSpeedChange}
                  className="w-2/3 text-xl p-2 border-2 border-gray-300 rounded-lg focus:border-indigo-500 focus:ring focus:ring-indigo-200 focus:ring-opacity-50"
                  min="0"
                />
              </div>
            ))}
          </div>
          <div className="mb-8">
            <label className="block text-3xl font-semibold text-gray-800 mb-4">
              Ability Score Bonuses
            </label>
            {Object.keys(abilityScores).map((ability) => (
              <div key={ability} className="mb-2 flex items-center">
                <label
                  htmlFor={ability}
                  className="w-1/3 text-xl text-right mr-2"
                >
                  {ability}:
                </label>
                <input
                  type="number"
                  id={ability}
                  name={ability}
                  value={abilityScores[ability]}
                  onChange={handleAbilityScoreChange}
                  className="w-2/3 text-xl p-2 border-2 border-gray-300 rounded-lg focus:border-indigo-500 focus:ring focus:ring-indigo-200 focus:ring-opacity-50"
                />
              </div>
            ))}
          </div>
          <div className="mb-8">
            <label className="block text-3xl font-semibold text-gray-800 mb-4">
              Languages
            </label>
            {languages.map((language, index) => (
              <div key={index} className="flex items-center mb-2">
                <input
                  type="text"
                  value={language}
                  onChange={(e) => handleLanguageChange(index, e.target.value)}
                  className="flex-grow text-xl p-2 border-2 border-gray-300 rounded-lg focus:border-indigo-500 focus:ring focus:ring-indigo-200 focus:ring-opacity-50 mr-2"
                />
                {index === languages.length - 1 ? (
                  <button
                    type="button"
                    onClick={addLanguageField}
                    className="bg-green-500 text-white p-2 rounded-full hover:bg-green-600 focus:outline-none focus:ring-2 focus:ring-green-500 focus:ring-opacity-50"
                  >
                    +
                  </button>
                ) : (
                  <button
                    type="button"
                    onClick={() => removeLanguageField(index)}
                    className="bg-red-500 text-white p-2 rounded-full hover:bg-red-600 focus:outline-none focus:ring-2 focus:ring-red-500 focus:ring-opacity-50"
                  >
                    x
                  </button>
                )}
              </div>
            ))}
          </div>
          <div className="mb-8">
            <label className="block text-3xl font-semibold text-gray-800 mb-4">
              Age Range
            </label>
            <div className="flex items-center space-x-4">
              <div>
                <label htmlFor="minAge" className="block text-xl mb-2">
                  Min Age:
                </label>
                <input
                  type="number"
                  id="minAge"
                  name="min"
                  value={ageRange.min}
                  onChange={handleAgeRangeChange}
                  className="w-full text-xl p-2 border-2 border-gray-300 rounded-lg focus:border-indigo-500 focus:ring focus:ring-indigo-200 focus:ring-opacity-50"
                  min="0"
                />
              </div>
              <div>
                <label htmlFor="maxAge" className="block text-xl mb-2">
                  Max Age:
                </label>
                <input
                  type="number"
                  id="maxAge"
                  name="max"
                  value={ageRange.max}
                  onChange={handleAgeRangeChange}
                  className="w-full text-xl p-2 border-2 border-gray-300 rounded-lg focus:border-indigo-500 focus:ring focus:ring-indigo-200 focus:ring-opacity-50"
                  min="0"
                />
              </div>
            </div>
          </div>
          <div className="mb-8">
            <label className="block text-3xl font-semibold text-gray-800 mb-4">
              Flavour Text
            </label>
            {flavourText.map((text, index) => (
              <div
                key={index}
                className="mb-4 p-4 border-2 border-gray-300 rounded-lg"
              >
                <div className="mb-2">
                  <input
                    type="text"
                    value={text.heading}
                    onChange={(e) =>
                      handleFlavourTextChange(index, "heading", e.target.value)
                    }
                    placeholder="Heading"
                    className="w-full text-xl p-2 border-2 border-gray-300 rounded-lg focus:border-indigo-500 focus:ring focus:ring-indigo-200 focus:ring-opacity-50"
                  />
                </div>
                <div className="mb-2">
                  <textarea
                    value={text.description}
                    onChange={(e) =>
                      handleFlavourTextChange(
                        index,
                        "description",
                        e.target.value
                      )
                    }
                    placeholder="Description"
                    className="w-full text-xl p-2 border-2 border-gray-300 rounded-lg focus:border-indigo-500 focus:ring focus:ring-indigo-200 focus:ring-opacity-50"
                    rows="3"
                  />
                </div>
                <div className="flex justify-end">
                  {index === flavourText.length - 1 ? (
                    <button
                      type="button"
                      onClick={addFlavourText}
                      className="bg-green-500 text-white p-2 rounded-full hover:bg-green-600 focus:outline-none focus:ring-2 focus:ring-green-500 focus:ring-opacity-50 mr-2"
                    >
                      +
                    </button>
                  ) : null}
                  {flavourText.length > 1 ? (
                    <button
                      type="button"
                      onClick={() => removeFlavourText(index)}
                      className="bg-red-500 text-white p-2 rounded-full hover:bg-red-600 focus:outline-none focus:ring-2 focus:ring-red-500 focus:ring-opacity-50"
                    >
                      x
                    </button>
                  ) : null}
                </div>
              </div>
            ))}
          </div>
          <div className="mb-8">
            <label className="block text-3xl font-semibold text-gray-800 mb-4">
              Other Boosts
            </label>
            {otherBoosts.map((boost, index) => (
              <div key={index} className="mb-4 flex items-center">
                <input
                  type="text"
                  value={boost.key}
                  onChange={(e) =>
                    handleOtherBoostChange(index, "key", e.target.value)
                  }
                  placeholder="Boost Name"
                  className="flex-1 text-xl p-2 border-2 border-gray-300 rounded-lg focus:border-indigo-500 focus:ring focus:ring-indigo-200 focus:ring-opacity-50 mr-2"
                />
                <input
                  type="text"
                  value={boost.value}
                  onChange={(e) =>
                    handleOtherBoostChange(index, "value", e.target.value)
                  }
                  placeholder="Boost Value"
                  className="flex-1 text-xl p-2 border-2 border-gray-300 rounded-lg focus:border-indigo-500 focus:ring focus:ring-indigo-200 focus:ring-opacity-50 mr-2"
                />
                {index === otherBoosts.length - 1 ? (
                  <button
                    type="button"
                    onClick={addOtherBoost}
                    className="bg-green-500 text-white p-2 rounded-full hover:bg-green-600 focus:outline-none focus:ring-2 focus:ring-green-500 focus:ring-opacity-50 mr-2"
                  >
                    +
                  </button>
                ) : null}
                {otherBoosts.length > 1 ? (
                  <button
                    type="button"
                    onClick={() => removeOtherBoost(index)}
                    className="bg-red-500 text-white p-2 rounded-full hover:bg-red-600 focus:outline-none focus:ring-2 focus:ring-red-500 focus:ring-opacity-50"
                  >
                    x
                  </button>
                ) : null}
              </div>
            ))}
          </div>
          <div className="mb-8">
            <label className="block text-3xl font-semibold text-gray-800 mb-4">
              Spells
            </label>
            {spells.map((spell, index) => (
              <div key={index} className="mb-4 flex items-center">
                <input
                  type="text"
                  value={spell.name}
                  onChange={(e) =>
                    handleSpellChange(index, "name", e.target.value)
                  }
                  placeholder="Spell Name"
                  className="flex-grow text-xl p-2 border-2 border-gray-300 rounded-lg focus:border-indigo-500 focus:ring focus:ring-indigo-200 focus:ring-opacity-50 mr-2"
                />
                <input
                  type="number"
                  value={spell.level}
                  onChange={(e) =>
                    handleSpellChange(index, "level", e.target.value)
                  }
                  placeholder="Level"
                  className="w-20 text-xl p-2 border-2 border-gray-300 rounded-lg focus:border-indigo-500 focus:ring focus:ring-indigo-200 focus:ring-opacity-50 mr-2"
                  min="0"
                />
                {index === spells.length - 1 ? (
                  <button
                    type="button"
                    onClick={addSpell}
                    className="bg-green-500 text-white p-2 rounded-full hover:bg-green-600 focus:outline-none focus:ring-2 focus:ring-green-500 focus:ring-opacity-50 mr-2"
                  >
                    +
                  </button>
                ) : null}
                {spells.length > 1 ? (
                  <button
                    type="button"
                    onClick={() => removeSpell(index)}
                    className="bg-red-500 text-white p-2 rounded-full hover:bg-red-600 focus:outline-none focus:ring-2 focus:ring-red-500 focus:ring-opacity-50"
                  >
                    x
                  </button>
                ) : null}
              </div>
            ))}
          </div>
          <div className="mb-8">
            <label className="block text-3xl font-semibold text-gray-800 mb-4">
              Attacks
            </label>
            {attacks.map((attack, index) => (
              <div
                key={index}
                className="mb-4 p-4 border-2 border-gray-300 rounded-lg"
              >
                <div className="grid grid-cols-2 gap-4 mb-2">
                  <input
                    type="text"
                    value={attack.name}
                    onChange={(e) =>
                      handleAttackChange(index, "name", e.target.value)
                    }
                    placeholder="Attack Name"
                    className="w-full text-xl p-2 border-2 border-gray-300 rounded-lg focus:border-indigo-500 focus:ring focus:ring-indigo-200 focus:ring-opacity-50"
                  />
                  <select
                    value={attack.attribute}
                    onChange={(e) =>
                      handleAttackChange(index, "attribute", e.target.value)
                    }
                    className="w-full text-xl p-2 border-2 border-gray-300 rounded-lg focus:border-indigo-500 focus:ring focus:ring-indigo-200 focus:ring-opacity-50"
                  >
                    <option value="">Select Attribute</option>
                    {abilityNames.map((ability) => (
                      <option key={ability} value={ability}>
                        {ability}
                      </option>
                    ))}
                  </select>
                </div>
                <div className="grid grid-cols-3 gap-4 mb-2">
                  <select
                    value={attack.damageType}
                    onChange={(e) =>
                      handleAttackChange(index, "damageType", e.target.value)
                    }
                    className="w-full text-xl p-2 border-2 border-gray-300 rounded-lg focus:border-indigo-500 focus:ring focus:ring-indigo-200 focus:ring-opacity-50"
                  >
                    <option value="">Select Damage Type</option>
                    {damageTypes.map((type) => (
                      <option key={type} value={type}>
                        {type}
                      </option>
                    ))}
                  </select>
                  <input
                    type="text"
                    value={attack.damageDice}
                    onChange={(e) =>
                      handleAttackChange(index, "damageDice", e.target.value)
                    }
                    placeholder="Damage Dice (e.g., 2)"
                    className="w-full text-xl p-2 border-2 border-gray-300 rounded-lg focus:border-indigo-500 focus:ring focus:ring-indigo-200 focus:ring-opacity-50"
                  />
                  <select
                    value={attack.dieSize}
                    onChange={(e) =>
                      handleAttackChange(index, "dieSize", e.target.value)
                    }
                    className="w-full text-xl p-2 border-2 border-gray-300 rounded-lg focus:border-indigo-500 focus:ring focus:ring-indigo-200 focus:ring-opacity-50"
                  >
                    <option value="">Select Die Size</option>
                    {dieSizes.map((size) => (
                      <option key={size} value={size}>
                        {size}
                      </option>
                    ))}
                  </select>
                </div>
                <div className="mb-2">
                  <input
                    type="number"
                    value={attack.range}
                    onChange={(e) =>
                      handleAttackChange(index, "range", e.target.value)
                    }
                    placeholder="Range"
                    className="w-full text-xl p-2 border-2 border-gray-300 rounded-lg focus:border-indigo-500 focus:ring focus:ring-indigo-200 focus:ring-opacity-50"
                    min="0"
                  />
                </div>
                <div className="flex justify-end">
                  {index === attacks.length - 1 ? (
                    <button
                      type="button"
                      onClick={addAttack}
                      className="bg-green-500 text-white p-2 rounded-full hover:bg-green-600 focus:outline-none focus:ring-2 focus:ring-green-500 focus:ring-opacity-50 mr-2"
                    >
                      +
                    </button>
                  ) : null}
                  {attacks.length > 1 ? (
                    <button
                      type="button"
                      onClick={() => removeAttack(index)}
                      className="bg-red-500 text-white p-2 rounded-full hover:bg-red-600 focus:outline-none focus:ring-2 focus:ring-red-500 focus:ring-opacity-50"
                    >
                      x
                    </button>
                  ) : null}
                </div>
              </div>
            ))}
          </div>
          <div className="mb-8">
            <label
              htmlFor="source"
              className="block text-3xl font-semibold text-gray-800 mb-4"
            >
              Source
            </label>
            <select
              id="source"
              value={selectedSource}
              onChange={handleSourceChange}
              className="w-full text-xl p-2 border-2 border-gray-300 rounded-lg focus:border-indigo-500 focus:ring focus:ring-indigo-200 focus:ring-opacity-50"
            >
              <option value="">Select a source</option>
              {sources.map((source, index) => (
                <option key={index} value={source}>
                  {source}
                </option>
              ))}
            </select>
          </div>
          <div className="mb-8">
            <label
              htmlFor="image"
              className="block text-3xl font-semibold text-gray-800 mb-4"
            >
              Race Image
            </label>
            <input
              type="file"
              id="image"
              accept="image/*"
              onChange={handleImageUpload}
              className="w-full text-xl p-2 border-2 border-gray-300 rounded-lg focus:border-indigo-500 focus:ring focus:ring-indigo-200 focus:ring-opacity-50"
            />
            {image.preview && (
              <div className="mt-4">
                <img
                  src={image.preview}
                  alt="Preview"
                  className="max-w-full h-auto"
                />
              </div>
            )}
          </div>
          <div className="mt-8">
            <button
              type="submit"
              onClick={handleSubmit}
              className="w-full text-xl font-semibold py-3 px-8 border border-transparent rounded-lg text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500 transition duration-150 ease-in-out"
            >
              Create Race
            </button>
          </div>
        </form>
      </div>
    </div>
  );
};
export default Race;
