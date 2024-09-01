import { useState } from "react";

import CreateItemsPage from "../assets/createitemspage.jpg";

const CreateItems = () => {
  const [name, setName] = useState("");
  const [typeTags, setTypeTags] = useState([]);
  const [itemType, setItemType] = useState("");
  const [requiresAttunement, setRequiresAttunement] = useState(false);
  const [cost, setCost] = useState("");
  const [weight, setWeight] = useState("");

  const handleNameChange = (event) => {
    setName(event.target.value);
  };

  const handleTypeTagChange = (event) => {
    const { value, checked } = event.target;
    if (checked) {
      setTypeTags([...typeTags, value]);
    } else {
      setTypeTags(typeTags.filter((type) => type !== value));
    }
  };

  const handleItemTypeChange = (event) => {
    setItemType(event.target.value);
  };

  const handleRequiresAttunementChange = (event) => {
    setRequiresAttunement(event.target.checked);
  };

  const handleCostChange = (event) => {
    setCost(event.target.value);
  };

  const handleWeightChange = (event) => {
    setWeight(event.target.value);
  };

  const handleSubmit = (event) => {
    event.preventDefault();
    console.log("Form submitted", {
      name,
      typeTags,
      itemType,
      requiresAttunement,
      cost,
      weight,
    });
  };

  return (
    <div className="create-class-container grid grid-cols-12 h-screen">
      <div className="col-span-3 h-full">
        <img
          src={CreateItemsPage}
          className="w-full object-cover h-full"
          alt="Create an item"
        />
      </div>
      <div className="col-span-9 p-8 flex items-center justify-center">
        <div className="w-full max-w-2xl text-center">
          <h1 className="text-4xl font-bold mb-8">Create New D&D Item</h1>
          <form className="w-full">
            <div className="mb-6">
              <label
                htmlFor="name"
                className="block text-xl font-semibold text-gray-700 mb-2"
              >
                Name
              </label>
              <input
                type="text"
                id="name"
                value={name}
                onChange={handleNameChange}
                className="w-full text-lg p-2 border-2 border-gray-300 rounded-lg focus:border-indigo-500 focus:ring focus:ring-indigo-200 focus:ring-opacity-50"
              />
            </div>
            <div className="mb-6">
              <label className="block text-xl font-semibold text-gray-700 mb-2">
                Type
              </label>
              <div className="grid grid-cols-3 gap-2">
                {[
                  "AdventuringGear",
                  "Ammunation",
                  "Trade Goods",
                  "Others",
                  "Vehicle (Air)",
                  "Vehicle (Land)",
                  "Vehicle (Water)",
                  "Vehicle (Space)",
                  "Artisian Tools",
                  "Food and Drink",
                  "Spellcasting Focus",
                  "Mount",
                  "Poison",
                  "Instrument",
                  "Tack and Harness",
                  "Martial Weapon",
                  "Simple Weapon",
                  "Firearm",
                  "Ranged Weapon",
                  "Melee Weapon",
                  "Light Armor",
                  "Medium Armor",
                  "Heavy Armor",
                  "Shields",
                  "Explosive",
                  "Gaming Set",
                  "Wondrous Item",
                  "Staff",
                  "Wand",
                  "Tattoo",
                  "Rod",
                  "Ring",
                  "Scroll",
                ].map((type) => (
                  <label key={type} className="inline-flex items-center">
                    <input
                      type="checkbox"
                      value={type}
                      checked={typeTags.includes(type)}
                      onChange={handleTypeTagChange}
                      className="form-checkbox h-5 w-5 text-indigo-600"
                    />
                    <span className="ml-2 text-sm">{type}</span>
                  </label>
                ))}
              </div>
            </div>
            <div className="mb-6">
              <label
                htmlFor="itemType"
                className="block text-xl font-semibold text-gray-700 mb-2"
              >
                Item Type
              </label>
              <select
                id="itemType"
                value={itemType}
                onChange={handleItemTypeChange}
                className="w-full text-lg p-2 border-2 border-gray-300 rounded-lg focus:border-indigo-500 focus:ring focus:ring-indigo-200 focus:ring-opacity-50"
              >
                <option value="">Select Item Type</option>
                <option value="Mundane">Mundane</option>
                <option value="Common">Common</option>
                <option value="Uncommon">Uncommon</option>
                <option value="Rare">Rare</option>
                <option value="VeryRare">Very Rare</option>
                <option value="Legendary">Legendary</option>
                <option value="Artifact">Artifact</option>
              </select>
            </div>
            <div className="mb-6">
              <label className="inline-flex items-center">
                <input
                  type="checkbox"
                  checked={requiresAttunement}
                  onChange={handleRequiresAttunementChange}
                  className="form-checkbox h-5 w-5 text-indigo-600"
                />
                <span className="ml-2 text-xl font-semibold text-gray-700">
                  Requires Attunement
                </span>
              </label>
            </div>
            <div className="mb-6">
              <label
                htmlFor="cost"
                className="block text-xl font-semibold text-gray-700 mb-2"
              >
                Cost
              </label>
              <input
                type="text"
                id="cost"
                value={cost}
                onChange={handleCostChange}
                className="w-full text-lg p-2 border-2 border-gray-300 rounded-lg focus:border-indigo-500 focus:ring focus:ring-indigo-200 focus:ring-opacity-50"
                placeholder="e.g. 10 gp"
              />
            </div>

            <div className="mb-6">
              <label
                htmlFor="weight"
                className="block text-xl font-semibold text-gray-700 mb-2"
              >
                Weight
              </label>
              <input
                type="text"
                id="weight"
                value={weight}
                onChange={handleWeightChange}
                className="w-full text-lg p-2 border-2 border-gray-300 rounded-lg focus:border-indigo-500 focus:ring focus:ring-indigo-200 focus:ring-opacity-50"
                placeholder="e.g. 2 lbs"
              />
            </div>
            <div className="mt-8">
              <button
                type="submit"
                className="w-full text-xl font-semibold py-3 px-8 border border-transparent rounded-lg text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500 transition duration-150 ease-in-out"
              >
                Create Item
              </button>
            </div>
          </form>
        </div>
      </div>
    </div>
  );
};

export default CreateItems;
