import React from 'react';
import './Button.css';

const Button = ({ label, onClick, type = 'button', variant = 'primary' }) => {
  const getClassName = () => {
    switch (variant) {
    case 'registration':
        return 'button button-registration';
    case 'login':
        return 'button button-login';
    default:
        return 'button';
}
  };

  return (
    <button type={type} onClick={onClick} className={getClassName()}>
      {label}
    </button>
  );
};

export default Button;