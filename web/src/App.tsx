import { BrowserRouter as Router, Routes, Route} from "react-router-dom";
import TicketCard from "./components/atomic/TicketCard";

function App() {
  const handleClick = () => alert('テストボタンが押されました')
  return (
    <div>
      <TicketCard
        title="テストイベント"
        time="12:00～"
        location="テスト会場"
        onClick={handleClick}
      />
    </div>
  );
}

export default App;
