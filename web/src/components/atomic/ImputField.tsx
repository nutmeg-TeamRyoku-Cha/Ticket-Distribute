// src/components/Atoms/InputField.tsx

import React from "react";
import Label from "./Label";

type InputFieldProps = {
  label: string;
  name: string;
  type: string;
  register: any; // react-hook-form の register() 呼び出し結果
  error?: { message?: string };
  placeholder?: string;
  className?: string;
  [key: string]: any; // その他のprops（min, maxなど）
};

const InputField: React.FC<InputFieldProps> = ({
  label,
  name,
  type,
  register,
  error,
  placeholder,
  className = "",
  ...rest
}) => {
  return (
    <div className={`mb-4 ${className}`}>
      <Label text={label} htmlFor={name} />
      <input
        id={name}
        name={name}
        type={type}
        placeholder={placeholder}
        {...register}
        {...rest}
        className="border border-gray-300 rounded px-3 py-2 w-full"
      />
      {error?.message && <p className="text-red-500 text-sm mt-1">{error.message}</p>}
    </div>
  );
};

export default InputField;