import React from "react";
import "./InputField.css"

type InputFieldProps = {
  label: string;
  type?: string;
  placeholder?: string;
  value: string | number;
  onChange: (e: React.ChangeEvent<HTMLInputElement>) => void;
  required?: boolean;
  error?: string;
  className?: string;       // ← ラッパーへ
  inputClassName?: string;  // ← inputへ
  size?: "sm" | "md" | "lg";
};

const InputField: React.FC<InputFieldProps> = ({
  label,
  type = "text",
  placeholder = "",
  value,
  onChange,
  required = false,
  error,
  className = "",
  inputClassName = "",
  size = "md",
}) => {
  const sizeMap = { sm: "h-9", md: "h-11", lg: "h-12" };

  return (
    <div className={`flex flex-col ${className}`}>
      <label className="mb-2 text-sm font-medium text-gray-700">
        {label}
        {required && <span className="text-red-500 ml-1">*</span>}
      </label>
      <input
        type={type}
        placeholder={placeholder}
        value={value}
        onChange={onChange}
        className={`nf-input ${sizeMap[size]} ${error ? "error" : ""} ${inputClassName}`}
        aria-invalid={!!error}
        aria-describedby={error ? `${label}-error` : undefined}
      />
      {error && (
        <span id={`${label}-error`} className="error-text mt-1">
          {error}
        </span>
      )}
    </div>
  );
};

export default InputField;
