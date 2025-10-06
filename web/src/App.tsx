import { BrowserRouter as Router, Routes, Route} from "react-router-dom";
import Login from "./pages/login"
import Label from "./components/atomic/Label";

function App() {
  return (
    <Router>
      <Routes>
        <Route path="/login" element={<Login />} />
        <Route path="/InputField" element={<InputField />} />
        <Route path="/Label" element={<Label text="ニックネーム" htmlFor="nickname"/>} />
      </Routes>
    </Router>
      
  )
}

export default App
