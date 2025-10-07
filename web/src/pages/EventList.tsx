import React, {useState} from "react"

import Label from "../components/atomic/Label"
import TicketCard from "../components/atomic/TicketCard"
import Button from "../components/atomic/Button"
import Header from "../components/Header"
import Footer from "../components/Footer"

import "./EventList.css"

const TicketList: React.FC = () => {
  const [regNickname, setRegNickname] = useState("");
  const [regBirthDate, setRegBirthDate] = useState("");
  const [regPartySize, setRegPartySize] = useState("");

  const [loginNickname, setLoginNickname] = useState("");
  const [loginBirthDate, setLoginBirthDate] = useState("");

  return (
    <>
      <Header title="企画一覧" />
      <div className="EventList-container">
        <div className="EventList-frame">
          <div style={{ textAlign: "center", lineHeight: 1.35, marginBottom: 12 }}>
            <Label text="なっちゃん  さん" fontSize={20} color="#222" /> <br></br>
            <Label text="来場者人数 : 1人" fontSize={14} color="#666" /> <br></br> <br></br><br></br>
            <Label text="企画一覧" fontSize={20} color="#222" /> <br></br>
            <Label text="9/14(土)" fontSize={14} color="#666" /> <br></br> <br></br>
          </div>
        </div>
    </div>
    <Footer/>
    </>
  );
};

export default TicketList;