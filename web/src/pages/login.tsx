import React, {useState} from "react"

import Label from "../components/atomic/Label"
import InputField from "../components/atomic/InputField"
import Button from "../components/atomic/Button"
import Header from "../components/Header"

import "./login.css"

const LoginPage: React.FC = () => {
  const [regNickname, setRegNickname] = useState("");
  const [regBirthDate, setRegBirthDate] = useState("");
  const [regPartySize, setRegPartySize] = useState("");

  const [loginNickname, setLoginNickname] = useState("");
  const [loginBirthDate, setLoginBirthDate] = useState("");

  return (
    <>
      <Header title="ログイン2" />
      <div className="login-container">
        <div className="login-frame">
          <div style={{ textAlign: "center", lineHeight: 1.35, marginBottom: 12 }}>
            <Label text="ようこそ！技大祭へ！" fontSize={32} color="#222" /> <br></br> <br></br>
            <Label text="新規の方は来場者登録をお願いします" fontSize={16} color="#666" /> <br></br> <br></br>
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
              <Label text="アカウントをお持ちの方" fontSize={16} color="#666" />
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
    </>
  );
};

export default LoginPage;