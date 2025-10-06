import { BrowserRouter as Router, Routes, Route} from "react-router-dom";
import Login from "./pages/login"
import Button from "./components/atomic/Button"

function App() {
  return (
    <Router>
      <Routes>
        <Route path="/login" element={<Login />} />
        <Route path="/button" element={<Button label="登録" variant="registration"/>} />
      </Routes>
    </Router>
  )
}

export default App