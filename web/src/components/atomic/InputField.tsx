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
    <div style={{ display: "50%", flexDirection: "column", marginBottom: "8px", fontSize: "20px" }}>
      <label style={{ marginBottom: "4px" }}>{label}<br></br></label>
      <input
        type={type}
        value={value}
        onChange={onChange}
        placeholder={placeholder}
        style={{ padding: "4px", width: "400px", height: "40px", fontSize: "20px"}}
      />
    </div>
  );
};

export default InputField;