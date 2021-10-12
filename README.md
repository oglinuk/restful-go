# REST API in Golang

This is an example REST API implementation using primarily the standard
library. The exceptions are as follows.

```
github.com/gorilla/mux v1.8.0 (for the router)
github.com/mattn/go-sqlite3 v1.14.8 (for the database)
github.com/stretchr/testify v1.7.0 (for testing)
gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b (for configuration)
```

## Getting Started

`./init app`

### With `docker`

`./init dapi && ./init ui`

### With `docker-compose`

`./init dcompose`


Once both are running go to [`localhost:9042`](http://localhost:9042). If
you want to access the page from other devices on the network, goto
[`localhost:9001`](http://localhost:9001) and replace `localhost` with
the `ip` given as the response.
