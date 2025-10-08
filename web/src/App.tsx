import { BrowserRouter as Router, Routes, Route} from "react-router-dom";

//Pageのみ配置
import LoginPage from "./pages/Account";
import TicketList from "./pages/TicketList";
import GetTicket from "./pages/GetTicket";
import EventList from "./pages/EventList";

function App() {  
  return (
    <Router>
      <Routes>
        <Route path="/account" element={<LoginPage />} />
        <Route path="/ticket" element={<TicketList />} />
        <Route path="/event" element={<EventList />} />
        <Route path="/getticket" element={<GetTicket />} />
        <Route path="/event" element={<EventList />} />
      </Routes>
    </Router>
  );
}

export default App;
