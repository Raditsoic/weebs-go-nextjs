import { cn } from '@/utils/lib/cn';
import React from 'react'

interface InputFieldProps {
  inputStyles?: string;
  placeholder?: string;
  variant: "form" | "search";
} 

const InputField = ({ inputStyles, placeholder, variant }: InputFieldProps ) => {
  return (
    <div className={cn()}>
      <input
        type="text"
        className={`${inputStyles}`}
        placeholder={`${placeholder}`}
      />
    </div>
  )
}

export default InputField