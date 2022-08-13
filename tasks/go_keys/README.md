## Интересные задачи на go

#### 1. Slice, capasity, append.
Что выведет следующий код?
```go
package main

import "fmt"

func main() {
    slice := []int{1, 2}
    slice = append(slice, 3)
    x := append(slice, 4)
    x = append(slice, 5)
    y := append(slice, 6)
    fmt.Println(x, y)
}
```
<details>
    <summary>Ответ</summary>
    [1 2 3 6] [1 2 3 6]
</details>