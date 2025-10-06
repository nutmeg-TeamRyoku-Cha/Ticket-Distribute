import { BrowserRouter as Router, Routes, Route} from "react-router-dom";
import Login from "./pages/login"
import Label from "./components/atomic/Label"

function App() {
  return (
    <Router>
      <Routes>
        <Route path="/login" element={<Login />} />
        <Route path="/TicketList" element={<TicketList />} />
        <Route path="/label" element={<Label text = "登録"/>} />
      </Routes>
    </Router>
      
  )
}

export default App
