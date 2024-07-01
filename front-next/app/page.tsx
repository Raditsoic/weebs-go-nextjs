"use client"
import Link from 'next/link'
import { Mangas, Animes } from '@/utils/types/api'
import { getAnimes, getMangas } from "@/utils/client/apollo-client";
import { useEffect, useState } from "react"

function AnimeCard() {
  const [animes, setAnimes] = useState<Animes | undefined>(undefined);
  useEffect(() => {
    const fetchedAnime = async () => {
      try {
        const response = await getAnimes();
        setAnimes(response);
      } catch (err) {
        console.error(err);
      }
    };

    fetchedAnime();
  }, []);

  if (animes === undefined) {
    return <div>Loading...</div>; 
  }

  return (
    <div>
      <div className="flex justify-center items-center p-3">
        <p className='text-xl font-bold'>Anime Collection</p>
      </div>
      <div className='p-5'>
        <div className="flex flex-col">
          <div className="grid lg:grid-cols-2 grid-cols-1 gap-8 items-center justify-center">
          {animes.map((anime) => (
            <div className={`container max-w-s justify-self-auto flex flex-col h-full`} key={anime.id}>
              <Link href="#" className="flex flex-col items-center bg-white border border-gray-200 rounded-lg shadow hover:bg-gray-100 dark:border-gray-700 dark:bg-gray-800 dark:hover:bg-gray-700 h-full">
                <img className="object-cover md:h-64 rounded-t-lg" src={anime.image} alt={anime.title} />
                <div className="flex flex-col justify-between p-4 leading-normal w-full">
                    <h5 className="mb-2 text-2xl font-bold tracking-tight text-gray-900 dark:text-white">{anime.title}</h5>
                    <p className="mb-3 font-normal text-gray-700 dark:text-gray-400">
                        {anime.description.length > 316 ? `${anime.description.substring(0, 316)}...` : anime.description}
                    </p>
                </div>
              </Link>
            </div>
          ))}
          </div>
        </div>
      </div>
    </div>
  )
}

function MangaCard() {
  const [mangas, setMangas] = useState<Mangas | undefined>(undefined);
  useEffect(() => {
    const fetchedManga = async () => {
      try {
        const response = await getMangas();
        setMangas(response);
      } catch (err) {
        console.error(err);
      }
    };

    fetchedManga();
  }, []);

  if (mangas === undefined) {
    return <div>Loading...</div>; 
  }

  return (
    <div>
      <div className="flex justify-center items-center p-3">
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
    </div>
    </div>
  )
}

export default function Home() {
  return (
    <main>
      <div className="flex flex-col p-10 my-2 mx-2">
        <p className='text-2xl font-semibold'>Welcome Home!</p>
        <p className='text-l font-normal mt-2'>
         
        </p>
      </div>
      <div className="grid grid-cols-2 p-6 justify-center divide-x divide-gray-700">
          <AnimeCard />
          <MangaCard />
      </div>
    </main>
  )
}

