# Pages Framework Object Oriented With Go

GoPages is an experimental prototype of a web framework, inspired by [Yegor Bugayenko](https://www.yegor256.com/) with the project [jpages](https://github.com/yegor256/jpages).


Here is an example of how a page is created and output.
```go
package main

import (
	"github.com/schillermann/gopages"
)

func main() {
	page := gopages.Page{
		[]gopages.Metadata{
			{
				"Body",
				"Hello World!",
			},
		},
	}

	page.Metadata().Print()
}
```