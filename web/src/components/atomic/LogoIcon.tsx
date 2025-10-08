import React from 'react';
import './LogoIcon.css';

type IconProps = {
  icon: string;
  alt?: string;
  onClick?: () => void;
  variant?: 'logo' | 'ticket'| 'event'| 'setting'| 'default';
  type?: 'button' | 'submit' | 'reset';
};

const Icon: React.FC<IconProps> = ({ icon, alt = '', onClick = () => {}, variant = 'default', type = 'button' }) => {
  const getClassName = () => {
    switch (variant) {
      case 'logo':
        return 'icon icon-logo';

      case 'ticket':
        return 'icon icon-ticket';
      case 'event':
        return 'icon icon-event';
      case 'setting':
        return 'icon icon-setting';
      default:
        return 'icon';
    }
  };

  return (
    <button type={type} onClick={onClick} className={getClassName()}>
      <img src={icon} alt={alt} className="icon-logo" />
    </button>
  );
};

export default Icon;