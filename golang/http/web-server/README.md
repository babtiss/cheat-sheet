## HTTP-сервер

- Простейший способ запустить HTTP-сервер — вызвать `http.ListenAndServe`:

```go
package main

import (
"log"
"net/http"
)

type Handler struct {}

func (h *Handler) ServeHTTP(w http.ResponseWriter, req *http.Request){
_,err := w.Write([]byte("hello"))
if err != nil {
log.Println(err)
}
}

func main() {
log.Println(http.ListenAndServe("127.0.0.1:9090", &Handler{}))
}
```
