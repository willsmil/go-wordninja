# go-wordninja
## Getting Start
```bash
go get github.com/willsmil/go-wordninja
```

## Usage
```go
package main

import (
	"github.com/willsmil/go-wordninja"
	"fmt"
)

func main()  {
	// only English characters
	eng := "thisisatest"
	fmt.Println(wordninja.CutEnglish(eng))
	
	// multi characters
	mul := "this哈isa,test"
	fmt.Println(wordninja.Cut(mul))
}
```
