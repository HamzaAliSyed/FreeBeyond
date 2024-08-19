import Header from "./components/Header"
import MainBody from "./components/MainBody";
import CreateAccountPage from "./pages/CreateAccountPage";
import CharacterName from "./pages/CharacterName";
import SignInPage from "./pages/SignInPage";
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
function App() {
  return (
    <Router>
      <Routes>
        <Route exact path="/" element={<Header />}/>
        <Route path="/create-account" element={<CreateAccountPage/>}/>
        <Route path="/sign-in" element={<SignInPage/>} />
        <Route path="/create-character" element={<CharacterName/>}/>
      </Routes>
      <MainBody />
    </Router>
  )
}

export default App
