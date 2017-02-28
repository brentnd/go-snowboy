# go-snowboy

The Go bindings for snowboy audio detection (https://github.com/Kitt-AI/snowboy) are generated using swig which
 creates a lot of extra types and uses calls with variable arguments. This makes writing integrations in golang difficult
 because the types aren't explicit. go-snowboy is intended to be a wrapper around the swig-generated Go code which will
 provide Go-style usage.

## Docs
See https://godoc.org/github.com/brentnd/go-snowboy

## Dependencies
* SWIG (v 3.0.12 recommended)

### Go Packages
* github.com/Kitt-AI/snowboy/swig/Go

## Example

Example usage in `example/cmd.go`, `example/fixed.go`.

### Building
```
go build -o build/detect example/cmd.go
```

### Running
```
usage: ./build/detect <resource> <keyword.umdl> <audio file>
```

### See Also
`Makefile` has some standard targets to do `go build` steps

## TODO
* Support other functions of SnowboyDetect (Reset, GetSensitivity, UpdateModel, NumHotwords, SampleRate, NumChannels, BitsPerSample)
* Support overloads for RunDetection, Go-style
* Add DetectFrom(io.Reader) chan Keyword for continuous async detection from a stream, publishing keywords detected onto channel
* If possible, clean up the way Keywords are declared and returned?