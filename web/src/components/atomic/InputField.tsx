import React from "react";

type InputFieldProps = {
  label: string;
  value: string;
  onChange: (e: React.ChangeEvent<HTMLInputElement>) => void;
  type?: string;
  placeholder?: string;
};

const InputField: React.FC<InputFieldProps> = ({
  label,
  value,
  onChange,
  type = "text",
  placeholder = "",
}) => {
  return (
    <div style={{ display: "flex", flexDirection: "column", marginBottom: "8px" }}>
      <label style={{ marginBottom: "4px", fontSize: "14px", fontWeight: 500 }}>{label}</label>
      <input
        type={type}
        value={value}
        onChange={onChange}
        placeholder={placeholder}
        style={{ width: "100%", height: "44px", padding: "10px 12px", fontSize: "16px", boxSizing: "border-box" }}
      />
    </div>
  );
};

export default InputField;