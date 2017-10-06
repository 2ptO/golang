# Remember to BET
- Benchmark
- Example
- Test

Syntax
```
BenchmarkXxx(b *testing.B)
ExampleXxx()
TestXxx(t *testing.T)
```

# Commands

```
go test
go test -bench .
go test -cover
go test -coverprofile <proflie file name>
go tool cover -html=<profile file name>

go doc -http=:8080
```