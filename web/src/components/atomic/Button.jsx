import React from 'react';
import './Button.css';

const Button = ({ label, onClick, type = 'button', variant = 'primary', icon }) => {
  const getClassName = () => {
    switch (variant) {
      case 'registration':
        return 'button button-registration';
      case 'login':
        return 'button button-login';
      case 'ticket':
        return 'button button-ticket';
      case 'event':
        return 'button button-event';
      case 'setting':
        return 'button button-setting';
      default:
        return 'button';
    }
  };

  return (
    <button type={type} onClick={onClick} className={getClassName()}>
      {icon && <img src={icon} alt="アイコン" className="button-icon" />}
      {label}
    </button>
  );
};

export default Button;