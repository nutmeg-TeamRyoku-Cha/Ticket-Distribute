import React from "react";
import "./Button.css";

export type ButtonVariant =
  | "primary"
  | "registration"
  | "login"
  | "chip" ;

export type ButtonProps = {
  label: React.ReactNode;
  onClick?: React.MouseEventHandler<HTMLButtonElement>;
  type?: "button" | "submit" | "reset";
  variant?: ButtonVariant;
  icon?: string | React.ReactNode;
  className?: string;
  disabled?: boolean;
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