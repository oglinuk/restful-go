# REST API in Golang

This is an example REST API implementation using primarily the standard
library, Docker, and Docker Compose. The exceptions are as follows.

```
github.com/go-chi/chi/v5 v5.0.7 (for the router)
github.com/go-chi/cors v1.2.0 (for CORS)
github.com/mattn/go-sqlite3 v1.14.8 (for the database)
github.com/stretchr/testify v1.7.0 (for testing)
gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b (for configuration)
```

## Getting Started

`./init app`

`./init clean` to stop the application

### With `docker`

`./init dapi && ./init ui`

`./init dclean` to stop/rm the docker container

### With `docker-compose`

`docker-compose up --build` || `./init dcompose`

Once both are running go to [`localhost:9042`](http://localhost:9042). If
you want to access the page from other devices on the network, goto
[`localhost:9001`](http://localhost:9001) and replace `localhost` with
the `ip` given as the response.

## TODO

* [X] Implement simple CRUD operations
	* [X] Insert
	* [X] SelectAll
	* [X] Select
	* [X] Update
	* [X] Delete
* [X] [Code Review](docs/code-reviews/1642914462)
* [ ] Implement API tokens using JWT
* [ ] Implement a database abstraction layer
* [ ] Add CI/CD pipeline
* [ ] Add `ARCHITECTURE.md`
