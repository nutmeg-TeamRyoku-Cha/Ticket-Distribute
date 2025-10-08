import React, { useEffect, useMemo, useState } from "react";
import { useNavigate } from "react-router-dom";

import Header from "../components/Header";
import Footer from "../components/Footer";
import Label from "../components/atomic/Label";
import NewTicketCard from "../components/atomic/NewTicketCard";

import "./EventList.css";
import "./TicketList.css";

type BuildingRes = {
  building_id: number;
  building_name: string;
  latitude?: string;
  longitude?: string;
};

export type ProjectResolvedRes = {
  project_id: number;
  project_name: string;
  building: BuildingRes;
  requires_ticket: boolean;
  start_time: string;
  end_time?: string;
};

type Visitor = {
  visitor_id: number;
  nickname: string;
  birth_date: string;
  party_size: number;
};

const API_BASE = import.meta.env.VITE_API_BASE_URL ?? "http://localhost:8080";

const formatDate = (dateString: string): string => {
  const date = new Date(dateString);
  const month = date.getMonth() + 1;
  const day = date.getDate();
  const dayOfWeek = ["日", "月", "火", "水", "木", "金", "土"][date.getDay()];
  return `${month}/${day}(${dayOfWeek})`;
};
const formatTime = (isoString: string | null | undefined): string => {
  if (!isoString) return "";
  const d = new Date(isoString);
  const h = d.getUTCHours().toString().padStart(2, "0");
  const m = d.getUTCMinutes().toString().padStart(2, "0");
  return `${h}:${m}`;
};

const EventList: React.FC = () => {
  const navigate = useNavigate();
  const [items, setItems] = useState<ProjectResolvedRes[]>([]);
  const [loading, setLoading] = useState(true);
  const [errorMessage, setErrorMessage] = useState<string | null>(null);
  const [visitor, setVisitor] = useState<Visitor | null>(null);
  const visitorId = 2;

  useEffect(() => {
    (async () => {
      try {
        const res = await fetch(`${API_BASE}/projects/resolved`, { credentials: "omit" });
        if (!res.ok) throw new Error(`${res.status} ${res.statusText}`);
        const json: ProjectResolvedRes[] = await res.json();
        setItems(json);
      } catch (e: any) {
        setErrorMessage(e?.message ?? "fetch error");
      } finally {
        setLoading(false);
      }
    })();
    (async () => {
      try {
        const res = await fetch(`${API_BASE}/visitors/${visitorId}`, { credentials: "omit" });
        if (!res.ok) return;
        const v: Visitor = await res.json();
        setVisitor(v);
      } catch {
        /* noop */
      }
    })();
  }, []);

  const grouped = useMemo(() => {
    const map = new Map<string, ProjectResolvedRes[]>();
    for (const p of items) {
      // TicketList と同じく ISO を文字列のまま分割して日付キー化
      const key = p.start_time.split("T")[0];
      if (!map.has(key)) map.set(key, []);
      map.get(key)!.push(p);
    }
    for (const arr of map.values()) {
      arr.sort((a, b) => +new Date(a.start_time) - +new Date(b.start_time));
    }
    return Array.from(map.entries()).sort(([a], [b]) => (a < b ? -1 : 1));
  }, [items]);

  return (
    <>
      <Header title="企画一覧" />
      <main className="EventList-container">
        <section className="EventList-frame">
          <div style={{ textAlign: "center", lineHeight: 1.35, marginBottom: 12 }}>
            <Label text={`${visitor?.nickname ?? "..."} さん`} fontSize={20} color="#222" /> <br /><br />
            <Label text={`来場者人数 : ${visitor?.party_size ?? "..."}人`} fontSize={14} color="#666" /> <br /><br />
          </div>
          <div style={{ textAlign: "center", lineHeight: 1.35, marginBottom: 12 }}>
            <Label text="企画一覧" fontSize={20} color="#222" />
          </div>

          {loading && <div className="EventList-status">読み込み中...</div>}
          {errorMessage && <div className="EventList-status error">読み込み失敗: {errorMessage}</div>}
          {!loading && !errorMessage && items.length === 0 && (
            <div className="EventList-status">企画がありません</div>
          )}

          {!loading && !errorMessage && grouped.map(([dateKey, list]) => (
            <div key={dateKey} className="EventList-date-group">
              <div style={{ marginBottom: 12 }}>
                <Label text={formatDate(dateKey)} fontSize={20} color="#000" />
              </div>
              <div className="EventList-list-wrapper">
                {list.map((p) => {
                  const title = p.project_name;
                  const time = `${formatTime(p.start_time)}～${formatTime(p.end_time)}`;
                  const location = p.building.building_name;

                  return (
                    <NewTicketCard
                      key={p.project_id}
                      title={title}
                      time={time}
                      location={location}
                      isUsed={false}
                      data-projectid={p.project_id}
                    >
                      <button
                        type="button"
                        className="use-ticket-button"
                        aria-label={`${title} の整理券を選ぶ`}
                        onClick={() => {
                          // URLは固定 /getticket。IDは state と sessionStorage に保存（リロード耐性）
                          sessionStorage.setItem("selectedProjectId", String(p.project_id));
                          navigate("/getticket", { state: { projectId: p.project_id } });
                        }}
                      >
                        {p.requires_ticket ? "get" : "detail"}
                      </button>
                    </NewTicketCard>
                  );
                })}
              </div>
            </div>
          ))}
        </section>
      </main>
      <Footer />
    </>
  );
};

export default EventList;
