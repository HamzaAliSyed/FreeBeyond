import CCTabBackground from "./pagespecificcomponents/cctabbackground";
import CCTabClass from "./pagespecificcomponents/cctabclass";
import CCTabRace from "./pagespecificcomponents/cctabrace";
import CCTabNameStats from "./pagespecificcomponents/cctabnamestats";
import { useState } from "react";

const CreateCharacterPage = () => {
  const tabs = [
    { id: "namestats", label: "Name and Stats", component: CCTabNameStats },
    { id: "background", label: "Background", component: CCTabBackground },
    { id: "race", label: "Race", component: CCTabRace },
    { id: "class", label: "Class", component: CCTabClass },
  ];
  const [activeTab, setActiveTab] = useState("namestats");

  return (
    <div className="max-w-4xl mx-auto mt-8">
      <div className="border-b border-gray-200">
        <nav className="-mb-px flex">
          {tabs.map((tab) => (
            <button
              key={tab.id}
              onClick={() => setActiveTab(tab.id)}
              className={`py-2 px-4 text-sm font-medium ${
                activeTab === tab.id
                  ? "border-b-2 border-blue-500 text-blue-600"
                  : "text-gray-500 hover:text-gray-700"
              }`}
            >
              {tab.label}
            </button>
          ))}
        </nav>
      </div>
      <div className="mt-4">
        {tabs.find((tab) => tab.id === activeTab)?.component()}
      </div>
    </div>
  );
};

export default CreateCharacterPage;
