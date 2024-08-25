import SpellCreate from "../assets/spellcreate.jpg";
import { useState, useEffect, useCallback } from "react";

const CreateSpell = () => {
  const [spellName, setSpellName] = useState("");
  const [source, setSourceName] = useState({});
  const [sourceLoaded, setSourceLoaded] = useState(false);
  const [selectedSource, setSelectedSource] = useState("");

  const getSources = useCallback(async () => {
    try {
      const response = await fetch("http://localhost:2712/api/sources/getall", {
        method: "GET",
      });

      if (response.ok) {
        const data = await response.json();
        const sourcesDict = data.reduce(
          (acc, item) => ({ ...acc, [item]: item }),
          {}
        );

        setSourceName(sourcesDict);
        setSourceLoaded(true);
      } else {
        const errorText = await response.text();
        console.error(`Error fetching sources: ${errorText}`);
      }
    } catch (error) {
      console.error("Fetch error:", error);
    }
  }, []);

  useEffect(() => {
    if (!sourceLoaded) {
      getSources();
    }
  }, [getSources, sourceLoaded]);

  const handleSourceChange = (e) => {
    setSelectedSource(e.target.value);
  };

  const handleSpellName = (e) => {
    e.preventDefault();
    setSpellName(e.target.value);
  };

  const handleSubmit = async (e) => {
    e.preventDefault();

    try {
    } catch (error) {
      console.error("Submit error:", error);
    }
  };

  return (
    <div className="grid grid-cols-12 h-screen">
      <div className="col-span-3">
        <img
          src={SpellCreate}
          alt="Spell Create"
          className="w-full h-full object-cover"
        />
      </div>
      <div className="col-span-9 flex items-center justify-center">
        <form
          className="w-2/3 bg-white p-8 shadow-lg rounded-lg"
          onSubmit={handleSubmit}
        >
          <div className="mb-4">
            <label className="block text-gray-700 text-sm font-bold mb-2">
              Spell Name
            </label>
            <input
              className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
              type="text"
              placeholder="Spell Name"
              value={spellName}
              onChange={handleSpellName}
              required
            />
          </div>
          <div className="mb-4">
            <label className="block text-gray-700 text-sm font-bold mb-2">
              Level
            </label>
            <select className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline">
              {[...Array(10).keys()].map((value) => (
                <option key={value} value={value}>
                  {value}
                </option>
              ))}
            </select>
          </div>
          <div className="mb-4">
            <label className="block text-gray-700 text-sm font-bold mb-2">
              Source
            </label>
            <select
              className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
              onChange={handleSourceChange}
              value={selectedSource}
              required
            >
              <option value="" disabled>
                Select a Source
              </option>
              {Object.keys(source).map((key) => (
                <option key={key} value={key}>
                  {source[key]}
                </option>
              ))}
            </select>
          </div>

          <div className="flex items-center justify-center">
            <button
              className="bg-[#BC0F0F] hover:bg-red-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline"
              type="submit"
            >
              Create
            </button>
          </div>
        </form>
      </div>
    </div>
  );
};

export default CreateSpell;
