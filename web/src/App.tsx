import { BrowserRouter as Router, Routes, Route} from "react-router-dom";
import { useState } from "react";

//Pageのみ配置
import LoginPage from "./pages/Account";
import TicketList from "./pages/TicketList";
import GetTicket from "./pages/GetTicket";
import EventList from "./pages/EventList";

function App() {
  const [text, setText] = useState("");
  
  return (
    <Router>
      <Routes>
        <Route path="/account" element={<LoginPage />} />
        <Route path="/tickets" element={<TicketList />} />
        <Route path="/getticket" element={<GetTicket />} />
        <Route path="/events" element={<EventList />} />
      </Routes>
    </Router>
  );
}

export default App;
