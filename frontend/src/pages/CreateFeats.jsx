import { useState, useEffect } from "react";
import FeatsSourceImage from "../assets/createfeatspage.jpg";

const CreateFeats = () => {
  const [sources, setSources] = useState([]);
  const [selectedSource, setSelectedSource] = useState("");

  const getsources = async () => {
    try {
      const response = await fetch(
        "http://localhost:2712/api/components/getallsourcesnames",
        {
          method: "GET",
          headers: {
            "Content-Type": "application/json",
          },
        }
      );

      if (response.ok) {
        const data = await response.json();
        setSources(data);
      } else {
        throw new Error("Internal Server Error");
      }
    } catch (error) {
      console.error("Error fetching sources:", error);
    }
  };

  useEffect(() => {
    getsources();
  }, []);

  return (
    <div className="create-feats-container grid grid-cols-12 h-screen">
      <div className="col-span-3 h-full">
        <img
          src={FeatsSourceImage}
          alt="Feats COver Image"
          className="w-full object-cover h-full"
        />
      </div>
      <div className="col-span-9 p-8 flex items-center justify-center">
        <form className="w-full max-w-2xl text-center">
          <div className="mb-8">
            <label
              htmlFor="name"
              className="block text-3xl font-semibold text-gray-800 mb-4"
            >
              Name
            </label>
            <input
              type="text"
              id="name"
              className="w-full px-4 py-2 border rounded-md"
              placeholder="Enter feat name"
            />
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
              onChange={(e) => setSelectedSource(e.target.value)}
              className="w-full px-4 py-2 border rounded-md"
            >
              <option value="">Select a source</option>
              {sources.map((source, index) => (
                <option key={index} value={source}>
                  {source}
                </option>
              ))}
            </select>
          </div>
        </form>
      </div>
    </div>
  );
};
export default CreateFeats;
