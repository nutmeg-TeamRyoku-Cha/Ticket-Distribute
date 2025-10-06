import { BrowserRouter as Router, Routes, Route} from "react-router-dom";
import Login from "./pages/login"
import TicketList from "./pages/TicketList"

function App() {
  return (
    <Router>
      <Routes>
        <Route path="/login" element={<Login />} />
        <Route path="/TicketList" element={<TicketList />} />
      </Routes>
    </Router>
      
  )
}

export default App
