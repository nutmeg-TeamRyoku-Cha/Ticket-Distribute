import { BrowserRouter as Router, Routes, Route} from "react-router-dom";

function App() {
  return (
    <Router>
      <Routes>
        <Route path="/logo" element={<img src="/images/logo-icon.png"/>} />
      </Routes>
    </Router>
  );
}

export default App;
