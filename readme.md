Alexa Ranking for go [![Build Status](https://travis-ci.org/negah/alexa.svg?branch=master)](https://travis-ci.org/negah/alexa)
====================
Use Alexa public rank api in golang.

Install
-------
```shell
go get github.com/negah/alexa
```

Example
-------
```go
package main

import(
  "github.com/negah/alexa"
  "fmt"
)

func main(){
  url := "google.com"

  globalRank, err := alexa.GlobalRank(url)

  if err != nil {
    fmt.Printf("%s", err)
  } else {
    fmt.Printf("%s rank in alexa is %s\n", url, globalRank)
  }

  countryRank, countryName, _, err := alexa.CountryRank(url)

  if err != nil {
    fmt.Printf("%s", err)
  } else {
    fmt.Printf("%s has rank %s in %s", url, countryRank, countryName )
  }
}

```

License
-------
MIT

Author
------
Dariush Abbasi (a.k.a [@dariubs](https://github.com/dariubs) )
