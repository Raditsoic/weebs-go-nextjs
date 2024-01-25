import { Manga } from "./manga";

export type Mangas = {
    map(arg0: (manga: any) => import("react").JSX.Element): import("react").ReactNode;
    manga: Manga[]
}