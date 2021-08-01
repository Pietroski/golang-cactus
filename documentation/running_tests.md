# Running Tests

For running controller tests, go to `/src/data/usecontrollers/` and simply type on terminal:

```shell
go test -v -coverprofile ./test_results/cover.out ./...
```

```shell
go tool cover -html=./test_results/cover.out -o ./test_results/cover.html
```
