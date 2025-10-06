import { BrowserRouter as Router, Routes, Route} from "react-router-dom";
import Login from "./pages/login"

function App() {
  return (
    <Router>
      <Routes>
        <Route path="/login" element={<Login />} />
        <Route path="/ImputField" element={<ImputField />} />
      </Routes>
    </Router>
      
  )
}

export default App
