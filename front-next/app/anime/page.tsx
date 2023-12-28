import Link from 'next/link'
import animeList from '@/utils/storage/data'
import { Anime } from '@/types/anime'

function SearchBar() {
    return (
      <form className="flex w-full items-center justify-center p-4"> {/* Set a fixed width for the form */}
        <div className="text-center">
        <label className="block text-xl font-medium leading-6 text-white ">Search Anime</label>
        <div className="relative mt-2 rounded-md shadow-sm">
          <input
            type="text"
            className="block w-64 h-10 rounded-md border-0 py-1.5 pl-3 pr-20 text-gray-900 ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
            placeholder="Attack on Titan...?"
          />
        </div>
      </div>
      </form>
    );
  }
  
export default function AnimePage() {
    return (
        <div>
            <SearchBar />
            <div className="grid lg:grid-cols-4 grid-cols-2 gap-8 grid-flow-row-dense p-5 justify-items-center">
                {animeList.map((anime: Anime) => (
                    <div className={`container max-w-s justify-self-auto flex flex-col h-full ${anime.title === 'Spirited Away' ? 'lg:col-span-2' : ''}`} key={anime.id}>
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
    )
}
