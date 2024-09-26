"use client"
import Link from 'next/link'
import { useEffect, useState } from "react"

import { Animes } from "@/utils/types/api";
import { getAnimes } from "@/utils/client/apollo-client";

import { cn } from '@/utils/lib/cn';

import SearchBar from "@/components/SearchBar";

// function SearchBar() {
//     return (
//       <div className="grid grid-cols-2 gap-8 w-full">
//         {/* Anime description */}
//         <div className='flex flex-col p-10'>
//           <p className='text-2xl font-semibold'>Anime</p>
//           <p className='text-l font-normal mt-2'>
//           is hand-drawn and computer-generated animation originating from Japan. 
//           Outside Japan and in English, anime refers specifically to animation produced in Japan. 
//           However, in Japan and Japanese, anime describes all animated works, regardless of style or origin. 
//           </p>
//         </div>
//         <form className="flex w-full items-center justify-center">
//           {/* Search input */}
//           <div className="flex flex-col w-3/4">
//             <input
//               type="text"
//               className="block w-full rounded-md mt-4 border-0 py-2 pl-3 pr-20 text-gray-900 ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
//               placeholder="Search Anime"
//             />
//           </div>
//       </form>
//     </div>
//     );
//   }
  
export default function AnimePage() {
  // const [animes, setAnimes] = useState<Animes | undefined>(undefined);
  // useEffect(() => {
  //   const fetchedAnime = async () => {
  //     try {
  //       const response = await getAnimes();
  //       setAnimes(response);
  //     } catch (err) {
  //       console.error(err);
  //     }
  //   };

  //   fetchedAnime();
  // }, []);

  // if (animes === undefined) {
  //   return <div>Loading...</div>; 
  // }

  return (
    <div>
      <div className={cn('flex items-center px-10 py-6 justify-between')}>
        <div className='flex flex-col w-2/4'>
          <p className='text-orange-400 text-2xl font-bold'>
            Anime
          </p>
            
          <p className='text-lg font-normal mt-2'>
            is hand-drawn and computer-generated animation originating from Japan. 
            Outside Japan and in English, anime refers specifically to animation produced in Japan. 
            However, in Japan and Japanese, anime describes all animated works, regardless of style or origin. 
          </p>
        </div>

        <SearchBar
          className={cn('my-4')}
          placeholder='Search Anime'
        />
      </div>
      
            {/* <div className="grid lg:grid-cols-4 grid-cols-2 gap-8 grid-flow-row-dense p-5 justify-items-center">
                {animes.map((anime) => (
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
            </div> */}
    </div>
  )
}
