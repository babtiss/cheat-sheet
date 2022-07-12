# Пакет net/http

- Пакет http предоставляет реализации HTTP клиента и сервера.

### Get, Head, Post, и PostForm выполняют HTTP (или HTTPS) запросы:

```go
resp, err := http.Get("http://www.01happy.com/demo/accept.php?id=1")
...
resp, err := http.Post("http://example.com/", "application/x-www-form-urlencoded", &buf)
...
resp, err := http.PostForm("http://example.com/form", url.Values{"key": {"Value"}, "id": {"0"}})
```

- Тело ответа должно быть закрыто клиентом после работы (в последующих примерах этот пункт будет сокращен)

```go
resp, err := http.Get("http://www.01happy.com/demo/accept.php?id=1")
defer resp.Body.Close()
body, err := ioutil.ReadAll(resp.Body)
if err != nil {
// handle error
}
```

### Client
- Для контроля над клиентскими HTTP заголовками, политикой перенаправлений, и другими настройками используют `Client`:

```go
client := &http.Client{
CheckRedirect: redirectPolicyFunc,
}
resp, err := client.Get("http://example.com")
...
```

### Сложные запросы
- Если вам нужно установить параметры заголовка запроса, куки и другие данные, используйте метод `http.Do`

```go
func httpDo() {
client := &http.Client{}
req, err := http.NewRequest("POST", "http://www.01happy.com/demo/accept.php", strings.NewReader("name=cjb"))
if err != nil {
// handle error
}
req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
req.Header.Set("Cookie", "name=anny")
resp, err := client.Do(req)
defer resp.Body.Close()
body, err := ioutil.ReadAll(resp.Body)
if err != nil {
// handle error
}
fmt.Println(string(body))
}
```


### Transport
- Для контроля над прокси, TLS конфигурацией, сохранения открытого соединения (keep-alive), сжатием, и другими настройками, используют Transport:

```go
tr := &http.Transport{
MaxIdleConns:       10,
IdleConnTimeout:    30 * time.Second,
DisableCompression: true,
}
client := &http.Client{Transport: tr}
resp, err := client.Get("https://example.com")
...
```

- Клиенты и транспорты безопасны для одновременного использования несколькими горутинами.
Для повышения эффективности их следует создавать только один раз и использовать повторно.

### Handler
- ListenAndServe стартует HTTP-сервер с заданным адресом и обработчиком. Обработчик обычно равен nil, что означает использовать `DefaultServeMux`(мультиплексор HTTP-запросов).
Handle и HandleFunc добавляют обработчики к `DefaultServeMux`:

```go
http.Handle("/foo", fooHandler)

http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
fmt.Fprintf(w, "Привет, %q", html.EscapeString (r.URL.Path))
})

log.Fatal(http.ListenAndServe(":8080", nil))
```

### Server

- Больше контроля над поведением сервера доступно при создании собственного сервера:
```go
s := &http.Server{ address: ":8080",
handler: myHandler,
ReadTimeout: 10 * time.Second,
WriteTimeout: 10 * time.Second,
MaxHeaderBytes: 1 << 20,
}
log.Fatal(s.ListenAndServe(s.ListenAndServe(s.ListenAndServe(s.ListenAndServe)))
```