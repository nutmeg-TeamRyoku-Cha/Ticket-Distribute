import React, { useState } from "react"
import { useNavigate } from "react-router-dom"

import Label from "../components/atomic/Label"
import InputField from "../components/atomic/InputField"
import Button from "../components/atomic/Button"
import Header from "../components/Header"

import "./login.css"

const API_BASE = import.meta.env.VITE_API_BASE_URL ?? "http://localhost:8080";

const LoginPage: React.FC = () => {
  const navigate = useNavigate();

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
      setRegNickname("");
      setRegBirthDate("");
      setRegPartySize("");
    } catch (e: any) {
      setMessage(`登録に失敗しました: ${e.message ?? e}`);
    } finally {
      setSubmitting(false);
    }
  };

  const handleLogin = async () => {
    if (!loginNickname || !loginBirthDate) {
      setMessage("ニックネームと生年月日を入力してください");
      return;
    }
    setSubmitting(true);
    setMessage(null);
    try {
      // 1) visitor_id を解決
      const r1 = await fetch(`${API_BASE}/visitors/resolve`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ nickname: loginNickname, birth_date: loginBirthDate }),
      });
      if (!r1.ok) {
        const t = await r1.text();
        throw new Error(t || `HTTP ${r1.status}`);
      }
      const { visitor_id } = await r1.json() as { visitor_id: number };

      // 2) セッション作成
      const r2 = await fetch(`${API_BASE}/sessions`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ visitor_id }),
        //credentials: "include",
      });
      if (!r2.ok) {
        const t = await r2.text();
        throw new Error(t || `HTTP ${r2.status}`);
      }
      const session = await r2.json() as { token?: string; expires_at?: string };
      if (session.token) localStorage.setItem("sessionToken", session.token);
      
      setMessage("ログインしました");
      navigate("/TicketList", { replace: true });
      setLoginNickname("");
      setLoginBirthDate("");
    } catch (e: any) {
      setMessage(`ログインに失敗しました: ${e.message ?? e}`);
    } finally {
      setSubmitting(false);
    }
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
            <InputField label="ニックネーム" value={regNickname} onChange={(e) => setRegNickname(e.target.value)} />
            <InputField label="生年月日" type="date" value={regBirthDate} onChange={(e) => setRegBirthDate(e.target.value)} placeholder="1995-01-01" />
            <InputField label="来場者人数" type="number" value={regPartySize} onChange={(e) => setRegPartySize(e.target.value)} />
            <div className="login-submitrow">
              <Button type="button" label={submitting ? "送信中..." : "登録する"} onClick={handleRegister} disabled={submitting} />
            </div>
          </div>

          <div className="login-divider">
            <div className="login-line"> </div>
              <Label text="アカウントをお持ちの方" fontSize={14} color="#666" />
            <div className="login-line"> </div>
          </div>

          <div className="login-form">
            <InputField label="ニックネーム" value={loginNickname} onChange={(e) => setLoginNickname(e.target.value)} placeholder="例: Taro" />
            <InputField label="生年月日" type="date" value={loginBirthDate} onChange={(e) => setLoginBirthDate(e.target.value)} placeholder="1995-01-01" />
            <div className="login-submitrow">
              <Button type="button" label={submitting ? "送信中..." : "ログイン"} onClick={handleLogin} disabled={submitting} />
            </div>
          </div>
          {message && (
            <div style={{ marginTop: 12, textAlign: "center", color: "red" }}>
              {message}
            </div>
          )} 
        </div>
      </div>
    </>
  );
};

export default LoginPage;