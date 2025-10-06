import React from "react";
import "./Button.css";

export type ButtonVariant =
  | "primary"
  | "registration"
  | "login"
  | "ticket"
  | "event"
  | "setting";

export type ButtonProps = {
  label: React.ReactNode;
  onClick?: React.MouseEventHandler<HTMLButtonElement>;
  type?: "button" | "submit" | "reset";
  variant?: ButtonVariant;
  icon?: string | React.ReactNode;
  className?: string;
};

const Button: React.FC<ButtonProps> = ({
  label,
  onClick,
  type = "button",
  variant = "primary",
  icon,
  className = "",
}) => {
  const getClassName = (): string => {
    switch (variant) {
      case "registration":
        return `button button-registration ${className}`;
      case "login":
        return `button button-login ${className}`;
      case "ticket":
        return `button button-ticket ${className}`;
      case "event":
        return `button button-event ${className}`;
      case "setting":
        return `button button-setting ${className}`;
      default:
        return `button button-primary ${className}`;
    }
  };

  return (
    <button type={type} onClick={onClick} className={getClassName()}>
      {icon &&
        (typeof icon === "string" ? (
          <img src={icon} alt="" className="button-icon" />
        ) : (
          icon
        ))}
      {label}
    </button>
  );
};

export default Button;