"use client"
import { useEffect, useState } from "react"
import { Anime } from "@/types/anime";
import { getAnime } from "@/utils/client/apollo-client";

export default function DedicatedAnime({ params }: { params: { animepage: any; slug:string 
}}) {
    const routerID = params.animepage;
    const [anime, setAnime] = useState<Anime | undefined>(undefined);
    useEffect(() => {
        const fetchedAnime = async () => {
            try {
                const response = await getAnime(routerID);
                setAnime(response)
            } catch (err) {
                console.error(err)
            }
        }

        fetchedAnime()
    }, [])

    return (
        <div>
            {anime?.title}
        </div>
    )
}