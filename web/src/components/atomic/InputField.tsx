import React from "react";

type InputFieldProps = {
  label: string;               // ラベル表示用
  type?: string;               // "text" / "number" / "date" など
  placeholder?: string;        // プレースホルダー
  value: string | number;      // 現在値（制御コンポーネント）
  onChange: (e: React.ChangeEvent<HTMLInputElement>) => void; // 値変更ハンドラ
  required?: boolean;          // 必須入力フラグ
  error?: string;              // エラーメッセージ（任意）
};

const InputField: React.FC<InputFieldProps> = ({
  label,
  type = "text",
  placeholder = "",
  value,
  onChange,
  required = false,
  error,
}) => {
  return (
    <div className="flex flex-col mb-4">
      <label className="mb-1 text-sm font-medium text-gray-700">
        {label}
        {required && <span className="text-red-500 ml-1">*</span>}
      </label>

      <input
        type={type}
        placeholder={placeholder}
        value={value}
        onChange={onChange}
        className={`border rounded-lg px-3 py-2 text-sm focus:outline-none focus:ring-2 focus:ring-orange-400 ${
          error ? "border-red-500" : "border-gray-300"
        }`}
      />

      {error && <span className="text-xs text-red-500 mt-1">{error}</span>}
    </div>
  );
};

export default InputField;
