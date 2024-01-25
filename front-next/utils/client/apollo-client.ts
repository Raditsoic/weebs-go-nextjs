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

export const getAnime = async (): Promise<Anime> => {
  const query = gql`
  query {
    oneAnime(id: "8414431888851029825") {
      id
      title
      image
      description
    } 
  }`

  try {
    const { anime }: { anime: Anime } = await graphQLClient.request(query);
    return anime
  } catch(err) {
    console.error(err)
    throw err
  }
}

export const getManga = async (): Promise<Manga> => {
  const query = gql`
  query {
    oneManga(id: "1214583687922155472") {
      id
      title
      image
      description
    } 
  }`

  try {
    const { manga }: { manga: Manga } = await graphQLClient.request(query);
    return manga
  } catch(err) {
    console.error(err)
    throw err
  }
}