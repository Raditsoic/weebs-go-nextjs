"use client"
import { useEffect, useState } from "react"
import { Manga } from "@/utils/types/api";
import { getManga } from "@/utils/client/apollo-client";

export default function DedicatedManga({ params }: { params: { mangapage: any; slug:string 
}}) {
    const routerID = params.mangapage;
    const [manga, setManga] = useState<Manga | undefined>(undefined);
    useEffect(() => {
        const fetchedManga = async () => {
            try {
                const response = await getManga(routerID);
                setManga(response)
            } catch (err) {
                console.error(err)
            }
        }

        fetchedManga()
    }, [])

    return (
        <div>
            {manga?.title}
        </div>
    )
}