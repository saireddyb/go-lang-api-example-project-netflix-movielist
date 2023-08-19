# go-lang-api-example-netflix-movielist


create a documentation for this project to start locally

## Installation
- Clone the repo
- Install the dependencies
- Run the server

### Clone the repo
```shell
git clone https://github.com/saireddyb/go-lang-api-example-project-netflix-movielist.git
cd go-lang-api-example-project-netflix-movielist

```
### Install the dependencies
```shell
go get -u github.com/gorilla/mux

```

### Run the server
```shell
go run main.go

```

## Usage
- Run the server
- Navigate to `http://localhost:8000/movies`
- You can see the list of movies
- You can see the details of a movie by navigating to `http://localhost:8000/movie/{id}`
- You can add a movie by navigating to `http://localhost:8000/movie/{id}` with method `POST`
- You can update a movie by navigating to `http://localhost:8000/movie/{id}` with method `PUT`
- You can delete a movie by navigating to `http://localhost:8000/movie/{id}` with method `DELETE`
- You can update the rating of a movie by navigating to `http://localhost:8000/movie/{id}/rating` with method `PUT`