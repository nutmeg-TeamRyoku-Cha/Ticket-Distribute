import { BrowserRouter as Router, Routes, Route} from "react-router-dom";
import { useState } from "react";

//Header & Footer
import Header from "./components/Header";
import Footer from "./components/Footer";

//Page
import LoginPage from "./pages/login";

//Component
import Button from "./components/atomic/Button";
import InputField from "./components/atomic/InputField";
import Label from "./components/atomic/Label"
import LogoIcon from "./components/atomic/LogoIcon"
import TicketCard from "./components/atomic/TicketCard";

function App() {
  const [text, setText] = useState("");
  
  return (
    <Router>
      <Routes>
        {/*Button*/}
        <Route
          path="/button"
          element={
            <div style={{ padding: 20 }}>
              <Button
                label="登録"
                variant="registration"
                onClick={() => alert("登録ボタンが押されました")}
              />
            </div>
          }
        />
        {/* InputField */}
        <Route
          path="/input"
          element={
            <div style={{ padding: 20 }}>
              <InputField
                label="名前"
                value={text}
                onChange={(e) => setText(e.target.value)}
                type="text"
                placeholder="山田太郎"
              />
            </div>
          }
        />
        {/*label*/}
        <Route
          path="/label"
          element={
            <div style={{ padding: 20 }}>
              <Label
                text="ニックネーム"
                htmlFor="nickname"
                required
                fontSize={24}
                color="#1f2937"
              />
              {/* 確認用：別サイズ・別色 */}
              <div style={{ marginTop: 16 }}>
                <Label
                  text="メールアドレス"
                  htmlFor="email"
                  fontSize="18px"
                  color="orange"
                />
              </div>
            </div>
          }
        />
        {/*logoicon*/}
        <Route
            path="/logo"
            element={
              <div style={{ padding: 20 }}>
                <LogoIcon
                  icon="/images/logo-icon.png"   // public/images にある画像ファイルを指定
                  alt="アプリのロゴ"
                  variant="logo"
                  type="button"
                  onClick={() => alert("ロゴがクリックされました")}
                />
              </div>
            }
        />

        <Route
            path="/ticketcard"
            element={
              <div style={{ padding: 20 }}>
                <TicketCard
                  title="テストイベント"
                  time="12:00～"
                  location="テスト会場"
                  onClick={() => alert("カードがクリックされました")}
                />
              </div>
            }
        />
        <Route
            path="/icon"
            element={
              <div style={{ padding: 20 }}>
                <LogoIcon
                  icon="/images/setting-icon.png"   // public/images にある画像ファイルを指定
                  alt="アプリのロゴ"
                  variant="logo"
                  type="button"
                  onClick={() => alert("ロゴがクリックされました")}
                />
              </div>
            }
        />
        <Route path="/login" element={<LoginPage />} />
      </Routes>
    </Router>
  );
}

export default App;
