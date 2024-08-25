import CreateCharacterPage from "./pages/CreateCharacterPage";
import { BrowserRouter as Router, Route, Routes } from "react-router-dom";
function App() {
  return (
    <Router>
      <Routes>
        <Route path="/create-character" element={<CreateCharacterPage />} />
      </Routes>
    </Router>
  );
}

export default App;
