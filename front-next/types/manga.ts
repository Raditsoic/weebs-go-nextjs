import { genre } from "./genre";

export type Manga = {
    id: string;
    title: string;
    image: string;
    description: string;
    genre: genre
}