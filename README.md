# go-lang-api-example-netflix-movielist

## Table of Contents
- [Description](#description)
- [Installation](#installation)
- [Clone the repo](#clone-the-repo)
- [Install the dependencies](#install-the-dependencies)
- [Run the server](#run-the-server)
## Description
This is a simple example of a REST API written in Go. It is a simple movie list API. It has the following features:
- List all movies
- Get a movie by id
- Add a movie
- Update a movie
- Delete a movie
- Update the rating of a movie
- List all directors
- Get a director by id
- Add a director
- Add initial data to the database


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
- You can see the list of movies by Navigate to `http://localhost:8000/api/movies`
- You can see the details of a movie by navigating to `http://localhost:8000/api/movie/{id}`
- You can add a movie by navigating to `http://localhost:8000/api/movie/{id}` with method `POST`
- You can update a movie by navigating to `http://localhost:8000/api/movie/{id}` with method `PUT`
- You can delete a movie by navigating to `http://localhost:8000/api/movie/{id}` with method `DELETE`
- You can update the rating of a movie by navigating to `http://localhost:8000/api/movie/{id}/rating` with method `PUT`
- You can see the list of directors by navigating to `http://localhost:8000/api/directors`
- You can see the details of a director by navigating to `http://localhost:8000/api/director/{id}`
- You can add a director by navigating to `http://localhost:8000/api/director/{id}` with method `POST`
- You can see the list of countries by navigating to `http://localhost:8000/api/countries`
- You can see the details of a country by navigating to `http://localhost:8000/api/country/{id}`
- You can add a country by navigating to `http://localhost:8000/api/country` with method `POST`
- You can add initial data to the database by navigating to `http://localhost:8000/api/initialdata` with method `POST`