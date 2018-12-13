# golang-circuit-breaker
Hystrix Circuit Breaker implemented in Golang

### Prerequisites
* [Go](https://golang.org)
* [Dep](https://github.com/golang/dep) - Go dependency management tool 

## Installing

Install dependencies

```bash
$ make deps
```

Next, start both downstream and upstream services

```
$ go run downstream/main.go
$ go run upstream/main.go
```

Finally, you can try how circuit breaker work by shutting down the downstream service

## Author

By [Reawpai Chunsoi](https://github.com/phaicom/)

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details

## Acknowledgments
* [hystrix-go](https://github.com/afex/hystrix-go)
* [Circuit Breaker and Retry by Dan Tran](https://medium.com/@trongdan_tran/circuit-breaker-and-retry-64830e71d0f6)