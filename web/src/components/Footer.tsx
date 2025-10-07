import React from 'react'
import { useNavigate, useLocation } from 'react-router-dom';
import './Footer.css';
import Label from './atomic/Label';
import LogoIcon from './atomic/LogoIcon';

const Footer = () => {
  const navigate = useNavigate();
  const { pathname } = useLocation();
  const go = (path: string, replace = false) => navigate(path, { replace });
  const handleKey = (e: React.KeyboardEvent<HTMLDivElement>, path: string) => {
    if (e.key === 'Enter' || e.key === ' ') {
      e.preventDefault();
      go(path);
    }
  };
  return (
    <footer className="footer">
      <div
        className={`footer-item ${pathname === '/events' ? 'is-active' : ''}`}
        role="button"
        tabIndex={0}
        aria-current={pathname === '/events' ? 'page' : undefined}
        onClick={() => go('/events')}
        onKeyDown={(e) => handleKey(e, '/events')}
      >
        <LogoIcon icon="/images/event-icon.png" alt="イベントアイコン" variant="footer" />
        <Label text="EVENT" variant="footer" color='#875318'/>
      </div>
      <div
        className={`footer-item ${pathname === '/tickets' ? 'is-active' : ''}`}
        role="button"
        tabIndex={0}
        aria-current={pathname === '/tickets' ? 'page' : undefined}
        onClick={() => go('/tickets')}
        onKeyDown={(e) => handleKey(e, '/tickets')}
      >
        <LogoIcon icon="/images/ticket-icon.png" alt="チケットアイコン" variant="footer" />
        <Label text="TICKET" variant="footer" color='#875318'/>
      </div>
      <div
        className={`footer-item ${pathname === '/settings' ? 'is-active' : ''}`}
        role="button"
        tabIndex={0}
        aria-current={pathname === '/settings' ? 'page' : undefined}
        onClick={() => go('/settings')}
        onKeyDown={(e) => handleKey(e, '/settings')}
      >
        <LogoIcon icon="/images/setting-icon.png" alt="設定アイコン" variant="footer" />
        <Label text="SETTING" variant="footer" color='#875318'/>
      </div>
    </footer>
  );
};
export default Footer;