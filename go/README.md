# GO Refctor

- Run :

```shell
go run texttest_fixture.go [<number-of-days>; default: 2]
```

- Run tests :

```shell
go test ./... -v
```

- Run tests and coverage :

```shell
go test ./... -coverprofile="coverage.out"
# to display in your browser
go tool cover -html="coverage.out"
# to write to coverage.html
go tool cover -html="coverage.out" -o coverage.html
```