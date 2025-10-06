// src/components/Atoms/Label.tsx

import React from "react";

type LabelProps = {
  text: string;
  htmlFor?: string;
  required?: boolean;
  className?: string;
  fontSize?: string | number;   // ← 追加
  color?: string;               // ← 追加
};

const Label: React.FC<LabelProps> = ({ text, htmlFor, required = false, className = "", fontSize, color }) => {
  return (
    <label
      htmlFor={htmlFor}
      className={`label ${className}`}
      style={{
        fontSize: fontSize ?? undefined,
        color: color ?? undefined,
      }}
    >
      {text}
      {required && <span className="text-red-500 ml-1">*</span>}
    </label>
  );
};

export default Label;