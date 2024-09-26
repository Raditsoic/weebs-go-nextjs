import React from 'react'

import {
  Carousel,
  CarouselContent,
  CarouselItem,
  CarouselNext,
  CarouselPrevious,
} from "@/components/ui/carousel"

type CardCarouselProps = {
  title?: string,
  description?: string,
  image?: string
}

const CardCarousel = ({ title, description, image }: CardCarouselProps) => {
  return (
    <Carousel>
      <CarouselPrevious />
      <CarouselContent>
        <CarouselItem>
          <div className="flex flex-col w-64 h-96 bg-gray-200 rounded-lg p-4">
            <img src={image} alt={title} className="w-full h-48 object-cover rounded-lg" />
            <div className="flex flex-col justify-between h-full">
              <div>
                <h2 className="text-lg font-bold">{title}</h2>
                <p className="text-sm text-gray-500">{description}</p>
              </div>
              <button className="bg-blue-500 text-white px-4 py-2 rounded-md">Read More</button>
            </div>
          </div>
        </CarouselItem>
      </CarouselContent>
      <CarouselNext />
    </Carousel>
  )
}

export default CardCarousel