import Header from "./components/Header"
import CreateAccountPage from "./pages/CreateAccountPage";
import SignInPage from "./pages/SignInPage";
import MainPage from "./pages/MainPage";
import CreateASource from "./pages/CreateSource";
import CreateSpell from "./pages/CreateSpell";
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
function App() {
  return (
    <Router>
      <Routes>
        <Route exact path="/" element={
          <>
            {<Header/>}
            {<MainPage/>}
          </>
        }/>
        <Route path="/create-account" element={<CreateAccountPage/>}/>
        <Route path="/sign-in" element={<SignInPage/>} />
        <Route path="create-source" element={<CreateASource/>} />
        <Route path="create-spell" element={<CreateSpell/>} />
      </Routes>
    </Router>
  )
}

export default App
