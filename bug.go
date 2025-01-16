```go
func main() {


    i := 0
    for {
        select {
        case <-time.After(1 * time.Second):
            fmt.Println(i)
            i++
        }
    }
}