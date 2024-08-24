import { useState } from "react";
import { useNavigate } from "react-router-dom";
import SourceCreate from "../assets/sourcecreate.jpg";

const CreateASource = () => {
    const homeNavigate = useNavigate();
    const [stateName, setStateName] = useState("");
    const [stateType, setStateType] = useState("");
    const [statePublishDate, setStatePublishDate] = useState("");

    // 1. Create a dictionary to map dropdown options to backend values
    const typeOptions = {
        "Core": "Core",
        "Supplements": "Supplements",
        "Settings": "Settings",
        "Extras" : "Extras",
        "Screens" : "Screens",
        "Miscellaneous" : "Miscellaneous",
    };

    const handleName = (e) => {
        e.preventDefault();
        setStateName(e.target.value);
    };

    // 2. Update handleType to use the dictionary
    const handleType = (e) => {
        e.preventDefault();
        const selectedType = e.target.value;
        setStateType(typeOptions[selectedType]); // Get the backend value from the dictionary
    };

    const handlePublishDate = (e) => {
        e.preventDefault();
        setStatePublishDate(e.target.value);
    };

    const OnSubmit = async (e) => {
        e.preventDefault();

        // Format the date as "MonthName-Day-Year"
        const formattedDate = new Date(statePublishDate).toLocaleDateString('en-US', {
            month: 'long',
            day: 'numeric',
            year: 'numeric'
        });

        const SourceCreation = {
            Name: stateName,
            Type: stateType, // This now contains the mapped value
            PublishDate: formattedDate,
        };

        try {
            const response = await fetch("http://localhost:2712/api/source/create", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify(SourceCreation),
            });

            if (response.ok) {
                alert("New Source Added");
                homeNavigate("/");
            } else {
                const errorText = await response.text();
                alert(`Couldn't create new source: ${errorText}`);
            }
        } catch (error) {
            console.error("Server error", error);
            alert(`There is an error in backend server: ${error}`);
        }
    };

    return (
        <div className="grid grid-cols-12 h-screen">
            <div className="col-span-3">
                <img
                    src={SourceCreate}
                    alt="Source Create"
                    className="w-full h-full object-cover"
                />
            </div>
            <div className="col-span-9 flex items-center justify-center">
                <form
                    className="w-2/3 bg-white p-8 shadow-lg rounded-lg"
                    onSubmit={OnSubmit}
                >
                    <div className="mb-4">
                        <label className="block text-gray-700 text-sm font-bold mb-2">
                            Name
                        </label>
                        <input
                            className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                            type="text"
                            placeholder="Name"
                            value={stateName}
                            onChange={handleName}
                            required
                        />
                    </div>
                    <div className="mb-4">
                        <label className="block text-gray-700 text-sm font-bold mb-2">
                            Type
                        </label>
                        {/* 3. Change the input to a select (dropdown) */}
                        <select
                            className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                            value={stateType}
                            onChange={handleType}
                            required
                        >
                            <option value="" disabled>Select a Type</option>
                            {Object.keys(typeOptions).map((key) => (
                                <option key={key} value={key}>{key}</option>
                            ))}
                        </select>
                    </div>
                    <div className="mb-4">
                        <label className="block text-gray-700 text-sm font-bold mb-2">
                            Publish Date
                        </label>
                        <input
                            className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                            type="date"
                            placeholder="Publish Date"
                            value={statePublishDate}
                            onChange={handlePublishDate}
                            required
                        />
                    </div>
                    <div className="flex items-center justify-between">
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

export default CreateASource;
