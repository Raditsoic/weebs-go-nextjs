import React from 'react'
import Link from 'next/link'
import { cn } from '@/utils/lib/cn'
import { FaGithub, FaInstagram, FaLinkedin } from "react-icons/fa";
import { BsTwitterX } from "react-icons/bs";

const Links = () => {
  return (
    <div className={cn("flex items-center justify-center space-x-2")}>
      <Link href="https://www.instagram.com/frditya.a/" target="_blank" rel="noopener noreferrer">
        <FaInstagram className={cn('h-6 w-6', 'hover:fill-orange-400')} />
      </Link>

      <Link href="https://github.com/Raditsoic" target="_blank" rel="noopener noreferrer">
        <FaGithub className={cn('h-6 w-6', 'hover:fill-orange-400')} />
      </Link>

      <Link href="https://www.linkedin.com/in/awang-fraditya-586b21248/" target="_blank" rel="noopener noreferrer">
        <FaLinkedin className={cn('h-6 w-6', 'hover:fill-orange-400')} />
      </Link>
      
      <Link href="https://X.com/raditsoic" target="_blank" rel="noopener noreferrer">
        <BsTwitterX className={cn('h-6 w-6', 'hover:fill-orange-400')} />
      </Link>
    </div>
  )
}

export default Links