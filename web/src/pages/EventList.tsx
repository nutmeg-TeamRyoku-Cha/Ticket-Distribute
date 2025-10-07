import React, { useEffect, useMemo, useState } from "react";

import Header from "../components/Header";
import Footer from "../components/Footer";
import Label from "../components/atomic/Label";
import NewTicketCard from "../components/atomic/NewTicketCard";

import "./EventList.css";

type BuildingRes = {
  building_id: number;
  building_name: string;
  latitude?: string;
  longitude?: string;
};

type ProjectResolvedRes = {
  project_id: number;
  project_name: string;
  building: BuildingRes;
  requires_ticket: boolean;
  start_time: string;
  end_time?: string;
};

const API_BASE = import.meta.env.VITE_API_BASE_URL ?? "http://localhost:8080";

const formatDateTimeJST = (iso: string) => {
  const d = new Date(iso);
  const date = d.toLocaleDateString("ja-JP");
  const time = d.toLocaleTimeString("ja-JP", { hour: "2-digit", minute: "2-digit" });
  return `${date} ${time}`;
};

const EventList: React.FC = () => {
  const [projects, setProjects] = useState<ProjectResolvedRes[]>([]);
  const [loading, setLoading] = useState(true);
  const [errorMessage, setErrorMessage] = useState<string | null>(null);

  useEffect(() => {
    const fetchProjects = async () => {
      try {
        const res = await fetch(`${API_BASE}/projects/resolved`, { credentials: "omit" });
        if (!res.ok) throw new Error(`${res.status} ${res.statusText}`);
        const data: ProjectResolvedRes[] = await res.json();
        setProjects(data);
      } catch (e: any) {
        setErrorMessage(e?.message ?? "fetch error");
      } finally {
        setLoading(false);
      }
    };
    fetchProjects();
  }, []);

  const projectsByDate = useMemo(() => {
    const map = new Map<string, ProjectResolvedRes[]>();
    for (const p of projects) {
      const ymd = new Date(p.start_time).toISOString().slice(0, 10);
      if (!map.has(ymd)) map.set(ymd, []);
      map.get(ymd)!.push(p);
    }
    for (const arr of map.values()) {
      arr.sort((a, b) => +new Date(a.start_time) - +new Date(b.start_time));
    }
    return Array.from(map.entries()).sort(([a], [b]) => (a < b ? -1 : 1));
  }, [projects]);

  return (
    <>
      <Header title="企画一覧" />
      <div className="EventList-container">
        <div className="EventList-frame">
          <div style={{ textAlign: "center", lineHeight: 1.35, marginBottom: 12 }}>
            <Label text="企画一覧" fontSize={20} color="#222" />
          </div>

          {loading && <div className="EventList-status">読み込み中...</div>}
          {errorMessage && <div className="EventList-status error">読み込み失敗: {errorMessage}</div>}
          {!loading && !errorMessage && projects.length === 0 && (
            <div className="EventList-status">企画がありません</div>
          )}

          {!loading && !errorMessage && projectsByDate.map(([ymd, list]) => (
            <section key={ymd} className="EventList-date-group">
              <Label
                text={new Date(ymd).toLocaleDateString("ja-JP")}
                fontSize={16}
                color="#666"
              />
              <div className="EventList-list-wrapper">
                {list.map((p) => {
                  const title = p.project_name;
                  const time = p.end_time
                    ? `${formatDateTimeJST(p.start_time)} 〜 ${formatDateTimeJST(p.end_time)}`
                    : `${formatDateTimeJST(p.start_time)}`;
                  const location = p.building.building_name;
                  return (
                    <NewTicketCard
                      key={p.project_id}
                      title={title}
                      time={time}
                      location={location}
                      isUsed={false}
                    >
                      {p.requires_ticket ? (
                        <span className="EventList-badge">整理券</span>
                      ) : (
                        <span style={{ fontSize: ".85rem", color: "#666" }}>整理券不要</span>
                      )}
                    </NewTicketCard>
                  );
                })}
              </div>
            </section>
          ))}
        </div>
      </div>
      <Footer />
    </>
  );
};

export default EventList;
