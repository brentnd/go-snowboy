# go-snowboy

The Go bindings for snowboy audio detection (https://github.com/Kitt-AI/snowboy) are generated using swig which
 creates a lot of extra types and uses calls with variable arguments. This makes writing integrations in golang difficult
 because the types aren't explicit. go-snowboy is intended to be a wrapper around the swig-generated Go code which will
 provide Go-style usage.

## Dependencies
* SWIG (v 3.0.12 recommended)
## Go Packages
* github.com/Kitt-AI/snowboy/swig/Go

## Example

Example usage in `example/main.go`.

### Building
```
go build -o build/detect example/main.go
```

### Running
```
usage: ./build/detect <resource> <keyword.umdl> <audio file>
```