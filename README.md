## wktfile - Go package for easy reading of WKT files
[![Go Report Card](https://goreportcard.com/badge/github.com/tomplex/wktfile)](https://goreportcard.com/report/github.com/tomplex/wktfile)


Package `wktfile` aims to make it easy to ingest and work with WKT files in Go.

Read the [godoc here](https://godoc.org/github.com/tomplex/wktfile).

### Quickstart

First, run:

`$ go get github.com/tomplex/wktfile`

then, create some data:

testdata.wkt
```
point_id|wkt
1|POINT(-72.549738631281 43.724067830418)
2|POINT(-73.3208999067471 44.9213993671176)
3|POINT(-72.3267250008435 44.3847434287252)
```


main.go
```go
package main

import (
    "fmt"
    "github.com/tomplex/wktfile"
)

func main() {
    wkt, err := wktfile.Read("./testdata.wkt")

    if err != nil {
        fmt.Println("error reading file")
    }

    fmt.Println(wkt.Header) // access the file's header as a []string

    // the file's rows are stored as a [][]string
    for _, row := range wkt.Rows {
        wktgeom := row[1]
        // do stuff with the geometry
    }

    fmt.Println(wkt.FilePath)

}

```

That's all there is to it, and also all the functionality this package provides at this time. The package defaults assume that the file is pipe ( "|" ) delimited and contains a header.

### Customization

Sometimes data doesn't follow normal standards. `wktfile` provides some useful helpers to account for this. The function `wktfile.Read()` accepts a number of functional parameters that describe the file being read, all of which currently pertain to the file's delimiter or presence of a file header. As en example:

```go
// Specify a comma delimiter
commaWKT, err := wktfile.Read("./testdata.wkt", wktfile.CommaDelimited)

// Specify a custom delimiter
customWKT, err := wktfile.Read("./testdata.wkt", wktfile.CustomDelimiter(';'))

// Speficy that the file has no header
headlessWKT, err := wktfile.Read("./testdata.wkt", wktfile.HasNoHeader)
```

More information can be found in the [godoc](https://godoc.org/github.com/tomplex/wktfile).

### TODO

 - Allow file contents to be represented as a `[]map[string]string`
 - Custom adapters to create Go `geometry` types from the file's data?