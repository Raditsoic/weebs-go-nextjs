"use client"
import { useEffect, useState } from "react"
import { Anime } from '@/utils/types/api'
import { getAnime } from "@/utils/client/apollo-client";

export default function DedicatedAnime({ params }: { params: { animepage: any; slug:string 
}}) {
    const routerID = params.animepage;
    const [anime, setAnime] = useState<Anime | undefined>(undefined);
    const [error, setError] = useState<string | null>(null);
    
    useEffect(() => {
        const fetchedAnime = async () => {
            try {
                const response = await getAnime(routerID);
                setAnime(response)
            } catch (err) {
                setError("Item Can't be Found")
            }
        }

        fetchedAnime()
    }, [])

    if (!error) {
        throw new Error("error")
    }

    if (anime == undefined) {
        return (
            <div>Loading...</div>
        )
    }

    return (
        <div>
            {anime?.title}
        </div>
    )
}