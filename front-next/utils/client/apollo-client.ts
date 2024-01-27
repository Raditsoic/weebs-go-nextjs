import { GraphQLClient, gql } from "graphql-request";
import { Animes } from "@/types/animes";
import { Mangas } from "@/types/mangas";
import { Anime } from "@/types/anime";
import { Manga } from "@/types/manga";

const endpoint = 'http://localhost:8080/query'

const graphQLClient = new GraphQLClient(endpoint)

export const getAnimes = async (): Promise<Animes> => {
  const query = gql`
  query {
    animes {
      id
      title
      image
      description
      genre {
        id
        name
      }
    }
  }
  `

  try {
    const { animes }: { animes: Animes } = await graphQLClient.request(query);
    return animes;
  } catch (error) {
    console.error(error)
    throw error
  }
}

export const getMangas = async (): Promise<Mangas> => {
  const query = gql`
  query {
    mangas {
      id
      title
      image
      description
      genre {
        id
        name
      }
    }
  }
  `

  try {
    const { mangas }: { mangas: Mangas } = await graphQLClient.request(query);
    return mangas;
  } catch (error) {
    console.error(error)
    throw error
  }
}

export const getAnime = async (id: string): Promise<Anime | undefined> => {
  const query = gql`
  query{
    oneAnime(id: "${id}") {
      id
      title
      image
      description
    } 
  }`

  try {
    const { oneAnime }: { oneAnime: Anime } = await graphQLClient.request(query);
    return oneAnime
  } catch(err) {
    return undefined
  }
}

export const getManga = async (id: string): Promise<Manga | undefined> => {
  const query = gql`
  query {
    oneManga(id: "${id}") {
      id
      title
      image
      description
    } 
  }`

  try {
    const { oneManga }: { oneManga: Manga } = await graphQLClient.request(query);
    return oneManga
  } catch(err) {
    return undefined
  }
}