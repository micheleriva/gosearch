![GoSearch](/docs/cover.png)


## Motivation

I just wanted to learn how to write a search engine from scratch without any prior experience.

## Features

- [x] Index content
- [x] Search content
- [x] Index content REST API
- [x] Search content REST API
- [ ] Delete content
- [ ] Delete content REST API
- [X] Update content
- [X] Update content REST API

## REST API

**Index a new document:** <br />
`curl -X POST http://localhost:8080/v1/insert -d "{\"content\": \"lorem ipsum\"}"`

**Search through all documents:** <br />
`curl -X GET http://localhost:8080/v1/search?q=lorem`

## Golang API

First install the package via Go Mod:

```
go get github.com/micheleriva/gosearch
```

Then you can use its native Golang APIs to index and search for documents:

```go
package main

import (
	"fmt"
	"github.com/micheleriva/gosearch"
)

func main() {
  gosearch.IndexDocument("Love is old, love is new, love is all, love is you")
  gosearch.IndexDocument("What is love? Baby don't hurt me, no more.")
  
  results := gosearch.Search("love")
  fmt.Println(results)
  // => ["Love is old, love is new, love is all, love is you", "What is love? Baby don't hurt me, no more."]
}
```

# License
gosearch is licensed under the [MIT](/LICENSE.md) license, but seriosuly, don't use it. Or do it at your own risk.
