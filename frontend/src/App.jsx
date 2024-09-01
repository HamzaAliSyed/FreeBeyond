import MainPage from "./pages/MainPage";
import CreateCharacterPage from "./pages/CreateCharacterPage";
import CreateSource from "./pages/CreateSource";
import CreateSpells from "./pages/CreateSpell";
import CreateFeats from "./pages/CreateFeats";
import CreateClass from "./pages/CreateClass";
import CreateItems from "./pages/CreateItem";
import { BrowserRouter as Router, Route, Routes } from "react-router-dom";
function App() {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<MainPage />} />
        <Route path="/create-character" element={<CreateCharacterPage />} />
        <Route path="/create-source" element={<CreateSource />} />
        <Route path="/create-feat" element={<CreateFeats />} />
        <Route path="/create-spell" element={<CreateSpells />} />
        <Route path="/create-class" element={<CreateClass />} />
        <Route path="/create-item" element={<CreateItems />} />
      </Routes>
    </Router>
  );
}

export default App;
