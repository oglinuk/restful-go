# 1642914462 Code Review

Part of this projects purpose is to practice code reviewing. I am the
sole programmer of this project at the moment, so I will have to do a
self-review. I must admit that I have only ever done one pseudo
code-review with my friend, so my experience is little.

## Dockerfile

This implementation can be improved significantly by changing to a
multi-stage build and using a `scratch` image. The current image size is
`986MB`. Since the database used is SQLite, `cgo` is required, which is
required to be disabled for using a multi-stage build. Need to remove the
direct dependency to `SQLite`.


**TODO**

* [ ] Implement a data abstraction layer
* [ ] Refactor to using multi-stage build with `scratch` image

## init

Missing tab-completion. The results of running the script through
`shellcheck` are as follows.

```BASH
In ../../../init line 4:
	cd ./cmd/api
        ^----------^ SC2164: Use 'cd ... || exit' or 'cd ... || return' in case cd fails.

Did you mean: 
	cd ./cmd/api || exit


In ../../../init line 15:
	cd ./cmd/ui
        ^---------^ SC2164: Use 'cd ... || exit' or 'cd ... || return' in case cd fails.

Did you mean: 
	cd ./cmd/ui || exit


In ../../../init line 39:
$*
^-- SC2048: Use "$@" (with quotes) to prevent whitespace problems.

For more information:
  https://www.shellcheck.net/wiki/SC2048 -- Use "$@" (with quotes) to prevent...
  https://www.shellcheck.net/wiki/SC2164 -- Use 'cd ... || exit' or 'cd ... |...
```

**TODO**

* [X] Implement tab-completion
* [X] Refactor all relative paths be absolute paths
* [X] Refactor `$*` to `"$@"`
* [X] Refactor `cd` to `cd || exit`

## cmd/api/main.go

Since there is only one import, change from using `()` to a single line.

**TODO**

* [X] Refactor `import` to one line

## cmd/ui/Dockerfile

This implementation can be improved significantly by changing to a
multi-stage build and using a `scratch` image. The current image size is
`948MB`.

**TODO**

* [X] Refactor to using multi-stage build with `scratch` image

## cmd/ui/main.go

The `addr` variable is hardcoded. The implementation of route handlers is
split between `main.go` and `handlers.go`. No check to ensure `tpl`
variable is not `nil`. Missing route for `getHeartbeat` handler.

**TODO**

* [X] Add check to see if `tpl` is `nil`
* [X] Refactor `addr` to get `HOST` and `PORT` env variables
* [X] Refactor route handler logic entirely to `handlers.go`
* [X] Add `getHeartbeat` to the router

## cmd/ui/models.go

`cmd/ui` is a minimal web UI to showcase a frontend for the REST
interface, so nothing should be public.

**TODO**

* [ ] Create `models_test.go`
* [ ] Refactor all structs to be private
	* [ ] `BookResp`
	* [ ] `BooksResp`
	* [ ] `HeartbeatResp`

## cmd/ui/handlers.go

Missing test file. The `localIP` and `dockerIP` variables are hardcoded.
The `client` variable is using `InsecureSkipVerify: true`, which "should
only be used for testing or in combination with VerifyConnection or
VerifyPeerCertificate". [[source](https://pkg.go.dev/crypto/tls#Config)]
Apply DRY to for the request logic.

**TODO**

* [ ] Create `handlers_test.go`
* [ ] Refactor `localIP` to get `localHOST` and `localPORT` env variables
* [ ] Refactor `dockerIP` to get `dockerHOST` and `dockerPORT` env variables
* [ ] Remove `Transport` from `client`
* [ ] Refactor the logic for marshaling the JSON data, creating an
`http.Request`, setting the `Content-Type`, and making the request into a
utility function

## cmd/ui/utils.go

Missing test file. There should be a check of both `v` and `body` to
ensure they are not `nil` in `decodeJSON`.

**TODO**

* [ ] Create `utils_test.go`
* [ ] Add check of `v` and `body` if `nil` in `decodeJSON`

## cmd/ui/static/js/index.js

The `backend` variable is hardcoded to `http://localhost:9001`. User
should be notified if an error occurs in `deleteReq`.

**TODO**

* [ ] Implement scanning of localhost IP addresses and assign active
address to `backend` variable
* [ ] Add `alert` to `catch` in `deleteReq`

## cmd/ui/templates/header.html

Title tag should be dynamic. The `favicon.svg` file should be inlined,
which will remove the current `~85 ms` load time. It will also remove the
double request since the current implementation is using `golang.org`
which redirects to `go.dev`.

**TODO**

* [ ] Refactor `<title></title>` to a Go `text/template` Action
* [ ] Inline `favicon.svg`

## internal/api/api.go

You wrote a TODO here. Rather than cluttering source code with TODOs,
complete it if it is small, or add it to a proper backlog.

**TODO**

* [ ] Refactor `TestRun` to iterate and make a request to all routes

## internal/api/resources/resources.go

Missing test file. Missing comments.

**TODO**

* [ ] Create `resources_test.go`
* [ ] Add comments
	* [ ] `Env`
	* [ ] `NewEnv`

## internal/api/resources/books_test.go

Missing comments. All tests should be checking the `resp.Body`.

**TODO**

* [ ] Add comments
	* [ ] `TestCreateBook`
	* [ ] `TestBookList`
	* [ ] `TestNoIdBookById`
	* [ ] `TestGetBookById`
	* [ ] `TestUpdateBookById`
	* [ ] `TestDeleteBookById`
* [ ] Check `resp.Body`
	* [ ] `TestCreateBook`
	* [ ] `TestBookList`
	* [ ] `TestNoIdBookById`
	* [ ] `TestGetBookById`
	* [ ] `TestUpdateBookById`
	* [ ] `TestDeleteBookById`


## internal/api/resources/books.go

The comment for `CreateBook` needs to be more informative. The responses
need to follow a standard. Remove unnecessary `else` clause in
`BookList`.

**TODO**

* [ ] Improve comment for `CreateBook`
* [ ] Refactor all `JSONIFY` calls to follow a defined standard
* [ ] Remove `else` clause in `BookList` and return early in error check

## internal/api/resources/heartbeat_test.go

Missing comment. `TestHeartbeat` only tests the status code, it should
also check to ensure the response isn't an empty string.

**TODO**

* [ ] Add comment
	* [ ] `TestHeartbeat`
* [ ] Add check of `resp.Body` to ensure `ip` is not an empty string

## internal/api/resources/heartbeat.go

Remove unnecessary `else` clause.

**TODO**

* [ ] Remove `else` clause in `Heartbeat` and return early in error check

## internal/api/resources/utils_test.go

Missing comments. All tests should be checking `resp.Body`. The name
`TestChiURLParam` should be `TestChiURLParams`. Cleanup should be moved
to `TestChiURLParams`, since it is the last called test.

**TODO**

* [ ] Add comments
	* [ ] `TestJSONIFY`
	* [ ] `TestRecord`
	* [ ] `TestChiURLParam`
* [ ] Check `resp.Body`
	* [ ] `TestJSONIFY`
	* [ ] `TestRecord`
	* [ ] `TestChiURLParam`
* [ ] Rename `TestChiURLParam` to `TestChiURLParams`
* [ ] Refactor `t.Cleanup` to `TestChiURLParams`


## internal/api/resources/utils.go

Missing comments. All functions should be checking if their input
parameters are valid.

**TODO**

* [ ] Add comments
	* [ ] `JSON`
	* [ ] `JSONIFY`
	* [ ] `Record`
	* [ ] `ChiURLParams`
* [ ] Add input checks
	* [ ] `JSONIFY`
	* [ ] `Record`
	* [ ] `ChiURLParams`

## internal/api/router/router_test.go

Missing comments. `TestRouterRoutes` should be checking `SubRoutes` and
`Handlers` as well. Should be checking for expected middlewares.

**TODO**

* [ ] Add comments
	* [ ] `TestNewRouter`
	* [ ] `TestRouterRoutes`
* [ ] Add `SubRoutes` check to `TestRouterRoutes`
* [ ] Add `Handlers` check to `TestRouterRoutes`
* [ ] Add middleware checks
	* [ ] `RequestID`
	* [ ] `RealIP`
	* [ ] `Logger`
	* [ ] `Recoverer`
	* [ ] `Timeout`
	* [ ] `cors`

## internal/api/router/router.go

Missing comments. The `AllowedOrigins` option will need to be refactored,
since it will cause errors if/when we need to use [Request.credentials with
Fetch](https://developer.mozilla.org/en-US/docs/Web/HTTP/CORS/Errors/CORSNotSupportingCredentials).

**TODO**

* [ ] Add comments
	* [ ] `NewRouter`
* [ ] Refactor `AllowedOrigins` option

## internal/pkg/config/config.go

Missing comments. Not sure how I feel about the singleton approach, and
since I am self-reviewing, it's basically a blind leading the blind
situation. Need to ask an experienced Go programmer their opinion.

**TODO**

* [ ] Add comments
	* [ ] `Config`
	* [ ] `ServerConfig`
	* [ ] `DatabaseConfig`
* [ ] Ask experience Go programmer about use of singleton for config

## internal/pkg/database/database.go

Refactor `Open` to `OpenSQLite`. Implement functionality for `redis`,
`postgres`, and `mongodb`.

**TODO**

* [ ] Refactor `Open` to `OpenSQLite`
* [ ] Implement `OpenRedis`
* [ ] Implement `OpenPostgreSQL`
* [ ] Implement `OpenMongodb`

## internal/pkg/models/models.go

Missing comment for `Book` and comment for `NewBook` should be more
informative.

**TODO**

* [ ] Add comment for `Book`
* [ ] Improve comment for `NewBook`

## internal/pkg/repositories/books_test.go

Missing comments. The `database.Open(bookSchema)` calls are unnecessary
since `NewBooksRepo` calls `database.Open(bookSchema)` if it is passed a
nil `*sql.DB`.

**TODO**

* [ ] Add comments
	* [ ] `TestNewBooksRepo`
	* [ ] `TestInsertBook`
	* [ ] `TestSelectAllBooks`
	* [ ] `TestRetrieveBook`
	* [ ] `TestUpdateBook`
	* [ ] `TestDeleteBook`
* [ ] Refactor unnecessary `database.Open(bookSchema)` calls to `nil`
	* [ ] `TestInsertBook`
	* [ ] `TestSelectAllBooks`
	* [ ] `TestRetrieveBook`
	* [ ] `TestUpdateBook`
	* [ ] `TestDeleteBook`

## internal/pkg/repositories/books.go

Missing comment for `BooksRepo`. We should never invoke an API like
`log.Fatalf`, which will cause the program to stop. All database queries
should be transactions.

**TODO**

* [ ] Add comments
	* [ ] `BooksRepo`
* [ ] Refactor all instances of `log.Fatalf` to return the error
* [ ] Refactor all database queries to be transactions
