# Movie Search Solution

Movie Search Solution

![Flow](MicroserviceDesign.png)

## Run Project

### Run Go Mod Vendor

```bash
make vendor
```

Or

```bash
go mod vendor
```

### Run API Gateway

```bash
go run cmd\main.go
```

### Run Movie Service

```bash
go run serviceMovieServer\server.go
```

## Run Test

```bash
make test
```

Or

```bash
go test -v ./... -cover
```

## Example Request and Response

### Movie Service

#### Get Movie List

| Method | URL                                                                           |
| ------ | ----------------------------------------------------------------------------- |
| `GET`  | `http://localhost:8080/search?pagination=[PageNumber]&searchword=[MovieName]` |

##### Parameter

| Method        | Valid options | Description                  |
| ------------- | ------------- | ---------------------------- |
| `pagination`  | `1-100`       | `Page number to return.`     |
| `searchword`  | ` `           | `Movie title to search for.` |


##### Response

###### Success (200)

```json
{
    "status": "OK",
    "data": {
        "Search": [
            {
                "Title": "Batman Begins",
                "Year": "2005",
                "imdbID": "tt0372784",
                "Type": "movie",
                "Poster": "https://m.media-amazon.com/images/M/MV5BOTY4YjI2N2MtYmFlMC00ZjcyLTg3YjEtMDQyM2ZjYzQ5YWFkXkEyXkFqcGdeQXVyMTQxNzMzNDI@._V1_SX300.jpg"
            },
            {
                "Title": "Batman v Superman: Dawn of Justice",
                "Year": "2016",
                "imdbID": "tt2975590",
                "Type": "movie",
                "Poster": "https://m.media-amazon.com/images/M/MV5BYThjYzcyYzItNTVjNy00NDk0LTgwMWQtYjMwNmNlNWJhMzMyXkEyXkFqcGdeQXVyMTQxNzMzNDI@._V1_SX300.jpg"
            },
            {
                "Title": "Batman",
                "Year": "1989",
                "imdbID": "tt0096895",
                "Type": "movie",
                "Poster": "https://m.media-amazon.com/images/M/MV5BMTYwNjAyODIyMF5BMl5BanBnXkFtZTYwNDMwMDk2._V1_SX300.jpg"
            },
            {
                "Title": "Batman Returns",
                "Year": "1992",
                "imdbID": "tt0103776",
                "Type": "movie",
                "Poster": "https://m.media-amazon.com/images/M/MV5BOGZmYzVkMmItM2NiOS00MDI3LWI4ZWQtMTg0YWZkODRkMmViXkEyXkFqcGdeQXVyODY0NzcxNw@@._V1_SX300.jpg"
            },
            {
                "Title": "Batman Forever",
                "Year": "1995",
                "imdbID": "tt0112462",
                "Type": "movie",
                "Poster": "https://m.media-amazon.com/images/M/MV5BNDdjYmFiYWEtYzBhZS00YTZkLWFlODgtY2I5MDE0NzZmMDljXkEyXkFqcGdeQXVyMTMxODk2OTU@._V1_SX300.jpg"
            },
            {
                "Title": "Batman & Robin",
                "Year": "1997",
                "imdbID": "tt0118688",
                "Type": "movie",
                "Poster": "https://m.media-amazon.com/images/M/MV5BMGQ5YTM1NmMtYmIxYy00N2VmLWJhZTYtN2EwYTY3MWFhOTczXkEyXkFqcGdeQXVyNTA2NTI0MTY@._V1_SX300.jpg"
            },
            {
                "Title": "The Lego Batman Movie",
                "Year": "2017",
                "imdbID": "tt4116284",
                "Type": "movie",
                "Poster": "https://m.media-amazon.com/images/M/MV5BMTcyNTEyOTY0M15BMl5BanBnXkFtZTgwOTAyNzU3MDI@._V1_SX300.jpg"
            },
            {
                "Title": "Batman: The Animated Series",
                "Year": "1992–1995",
                "imdbID": "tt0103359",
                "Type": "series",
                "Poster": "https://m.media-amazon.com/images/M/MV5BOTM3MTRkZjQtYjBkMy00YWE1LTkxOTQtNDQyNGY0YjYzNzAzXkEyXkFqcGdeQXVyOTgwMzk1MTA@._V1_SX300.jpg"
            },
            {
                "Title": "Batman: Under the Red Hood",
                "Year": "2010",
                "imdbID": "tt1569923",
                "Type": "movie",
                "Poster": "https://m.media-amazon.com/images/M/MV5BNmY4ZDZjY2UtOWFiYy00MjhjLThmMjctOTQ2NjYxZGRjYmNlL2ltYWdlL2ltYWdlXkEyXkFqcGdeQXVyNTAyODkwOQ@@._V1_SX300.jpg"
            },
            {
                "Title": "Batman: The Dark Knight Returns, Part 1",
                "Year": "2012",
                "imdbID": "tt2313197",
                "Type": "movie",
                "Poster": "https://m.media-amazon.com/images/M/MV5BMzIxMDkxNDM2M15BMl5BanBnXkFtZTcwMDA5ODY1OQ@@._V1_SX300.jpg"
            }
        ],
        "totalResults": "393",
        "Response": "True"
    },
    "accessTime": "22-10-2020 14:54:45"
}
```

###### Error Bad Request (400)

```json
{
  "status": "Bad Request",
  "data": <error-message>,
  "accessTime": "11-10-2020 16:58:09"
}
```

###### Error Internal Server Error (500)

```json
{
  "status": "Internal Server Error",
  "data": <error-message>,
  "accessTime": "11-10-2020 16:58:09"
}
```