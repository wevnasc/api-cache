# api-cache
This is an example of a use case for caching, for this example, I am using Redis as the cache data source and a web server wroten in go 

the only responsibility of this service is to provide the profile info from a user using the GitHub API as a source of truth, the response of this API is cached to avoid making unnecessary requests to the GitHub API.

## Stack
- go
- redis
- docker

## how to run
```sh 
$ docker compose build
$ docker compose up -d
``` 

go to [http://localhost:3000/?username=wevnasc](http://localhost:3000/?username=wevnasc)