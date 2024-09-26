import React from 'react'
import { LuSearch } from "react-icons/lu";

import { cn } from '@/utils/lib/cn'

type SearchBarProps = {
  placeholder?: string
  className?: string
}

const SearchBar = ({ placeholder, className }: SearchBarProps) => {
  return (
    <div className={cn(
      'flex items-center',
      'w-2/5 px-4 rounded-lg bg-slate-100'
    )}>
      <form className={cn('flex-grow')}>
        <input
          type="text"
          className={cn(
            'w-full bg-transparent focus:outline-none',
            'text-gray-900 font-semibold',
            className
          )}
          placeholder={placeholder}
        />
      </form>

      <LuSearch className={cn('h-6 w-6 flex-shrink-0 ml-4', 'text-gray-900')}/>
    </div>
  )
}

export default SearchBar