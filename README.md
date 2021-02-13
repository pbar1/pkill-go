# pkill-go

Like `pkill` and `pgrep`, in pure Go without using `exec.Cmd`.

```sh
go get github.com/pbar1/pkill-go
```

### Usage

```go
package main

import (
  "fmt"

  "github.com/pbar1/pkill-go"
)

func main() {
  pids, err := pkill.Pgrep("^go$")
  if err != nil {
    fmt.Println(err)
  }
  for _, pid := range pids {
    fmt.Println(pid)
  }
}
```
