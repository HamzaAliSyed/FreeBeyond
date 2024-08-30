import { useState } from "react";
import { useNavigate } from "react-router-dom";
import CreateSourceImage from "../assets/createsourcepage.jpg";

const typeOptions = [
  { value: "", label: "Please select a source" },
  { value: "Core", label: "Core" },
  { value: "Supplements", label: "Supplements" },
  { value: "Settings", label: "Settings" },
  { value: "Extras", label: "Extras" },
  { value: "Miscellaneous", label: "Miscellaneous" },
];

const monthOptions = [
  { value: "", label: "Month" },
  { value: "01", label: "January" },
  { value: "02", label: "February" },
  { value: "03", label: "March" },
  { value: "04", label: "April" },
  { value: "05", label: "May" },
  { value: "06", label: "June" },
  { value: "07", label: "July" },
  { value: "08", label: "August" },
  { value: "09", label: "September" },
  { value: "10", label: "October" },
  { value: "11", label: "November" },
  { value: "12", label: "December" },
];

const CreateSource = () => {
  const navigate = useNavigate();

  const [name, setName] = useState("");
  const handleNameChange = (event) => {
    setName(event.target.value);
  };

  const [type, setType] = useState("");
  const handleTypeChange = (event) => {
    const newType = event.target.value;
    setType(newType);
    console.log("Selected type:", newType);
  };

  const [date, setDate] = useState({ month: "", day: "", year: "" });

  const handleDateChange = (event) => {
    const { name, value } = event.target;
    setDate((prevDate) => ({
      ...prevDate,
      [name]: value,
    }));
  };

  const handleCreateSource = async (event) => {
    event.preventDefault();

    console.log("Current state - Name:", name, "Type:", type, "Date:", date);

    if (!name || !type || !date.month || !date.day || !date.year) {
      alert("Please fill in all required fields");
      return;
    }

    const monthNames = [
      "January",
      "February",
      "March",
      "April",
      "May",
      "June",
      "July",
      "August",
      "September",
      "October",
      "November",
      "December",
    ];
    const monthName = monthNames[parseInt(date.month) - 1];
    const formattedDate = `${monthName}-${date.day.padStart(2, "0")}-${
      date.year
    }`;
    const sourceData = {
      name: name,
      type: type,
      publishdate: formattedDate,
    };

    console.log("Sending source data:", sourceData);

    try {
      const response = await fetch(
        "http://localhost:2712/api/components/createsource",
        {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify(sourceData),
        }
      );

      console.log("Response status:", response.status);
      const responseData = await response.text();
      console.log("Response data:", responseData);

      if (response.ok) {
        alert("Source created successfully!");
        navigate("/");
      } else {
        throw new Error(`Failed to create source: ${responseData}`);
      }
    } catch (error) {
      console.error("Error creating source:", error);
      if (error.response) {
        console.error("Response status:", error.response.status);
        console.error("Response data:", error.response.data);
      }
      alert("Failed to create new source");
    }
  };

  return (
    <div className="create-source-container grid grid-cols-12 h-screen">
      <div className="col-span-3 h-full">
        <img
          src={CreateSourceImage}
          className="w-full object-cover h-full"
          alt="People near a campfire"
        />
      </div>
      <div className="col-span-9 p-8 flex items-center justify-center">
        <form
          onSubmit={handleCreateSource}
          className="w-full max-w-2xl text-center"
        >
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
              value={name}
              onChange={handleNameChange}
              className="w-full text-2xl p-3 border-2 border-gray-300 rounded-lg focus:border-indigo-500 focus:ring focus:ring-indigo-200 focus:ring-opacity-50"
            />
          </div>
          <div className="mb-8">
            <label
              htmlFor="type"
              className="block text-3xl font-semibold text-gray-800 mb-4"
            >
              Type
            </label>
            <select
              id="type"
              name="type"
              value={type}
              onChange={handleTypeChange}
              className="w-full text-2xl p-3 border-2 border-gray-300 rounded-lg focus:border-indigo-500 focus:ring focus:ring-indigo-200 focus:ring-opacity-50"
            >
              {typeOptions.map((option) => (
                <option key={option.value} value={option.value}>
                  {option.label}
                </option>
              ))}
            </select>
          </div>
          <div className="mb-8">
            <label className="block text-3xl font-semibold text-gray-800 mb-4">
              Date
            </label>
            <div className="flex space-x-4">
              <select
                name="month"
                value={date.month}
                onChange={handleDateChange}
                className="flex-1 text-2xl p-3 border-2 border-gray-300 rounded-lg focus:border-indigo-500 focus:ring focus:ring-indigo-200 focus:ring-opacity-50"
              >
                {monthOptions.map((option) => (
                  <option key={option.value} value={option.value}>
                    {option.label}
                  </option>
                ))}
              </select>
              <input
                type="text"
                name="day"
                value={date.day}
                onChange={handleDateChange}
                placeholder="Day"
                className="flex-1 text-2xl p-3 border-2 border-gray-300 rounded-lg focus:border-indigo-500 focus:ring focus:ring-indigo-200 focus:ring-opacity-50"
              />
              <input
                type="text"
                name="year"
                value={date.year}
                onChange={handleDateChange}
                placeholder="Year"
                className="flex-1 text-2xl p-3 border-2 border-gray-300 rounded-lg focus:border-indigo-500 focus:ring focus:ring-indigo-200 focus:ring-opacity-50"
              />
            </div>
          </div>
          <button
            type="submit"
            className="text-xl font-semibold py-3 px-8 border border-transparent rounded-lg text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500 transition duration-150 ease-in-out"
          >
            Create Source
          </button>
        </form>
      </div>
    </div>
  );
};
export default CreateSource;
