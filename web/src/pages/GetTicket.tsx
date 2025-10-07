import React, {useState} from "react"

import Label from "../components/atomic/Label"
import Button from "../components/atomic/Button"
import Header from "../components/Header"
import Footer from "../components/Footer"

import "./GetTicket.css"

const GetTicket: React.FC = () => {
  const [regNickname, setRegNickname] = useState("");
  const [regBirthDate, setRegBirthDate] = useState("");
  const [regPartySize, setRegPartySize] = useState("");

  const [loginNickname, setLoginNickname] = useState("");
  const [loginBirthDate, setLoginBirthDate] = useState("");

  return (
    <>
      <Header title="整理券一覧" />
      <div className="TicketList-container">
        <div className="TicketList-frame">
          <div style={{ textAlign: "center", lineHeight: 1.35, marginBottom: 12 }}>
            <Label text="なっちゃん  さん" fontSize={20} color="#222" /> <br></br>
            <Label text="来場者人数 : 1人" fontSize={14} color="#666" /> <br></br> <br></br><br></br>
            <Label text="整理券一覧" fontSize={20} color="#222" /> <br></br>
            <Label text="9/14(土)" fontSize={14} color="#666" /> <br></br> <br></br>
          </div>

          <div className="Ticket-Detail">
            <Label text="お化け屋敷" fontSize={18} color="#222" />
            <Label text="時間 : 10:00～16:00" fontSize={14} color="#666" />
            <Label text="場所 : 総合研究棟2F" fontSize={14} color="#666" /> <br></br> <br></br>
          </div>

          <div className="Ticket-Member">
            <Label text="登録人数 : 1 名" fontSize={18} color="#222" /> <br></br><br></br>
          </div>

          <div>
            <div className="GetTicket-Button">
              <Button type="button" label={"整理券獲得"}/>
              <Button type="button" label={"戻る　　　"}/><br></br> <br></br>
            </div>
          </div>
        </div>
      </div>
      <Footer/>
    </>
  );
};

export default GetTicket;