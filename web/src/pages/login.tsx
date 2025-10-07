import React, { useState } from "react"
import Label from "../components/atomic/Label"
import InputField from "../components/atomic/InputField"
import Button from "../components/atomic/Button"
import Header from "../components/Header"
import Footer from "../components/Footer"
import "./login.css"

const API_BASE = import.meta.env.VITE_API_BASE_URL ?? "http://localhost:8080";

const LoginPage: React.FC = () => {
  const [regNickname, setRegNickname] = useState("");
  const [regBirthDate, setRegBirthDate] = useState("");
  const [regPartySize, setRegPartySize] = useState("");
  const [loginNickname, setLoginNickname] = useState("");
  const [loginBirthDate, setLoginBirthDate] = useState("");
  const [submitting, setSubmitting] = useState(false);
  const [message, setMessage] = useState<string | null>(null);

  const handleRegister = async () => {
    if (!regNickname || !regBirthDate || !regPartySize) {
      setMessage("未入力があります");
      return;
    }
    const partySize = Number(regPartySize);
    if (!Number.isInteger(partySize) || partySize < 1) {
      setMessage("来場者人数は1以上の整数にしてください");
      return;
    }
    setSubmitting(true);
    setMessage(null);
    try {
      const res = await fetch(`${API_BASE}/visitors`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({
          nickname: regNickname,
          birth_date: regBirthDate,
          party_size: partySize,
        }),
      });
      if (!res.ok) {
        const text = await res.text();
        throw new Error(text || `HTTP ${res.status}`);
      }
      const data: { visitor_id: number } = await res.json();
      setMessage(`登録しました（ID: ${data.visitor_id}）`);
    } catch (e: any) {
      setMessage(`登録に失敗しました: ${e.message ?? e}`);
    } finally {
      setSubmitting(false);
    }
  };

  const handleLogin = async () => {
    setMessage("ログインAPI接続は後で実装します");
  };

  return (
    <>
      <Header title="ログイン" />
      <div className="login-container">
        <div className="login-frame">
          <div style={{ textAlign: "center", lineHeight: 1.35, marginBottom: 12 }}>
            <Label text="ようこそ！技大祭へ！" fontSize={20} color="#222" /> <br></br> <br></br>
            <Label text="新規の方は来場者登録をお願いします" fontSize={14} color="#666" /> <br></br> <br></br>
          </div>

          <div className="login-form">
            <InputField label="ニックネーム" value="" onChange={(e) => setRegNickname(e.target.value)} />
            <InputField label="生年月日" value="" onChange={(e) => setRegBirthDate(e.target.value)} placeholder="1995/1/1" />
            <InputField label="来場者人数" type="number" value="" onChange={(e) => setRegPartySize(e.target.value)} />
            <div className="login-submitrow">
              <Button type="button" label="登録する" />
            </div>
          </div>

          <div className="login-divider">
            <div className="login-line"> </div>
              <Label text="アカウントをお持ちの方" fontSize={14} color="#666" />
            <div className="login-line"> </div>
          </div>

          <div className="login-form">
            <InputField label="ニックネーム" value="" onChange={(e) => setLoginNickname(e.target.value)} />
            <InputField label="生年月日" value="" onChange={(e) => setLoginBirthDate(e.target.value)} placeholder="1995/1/1" />
            <div className="login-submitrow">
              <Button type="button" label="ログイン" />
            </div>
          </div>  
        </div>
      </div>
      <Footer/>
    </>
  );
};

export default LoginPage;