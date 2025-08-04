# random-person-selector
A simple tool to randomly and fairly select a name from a list

## Language

I figured this was a nice chance to have a play with [Go](https://go.dev/doc/tutorial/getting-started).

You'll need go installed: 

```bash
brew install go
```

and you can then check the success of this with: 

```bash
go version
```

## Running

To run a go script, you can run: 

```bash
go run .
```

This relies the `go.mod` file to infer the entry point and what to run. 

## Formatting

Go is quite nice for formatting. 

You can run: 

```bash
gofmt -w main.go
```

to format a file. The `-w` flag will write to the source file, otherwise it goes to STDOUT. 

## Testing

To run the tests, you just run: 

```bash
go test .
```

The naming convention is `*_test.go`. 