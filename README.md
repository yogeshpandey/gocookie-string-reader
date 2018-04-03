# Cookie string parser for go

Reads cookie string into http.Cookie 

## Getting Started

Get library
```bash
go get github.com/yogeshpandey/gocookie-string-reader
```

Use library

```go
import cookieParser "github.com/yogeshpandey/gocookie-string-reader"
import "fmt"



c:=cookieParser.ParseToCookie("Cookie-1=v$1; c2=v2")
fmt.Println("%#v",c)

```