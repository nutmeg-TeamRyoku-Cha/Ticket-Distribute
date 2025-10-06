import { BrowserRouter as Router, Routes, Route} from "react-router-dom";
import Login from "./pages/login"
import Label from "./components/atomic/Label";

function App() {
  return (
    <Router>
      <Routes>
        <Route path="/login" element={<Login />} />
        <Route path="/ImputField" element={<ImputField />} />
        <Route path="/Label" element={<Label />} />
      </Routes>
    </Router>
      
  )
}

export default App
