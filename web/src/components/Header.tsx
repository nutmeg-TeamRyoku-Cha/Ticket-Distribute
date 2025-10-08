import React from 'react';
import './Header.css';
import Label from './atomic/Label';
import LogoIcon from './atomic/LogoIcon';

type HeaderProps = {
  title?: string;
};

const Header: React.FC<HeaderProps> = ({ title = '整理券一覧' }) => {
  return (
    <header className="header">
      <div className="header-content">
        <Label text={title} fontSize={18}/>
        <LogoIcon icon="/images/logo-icon.png" alt="45thロゴ" variant="logo" />
      </div>
      <div className="header-line" />
    </header>
  );
};

export default Header;
