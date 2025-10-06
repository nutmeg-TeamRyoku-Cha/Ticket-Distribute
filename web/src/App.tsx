import { BrowserRouter as Router, Routes, Route} from "react-router-dom";
import LogoIcon from "./pages/LogoIcon";

function App() {
  return (
    <Router>
      <Routes>
        <Route path="/logo" element={<LogoIcon />} />
      </Routes>
    </Router>
  );
}

export default App;
