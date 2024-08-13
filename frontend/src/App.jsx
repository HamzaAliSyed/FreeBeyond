import Header from "./components/Header"
import CreateAccountPage from "./pages/CreateAccountPage";
import SignInPage from "./pages/SignInPage";
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
function App() {
  return (
    <Router>
      <Routes>
        <Route exact path="/" element={<Header />}/>
        <Route path="/create-account" element={<CreateAccountPage/>}/>
        <Route path="/sign-in" element={<SignInPage/>} />
      </Routes>
    </Router>
  )
}

export default App
