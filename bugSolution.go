```go
package main

import (

    "fmt"
    "os"
    "os/signal"
    "syscall"
    "time"
)

func main() {
    i := 0
    done := make(chan os.Signal, 1)
    signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)

    for {
        select {
        case <-time.After(1 * time.Second):
            fmt.Println(i)
            i++
        case <-done:
            fmt.Println("Exiting...")
            return
        }
    }
}
```