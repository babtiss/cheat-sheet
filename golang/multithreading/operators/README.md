### SELECT
Оператор `select` позволяет go-процедуре находиться в ожидании нескольких операций передачи данных.
`select` блокируется до тех пор, пока один из его блоков case не будет готов к запуску, а затем выполняет этот блок.
Если сразу несколько блоков могут быть запущены, то выбирается произвольный.
Блок `default` в `select` запускается, если никакой другой блок не готов.

Рассмотрим работу оператора на примере:
```go
package main

import (
    "fmt"
    "time"
)

func main() {

    c1 := make(chan string)
    c2 := make(chan string)

    go func() {
        time.Sleep(1 * time.Second)
        c1 <- "one second"
    }()
    go func() {
        time.Sleep(3 * time.Second)
        c2 <- "three second"
    }()

    for {
        select {
            case msg1 := <-c1:
                fmt.Println("received", msg1)
            case msg2 := <-c2:
                fmt.Println("received", msg2)
        }
    }
}
```
В таком случае мы получим последовательно строки:
```
received one second
received one second
received one second
received three second
received one second
received one second
received one second
received three second
...

```
Это происходит, потомучто одна горутина выполняет запись и чтение, пока другая спит.
