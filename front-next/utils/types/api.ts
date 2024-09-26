export type genre = {
    id: string
    name: string
}

export type Anime = {
    id: string;
    title: string;
    image: string;
    description: string;
    genre: genre
}

export type Manga = {
    id: string;
    title: string;
    image: string;
    description: string;
    genre: genre
}

export type Animes = {
    map(arg0: (anime: any) => import("react").JSX.Element): import("react").ReactNode;
    anime: Anime[]
}

export type Mangas = {
    map(arg0: (manga: any) => import("react").JSX.Element): import("react").ReactNode;
    manga: Manga[]
}