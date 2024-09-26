"use client"

import Link from 'next/link'
import { useEffect, useState } from "react"

import { Mangas, Animes } from '@/utils/types/api'
import { getAnimes, getMangas } from "@/utils/client/apollo-client";
import { cn } from '@/utils/lib/cn'

import Hero from '@/components/Hero'
import CardCarousel from '@/components/CardCarousel';

const animes = [
  {
    id: 1,
    title: 'Naruto',
    description: 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Duis ullamcorper lorem mauris, ac consectetur augue vulputate finibus. Aliquam convallis enim vitae ex iaculis fermentum. Etiam tincidunt vel ante ut laoreet. Aliquam varius eget mauris in pretium. Suspendisse in lectus at diam ultrices viverra sed sed velit.',
    image: 'bjir',
  },
  {
    id: 2,
    title: 'Naruto',
    description: 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Duis ullamcorper lorem mauris, ac consectetur augue vulputate finibus. Aliquam convallis enim vitae ex iaculis fermentum. Etiam tincidunt vel ante ut laoreet. Aliquam varius eget mauris in pretium. Suspendisse in lectus at diam ultrices viverra sed sed velit.',
    image: 'bjir',
  },
  {
    id: 3,
    title: 'Naruto',
    description: 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Duis ullamcorper lorem mauris, ac consectetur augue vulputate finibus. Aliquam convallis enim vitae ex iaculis fermentum. Etiam tincidunt vel ante ut laoreet. Aliquam varius eget mauris in pretium. Suspendisse in lectus at diam ultrices viverra sed sed velit.',
    image: 'bjir',
  }
]

const mangas = [
  {
    id: 1,
    title: 'Naruto',
    description: 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Duis ullamcorper lorem mauris, ac consectetur augue vulputate finibus. Aliquam convallis enim vitae ex iaculis fermentum. Etiam tincidunt vel ante ut laoreet. Aliquam varius eget mauris in pretium. Suspendisse in lectus at diam ultrices viverra sed sed velit.',
    image: 'bjir',
  },
  {
    id: 2,
    title: 'Naruto',
    description: 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Duis ullamcorper lorem mauris, ac consectetur augue vulputate finibus. Aliquam convallis enim vitae ex iaculis fermentum. Etiam tincidunt vel ante ut laoreet. Aliquam varius eget mauris in pretium. Suspendisse in lectus at diam ultrices viverra sed sed velit.',
    image: 'bjir',
  },
  {
    id: 3,
    title: 'Naruto',
    description: 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Duis ullamcorper lorem mauris, ac consectetur augue vulputate finibus. Aliquam convallis enim vitae ex iaculis fermentum. Etiam tincidunt vel ante ut laoreet. Aliquam varius eget mauris in pretium. Suspendisse in lectus at diam ultrices viverra sed sed velit.',
    image: 'bjir',
  }
]

export default function Home() {
  return (
    <main>
      <Hero />

      <CardCarousel />

      {/* <div className="grid grid-cols-2 p-6 justify-center divide-x divide-gray-700">
        <AnimeCard/>
        <MangaCard />
      </div> */}
    </main>
  )
}

