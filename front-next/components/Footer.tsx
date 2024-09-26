import React from 'react'
import Link from 'next/link'
import Links from './Links'
import { cn } from '@/utils/lib/cn'

const Footer = () => {
  return (
    <footer className={cn('bg-gray-900', 'flex justify-between items-center', 'px-5 py-4 md:px-16')}>
      <Link href='/' className={cn('flex justify-center items-center space-x-3')}>
        <img src='/images/kikicat.jpg' alt='Weebs' className={cn('h-10 w-10')} />

        <p className={cn('text-md')}>
          2023 Weebs, Inc
        </p>
      </Link>

      <p className={cn('text-md font-bold')}>
        by raditsoic & harvdt
      </p>

      <Links />
    </footer>
  )
}

export default Footer