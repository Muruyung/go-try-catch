## Getting Started

### Installation

```sh
go get github.com/Muruyung/go-try-catch@latest
```

### Usage
There is example usage using try and catch function

```go
Try(func(ThrowNewException func(any)) {
  ...
  ThrowNewException(ce)
  ...
  ThrowNewException(err)
  ...
}).Catch(func(ce CustomError) {
  ...
}).Catch(func(e error, st StackTrace) {
  ...
  st.Print()
})
```

There is example usage using try and catch block struct
```go
Block {
  Try: func(ThrowNewException func(any)) {
    ...
    ThrowNewException(err)
  },
  Catch: func(e error, st StackTrace) {
    ...
    st.Print()
  }
}.Do()
```

### Functions

| Name | Description |
| - | - |
| `Try()` | Takes `func ()`, wrap your code here! |
| `Catch()` | Takes `func (any)` or `func (any, StackTrace)`, and it will only accept the error type you have declared. You can accept second parameter, which is the stack trace begin from the last `ThrowNewException()`. |
| `ThrowNewException()` | Takes `any`. **Will only throw an error when the parameter is not** `nil`. |
| `Error()` | Return `error` value from `ThrowNewException()`. |
| `GetException()` | Return `any` value from `ThrowNewException()`. |
| `GetStackTrace()` | Return **stack trace** value from `ThrowNewException()`. |
| `st.Print()` | If you have declared parameter `st StackTrace`, you can print the stack trace using `st.Print()`. |
| `st.String()` | If you have declared parameter `st StackTrace`, you can get value of stack trace using `st.String()`. |

### Example

Let's say you want to fetch JSON from a url and unmarshal it, you can simply
write it like this:

```go
import (
  "encoding/json"
  "fmt"
  "io/ioutil"
  "net/http"

  gotry "github.com/Muruyung/go-try-catch"
)

func main() {
  // Example using function try and catch
  e := gotry.Try(func(ThrowNewException func(any)) {
    resp, err := http.Get("https://jsonplaceholder.typicode.com/posts")
    ThrowNewException(err)
    defer resp.Body.Close()

    b, err := ioutil.ReadAll(resp.Body)
    ThrowNewException(err)

    var data []map[string]interface{}
    err = json.Unmarshal(b, &data)
    ThrowNewException(err)

    fmt.Println(data)
  }).Catch(func(e error, st gotry.StackTrace) {
    fmt.Println(e)
    st.Print()
  })
  if err := e.Error(); err != nil {
    fmt.Println(err)
  }

  // Example using block struct try and catch
  e = gotry.Block{
    Try: func(ThrowNewException func(any)) {
      err := errors.New("this is an error using block exception")
      ThrowNewException(err)
    },
    Catch: func(e error, st gotry.StackTrace) {
      fmt.Println(e)
      for _, val := range st.GetList() {
        fmt.Print(val)
      }
    },
  }.Do()
  if err := e.Error(); err != nil {
    fmt.Println(err)
  }
}
```

For more examples, head over to
<https://github.com/Muruyung/go-try-catch/tree/main/.examples/main.go>!

<p align="right">(<a href="#top">back to top</a>)</p>
