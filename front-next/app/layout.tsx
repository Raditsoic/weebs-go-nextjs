import type { Metadata } from 'next'
import { Poppins, Syncopate } from 'next/font/google'
import './globals.css'

import { cn } from '@/lib/utils';

const poppins = Poppins({
  subsets: ["latin"],
  weight: ["400", "500", "600", "700"],
  variable: "--font-poppins",
});

const syncopate = Syncopate({
  subsets: ["latin"],
  weight: ["400", "700"],
  variable: "--font-syncopate",
});

export const metadata: Metadata = {
  icons: {
    icon: '/images/kikicat.jpg'
  },
  title: 'Weebs',
  description: 'by Raditsoic',
}

export default function RootLayout({
  children,
}: {
  children: React.ReactNode
}) {
  return (
    <html lang="en">
      <body className={cn(`${poppins.variable} ${syncopate.variable}`, 'text-white bg-slate-950')}>
        {children}
      </body>
    </html>
  )
}
