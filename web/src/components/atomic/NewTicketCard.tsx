import React from 'react';
import './NewTicketCard.css';

// propsの型定義
type NewTicketCardProps = {
  title: string;
  time: string;
  location: string;
  isUsed: boolean; // 使用済みかどうかを判定するpropを追加
  children: React.ReactNode; // ボタンなどの子要素を受け取る
};

const NewTicketCard: React.FC<NewTicketCardProps> = ({ title, time, location, isUsed, children }) => {
  // isUsedの状態に応じてCSSクラスを動的に変更
  const cardClassName = `new-ticket-card ${isUsed ? 'used' : ''}`;

  return (
    <div className={cardClassName}>
      {/* --- チケット情報表示エリア --- */}
      <div className="new-ticket-info">
        <div className="new-ticket-title">{title}</div>
        <div className="new-ticket-time">{time}</div>
        <div className="new-ticket-location">{location}</div>
      </div>

      {/* --- ボタン表示エリア --- */}
      <div className="new-ticket-action">
        {children}
      </div>
    </div>
  );
};

export default NewTicketCard;