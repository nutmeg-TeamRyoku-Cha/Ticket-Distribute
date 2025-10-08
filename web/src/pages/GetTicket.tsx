import React, { useEffect, useState } from "react";
import { useLocation, useNavigate } from "react-router-dom";
import Label from "../components/atomic/Label";
import Button from "../components/atomic/Button";
import Header from "../components/Header";
import Footer from "../components/Footer";
import "./GetTicket.css";

type ProjectRes = {
  project_id: number;
  project_name: string;
  building_id: number;
  requires_ticket: boolean;
  remaining_tickets: number;
  start_time: string;
  end_time: string;
};

type VisitorRes = {
  visitor_id: number;
  nickname: string;
  birth_date: string;
  party_size: number;
};

type BuildingRes = {
  building_id: number;
  building_name: string;
  latitude?: string;
  longitude?: string;
};

const API_BASE = import.meta.env.VITE_API_BASE_URL ?? "http://localhost:8080";
const visitorId = 1;

const NINE_HOURS_MS = 9 * 60 * 60 * 1000;
const toMinus9h = (iso: string) => new Date(new Date(iso).getTime() - NINE_HOURS_MS);

const formatDateJst = (iso: string) => {
  const d = toMinus9h(iso);
  const m = d.getMonth() + 1;
  const day = d.getDate();
  const w = "日月火水木金土"[d.getDay()];
  return `${m}/${day}(${w})`;
};

const formatTimeJst = (iso: string) => {
  const d = toMinus9h(iso);
  const hh = d.getHours().toString().padStart(2, "0");
  const mm = d.getMinutes().toString().padStart(2, "0");
  return `${hh}:${mm}`;
};

const GetTicket: React.FC = () => {
  const { state } = useLocation() as { state?: { projectId?: number } };
  const projectId = state?.projectId;
  const nav = useNavigate();

  const [project, setProject] = useState<ProjectRes | null>(null);
  const [visitor, setVisitor] = useState<VisitorRes | null>(null);
  const [building, setBuilding] = useState<BuildingRes | null>(null);
  const [error, setError] = useState<string | null>(null);
  const [submitting, setSubmitting] = useState(false);
  const [ticketId, setTicketId] = useState<number | null>(null);

  useEffect(() => {
    if (!projectId) {
      setError("プロジェクトIDが指定されていません");
      return;
    }
    const run = async () => {
      try {
        const [p, v] = await Promise.all([
          fetch(`${API_BASE}/projects/${projectId}`),
          fetch(`${API_BASE}/visitors/${visitorId}`),
        ]);
        if (!p.ok) throw new Error(`[${p.status}] プロジェクト取得失敗`);
        if (!v.ok) throw new Error(`[${v.status}] 来場者取得失敗`);

        const pj: ProjectRes = await p.json();
        const vs: VisitorRes = await v.json();
        setProject(pj);
        setVisitor(vs);

        const b = await fetch(`${API_BASE}/buildings/${pj.building_id}`);
        setBuilding(b.ok ? await b.json() : null);
      } catch (e: any) {
        setError(e.message ?? "取得に失敗しました");
      }
    };
    run();
  }, [projectId]);

  const handleAcquire = async () => {
    if (!projectId || !visitor || !project) return;

    const decreaseCount = Math.max(1, visitor.party_size);
    if (project.remaining_tickets < decreaseCount) {
      setError(`残枚数不足（残り ${project.remaining_tickets} 枚 / 必要 ${decreaseCount} 枚）`);
      return;
    }

    try {
      setSubmitting(true);

      const patchRes = await fetch(`${API_BASE}/projects/${projectId}/remaining_tickets`, {
        method: "PATCH",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ decrease_tickets: decreaseCount }),
      });
      if (!patchRes.ok) throw new Error(`[${patchRes.status}] 残チケット更新に失敗`);

      const postRes = await fetch(`${API_BASE}/tickets`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({
          visitor_id: visitor.visitor_id,
          project_id: project.project_id,
          entry_start_time: project.start_time,
          entry_end_time: project.end_time,
        }),
      });
      if (!postRes.ok) throw new Error(`[${postRes.status}] チケット発行に失敗`);
      const { ticket_id } = await postRes.json();

      setTicketId(Number(ticket_id));
      setProject(prev =>
        prev ? { ...prev, remaining_tickets: Math.max(0, prev.remaining_tickets - decreaseCount) } : prev
      );
    } catch (e: any) {
      setError(e.message ?? "整理券獲得に失敗しました");
    } finally {
      setSubmitting(false);
    }
  };

  if (error) {
    return (
      <>
        <Header title="エラー" />
        <div className="TicketList-container">
          <div className="TicketList-frame" style={{ color: "red", textAlign: "center" }}>
            <p>{error}</p>
          </div>
        </div>
        <Footer />
      </>
    );
  }

  // 成功サンクス画面
  if (ticketId && project && visitor) {
    return (
      <>
        <Header title="整理券を発行しました" />
        <div className="TicketList-container">
          <div className="TicketList-frame" style={{ textAlign: "center" }}>
            <Label text="登録が完了しました！" fontSize={22} color="#222" />
            <br /><br />
            <Label text={`イベント : ${project.project_name}`} fontSize={16} color="#666" />
            <br />
            <Label
              text={`日時 : ${formatDateJst(project.start_time)} ${formatTimeJst(project.start_time)}～${formatTimeJst(project.end_time)}`}
              fontSize={16}
              color="#666"
            />
            <br />
            <Label
              text={`場所 : ${building?.building_name ?? `建物ID ${project.building_id}`}`}
              fontSize={16}
              color="#666"
            />
            <br /><br />
            <div className="GetTicket-Button">
              <Button
                type="button"
                label="整理券一覧へ"
                onClick={() => nav("/ticket", { state: { projectId: project.project_id } })}
              />
              <Button
                type="button"
                label="イベントに戻る"
                onClick={() => nav(-1)}
              />
            </div>
          </div>
        </div>
        <Footer />
      </>
    );
  }

  // 通常画面
  return (
    <>
      <Header title="整理券獲得" />
      <div className="TicketList-container">
        <div className="TicketList-frame">
          <div style={{ textAlign: "center", lineHeight: 1.35, marginBottom: 12 }}>
            <Label text={`${visitor?.nickname ?? "..." } さん`} fontSize={20} color="#222" /> <br />
            <Label text={`来場者人数 : ${visitor?.party_size ?? "..."}人`} fontSize={14} color="#666" /> <br /><br />
          </div>

          {project && (
            <>
              <div className="Ticket-Detail">
                <Label text={project.project_name} fontSize={18} color="#222" />
                <Label text={formatDateJst(project.start_time)} fontSize={14} color="#666" />
                <Label
                  text={`時間 : ${formatTimeJst(project.start_time)}～${formatTimeJst(project.end_time)}`}
                  fontSize={14}
                  color="#666"
                />
                <Label
                  text={`場所 : ${building?.building_name ?? `建物ID ${project.building_id}`}`}
                  fontSize={14}
                  color="#666"
                />
                <br /><br />
                <Label text={`残り枚数 : ${project.remaining_tickets}`} fontSize={14} color="#666" />
              </div>

              <div className="Ticket-Member">
                <Label text={`登録人数 : ${visitor?.party_size ?? "..." } 名`} fontSize={18} color="#222" /> <br /><br />
              </div>
            </>
          )}

          <div>
            <div className="GetTicket-Button">
              <Button
                type="button"
                label={submitting ? "処理中..." : "整理券獲得"}
                onClick={handleAcquire}
                disabled={submitting || !project || !visitor || project.remaining_tickets <= 0}
              />
              <Button type="button" label={"戻る　　　"} onClick={() => nav(-1)} />
              <br /><br />
            </div>
          </div>
        </div>
      </div>
      <Footer />
    </>
  );
};

export default GetTicket;
