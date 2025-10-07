import React, { useState, useEffect } from "react";

import Label from "../components/atomic/Label";
import TicketCard from "../components/atomic/TicketCard";
import Header from "../components/Header";
import Footer from "../components/Footer";

import "./TicketList.css";

// --- 型定義 (変更なし) ---
interface Building {
  building_id: number;
  building_name: string;
  latitude?: string;
  longitude?: string;
}

interface Project {
  project_id: number;
  project_name: string;
  building: Building;
  requires_ticket: boolean;
  start_time: string;
  end_time?: string;
}

interface Ticket {
  ticket_id: number;
  visitor_id: number;
  project: Project;
  status: string;
  entry_start_time: string | null;
  entry_end_time: string | null;
}

interface Visitor {
  visitor_id: number;
  nickname: string;
  birth_date: string;
  party_size: number;
}

// --- ヘルパー関数 (修正あり) ---
const formatDate = (dateString: string | null): string => {
  if (!dateString) return "日付未定";
  // "YYYY-MM-DDTHH:..." または "YYYY-MM-DD" の形式に対応
  const date = new Date(dateString);
  const month = date.getMonth() + 1;
  const day = date.getDate();
  const dayOfWeek = ["日", "月", "火", "水", "木", "金", "土"][date.getDay()];
  return `${month}/${day}(${dayOfWeek})`;
};

const formatTime = (isoString: string | null): string => {
  if (!isoString) return ""; // nullの場合は空文字を返す
  const date = new Date(isoString);
  const hours = date.getUTCHours().toString().padStart(2, "0");
  const minutes = date.getUTCMinutes().toString().padStart(2, "0");
  return `${hours}:${minutes}`;
};

// --- ここから追加 ---
// チケットを日付でグループ化するための型
interface GroupedTickets {
  [date: string]: Ticket[];
}
// --- ここまで追加 ---


const TicketList: React.FC = () => {
  const [visitor, setVisitor] = useState<Visitor | null>(null);
  // --- ここから修正：グループ化されたチケットを保持するState ---
  const [groupedTickets, setGroupedTickets] = useState<GroupedTickets>({});
  // --- ここまで修正 ---
  const [error, setError] = useState<string | null>(null);

  const visitorId = 1;

  useEffect(() => {
    const fetchVisitor = async () => {
      try {
        const res = await fetch(`http://localhost:8080/visitors/${visitorId}`);
        if (!res.ok) throw new Error(`[${res.status}] 来場者情報の取得に失敗`);
        const data: Visitor = await res.json();
        setVisitor(data);
      } catch (err: any) {
        setError(err.message);
        console.error(err);
      }
    };

    const fetchAndProcessTickets = async () => {
      try {
        const res = await fetch(`http://localhost:8080/tickets/visitor/${visitorId}`);
        if (!res.ok) throw new Error(`[${res.status}] 整理券情報の取得に失敗`);
        const data: Ticket[] = await res.json();
        
        // --- ここから修正：チケットの処理ロジック ---
        const processedTickets = data
          // 1. "issued" (有効)なチケットのみフィルタリング
          .filter(ticket => ticket.status === "issued")
          // 2. 日付ごとにグループ化
          .reduce((acc: GroupedTickets, ticket) => {
            // entry_start_timeがなければ、プロジェクトのstart_timeを基準にする
            const dateKey = (ticket.entry_start_time || ticket.project.start_time).split('T')[0];
            if (!acc[dateKey]) {
              acc[dateKey] = [];
            }
            acc[dateKey].push(ticket);
            return acc;
          }, {});

        setGroupedTickets(processedTickets);
        // --- ここまで修正 ---

      } catch (err: any) {
        setError(err.message);
        console.error(err);
      }
    };

    fetchVisitor();
    fetchAndProcessTickets();
  }, [visitorId]);

  if (error) {
    return (
        <>
            <Header title="エラー" />
            <div className="TicketList-container" style={{ color: "red", textAlign: "center" }}>
                <p>データの読み込みに失敗しました。</p>
                <p>{error}</p>
            </div>
            <Footer />
        </>
    )
  }

  return (
    <>
      <Header title="整理券一覧" />
      <div className="TicketList-container">
        <div className="TicketList-frame">
          <div style={{ textAlign: "center", lineHeight: 1.35, marginBottom: 24 }}>
            <Label text={`${visitor?.nickname ?? "..."} さん`} fontSize={20} color="#222" /> <br /><br />
            <Label text={`来場者人数 : ${visitor?.party_size ?? "..."}人`} fontSize={14} color="#666" />
          </div>
          
          {/* --- ここから修正：グループ化されたチケットの描画ロジック --- */}
          <div className="Ticket-list-wrapper">
            {Object.keys(groupedTickets).length > 0 ? (
              // 日付のキーでソートしてから表示
              Object.keys(groupedTickets).sort().map(date => (
                <div key={date} className="Ticket-date-group">
                  <div style={{ marginBottom: 12 }}>
                    <Label text={formatDate(date)} fontSize={20} color="#222" />
                  </div>
                  <div className="Ticket-form">
                    {groupedTickets[date].map(ticket => {
                      // 時間指定があればそれを、なければプロジェクト時間を表示
                      const startTime = formatTime(ticket.entry_start_time || ticket.project.start_time);
                      const endTime = formatTime(ticket.entry_end_time || ticket.project.end_time);
                      const timeRange = startTime ? `${startTime}～${endTime}` : "イベント時間内有効";

                      return (
                        <TicketCard
                          key={ticket.ticket_id}
                          title={ticket.project.project_name}
                          time={timeRange}
                          location={ticket.project.building.building_name}
                          onClick={() => alert(`チケットID: ${ticket.ticket_id}がクリックされました！`)}
                        />
                      );
                    })}
                  </div>
                </div>
              ))
            ) : (
              <p style={{textAlign: "center", color: "#666"}}>有効な整理券はありません。</p>
            )}
          </div>
          {/* --- ここまで修正 --- */}
        </div>
      </div>
      <Footer/>
    </>
  );
};

export default TicketList;