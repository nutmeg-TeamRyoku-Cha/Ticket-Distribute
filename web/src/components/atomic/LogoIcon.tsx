import React from 'react';
import './LogoIcon.css';

type IconProps = {
  icon: string;
  alt?: string;
  onClick: () => void;
  variant?: 'logo' | 'default';
  type?: 'button' | 'submit' | 'reset';
};

const Icon: React.FC<IconProps> = ({ icon, alt = '', onClick, variant = 'default', type = 'button' }) => {
  const getClassName = () => {
    switch (variant) {
      case 'logo':
        return 'icon icon-logo';
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