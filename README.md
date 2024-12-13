<img align=right width="168" src="docs/gouef_logo.png">

# gouef/mode
Mode of project

[![GoDoc](https://pkg.go.dev/badge/github.com/gouef/mode.svg)](https://pkg.go.dev/github.com/gouef/mode)
[![GitHub stars](https://img.shields.io/github/stars/gouef/mode?style=social)](https://github.com/gouef/mode/stargazers)
[![Go Report Card](https://goreportcard.com/badge/github.com/gouef/mode)](https://goreportcard.com/report/github.com/gouef/mode)


## Vesions
![Stable Version](https://img.shields.io/github/v/release/gouef/mode?label=Stable&labelColor=green)
![GitHub Release](https://img.shields.io/github/v/release/gouef/mode?label=RC&include_prereleases&filter=*rc*&logoSize=diago)
![GitHub Release](https://img.shields.io/github/v/release/gouef/mode?label=Beta&include_prereleases&filter=*beta*&logoSize=diago)

## Introduction

Mode of project, like Release, Debug, Testing

## Examples

### Basic
```go
package main
import "github.com/gouef/mode"

func main()  {
    m := mode.NewBasicMode()
    
    // some code
    if m.IsRelease() {
        // some code
    }
}
```


### Additional modes
```go
package main
import "github.com/gouef/mode"

modes := []string{"staging"}
		mode := mode.NewMode(modes)

func main()  {
	modes := []string{"staging"}
	m := mode.NewMode(modes)

	// some code
	if m.IsRelease() {
		m.EnableMode("staging")
	}
	
	if m.IsMode("staging") {
		// some code
	}
}
```


