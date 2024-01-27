import type { Metadata } from 'next'
import { Inter } from 'next/font/google'
import './globals.css'
import Link from 'next/link'

const inter = Inter({ subsets: ['latin'] })

export const metadata: Metadata = {
  icons: {
    icon: '/images/kikicat.jpg'
  },
  title: 'Weebs',
  description: 'by Raditsoic',
}

function Navbar() {
  return (
    <div className="max-w-screen-xl flex flex-wrap items-center justify-between mx-auto p-4">
      <Link href="/" className="flex items-center space-x-3 rtl:space-x-reverse">
        <img src="/images/kikicat.jpg" className="h-8" alt="Weebs Logo" />
        <span className="self-center text-2xl font-semibold whitespace-nowrap dark:text-white">Weebs</span>
      </Link>
      <button data-collapse-toggle="navbar-default" type="button" className="inline-flex items-center p-2 w-10 h-10 justify-center text-sm text-gray-500 rounded-lg md:hidden hover:bg-gray-100 focus:outline-none focus:ring-2 focus:ring-gray-200 dark:text-gray-400 dark:hover:bg-gray-700 dark:focus:ring-gray-600" aria-controls="navbar-default" aria-expanded="false">
        <span className="sr-only">Open main menu</span>
        <svg className="w-5 h-5" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 17 14">
          <path stroke="currentColor" strokeLinecap="round" strokeLinejoin="round" strokeWidth="2" d="M1 1h15M1 7h15M1 13h15"/>
        </svg>
      </button>
      <div className="hidden w-full md:block md:w-auto" id="navbar-default">
        <ul className="font-medium flex flex-col p-4 md:p-0 mt-4 border border-gray-100 rounded-lg bg-gray-50 md:flex-row md:space-x-8 rtl:space-x-reverse md:mt-0 md:border-0 md:bg-white dark:bg-gray-800 md:dark:bg-gray-900 dark:border-gray-700">
          <li>
            <Link href="/" className="block py-2 px-3 text-white bg-blue-700 rounded md:bg-transparent md:text-blue-700 md:p-0 dark:text-white md:dark:text-blue-500" aria-current="page">Home</Link>
          </li>
          <li>
            <Link href="/anime" className="block py-2 px-3 text-gray-900 rounded hover:bg-gray-100 md:hover:bg-transparent md:border-0 md:hover:text-blue-700 md:p-0 dark:text-white md:dark:hover:text-blue-500 dark:hover:bg-gray-700 dark:hover:text-white md:dark:hover:bg-transparent">Anime</Link>
          </li>
          <li>
            <Link href="/manga" className="block py-2 px-3 text-gray-900 rounded hover:bg-gray-100 md:hover:bg-transparent md:border-0 md:hover:text-blue-700 md:p-0 dark:text-white md:dark:hover:text-blue-500 dark:hover:bg-gray-700 dark:hover:text-white md:dark:hover:bg-transparent">Manga</Link>
          </li>
          <li>
            <Link href="/about" className="block py-2 px-3 text-gray-900 rounded hover:bg-gray-100 md:hover:bg-transparent md:border-0 md:hover:text-blue-700 md:p-0 dark:text-white md:dark:hover:text-blue-500 dark:hover:bg-gray-700 dark:hover:text-white md:dark:hover:bg-transparent">About</Link>
          </li>
        </ul>
      </div>
    </div>
  )
}

function Links() {
  return (
    <div className="flex flex-col items-center justify-center p-3">
      <div className='grid grid-cols-4'>
        <div className='flex flex-col p-1'>
            <Link href="https://www.instagram.com/frditya.a/" target="_blank" rel="noopener noreferrer">
              <div className='container flex justify-center p-1'>
                <img src="/images/logo/instagram.svg" alt="Instagram" className="h-6 w-6" />
              </div>
            </Link>
        </div>
          <div className='flex flex-col p-1'>
            <Link href="https://github.com/Raditsoic" target="_blank" rel="noopener noreferrer">
              <div className='container flex justify-center p-1'>
                <img src="/images/logo/github.svg" alt="GitHub" className="h-6 w-6" />
              </div>
            </Link>
          </div>
          <div className='flex flex-col p-1'>
            <Link href="https://www.linkedin.com/in/awang-fraditya-586b21248/" target="_blank" rel="noopener noreferrer">
              <div className='container flex justify-center p-1'>
                <img src="/images/logo/linkedin.svg" alt="LinkedIn" className="h-6 w-6" />
              </div>
            </Link>
          </div>
          <div className='flex flex-col p-1'>
            <Link href="https://X.com/raditsoic" target="_blank" rel="noopener noreferrer">
              <div className='container flex justify-center p-1'>
                <img src="/images/logo/twitter.svg" alt="Twitter" className="h-6 w-6" />
              </div>
            </Link>
          </div>
        </div>
      </div>
  )}

export default function RootLayout({
  children,
}: {
  children: React.ReactNode
}) {
  return (
    <html lang="en">
      <body>
        <div className='flex flex-col min-h-screen'>
        <header className="bg-white border-gray-200 dark:bg-gray-900">
          <Navbar></Navbar>
        </header>
        {children}
        <footer className="bg-white border-gray-200 dark:bg-gray-900">
          <div className='grid grid-cols-3 gap-8'>
            <div className='flex flex-col justify-center p-3'>
              <div className='container flex'>
                <Link href="/">
                  <div className='grid grid-cols-2 gap-2'>
                    <div className='flex flex-col justify-end items-end'>
                      <img src='/images/kikicat.jpg' alt='Weebs' className='h-10 w-10'/>
                    </div>
                    <div className='flex flex-col justify-center items-center'>
                      <p className='text-sm'>2023 Weebs, Inc</p>
                    </div>
                  </div>
                </Link>
              </div>
            </div>
            <div className='flex flex-col items-center justify-center p-3'>
              <div>
                <p className='text-3xl font-extrabold'>Created by raditsoic</p>
              </div>
            </div>
            <Links></Links>
          </div>
        </footer>
        </div>
      </body>
    </html>
  )
}
