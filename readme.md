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

  globalRank := alexa.GlobalRank(url)
  fmt.Printf("%s rank in alexa is %s\n", url, globalRank)

  countryRank, countryName, _ := alexa.CountryRank(url)
  fmt.Printf("%s has rank %s in %s", url, countryRank, countryName )
}

```

License
-------
MIT

Author
------
Dariush Abbasi (a.k.a [@dariubs](https://github.com/dariubs) )
