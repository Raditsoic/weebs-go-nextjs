"use client"

import { useEffect, useState } from "react";
import Link from 'next/link'

import { getMangas } from '@/utils/client/apollo-client';
import { Mangas } from '@/utils/types/api';
import { cn } from "@/utils/lib/cn";

import SearchBar from "@/components/SearchBar";


export default function Manga() {
    // const [mangas, setMangas] = useState<Mangas | undefined>(undefined);
    // useEffect(() => {
    // const fetchedManga = async () => {
    //   try {
    //     const response = await getMangas();
    //     setMangas(response);
    //   } catch (err) {
    //     console.error(err);
    //   }
    // };

    // fetchedManga();
    // }, []);

    // if (mangas === undefined) {
    //     return <div>Loading...</div>; 
    // }

  return (
    <div>
      <div className={cn('flex items-center px-10 py-6 justify-between')}>
        <div className='flex flex-col w-2/4'>
          <p className='text-orange-400 text-2xl font-bold'>
            Manga
          </p>
            
          <p className='text-lg font-normal mt-2'>
            are comics or graphic novels originating from Japan.
            Most manga conform to a style developed in Japan in the late 19th century.
            The term manga is used in Japan to refer to both comics and cartooning.
            Outside of Japan, the word is typically used to refer to comics originally published in Japan.
          </p>
        </div>

        <SearchBar
          className={cn('my-4')}
          placeholder='Search Anime'
        />
      </div>


      {/* <div className="flex justify-center items-center p-3">
        <p className='text-xl font-bold'>Manga Collection</p>
      </div>   
      <div className='p-5'>
      <div className="flex flex-col">
        <div className="grid lg:grid-cols-2 grid-cols-1 gap-8 items-center justify-center">
        {mangas.map((manga) => (
          <div className={`container max-w-s justify-self-auto flex flex-col h-full`} key={manga.id}>
            <Link href="#" className="flex flex-col items-center bg-white border border-gray-200 rounded-lg shadow hover:bg-gray-100 dark:border-gray-700 dark:bg-gray-800 dark:hover:bg-gray-700 h-full">
            <img className="object-cover md:h-64 rounded-t-lg" src={manga.image} alt={manga.title} />
              <div className="flex flex-col justify-between p-4 leading-normal w-full">
                  <h5 className="mb-2 text-2xl font-bold tracking-tight text-gray-900 dark:text-white">{manga.title}</h5>
                  <p className="mb-3 font-normal text-gray-700 dark:text-gray-400">
                      {manga.description.length > 316 ? `${manga.description.substring(0, 316)}...` : manga.description}
                  </p>
              </div>
            </Link>
          </div>
        ))}
        </div>
      </div>
    </div> */}
    </div>
  )
}