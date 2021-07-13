$ make clean build install run
go clean
go build httpd/newsfeed_service.go
go install httpd/newsfeed_service.go
go run httpd/newsfeed_service.go
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:	export GIN_MODE=release
 - using code:	gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /                         --> newsfeed/httpd/handler.GetStatus.func1 (3 handlers)
[GIN-debug] GET    /newsfeed                 --> newsfeed/httpd/handler.GetAllNewsfeed.func1 (3 handlers)
[GIN-debug] POST   /newsfeed                 --> newsfeed/httpd/handler.PostNewsfeed.func1 (3 handlers)
[GIN-debug] GET    /newsfeed/:id             --> newsfeed/httpd/handler.GetNewsfeed.func1 (3 handlers)
[GIN-debug] PUT    /newsfeed/:id             --> newsfeed/httpd/handler.PutNewsfeed.func1 (3 handlers)
[GIN-debug] DELETE /newsfeed/:id             --> newsfeed/httpd/handler.DeleteNewsfeed.func1 (3 handlers)
[GIN-debug] Listening and serving HTTP on :8080
[GIN] 2021/07/13 - 15:51:38 | 201 |     252.929µs |             ::1 | POST     "/newsfeed"
[GIN] 2021/07/13 - 15:51:48 | 201 |      45.137µs |             ::1 | POST     "/newsfeed"
[GIN] 2021/07/13 - 15:52:05 | 200 |      13.424µs |             ::1 | GET      "/newsfeed/2"
[GIN] 2021/07/13 - 15:52:18 | 200 |      46.569µs |             ::1 | PUT      "/newsfeed/2"
[GIN] 2021/07/13 - 15:52:30 | 200 |  500.515198ms |             ::1 | GET      "/newsfeed?start=1&end=5"
[GIN] 2021/07/13 - 15:54:08 | 404 |       1.743µs |             ::1 | DELETE   "/newsfeed/3"
[GIN] 2021/07/13 - 15:54:12 | 201 |      40.069µs |             ::1 | POST     "/newsfeed"
[GIN] 2021/07/13 - 15:54:15 | 200 |       2.064µs |             ::1 | DELETE   "/newsfeed/3"


For POSTMAN use the postman.json for all the 5 APIs


Inspect (JavaScript):
await fetch('/newsfeed', { method:'POST', headers: {'content-type':'application/json'}, body: JSON.stringify({title:'Hello', message:'World' }) })
