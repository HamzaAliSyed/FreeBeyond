import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import CreateSubClassPageImage from "../assets/createsubclasspage.jpg";

const CreateSubClass = () => {
  const [name, setName] = useState("");
  const [parentClass, setParentClass] = useState("");
  const [availableClasses, setAvailableClasses] = useState([]);
  const [source, setSource] = useState("");
  const [availableSources, setAvailableSources] = useState([]);
  const navigate = useNavigate();

  const handleSetName = (event) => {
    setName(event.target.value);
  };

  const handleParentClassChange = (event) => {
    setParentClass(event.target.value);
  };

  const handleSourceChange = (event) => {
    setSource(event.target.value);
  };

  const handleSubmit = async (event) => {
    event.preventDefault();
    const subclassData = {
      name,
      parentClass,
      source,
    };

    try {
      const response = await fetch(
        "http://localhost:2712/api/components/createsubclass",
        {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify(subclassData),
        }
      );

      if (response.ok) {
        alert("Subclass created successfully!");
        navigate("/");
      } else {
        throw new Error(`Failed to create subclass: ${responseText}`);
      }
    } catch (error) {
      console.error("Error creating subclass:", error);
      alert("Failed to create subclass");
    }
  };

  const fetchAllClass = async () => {
    try {
      const response = await fetch(
        "http://localhost:2712/api/components/getallclasses"
      );

      if (response.ok) {
        const classes = await response.json();
        setAvailableClasses(classes);
      } else {
        throw new Error("Failed to fetch classes");
      }
    } catch (error) {
      console.error("Error fetching classes:", error);
    }
  };

  const fetchAllSources = async () => {
    try {
      const response = await fetch(
        "http://localhost:2712/api/components/getallsourcesnames"
      );

      if (response.ok) {
        const sources = await response.json();
        setAvailableSources(sources);
      } else {
        throw new Error("Failed to fetch sources");
      }
    } catch (error) {
      console.error("Error fetching sources:", error);
    }
  };

  useEffect(() => {
    fetchAllClass();
    fetchAllSources();
  }, []);

  return (
    <div className="create-subclass-container grid grid-cols-12 h-screen">
      <div className="col-span-3 h-full">
        <img
          src={CreateSubClassPageImage}
          className="w-full object-cover h-full"
          alt="A specialist"
        />
      </div>
      <div className="col-span-9 p-8 flex items-center justify-center">
        <form className="w-full max-w-2xl text-center" onSubmit={handleSubmit}>
          <div className="mb-8">
            <label
              className="block text-3xl font-semibold text-gray-800 mb-4"
              htmlFor="name"
            >
              SubClass Name
            </label>
            <input
              className="w-full text-2xl p-3 border-2 border-gray-300 rounded-lg focus:border-indigo-500 focus:ring focus:ring-indigo-200 focus:ring-opacity-50"
              id="name"
              type="text"
              value={name}
              onChange={handleSetName}
            />
          </div>
          <div className="mb-8">
            <label
              className="block text-3xl font-semibold text-gray-800 mb-4"
              htmlFor="parentClass"
            >
              Parent Class
            </label>
            <select
              id="parentClass"
              value={parentClass}
              onChange={handleParentClassChange}
              className="w-full text-2xl p-3 border-2 border-gray-300 rounded-lg focus:border-indigo-500 focus:ring focus:ring-indigo-200 focus:ring-opacity-50"
            >
              <option value="">Select Parent Class</option>
              {availableClasses.map((classes) => (
                <option key={classes} value={classes}>
                  {classes}
                </option>
              ))}
            </select>
          </div>
          <div className="mb-8">
            <label
              htmlFor="Source"
              className="block text-3xl font-semibold text-gray-800 mb-4"
            >
              Source
            </label>
            <select
              id="source"
              value={source}
              onChange={handleSourceChange}
              className="w-full text-2xl p-3 border-2 border-gray-300 rounded-lg focus:border-indigo-500 focus:ring focus:ring-indigo-200 focus:ring-opacity-50"
            >
              <option value="">Select Source</option>
              {availableSources.map((source) => (
                <option key={source} value={source}>
                  {source}
                </option>
              ))}
            </select>
          </div>
          <button
            type="submit"
            className="mt-8 text-xl font-semibold py-3 px-8 border border-transparent rounded-lg text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500 transition duration-150 ease-in-out"
          >
            Create Subclass
          </button>
        </form>
      </div>
    </div>
  );
};
export default CreateSubClass;
