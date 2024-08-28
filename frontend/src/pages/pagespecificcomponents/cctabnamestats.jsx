import React from "react";

const CCTabNameStats = ({ data, onDataChange }) => {
  const handleCharacterPortrait = (event) => {
    const file = event.target.files[0];
    if (file) {
      const reader = new FileReader();
      reader.onload = (event2) =>
        onDataChange({ characterPortrait: event2.target.result });
      reader.readAsDataURL(file);
    }
  };

  const handleNameChange = (event) => {
    onDataChange({ name: event.target.value });
  };

  return (
    <div className="w-full h-[calc(100vh-48px)] bg-gray-200 rounded-md p-3">
      <h1 className="text-gray-800 text-center text-4xl mb-3">Name & Stats</h1>
      <div className="grid grid-cols-2 gap-4 h-[calc(100%-4rem)]">
        <div className="px-5 py-3">
          <input
            type="text"
            value={data.name || ""}
            onChange={handleNameChange}
            placeholder="Character Name"
            className="w-full p-2 border rounded"
          />
        </div>
        <div className="flex flex-col h-full">
          <h2 className="text-xl font-semibold mb-2">Character Portrait</h2>
          <label
            className="flex-grow bg-gray-300 flex items-center justify-center cursor-pointer rounded-md overflow-hidden"
            htmlFor="image-upload"
          >
            {data.characterPortrait ? (
              <img
                src={data.characterPortrait}
                alt="Character Portrait"
                className="w-full h-full object-cover"
              />
            ) : (
              <span className="text-black text-center p-4">
                Click here to upload character portrait
              </span>
            )}
          </label>
          <input
            id="image-upload"
            type="file"
            accept="image/*"
            className="hidden"
            onChange={handleCharacterPortrait}
          />
        </div>
      </div>
    </div>
  );
};

export default CCTabNameStats;
