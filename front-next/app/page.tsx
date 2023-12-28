import Image from 'next/image'
import Link from 'next/link'
import animeList, { mangaList } from '@/utils/storage/data'
import { Anime } from '@/types/anime'
import { Manga } from '@/types/manga'

function AnimeCard() {
  return (
    <div>
      <div className="flex justify-center items-center p-3">
        <p className='text-xl font-bold'>Anime Collection</p>
      </div>
      <div className='p-5'>
        <div className="flex flex-col">
          <div className="grid lg:grid-cols-2 grid-cols-1 gap-8 items-center justify-center">
          {animeList.map((anime: Anime) => (
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
  return (
    <div>
      <div className="flex justify-center items-center p-3">
        <p className='text-xl font-bold'>Manga Collection</p>
      </div>   
      <div className='p-5'>
      <div className="flex flex-col">
        <div className="grid lg:grid-cols-2 grid-cols-1 gap-8 items-center justify-center">
        {mangaList.map((manga: Manga) => (
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
      <div className="flex p-6">
        <p className='text-4xl font-bold'>Welcome to Weebs<br></br>Collections of Anime and Manga</p>
      </div>
      <div className="grid grid-cols-2 p-6 justify-center divide-x divide-gray-700">
          <AnimeCard></AnimeCard>
          <MangaCard></MangaCard>
      </div>
    </main>
  )
}
