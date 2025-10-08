import React, { useState, useEffect } from "react";
import Label from "../components/atomic/Label";
import NewTicketCard from "../components/atomic/NewTicketCard"; //ここがNewTicketCardに変わった
import Header from "../components/Header";
import Footer from "../components/Footer";
import "./TicketList.css";

// --- 型定義 ---
interface Building { building_id: number; building_name: string; latitude?: string; longitude?: string; }
interface Project { project_id: number; project_name: string; building: Building; requires_ticket: boolean; start_time: string; end_time?: string; }
interface Ticket { ticket_id: number; visitor_id: number; project: Project; status: string; entry_start_time: string | null; entry_end_time: string | null; }
interface Visitor { visitor_id: number; nickname: string; birth_date: string; party_size: number; }
interface GroupedTickets { [date: string]: Ticket[]; }

// --- ヘルパー関数 ---
const formatDate = (dateString: string | null): string => { if (!dateString) return "日付未定"; const date = new Date(dateString); const month = date.getMonth() + 1; const day = date.getDate(); const dayOfWeek = ["日", "月", "火", "水", "木", "金", "土"][date.getDay()]; return `${month}/${day}(${dayOfWeek})`; };
const formatTime = (isoString?: string | null): string => { if (!isoString) return ""; const date = new Date(isoString); const hours = date.getUTCHours().toString().padStart(2, "0"); const minutes = date.getUTCMinutes().toString().padStart(2, "0"); return `${hours}:${minutes}`; };

const API_BASE = import.meta.env.VITE_API_BASE_URL ?? "/api";

const TicketList: React.FC = () => {
  const [visitor, setVisitor] = useState<Visitor | null>(null);
  const [groupedTickets, setGroupedTickets] = useState<GroupedTickets>({});
  const [error, setError] = useState<string | null>(null);
  const visitorId = 1; //ここがvisitor１限定です

  const fetchAndProcessTickets = async () => {
    try {
      const res = await fetch(`${API_BASE}/tickets/visitor/${visitorId}`);
      if (!res.ok) throw new Error(`[${res.status}] 整理券情報の取得に失敗`);
      const data: Ticket[] = await res.json();
      const processedTickets = data.reduce((acc: GroupedTickets, ticket) => {
          const dateKey = (ticket.entry_start_time || ticket.project.start_time).split('T')[0];
          if (!acc[dateKey]) { acc[dateKey] = []; }
          acc[dateKey].push(ticket);
          return acc;
        }, {});
      setGroupedTickets(processedTickets);
    } catch (err: any) {
      setError(err.message);
      console.error(err);
    }
  };

  useEffect(() => {
    const fetchVisitor = async () => {
      try {
        const res = await fetch(`${API_BASE}/visitors/${visitorId}`);
        if (!res.ok) throw new Error(`[${res.status}] 来場者情報の取得に失敗`);
        const data: Visitor = await res.json();
        setVisitor(data);
      } catch (err: any) {
        setError(err.message);
      }
    };
    fetchVisitor();
    fetchAndProcessTickets();
  }, [visitorId]);

  const handleUseTicket = async (ticketId: number) => {
    try {
      const res = await fetch(`${API_BASE}/tickets/${ticketId}/status`, {
        method: 'PATCH',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ status: 'used' }),
      });
      if (!res.ok) { throw new Error(`[${res.status}] チケットの状態更新に失敗`); }
      const newGroupedTickets = { ...groupedTickets };
      for (const date in newGroupedTickets) {
        const ticketIndex = newGroupedTickets[date].findIndex(t => t.ticket_id === ticketId);
        if (ticketIndex !== -1) {
          newGroupedTickets[date][ticketIndex].status = 'used';
          break;
        }
      }
      setGroupedTickets(newGroupedTickets);
    } catch (err: any) {
      setError(err.message);
      console.error(err);
    }
  };

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
          <div style={{ textAlign: "center", lineHeight: 1.35, marginBottom: 36 }}>
            <Label text={`${visitor?.nickname ?? "..."} さん`} fontSize={20} color="#222" /> <br /><br />
            <Label text={`来場者人数 : ${visitor?.party_size ?? "..."}人`} fontSize={14} color="#666" />
          </div>
          <div style={{ textAlign: "center", lineHeight: 1.35, marginBottom: 12 }}>
            <Label text="お手持ちの整理券一覧" fontSize={20} color="#222" />
          </div>
          <div className="Ticket-list-wrapper">
            {Object.keys(groupedTickets).length > 0 ? (
              Object.keys(groupedTickets).sort().map(date => (
                <div key={date} className="Ticket-date-group">
                  <div style={{ marginBottom: 12 }}>
                    <Label text={formatDate(date)} fontSize={20} color="#222" />
                  </div>
                  <div className="Ticket-form">
                    {groupedTickets[date].map(ticket => {
                      const timeRange = `${formatTime(ticket.entry_start_time || ticket.project.start_time)}～${formatTime(ticket.entry_end_time || ticket.project.end_time)}`;
                      const isUsed = ticket.status === 'used';
                      return (
                        <NewTicketCard
                          key={ticket.ticket_id}
                          title={ticket.project.project_name}
                          time={timeRange}
                          location={ticket.project.building.building_name}
                          isUsed={isUsed}
                        >
                          <button
                            className="use-ticket-button"
                            onClick={() => handleUseTicket(ticket.ticket_id)}
                            disabled={isUsed}
                          >
                            {isUsed ? 'used' : 'use'}
                          </button>
                        </NewTicketCard>
                      );
                    })}
                  </div>
                </div>
              ))
            ) : (
              <p style={{textAlign: "center", color: "#666"}}>有効な整理券はありません。</p>
            )}
          </div>
        </div>
      </div>
      <Footer/>
    </>
  );
};

export default TicketList;