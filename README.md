<<<<<<< HEAD
# Weebs

Implementation of **NextJS - Go - GraphQL - MongoDB** in Anime & Manga List Web App

## Feature

- User Register and Login
- Anime & Manga List
- Anime & Manga Review
- Fully Authenticated App with JWT

## How to run the backend?

### Docker (Recommended)

#### Pre-requisite
- Docker

1. Navigate to the backend directory 
```bash
cd back-go
```
2. Copy the .env.example to .env file and change the variables according to your values
```bash
cp .env.example .env
```
3. Run inside the app inside docker container
```bash
docker-compose up --build
```
4. Access the app on http://localhost:8080

### Local

#### Pre-requisite
- Golang
- MongoDB

#### Run
1. Navigate to the backend directory 
```bash
cd back-go
```
2. Copy the .env.example to .env file and change the variables according to your values
```bash
cp .env.example .env
```
3. Start the MongoDB
```bash
sudo systemctl start mongod
```
4. Run the app on local
```bash
go run server.go
```
5. Access the app on http://localhost:8080
=======
# Weebs

## By Raditsoic

Implementation of **NextJS - React - Go - GraphQL - MongoDB** in Anime & Manga List Web App
>>>>>>> 9b9b4dc (Add: front-end)
