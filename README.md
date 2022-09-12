# movie-api


## How to run migrations
in order to run your migrations, you gotta have the golang-migrate tool properly installed, after that 
just run the simple command:

```bash
migrate -path=./migrations -database="postgres://<username>:<password>@localhost:5432/movie_api?sslmode=disable" up
```


