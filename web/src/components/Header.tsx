import React from 'react';
import './Header.css';
import Label from './atomic/Label';      // ← テキストコンポーネント
import LogoIcon from './atomic/LogoIcon';      // ← アイコンコンポーネント

const Header = () => {
  return (
    <header className="header">
      <div className="header-content">
        <Label text="整理券一覧" variant="header" />
        <LogoIcon icon="/images/logo-icon.png" alt="45thロゴ" variant="logo" />
      </div>
      <div className="header-line" />
    </header>
  );
};

export default Header;