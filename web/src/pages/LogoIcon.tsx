import React from 'react';
import LogoIconComponent from '../components/atomic/LogoIcon';

const images = () => {
  const handleLogo = () => {
    alert('25th-ロゴ');
  };

  return (
    <div>
      <LogoIconComponent
        icon="/images/logo-icon.png"
        alt="25周年ロゴ"
        variant="logo"
        onClick={handleLogo}
      />
    </div>
  );
};

export default images;