# Потокобезопасность

### Гонка горутин иногда приводит к изменениям состояния любого значения, хранящегося в адресе памяти без какого-либо соблюдения порядка.

*Пример*
```go
package main

import (
    "fmt"
    "time"
)

type Counter struct {
    Count int
}

func updateC(c *Counter, cur int) {
    c.Count += cur
    fmt.Printf("New Count is %v\n", c.Count)
}

func main() {
    totalCur := 50
    c := Counter{1000000}
    for i := 0; i <= totalCur; i++ {
        go updateC(&c, int(i))
    }
    time.Sleep(time.Second * 1)  // this is just so that we don't need channels
}
```

output:
```go
New Count is 1001199
New Count is 1001030
New Count is 1001206
...
New Count is 1001176
New Count is 1000957
New Count is 1001005
```

> Как раз потомучто горутины обращались к области памяти безпорядочно мы и получаем такой несуразный результат

