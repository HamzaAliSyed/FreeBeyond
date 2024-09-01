import { useState } from "react";
import ClassPage from "../assets/createclasspage.jpg";
import { useNavigate } from "react-router-dom";

const CreateClass = () => {
  const navigate = useNavigate();
  const [name, setName] = useState("");
  const [hitDie, setHitDie] = useState("");
  const [armorProficiency, setArmorProficiency] = useState([]);
  const [weaponProficiency, setWeaponProficiency] = useState([]);
  const [savingThrowProficiency, setSavingThrowProficiency] = useState([]);
  const [skillsCanChoose, setSkillsCanChoose] = useState(0);
  const [skillsChoiceList, setSkillsChoiceList] = useState([]);

  const hitDieOptions = ["d4", "d6", "d8", "d10", "d12", "d20"];

  const handleNameChange = (event) => {
    setName(event.target.value);
  };

  const handleHitDieChange = (event) => {
    setHitDie(event.target.value);
  };

  const handleArmorProficiencyChange = (event) => {
    const { value, checked } = event.target;
    if (checked) {
      setArmorProficiency([...armorProficiency, value]);
    } else {
      setArmorProficiency(armorProficiency.filter((item) => item !== value));
    }
  };

  const handleWeaponProficiencyChange = (event) => {
    const { value, checked } = event.target;
    if (checked) {
      setWeaponProficiency([...weaponProficiency, value]);
    } else {
      setWeaponProficiency(weaponProficiency.filter((item) => item !== value));
    }
  };

  const handleSavingThrowProficiencyChange = (event) => {
    const { value, checked } = event.target;
    if (checked) {
      setSavingThrowProficiency([...savingThrowProficiency, value]);
    } else {
      setSavingThrowProficiency(
        savingThrowProficiency.filter((item) => item !== value)
      );
    }
  };

  const handleSkillsCanChooseChange = (event) => {
    setSkillsCanChoose(parseInt(event.target.value));
  };

  const handleSkillsChoiceListChange = (event) => {
    const { value, checked } = event.target;
    if (checked) {
      setSkillsChoiceList([...skillsChoiceList, value]);
    } else {
      setSkillsChoiceList(skillsChoiceList.filter((skill) => skill !== value));
    }
  };

  const handleSubmit = async (event) => {
    event.preventDefault();

    const classData = {
      name,
      hitDie,
      armorProficiency,
      weaponProficiency,
      savingThrowProficiency,
      skillsCanChoose,
      skillsChoiceList,
    };

    try {
      const response = await fetch(
        "http://localhost:2712/api/components/createclass",
        {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify(classData),
        }
      );

      if (response.ok) {
        alert("Class created successfully!");
        navigate("/"); // Redirect to home page or class list
      } else {
        throw new Error("Failed to create class");
      }
    } catch (error) {
      console.error("Error creating class:", error);
      alert("Failed to create class");
    }
  };

  return (
    <div className="create-class-container grid grid-cols-12 h-screen">
      <div className="col-span-3 h-full">
        <img
          src={ClassPage}
          className="w-full object-cover h-full"
          alt="An adventurer"
        />
      </div>
      <div className="col-span-9 p-8 flex items-center justify-center">
        <form onSubmit={handleSubmit} className="w-full max-w-2xl text-center">
          <div className="mb-8">
            <label
              htmlFor="name"
              className="block text-3xl font-semibold text-gray-800 mb-4"
            >
              Class Name
            </label>
            <input
              type="text"
              id="name"
              value={name}
              onChange={handleNameChange}
              className="w-full text-2xl p-3 border-2 border-gray-300 rounded-lg focus:border-indigo-500 focus:ring focus:ring-indigo-200 focus:ring-opacity-50"
            />
          </div>
          <div className="mb-8">
            <label
              htmlFor="hitDie"
              className="block text-3xl font-semibold text-gray-800 mb-4"
            >
              Hit Die
            </label>
            <select
              id="hitDie"
              value={hitDie}
              onChange={handleHitDieChange}
              className="w-full text-2xl p-3 border-2 border-gray-300 rounded-lg focus:border-indigo-500 focus:ring focus:ring-indigo-200 focus:ring-opacity-50"
            >
              <option value="">Select Hit Die</option>
              {hitDieOptions.map((die) => (
                <option key={die} value={die}>
                  {die}
                </option>
              ))}
            </select>
          </div>
          <div className="mb-8">
            <label className="block text-3xl font-semibold text-gray-800 mb-4">
              Armor Proficiency
            </label>
            <div className="flex flex-wrap gap-4">
              {["Light", "Medium", "Heavy", "Shields"].map((armor) => (
                <label key={armor} className="inline-flex items-center">
                  <input
                    type="checkbox"
                    value={armor}
                    checked={armorProficiency.includes(armor)}
                    onChange={handleArmorProficiencyChange}
                    className="form-checkbox h-5 w-5 text-indigo-600"
                  />
                  <span className="ml-2 text-xl">{armor}</span>
                </label>
              ))}
            </div>
          </div>
          <div className="mb-8">
            <label className="block text-3xl font-semibold text-gray-800 mb-4">
              Weapon Proficiency
            </label>
            <div className="flex flex-wrap gap-4">
              {["Simple", "Martial", "Firearm"].map((weapon) => (
                <label key={weapon} className="inline-flex items-center">
                  <input
                    type="checkbox"
                    value={weapon}
                    checked={weaponProficiency.includes(weapon)}
                    onChange={handleWeaponProficiencyChange}
                    className="form-checkbox h-5 w-5 text-indigo-600"
                  />
                  <span className="ml-2 text-xl">{weapon}</span>
                </label>
              ))}
            </div>
          </div>
          <div className="mb-8">
            <label className="block text-3xl font-semibold text-gray-800 mb-4">
              Saving Throw Proficiency
            </label>
            <div className="flex flex-wrap gap-3">
              {[
                "Strength",
                "Dexterity",
                "Constitution",
                "Intelligence",
                "Wisdom",
                "Charisma",
              ].map((ability) => (
                <label key={ability} className="inline-flex items-center">
                  <input
                    type="checkbox"
                    value={ability}
                    checked={savingThrowProficiency.includes(ability)}
                    onChange={handleSavingThrowProficiencyChange}
                    className="form-checkbox h-5 w-5 text-indigo-600"
                  />
                  <span className="ml-1 text-xl">{ability}</span>
                </label>
              ))}
            </div>
          </div>
          <div className="mb-8">
            <label className="block text-3xl font-semibold text-gray-800 mb-4">
              Skill Proficiencies
            </label>
            <div className="mb-4">
              <label htmlFor="skillsCanChoose" className="block text-xl mb-2">
                Number of skills to choose:
              </label>
              <input
                type="number"
                id="skillsCanChoose"
                value={skillsCanChoose}
                onChange={handleSkillsCanChooseChange}
                min="0"
                className="w-full text-2xl p-3 border-2 border-gray-300 rounded-lg focus:border-indigo-500 focus:ring focus:ring-indigo-200 focus:ring-opacity-50"
              />
            </div>
            <div className="flex flex-wrap gap-4">
              {[
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
              ].map((skill) => (
                <label key={skill} className="inline-flex items-center">
                  <input
                    type="checkbox"
                    value={skill}
                    checked={skillsChoiceList.includes(skill)}
                    onChange={handleSkillsChoiceListChange}
                    className="form-checkbox h-5 w-5 text-indigo-600"
                  />
                  <span className="ml-2 text-xl">{skill}</span>
                </label>
              ))}
            </div>
          </div>
          <button
            type="submit"
            className="mt-8 text-xl font-semibold py-3 px-8 border border-transparent rounded-lg text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500 transition duration-150 ease-in-out"
          >
            Create Class
          </button>
        </form>
      </div>
    </div>
  );
};
export default CreateClass;
