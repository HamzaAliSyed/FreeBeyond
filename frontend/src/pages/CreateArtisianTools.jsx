import { useState, useEffect } from "react";
import { useNavigate } from "react-router-dom";
import CreateArtisianPage from "../assets/createartisiantoolspage.jpg";

const CreateArtisianTools = () => {
  const navigate = useNavigate();
  const [name, setName] = useState("");
  const [description, setDescription] = useState("");
  const [items, setItems] = useState([]);
  const [flavourText, setFlavourText] = useState([
    { title: "", flavourText: "" },
  ]);
  const [dcTable, setDcTable] = useState({});
  const [allItems, setAllItems] = useState([]);
  const [isLoading, setIsLoading] = useState(false);

  useEffect(() => {
    fetchItems();
  }, []);

  const fetchItems = async () => {
    try {
      const response = await fetch(
        "http://localhost:2712/api/components/getallitems"
      );
      if (!response.ok) {
        throw new Error("Network response was not ok");
      }
      const data = await response.json();
      setAllItems(data);
    } catch (error) {
      console.error("Error fetching items:", error);
    }
  };

  const handleInputChange = (setter) => (event) => {
    setter(event.target.value);
  };

  const handleItemsChange = (event) => {
    setItems(
      Array.from(event.target.selectedOptions, (option) => option.value)
    );
  };

  const handleFlavourTextChange = (index, field) => (event) => {
    const newFlavourText = [...flavourText];
    newFlavourText[index][field] = event.target.value;
    setFlavourText(newFlavourText);
  };

  const addFlavourText = () => {
    setFlavourText([...flavourText, { title: "", flavourText: "" }]);
  };

  const removeFlavourText = (index) => {
    setFlavourText(flavourText.filter((_, i) => i !== index));
  };

  const handleDcTableChange = (oldKey, newKey, value) => {
    const newDcTable = { ...dcTable };
    delete newDcTable[oldKey];
    if (newKey !== "" || value !== "") {
      newDcTable[newKey] = value === "" ? 0 : parseInt(value);
    }
    setDcTable(newDcTable);
  };

  const addDcTableEntry = () => {
    const newKey = `Task ${Object.keys(dcTable).length + 1}`;
    setDcTable({ ...dcTable, [newKey]: 0 });
  };

  const removeDcTableEntry = (key) => {
    const newDcTable = { ...dcTable };
    delete newDcTable[key];
    setDcTable(newDcTable);
  };

  const handleSubmit = async (event) => {
    event.preventDefault();
    setIsLoading(true);

    const artisianToolsData = {
      name,
      description,
      items,
      flavourText,
      dcTable,
    };

    try {
      const response = await fetch(
        "http://localhost:2712/api/components/createartisiantools",
        {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify(artisianToolsData),
        }
      );

      if (!response.ok) {
        throw new Error("Failed to create artisian tools");
      }

      const result = await response.json();
      alert(
        `Artisian tools "${result.name}" created successfully with ID: ${result.id}`
      );
      navigate("/");
    } catch (error) {
      console.error("Error creating artisian tools:", error);
      alert("Failed to create artisian tools. Please try again.");
    } finally {
      setIsLoading(false);
    }
  };

  return (
    <div className="create-artisian-tools-container grid grid-cols-12 h-screen">
      <div className="col-span-3 h-full">
        <img
          src={CreateArtisianPage}
          className="w-full object-cover h-full"
          alt="Create Artisian Tools"
        />
      </div>
      <div className="col-span-9 p-8 flex items-center justify-center">
        <div className="w-full max-w-2xl text-center">
          <h1 className="text-4xl font-bold mb-8">Create Artisian Tools</h1>
          <div className="max-h-[calc(100vh-200px)] overflow-y-auto">
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
                  onChange={handleInputChange(setName)}
                  className="w-full text-lg p-2 border-2 border-gray-300 rounded-lg focus:border-indigo-500 focus:ring focus:ring-indigo-200 focus:ring-opacity-50"
                />
              </div>
              <div className="mb-6">
                <label
                  htmlFor="description"
                  className="block text-xl font-semibold text-gray-700 mb-2"
                >
                  Description
                </label>
                <textarea
                  id="description"
                  value={description}
                  onChange={handleInputChange(setDescription)}
                  className="w-full text-lg p-2 border-2 border-gray-300 rounded-lg focus:border-indigo-500 focus:ring focus:ring-indigo-200 focus:ring-opacity-50"
                  rows="3"
                />
              </div>
              <div className="mb-6">
                <label className="block text-xl font-semibold text-gray-700 mb-2">
                  Items
                </label>
                <select
                  multiple
                  value={items}
                  onChange={handleItemsChange}
                  className="w-full text-lg p-2 border-2 border-gray-300 rounded-lg focus:border-indigo-500 focus:ring focus:ring-indigo-200 focus:ring-opacity-50"
                >
                  {allItems.map((item, index) => (
                    <option key={index} value={item}>
                      {item}
                    </option>
                  ))}
                </select>
              </div>
              <div className="mb-6">
                <label className="block text-xl font-semibold text-gray-700 mb-2">
                  Flavour Text
                </label>
                {flavourText.map((ft, index) => (
                  <div
                    key={index}
                    className="mb-4 p-4 border border-gray-300 rounded-lg relative"
                  >
                    <input
                      type="text"
                      value={ft.title}
                      onChange={handleFlavourTextChange(index, "title")}
                      className="w-full text-lg p-2 mb-2 border-2 border-gray-300 rounded-lg focus:border-indigo-500 focus:ring focus:ring-indigo-200 focus:ring-opacity-50"
                      placeholder="Enter title"
                    />
                    <textarea
                      value={ft.flavourText}
                      onChange={handleFlavourTextChange(index, "flavourText")}
                      className="w-full text-lg p-2 mb-2 border-2 border-gray-300 rounded-lg focus:border-indigo-500 focus:ring focus:ring-indigo-200 focus:ring-opacity-50"
                      placeholder="Enter flavour text"
                      rows="3"
                    />
                    <button
                      type="button"
                      onClick={() => removeFlavourText(index)}
                      className="absolute top-2 right-2 text-red-500 hover:text-red-700"
                    >
                      ×
                    </button>
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
                <label className="block text-xl font-semibold text-gray-700 mb-2">
                  DC Table
                </label>
                {Object.entries(dcTable).map(([key, value], index) => (
                  <div key={index} className="flex mb-2 relative">
                    <input
                      type="text"
                      value={key}
                      onChange={(e) =>
                        handleDcTableChange(key, e.target.value, value)
                      }
                      className="w-1/2 text-lg p-2 border-2 border-gray-300 rounded-l-lg focus:border-indigo-500 focus:ring focus:ring-indigo-200 focus:ring-opacity-50"
                      placeholder="Task"
                    />
                    <input
                      type="number"
                      value={value}
                      onChange={(e) =>
                        handleDcTableChange(key, key, e.target.value)
                      }
                      className="w-1/2 text-lg p-2 border-2 border-gray-300 rounded-r-lg focus:border-indigo-500 focus:ring focus:ring-indigo-200 focus:ring-opacity-50"
                      placeholder="DC"
                    />
                    <button
                      type="button"
                      onClick={() => removeDcTableEntry(key)}
                      className="absolute -right-6 top-1/2 transform -translate-y-1/2 text-red-500 hover:text-red-700"
                    >
                      ×
                    </button>
                  </div>
                ))}
                <button
                  type="button"
                  onClick={addDcTableEntry}
                  className="mt-2 px-4 py-2 bg-indigo-600 text-white rounded-lg hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:ring-opacity-50"
                >
                  + Add DC Entry
                </button>
              </div>
              <div className="mt-8">
                <button
                  type="submit"
                  disabled={isLoading}
                  className={`w-full text-xl font-semibold py-3 px-8 border border-transparent rounded-lg text-white ${
                    isLoading
                      ? "bg-indigo-400"
                      : "bg-indigo-600 hover:bg-indigo-700"
                  } focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500 transition duration-150 ease-in-out`}
                >
                  {isLoading ? "Creating..." : "Create Artisian Tools"}
                </button>
              </div>
            </form>
          </div>
        </div>
      </div>
    </div>
  );
};

export default CreateArtisianTools;
