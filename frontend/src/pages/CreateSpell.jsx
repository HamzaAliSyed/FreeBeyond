import { useState, useEffect } from "react";
import { useNavigate } from "react-router-dom";
import CreateSpell from "../assets/createspellpage.jpg";

const CreateSpells = () => {
  const [spellName, setSpellName] = useState("");
  const [spellLevel, setSpellLevel] = useState("Cantrip");
  const [castingTime, setCastingTime] = useState("Action");
  const [customCastingTime, setCustomCastingTime] = useState("");
  const [duration, setDuration] = useState("Instantaneous");
  const [customDuration, setCustomDuration] = useState("");
  const [school, setSchool] = useState("Abjuration");
  const [concentration, setConcentration] = useState(false);
  const [range, setRange] = useState("");
  const [components, setComponents] = useState([]);
  const [flavorText, setFlavorText] = useState("");
  const [classes, setClasses] = useState([]);
  const [subclasses, setSubclasses] = useState([]);
  const [sources, setSources] = useState([]);
  const [selectedSource, setSelectedSource] = useState("");
  const [isRangeBasedAOE, setIsRangeBasedAOE] = useState(false);
  const [aoeShape, setAoeShape] = useState("");
  const [aoeRadius, setAoeRadius] = useState("");
  const [saveAttribute, setSaveAttribute] = useState("");
  const [saveEffect, setSaveEffect] = useState("");
  const [customSaveEffect, setCustomSaveEffect] = useState("");
  const [damages, setDamages] = useState([{ type: "", number: "", die: "" }]);
  const [availableClasses, setAvailableClasses] = useState([]);
  const [availableSubclasses, setAvailableSubclasses] = useState([]);
  const navigate = useNavigate();

  const spellLevels = [
    "Cantrip",
    "1st",
    "2nd",
    "3rd",
    "4th",
    "5th",
    "6th",
    "7th",
    "8th",
    "9th",
  ];

  const castingTimeOptions = [
    "Action",
    "Bonus Action",
    "Reaction",
    "1 Minute",
    "10 Minutes",
    "1 Hour",
    "4 Hours",
    "8 Hours",
    "12 Hours",
    "1 Day",
    "Custom",
  ];

  const durationOptions = [
    "Instantaneous",
    "Until Dispelled",
    "1 Minute",
    "10 Minutes",
    "1 Hour",
    "Custom",
  ];

  const magicSchools = [
    "Abjuration",
    "Conjuration",
    "Divination",
    "Enchantment",
    "Evocation",
    "Illusion",
    "Necromancy",
    "Transmutation",
  ];

  const damageTypes = [
    "Acid",
    "Bludgeoning",
    "Cold",
    "Fire",
    "Force",
    "Lightning",
    "Necrotic",
    "Piercing",
    "Poison",
    "Psychic",
    "Radiant",
    "Slashing",
    "Thunder",
  ];

  const HandleSpellName = (e) => {
    setSpellName(e.target.value);
  };

  const HandleSpellLevel = (e) => {
    setSpellLevel(e.target.value);
  };

  const HandleCastingTime = (e) => {
    setCastingTime(e.target.value);
  };

  const HandleCustomCastingTime = (e) => {
    setCustomCastingTime(e.target.value);
  };

  const HandleDuration = (e) => {
    setDuration(e.target.value);
  };

  const HandleCustomDuration = (e) => {
    setCustomDuration(e.target.value);
  };

  const HandleSchool = (e) => {
    setSchool(e.target.value);
  };

  const HandleConcentration = (e) => {
    setConcentration(e.target.checked);
  };

  const HandleRange = (e) => {
    setRange(e.target.value);
  };

  const HandleComponents = (e) => {
    const component = e.target.value;
    if (e.target.checked) {
      setComponents([...components, component]);
    } else {
      setComponents(components.filter((c) => c !== component));
    }
  };

  const HandleFlavorText = (e) => {
    setFlavorText(e.target.value);
  };

  const HandleClassChange = (className) => {
    setClasses((prev) =>
      prev.includes(className)
        ? prev.filter((c) => c !== className)
        : [...prev, className]
    );
  };

  const HandleSubclassChange = (subclassName) => {
    setSubclasses((prev) =>
      prev.includes(subclassName)
        ? prev.filter((s) => s !== subclassName)
        : [...prev, subclassName]
    );
  };

  const HandleSourceChange = (e) => {
    setSelectedSource(e.target.value);
  };

  useEffect(() => {
    const fetchSources = async () => {
      try {
        const response = await fetch(
          "http://localhost:2712/api/components/getallsourcesnames"
        );
        if (!response.ok) {
          throw new Error("Network response was not ok");
        }
        const data = await response.json();
        setSources(data);
      } catch (error) {
        console.error("Error fetching sources:", error);
      }
    };

    const fetchClassesAndSubclasses = async () => {
      const classesResponse = await fetch(
        "http://localhost:2712/api/components/getallclasses"
      );
      const subclassesResponse = await fetch(
        "http://localhost:2712/api/components/getallsubclasses"
      );

      if (classesResponse.ok && subclassesResponse.ok) {
        setAvailableClasses(await classesResponse.json());
        setAvailableSubclasses(await subclassesResponse.json());
      }
    };

    fetchSources();
    fetchClassesAndSubclasses();
  }, []);

  const HandleRangeBasedAOE = (e) => {
    setIsRangeBasedAOE(e.target.checked);
  };

  const HandleAoeShape = (e) => {
    setAoeShape(e.target.value);
  };

  const HandleAoeRadius = (e) => {
    setAoeRadius(e.target.value);
  };

  const HandleSaveAttribute = (e) => {
    setSaveAttribute(e.target.value);
  };

  const HandleSaveEffect = (e) => {
    setSaveEffect(e.target.value);
  };

  const HandleCustomSaveEffect = (e) => {
    setCustomSaveEffect(e.target.value);
  };

  const HandleDamageChange = (index, field, value) => {
    const newDamages = [...damages];
    newDamages[index][field] = value;
    setDamages(newDamages);
  };

  const AddDamage = () => {
    setDamages([...damages, { type: "", number: "", die: "" }]);
  };

  const handleSubmit = async (e) => {
    e.preventDefault();

    const spellData = {
      name: spellName,
      level: parseInt(spellLevel.replace(/[^\d]/g, "")) || 0,
      castingtime: castingTime === "Custom" ? customCastingTime : castingTime,
      duration: duration === "Custom" ? customDuration : duration,
      school: school,
      concentration: concentration,
      range: range,
      components: components,
      flavourtext: flavorText,
      classes: classes,
      subclasses: subclasses,
      source: selectedSource,
      type: isRangeBasedAOE ? "AttackBasedRangeAOEAttack" : "Basic",
      aoeshape: aoeShape,
      aoeradius: parseInt(aoeRadius) || 0,
      saveattribute: saveAttribute,
      damage: damages.reduce((acc, damage) => {
        if (damage.type && damage.number && damage.die) {
          acc[damage.type] = `${damage.number}${damage.die}`;
        }
        return acc;
      }, {}),
      saveeffect: saveEffect === "Custom" ? customSaveEffect : saveEffect,
    };

    try {
      const response = await fetch(
        "http://localhost:2712/api/components/createspells",
        {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify(spellData),
        }
      );

      if (response.ok) {
        alert("Spell created successfully");
        navigate("/");
      } else {
        console.error("Failed to create spell");
        alert("Failed to create spell");
      }
    } catch (error) {
      console.error("Error:", error);
      // Handle network errors
    }
  };

  return (
    <div className="create-source-container grid grid-cols-12 h-screen">
      <div className="col-span-4 h-full">
        <img
          className="w-full object-cover h-full"
          src={CreateSpell}
          alt="Charging a random spell"
        />
      </div>
      <div className="col-span-8 p-8 flex items-center justify-center">
        <form className="w-full max-w-2xl text-center" onSubmit={handleSubmit}>
          <div className="mb-8">
            <label
              htmlFor="spellname"
              className="block text-3xl font-semibold text-gray-800 mb-4"
            >
              Spell Name
            </label>
            <input
              type="text"
              value={spellName}
              onChange={HandleSpellName}
              placeholder="Spell Name"
              className="w-full p-2 mb-4 border rounded"
            />
          </div>
          <div className="mb-8">
            <label
              htmlFor="spelllevel"
              className="block text-3xl font-semibold text-gray-800 mb-4"
            >
              Spell Level
            </label>
            <select
              id="spelllevel"
              value={spellLevel}
              onChange={HandleSpellLevel}
              className="w-full p-2 mb-4 border rounded"
            >
              {spellLevels.map((level) => (
                <option key={level} value={level}>
                  {level}
                </option>
              ))}
            </select>
          </div>
          <div className="mb-8">
            <label
              htmlFor="castingTime"
              className="block text-3xl font-semibold text-gray-800 mb-4"
            >
              Casting Time
            </label>
            <select
              id="castingTime"
              value={castingTime}
              onChange={HandleCastingTime}
              className="w-full p-2 mb-4 border rounded"
            >
              {castingTimeOptions.map((option) => (
                <option key={option} value={option}>
                  {option}
                </option>
              ))}
            </select>
            {castingTime === "Custom" && (
              <input
                type="text"
                value={customCastingTime}
                onChange={HandleCustomCastingTime}
                placeholder="Enter custom casting time"
                className="w-full p-2 mb-4 border rounded"
              />
            )}
          </div>
          <div className="mb-8">
            <label
              htmlFor="duration"
              className="block text-3xl font-semibold text-gray-800 mb-4"
            >
              Duration
            </label>
            <select
              id="duration"
              value={duration}
              onChange={HandleDuration}
              className="w-full p-2 mb-4 border rounded"
            >
              {durationOptions.map((option) => (
                <option key={option} value={option}>
                  {option}
                </option>
              ))}
            </select>
            {duration === "Custom" && (
              <input
                type="text"
                value={customDuration}
                onChange={HandleCustomDuration}
                placeholder="Enter custom duration"
                className="w-full p-2 mb-4 border rounded"
              />
            )}
          </div>
          <div className="mb-8">
            <label
              htmlFor="school"
              className="block text-3xl font-semibold text-gray-800 mb-4"
            >
              School of Magic
            </label>
            <select
              id="school"
              value={school}
              onChange={HandleSchool}
              className="w-full p-2 mb-4 border rounded"
            >
              {magicSchools.map((magicSchool) => (
                <option key={magicSchool} value={magicSchool}>
                  {magicSchool}
                </option>
              ))}
            </select>
          </div>
          <div className="mb-8 flex items-center">
            <input
              type="checkbox"
              id="concentration"
              checked={concentration}
              onChange={HandleConcentration}
              className="w-5 h-5 mr-3"
            />
            <label
              htmlFor="concentration"
              className="text-2xl font-semibold text-gray-800"
            >
              Concentration
            </label>
          </div>
          <div className="mb-8">
            <label
              htmlFor="range"
              className="block text-3xl font-semibold text-gray-800 mb-4"
            >
              Range
            </label>
            <input
              type="text"
              id="range"
              value={range}
              onChange={HandleRange}
              placeholder="Enter spell range"
              className="w-full p-2 mb-4 border rounded"
            />
          </div>
          <div className="mb-8">
            <label className="block text-3xl font-semibold text-gray-800 mb-4">
              Components
            </label>
            <div className="flex space-x-4">
              {["Verbal", "Somatic", "Material"].map((component) => (
                <div key={component} className="flex items-center">
                  <input
                    type="checkbox"
                    id={component.toLowerCase()}
                    value={component}
                    checked={components.includes(component)}
                    onChange={HandleComponents}
                    className="w-5 h-5 mr-2"
                  />
                  <label htmlFor={component.toLowerCase()} className="text-xl">
                    {component}
                  </label>
                </div>
              ))}
            </div>
          </div>
          <div className="mb-8">
            <label
              htmlFor="flavorText"
              className="block text-3xl font-semibold text-gray-800 mb-4"
            >
              Flavor Text
            </label>
            <textarea
              id="flavorText"
              value={flavorText}
              onChange={HandleFlavorText}
              placeholder="Enter spell flavor text"
              className="w-full p-2 mb-4 border rounded h-32 resize-y"
            />
          </div>
          <div className="mb-8">
            <label className="block text-3xl font-semibold text-gray-800 mb-4">
              Classes
            </label>
            <div className="grid grid-cols-3 gap-2">
              {availableClasses.map((className) => (
                <div key={className} className="flex items-center">
                  <input
                    type="checkbox"
                    id={`class-${className}`}
                    checked={classes.includes(className)}
                    onChange={() => HandleClassChange(className)}
                    className="w-5 h-5 mr-2"
                  />
                  <label htmlFor={`class-${className}`}>{className}</label>
                </div>
              ))}
            </div>
          </div>

          <div className="mb-8">
            <label className="block text-3xl font-semibold text-gray-800 mb-4">
              Subclasses
            </label>
            <div className="grid grid-cols-3 gap-2">
              {availableSubclasses.map((subclassName) => (
                <div key={subclassName} className="flex items-center">
                  <input
                    type="checkbox"
                    id={`subclass-${subclassName}`}
                    checked={subclasses.includes(subclassName)}
                    onChange={() => HandleSubclassChange(subclassName)}
                    className="w-5 h-5 mr-2"
                  />
                  <label htmlFor={`subclass-${subclassName}`}>
                    {subclassName}
                  </label>
                </div>
              ))}
            </div>
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
              onChange={HandleSourceChange}
              className="w-full p-2 mb-4 border rounded"
            >
              <option value="">Select a source</option>
              {sources.map((source) => (
                <option key={source} value={source}>
                  {source}
                </option>
              ))}
            </select>
          </div>
          <div className="mb-8">
            <div className="flex items-center">
              <input
                type="checkbox"
                id="rangeBasedAOE"
                checked={isRangeBasedAOE}
                onChange={HandleRangeBasedAOE}
                className="w-5 h-5 mr-3"
              />
              <label
                htmlFor="rangeBasedAOE"
                className="text-2xl font-semibold text-gray-800"
              >
                Range Based AOE
              </label>
            </div>
          </div>

          {isRangeBasedAOE && (
            <>
              <div className="mb-8">
                <label
                  htmlFor="aoeShape"
                  className="block text-2xl font-semibold text-gray-800 mb-2"
                >
                  AOE Shape
                </label>
                <select
                  id="aoeShape"
                  value={aoeShape}
                  onChange={HandleAoeShape}
                  className="w-full p-2 mb-4 border rounded"
                >
                  <option value="">Select AOE Shape</option>
                  {["Square", "Cone", "Line", "Sphere"].map((shape) => (
                    <option key={shape} value={shape}>
                      {shape}
                    </option>
                  ))}
                </select>
              </div>

              <div className="mb-8">
                <label
                  htmlFor="aoeRadius"
                  className="block text-2xl font-semibold text-gray-800 mb-2"
                >
                  AOE Radius
                </label>
                <input
                  type="text"
                  id="aoeRadius"
                  value={aoeRadius}
                  onChange={HandleAoeRadius}
                  className="w-full p-2 mb-4 border rounded"
                  placeholder="Enter AOE radius"
                />
              </div>

              <div className="mb-8">
                <label
                  htmlFor="saveAttribute"
                  className="block text-2xl font-semibold text-gray-800 mb-2"
                >
                  Save Attribute
                </label>
                <select
                  id="saveAttribute"
                  value={saveAttribute}
                  onChange={HandleSaveAttribute}
                  className="w-full p-2 mb-4 border rounded"
                >
                  <option value="">Select Save Attribute</option>
                  {[
                    "Strength",
                    "Dexterity",
                    "Constitution",
                    "Intelligence",
                    "Wisdom",
                    "Charisma",
                  ].map((attr) => (
                    <option key={attr} value={attr}>
                      {attr}
                    </option>
                  ))}
                </select>
              </div>

              <div className="mb-8">
                <label className="block text-2xl font-semibold text-gray-800 mb-2">
                  Damage
                </label>
                {damages.map((damage, index) => (
                  <div key={index} className="flex space-x-4 mb-4">
                    <select
                      value={damage.type}
                      onChange={(e) =>
                        HandleDamageChange(index, "type", e.target.value)
                      }
                      className="flex-1 p-2 border rounded"
                    >
                      <option value="">Select Damage Type</option>
                      {damageTypes.map((type) => (
                        <option key={type} value={type}>
                          {type}
                        </option>
                      ))}
                    </select>
                    <input
                      type="number"
                      value={damage.number}
                      onChange={(e) =>
                        HandleDamageChange(index, "number", e.target.value)
                      }
                      className="flex-1 p-2 border rounded"
                      placeholder="Number"
                    />
                    <select
                      value={damage.die}
                      onChange={(e) =>
                        HandleDamageChange(index, "die", e.target.value)
                      }
                      className="flex-1 p-2 border rounded"
                    >
                      <option value="">Select Die</option>
                      {["d4", "d6", "d8", "d10", "d12", "d20"].map((die) => (
                        <option key={die} value={die}>
                          {die}
                        </option>
                      ))}
                    </select>
                  </div>
                ))}
                <button
                  type="button"
                  onClick={AddDamage}
                  className="mt-2 px-4 py-2 bg-blue-500 text-white rounded hover:bg-blue-600"
                >
                  Add Damage
                </button>
              </div>

              <div className="mb-8">
                <label
                  htmlFor="saveEffect"
                  className="block text-2xl font-semibold text-gray-800 mb-2"
                >
                  Save Effect
                </label>
                <select
                  id="saveEffect"
                  value={saveEffect}
                  onChange={HandleSaveEffect}
                  className="w-full p-2 mb-4 border rounded"
                >
                  <option value="">Select Save Effect</option>
                  <option value="Halved">Halved</option>
                  <option value="No Damage">No Damage</option>
                  <option value="Custom">Custom</option>
                </select>
                {saveEffect === "Custom" && (
                  <input
                    type="text"
                    value={customSaveEffect}
                    onChange={HandleCustomSaveEffect}
                    className="w-full p-2 mb-4 border rounded"
                    placeholder="Enter custom save effect"
                  />
                )}
              </div>
            </>
          )}
          <button
            type="submit"
            className="bg-blue-500 text-white px-4 py-2 rounded"
          >
            Create Spell
          </button>
        </form>
      </div>
    </div>
  );
};

export default CreateSpells;
