## Gin middleware for Apex-log

Middleware for [gin-gonic](https://github.com/gin-gonic/gin) and [apex-log](https://github.com/apex/log)

[![GoDoc](https://godoc.org/github.com/thedanielforum/apex-gin?status.svg)](https://godoc.org/github.com/thedanielforum/apex-gin)

Example output
```text
INFO[0005] client_ip=127.0.0.1 latency=168.439199ms method=PUT path=/v1/user/info/123 status_code=200 message=test message
```

*Install*
```bash
go get https://github.com/thedanielforum/apex-gin-middleware
```

*How to use*
```go
// Configure apex-log handler
log.SetHandler(text.New(os.Stderr))

// Configure gin
r := gin.New()

// Tell gin to use apex_gin middleware
r.Use(apex_gin.Handler("test message"))

// Define routes
r.GET("/ping", func(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
})

// Start HTTP server
if err := r.Run(":8080"); err != nil {
	log.WithError(err).Fatal("server failed to start")
}
```

More info about the apex-log package [medium](https://medium.com/@tjholowaychuk/apex-log-e8d9627f4a9a)


TODO
* Apex log levels info/warning/error/fatal
* Unit-testing
