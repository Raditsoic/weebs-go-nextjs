import React, { useRef } from 'react';
import Link from 'next/link';

import { Swiper, SwiperSlide } from 'swiper/react';
import { Swiper as SwiperType } from 'swiper';
import 'swiper/css';

import { Rating } from 'react-simple-star-rating'

import { FaChevronLeft, FaChevronRight, FaArrowRight } from "react-icons/fa";
import { FaStar } from "react-icons/fa";

import { cn } from '@/utils/lib/cn';

const recommendations = [
  {
    id: 1,
    title: 'Castle in the Sky',
    genre: 'Gataw',
    description: 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Duis ullamcorper lorem mauris, ac consectetur augue vulputate finibus. Aliquam convallis enim vitae ex iaculis fermentum. Etiam tincidunt vel ante ut laoreet. Aliquam varius eget mauris in pretium. Suspendisse in lectus at diam ultrices viverra sed sed velit.',
    rating: 4.3,
    releaseYear: '2004',
    type: 'Anime',
    image: '/images/anilist/castle_in_the_sky.png',
  },
  {
    id: 2,
    title: 'Grave of the Fireflies',
    genre: 'Gataw',
    description: 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Duis ullamcorper lorem mauris, ac consectetur augue vulputate finibus. Aliquam convallis enim vitae ex iaculis fermentum. Etiam tincidunt vel ante ut laoreet. Aliquam varius eget mauris in pretium. Suspendisse in lectus at diam ultrices viverra sed sed velit.',
    rating: 4.7,
    releaseYear: '2009',
    type: 'Anime',
    image: '/images/anilist/grave_of_the_fireflies.jpg',
  },
  {
    id: 3,
    title: 'Howl Moving Castle',
    genre: 'Gataw',
    description: 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Duis ullamcorper lorem mauris, ac consectetur augue vulputate finibus. Aliquam convallis enim vitae ex iaculis fermentum. Etiam tincidunt vel ante ut laoreet. Aliquam varius eget mauris in pretium. Suspendisse in lectus at diam ultrices viverra sed sed velit.',
    rating: 3.8,
    releaseYear: '2011',
    type: 'Anime',
    image: '/images/anilist/howl_moving_castle.jpg',
  }
]

type SlideContentProps = {
  type: string
  title: string
  genre: string
  image: string
  rating: number
  description: string
  releaseYear: string
}

function SlideContent({ type, title, genre, image, rating, description, releaseYear }: SlideContentProps) {
  return (
    <div className="relative h-full w-full">
      <img
        src={image}
        className={cn('absolute inset-y-0 right-0 z-10 h-full w-auto object-cover')}
        alt={title}
      />

      <div className={cn(
        'relative z-20 h-full w-full',
        'text-white',
        'bg-gradient-to-r from-black from-45% via-black/10 via-70% to-transparent'
      )}>
        <div className={cn(
          'relative z-30 h-full w-1/2 pl-32 pr-16',
          'text-white flex flex-col justify-center'
        )}>
          <p className={cn('font-poppins font-bold text-gray-400 text-lg')}>
            {type}
          </p>

          <p className={cn('font-bold text-3xl')}>
            {title} <span>({releaseYear})</span>
          </p>

          <p className={cn('font-poppins font-semibold text-orange-400 mt-1')}>
            {genre}
          </p>
          
          <p className={cn('font-medium mt-6')}>
            {description}
          </p>

          <div className={cn('flex items-center gap-2 my-3')}>
            <FaStar
              size={22}
              className={cn('fill-orange-400')}
            />

            <p className={cn('font-poppins font-semibold')}>{rating}/5</p>
          </div>

          <Link href='/' className="z-40 mt-3">
            <button className={cn(
              'h-[52px] w-[160px] bg-slate-50 rounded-sm flex items-center justify-between px-5',
              'text-black font-poppins font-bold',
              'hover:text-white hover:bg-orange-400 transition-colors duration-200'
            )}>
              Go to page
          
              <FaArrowRight />
            </button>
          </Link>
        </div>
      </div>
    </div>
  )
}

const Hero = () => {
  const swiperRef = useRef<SwiperType | null>(null);

  return (
    <div className="relative h-[500px]">
      <Swiper
        onSwiper={(swiper) => (swiperRef.current = swiper)}
        spaceBetween={0}
        slidesPerView={1}
        className="bg-white text-black h-full"
      >
        {recommendations.map((recommendation) => (
          <SwiperSlide key={recommendation.id}>
            <SlideContent
              title={recommendation.title}
              image={recommendation.image}
              description={recommendation.description}
              genre={recommendation.genre}
              rating={recommendation.rating}
              releaseYear={recommendation.releaseYear}
              type={recommendation.type}
            />
          </SwiperSlide>
        ))}
      </Swiper>

      <div className={cn(
        'absolute bottom-0 flex items-center justify-between z-50 px-4',
        'w-full h-[500px] pointer-events-none',
        'opacity-0 hover:opacity-100 transition-opacity duration-300'
      )}>
        <div className={cn('h-[500px] w-28 pointer-events-auto flex items-center')}>
          <button
            onClick={() => swiperRef.current?.slidePrev()}
            className={cn(
              'h-10 w-10 rounded-full bg-slate-500/50 flex items-center justify-center text-white shadow-xl',
              'pointer-events-auto'
            )}
            aria-label='Previous slide'
          >
            <FaChevronLeft size={20} />
          </button>
        </div>

        <div className={cn('h-[500px] w-36 pointer-events-auto flex items-center justify-end')}>
          <button
            onClick={() => swiperRef.current?.slideNext()}
            className={cn(
              'h-10 w-10 rounded-full bg-slate-500/50 flex items-center justify-center text-white shadow-xl',
              'pointer-events-auto'
            )}
            aria-label='Next slide'
          >
            <FaChevronRight size={20} />
          </button>
        </div>
      </div>
    </div>
  );
};

export default Hero;