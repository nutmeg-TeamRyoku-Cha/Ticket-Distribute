import React from 'react';
import './Footer.css';
import Label from './atomic/Label';
import LogoIcon from './atomic/LogoIcon';
const Footer = () => {
  return (
    <footer className="footer">
      <div className="footer-item">
        <LogoIcon icon="/images/event-icon.png" alt="イベントアイコン" variant="footer" />
        <Label text="EVENT" variant="footer" color='#875318'/>
      </div>
      <div className="footer-item">
        <LogoIcon icon="/images/ticket-icon.png" alt="チケットアイコン" variant="footer" />
        <Label text="TICKET" variant="footer" color='#875318'/>
      </div>
      <div className="footer-item">
        <LogoIcon icon="/images/setting-icon.png" alt="設定アイコン" variant="footer" />
        <Label text="SETTING" variant="footer" color='#875318'/>
      </div>
    </footer>
  );
};
export default Footer;