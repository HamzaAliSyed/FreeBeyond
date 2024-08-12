import Header from "./components/Header"
import CreateAccountPage from "./pages/CreateAccountPage";
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
function App() {
  return (
    <Router>
      <Routes>
        <Route exact path="/" element={<Header />}/>
        <Route path="/create-account" element={<CreateAccountPage/>}/>
      </Routes>
    </Router>
  )
}

export default App
