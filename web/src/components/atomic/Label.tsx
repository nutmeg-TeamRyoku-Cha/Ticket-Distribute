// src/components/Atoms/Label.tsx

import React from "react";

type LabelProps = {
  text: string;
  htmlFor?: string;
  className?: string;
};

const Label: React.FC<LabelProps> = ({ text, htmlFor, className }) => {
  return (
    <label htmlFor={htmlFor} className={className}>
      {text}
    </label>
  );
};

export default Label;