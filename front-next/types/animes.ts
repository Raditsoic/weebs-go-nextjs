import {Anime} from "./anime";

export type Animes = {
    map(arg0: (anime: any) => import("react").JSX.Element): import("react").ReactNode;
    anime: Anime[]
}