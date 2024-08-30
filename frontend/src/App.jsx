import MainPage from "./pages/MainPage";
import CreateCharacterPage from "./pages/CreateCharacterPage";
import CreateSource from "./pages/CreateSource";
import { BrowserRouter as Router, Route, Routes } from "react-router-dom";
function App() {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<MainPage />} />
        <Route path="/create-character" element={<CreateCharacterPage />} />
        <Route path="/create-source" element={<CreateSource />} />
      </Routes>
    </Router>
  );
}

export default App;
