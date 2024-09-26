import Link from 'next/link'
import React from 'react'

import { RiArrowDropDownLine } from "react-icons/ri";
import { FaSignOutAlt } from "react-icons/fa";

import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu"

import { cn } from '@/utils/lib/cn'

type ProfileBarProps = {
  isLogged: boolean
  className?: string
}

const user = {
  username: 'Suboq-kun',
};

const ProfileBar = ({ isLogged, className }: ProfileBarProps) => {
  if (isLogged) {
    return (
      <div className={cn('flex justify-end items-center', className)}>
        <div className={cn(
          'h-10 w-10 rounded-full mr-3 bg-white',
          'hover:shadow-lg hover:shadow-black cursor-pointer '
        )}>
          <img
            src='/images/suboq-kun.jpg'
            className={cn('rounded-full')}
          />
        </div>

        <DropdownMenu>
          <DropdownMenuTrigger>
            <div className={cn('flex justify-center items-center hover:text-orange-200')}>
              <div className={cn('font-bold text-base font-poppins')}>
                {user.username}
              </div>

              <RiArrowDropDownLine size={36} />
              </div>
          </DropdownMenuTrigger>

          <DropdownMenuContent>
            <DropdownMenuLabel className={cn('font-semibold font-poppins')}>Your Account</DropdownMenuLabel>
            <DropdownMenuSeparator />

            <DropdownMenuItem className={cn('font-poppins')}>
              <Link href='/'>
                Profile
              </Link>
            </DropdownMenuItem>

            <DropdownMenuItem className={cn('font-poppins')}>
              <Link href='/'>
                Watch List
              </Link>
            </DropdownMenuItem>

            <DropdownMenuItem className={cn('font-poppins')}>
              <button className='flex items-center justify-between w-full'>
                Sign Out

                <FaSignOutAlt />
              </button>
            </DropdownMenuItem>
            
          </DropdownMenuContent>
        </DropdownMenu>
      </div>
    )
  } else {
    return (
      <div className={cn('flex justify-end items-center space-x-3', className)}>
        <Link href="/">
          <div className={cn(
            'bg-white rounded-sm h-9 w-[100px]',
            'text-orange-400 font-bold font-poppins flex justify-center items-center',
            'hover:border-2 hover:border-orange-400 hover:shadow-lg hover:shadow-black'
          )}>
            Login
          </div>
        </Link>

        <Link href='/'>
          <div className={cn(
          'bg-orange-400 rounded-sm h-9 w-[100px]',
          'text-white font-bold font-poppins flex justify-center items-center',
          'hover:border-2 hover:border-white hover:shadow-lg hover:shadow-black'
          )}>
            Register
          </div>
        </Link>
      </div>
    )
  }
}

export default ProfileBar