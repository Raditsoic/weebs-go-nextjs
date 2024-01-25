import { genre } from "./genre";

export type Anime = {
    id: string;
    title: string;
    image: string;
    description: string;
    genre: genre
}