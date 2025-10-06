// src/components/Atoms/Label.tsx

import React from "react";

type LabelProps = {
  /** ラベルに表示するテキスト */
  text: string;
  /** 関連づける入力要素のid（for属性に対応） */
  htmlFor?: string;
  /** 必須項目のとき true */
  required?: boolean;
  /** カスタムクラス（任意） */
  className?: string;
};

const Label: React.FC<LabelProps> = ({ text, htmlFor, required = false, className }) => {
  return (
    <label
      htmlFor={htmlFor}
      className={`text-sm font-medium text-gray-700 ${className ?? ""}`}
    >
      {text}
      {required && <span className="text-red-500 ml-1">*</span>}
    </label>
  );
};

export default Label;