# Пакет net/http

### Пакет http предоставляет реализации HTTP клиента и сервера.

* Get, Head, Post, и PostForm выполняют HTTP (или HTTPS) запросы: *

```go
resp, err := http.Get("http://example.com/")
...
resp, err := http.Post("http://example.com/upload", "image/jpeg", &buf)
...
resp, err := http.PostForm("http://example.com/form",
url.Values{"key": {"Value"}, "id": {"123"}})
```