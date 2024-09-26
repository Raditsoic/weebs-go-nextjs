'use client'

import React from 'react'
import Link from 'next/link'
import { usePathname } from 'next/navigation'

import { cn } from '@/utils/lib/cn'

import ProfileBar from './ProfileBar'

const Navbar = () => {
  const path = usePathname()

  return (
    <div className={cn(
      'flex items-center justify-between px-5 md:px-16 py-8',
      'h-[100px]',
      'bg-gradient-to-r from-gray-900 to-slate-800'
    )}>
      <Link href='/' className={cn(
        'flex items-center space-x-3'
      )}>
        <img 
          src='/images/kikicat.jpg'
          alt='Weebs Logo' 
          className={cn('rounded-full h-10')}
        />
        
        <span className={cn('font-syncopate text-2xl font-bold')}>
          Weebs
        </span>
      </Link>

      <div className={cn('flex-grow flex justify-center')}>
        <div className={cn('space-x-7 max-md:hidden')}>
          <Link href='/' className={cn(
            'font-medium font-poppins text-lg hover:text-orange-200',
            path === '/' && 'text-orange-400 font-bold text-xl'
          )}>
            Home
          </Link>

          <Link href='/anime' className={cn(
            'font-medium font-poppins text-lg hover:text-orange-200',
            path === '/anime' && 'text-orange-400 font-bold text-xl'
          )}>
            Anime
          </Link>
          
          <Link href='/manga' className={cn(
            'font-medium font-poppins text-lg hover:text-orange-200',
            path === '/manga' && 'text-orange-400 font-bold text-xl'
          )}>
            Manga
          </Link>

          <Link href='/about' className={cn(
            'font-medium font-poppins text-lg hover:text-orange-200',
            path === '/about' && 'text-orange-400 font-bold text-xl'
          )}>
            About
          </Link>
        </div>
      </div>

      <ProfileBar
        isLogged={false}
        className={cn('w-[260px]')}
      />
    </div>
  )
}

export default Navbar