import React from 'react';
import './TicketCard.css';
type TicketCardProps = {
  title: string;
  time: string;
  location: string;
  onClick: () => void;
};
const TicketCard: React.FC<TicketCardProps> = ({ title, time, location, onClick }) => {
  return (
    <button className="ticket-card" onClick={onClick}>
      <div className="ticket-title">{title}</div>
      <div className="ticket-time">{time}</div>
      <div className="ticket-location">{location}</div>
    </button>
  );
};
export default TicketCard;