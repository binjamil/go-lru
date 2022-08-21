# go-lru

A generic, thread-safe LRU cache library, implemented via a generic doubly linked list (also available separately in package "github.com/binjamil/go-lru/list")

## Documentation

For complete docs, check out [Go Packages](https://pkg.go.dev/github.com/binjamil/go-lru)

## Installation

```bash
go get github.com/binjamil/go-lru
```

## Usage

```go
import "github.com/binjamil/go-lru"

func Consumer() {
    lru, err := lru.New[int, string](69)
    if err != nil {
        panic(err)
    }

    for i := 0; i < 69; i++ {
        lru.Add(i, "Lalo Salamanca") // Type-safe Add
    }

    var val string
    val, ok := lru.Get(0) // Type-safe Get
    if ok {
        println(val)
    }
}
```
