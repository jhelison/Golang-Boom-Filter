# Golang bloom filter

This repository contains a basic implementation of a Bloom filter using Golang.

## How to use

```
package main

import (
    "bloom"
)

func main() {
    bloomFilter := bloom.NewBloom(1000, 5, bloom.FNVHashStrategy{})
    bloomFilter.Add([]byte("hello"))
    bloomFilter.Add([]byte("world"))

    // Check if elements are possibly in the set
    exists := bloomFilter.Check([]byte("hello")) // returns true
    notExists := bloomFilter.Check([]byte("unknown")) // returns false (likely)
}

```

## How to test

To run all tests, execute:

```
go test
```
