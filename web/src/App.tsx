import { BrowserRouter as Router, Routes, Route} from "react-router-dom";
import React, { useState } from "react";
import InputField from "./components/atomic/InputField";

function App() {
  const [value, setValue] = useState("");

  return (
    <Router>
      <Routes>
        <Route
          path="/input"
          element={
            <InputField
              label="ユーザー名"
              placeholder="山田太郎"
              value={value}
              onChange={(e) => setValue(e.target.value)}
              type="text"
            />
          }
        />
      </Routes>
    </Router>
  )
}

export default App;
