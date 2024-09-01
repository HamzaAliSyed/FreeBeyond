import { useState, useEffect } from "react";
import { useNavigate } from "react-router-dom";

import CreateItemsPage from "../assets/createitemspage.jpg";

const CreateItems = () => {
  const navigate = useNavigate();

  const [name, setName] = useState("");
  const [typeTags, setTypeTags] = useState([]);
  const [itemType, setItemType] = useState("");
  const [requiresAttunement, setRequiresAttunement] = useState(false);
  const [cost, setCost] = useState("");
  const [weight, setWeight] = useState("");
  const [flavourTexts, setFlavourTexts] = useState([
    { title: "", flavourtext: "" },
  ]);

  const [sources, setSources] = useState([]);
  const [selectedSource, setSelectedSource] = useState("");

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

  const handleSubmit = async (event) => {
    event.preventDefault();

    const itemData = {
      name,
      typeTags,
      itemType,
      requiresAttunement,
      cost,
      weight,
      flavourText: flavourTexts,
      source: selectedSource,
    };

    try {
      const response = await fetch(
        "http://localhost:2712/api/components/createitems",
        {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify(itemData),
        }
      );

      if (!response.ok) {
        throw new Error("Failed to create item");
      }

      alert("Item created successfully!");
      navigate("/");
    } catch (error) {
      console.error("Error creating item:", error);
      alert("Failed to create item. Please try again.");
    }
  };

  const addFlavourText = () => {
    setFlavourTexts([...flavourTexts, { title: "", flavourtext: "" }]);
  };

  const handleFlavourTextChange = (index, field, value) => {
    const newFlavourTexts = [...flavourTexts];
    newFlavourTexts[index][field] = value;
    setFlavourTexts(newFlavourTexts);
  };

  const handleSourceChange = (event) => {
    setSelectedSource(event.target.value);
  };

  useEffect(() => {
    fetchSources();
  }, []);

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
          <form className="w-full" onSubmit={handleSubmit}>
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
                  "Adventuring Gear",
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
            <div className="mb-6">
              <label className="block text-xl font-semibold text-gray-700 mb-2">
                Flavour Text
              </label>
              {flavourTexts.map((flavourText, index) => (
                <div key={index} className="mb-4">
                  <input
                    type="text"
                    value={flavourText.title}
                    onChange={(e) =>
                      handleFlavourTextChange(index, "title", e.target.value)
                    }
                    className="w-full text-lg p-2 mb-2 border-2 border-gray-300 rounded-lg"
                    placeholder="Title"
                  />
                  <textarea
                    value={flavourText.flavourtext}
                    onChange={(e) =>
                      handleFlavourTextChange(
                        index,
                        "flavourtext",
                        e.target.value
                      )
                    }
                    className="w-full text-lg p-2 border-2 border-gray-300 rounded-lg"
                    rows="3"
                    placeholder="Flavour Text"
                  />
                </div>
              ))}
              <button
                type="button"
                onClick={addFlavourText}
                className="mt-2 px-4 py-2 bg-indigo-600 text-white rounded-lg hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:ring-opacity-50"
              >
                + Add Flavour Text
              </button>
            </div>
            <div className="mb-6">
              <label
                htmlFor="source"
                className="block text-xl font-semibold text-gray-700 mb-2"
              >
                Source
              </label>
              <select
                id="source"
                value={selectedSource}
                onChange={handleSourceChange}
                className="w-full text-lg p-2 border-2 border-gray-300 rounded-lg focus:border-indigo-500 focus:ring focus:ring-indigo-200 focus:ring-opacity-50"
              >
                <option value="">Select a Source</option>
                {sources.map((sourceName, index) => (
                  <option key={index} value={sourceName}>
                    {sourceName}
                  </option>
                ))}
              </select>
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
