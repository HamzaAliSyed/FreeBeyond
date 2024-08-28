import React, { useState } from "react";

const CCTabNameStats = () => {
  const [characterPortrait, setCharacterPortrait] = useState(undefined);

  const handleCharacterPortrait = (event) => {
    const file = event.target.files[0];
    if (file) {
      const reader = new FileReader();
      reader.onload = (event2) => setCharacterPortrait(event2.target.result);
      reader.readAsDataURL(file);
    }
  };

  return (
    <div className="w-9/11 h-[calc(100vh-48px)] bg-gray-200 mx-auto rounded-md p-3">
      <h1 className="text-gray-800 text-center text-4xl mb-3">Name & Stats</h1>
      <div className="grid grid-cols-2 gap-4 h-[calc(100%-4rem)]">
        <div className="px-5 py-3">1</div>
        <div className="flex flex-col h-full px-6">
          <h2 className="text-xl font-semibold mb-2 text-center	">
            Character Portrait
          </h2>
          <label
            className="flex-grow bg-gray-300 flex items-center justify-center cursor-pointer rounded-md overflow-hidden"
            htmlFor="image-upload"
          >
            {characterPortrait ? (
              <img
                src={characterPortrait}
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
