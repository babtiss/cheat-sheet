## Простенький сервер
### На стороне сервера можно отследить всю информацию о запросах клиентов

```go
package main

import (
"fmt"
"log"
"net/http"
"strings"
)

func printerAdditionalInfo(r *http.Request) {
// printing server information
fmt.Println(r.Form)
fmt.Println("path", r.URL.Path)
fmt.Println("scheme", r.URL.Scheme)
for k, v := range r.Form {
fmt.Println("key:", k)
fmt.Println("val:", strings.Join(v, ""))
}
}

func simpleHandler(w http.ResponseWriter, r *http.Request) {
err := r.ParseForm()
if err != nil {
return
}

printerAdditionalInfo(r)

_, err = fmt.Fprintf(w, "Hello big dady!")
if err != nil {
return
}
}

func main() {
http.HandleFunc("/", simpleHandler)
err := http.ListenAndServe(":9090", nil)
if err != nil {
log.Fatal("ListenAndServe: ", err)
}
}

```