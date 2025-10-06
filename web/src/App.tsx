import { BrowserRouter as Router, Routes, Route} from "react-router-dom";
import { useState } from "react";
import Login from "./pages/login";
import Label from "./components/atomic/Label";
import InputField from "./components/atomic/InputField";

function App() {
  const [text, setText] = useState("");

  return (
    <Router>
      <Routes>
        <Route path="/login" element={<Login />} />
        <Route path="/inputfield" element={<InputField label = "山田太郎" type="text" placeholder="お名前を入力してください" value={text} onChange={(e) => setText(e.target.value)} required={true} error={text === "" ? "必須項目です" : undefined} size="md" />} />
        <Route path="/label" element={<Label text="ニックネーム" htmlFor="nickname" fontSize = {50} color = "blue"/>} />
      </Routes>
    </Router>
      
  )
}

export default App
