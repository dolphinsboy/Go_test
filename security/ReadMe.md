### Code

```
package main
import (
    "net/http"
    "fmt"
)

const(
    PORT = 8080
    DOMAIN = "localhost"
    TEMPLATE = "hello"
)

func rootHandler(w http.ResponseWriter, r*http.Request){
    w.Header().Set("Content-Type", "text/html")
    w.Header().Set("Content-Length", fmt.Sprint(len(TEMPLATE)))
    w.Write([]byte(TEMPLATE))
}

func main() {
    http.HandleFunc(fmt.Sprintf("%s:%d/", DOMAIN, PORT), rootHandler)
    err :=http.ListenAndServeTLS(fmt.Sprintf(":%d", PORT), "server.crt", "server.key", nil)

    if err != nil {
        fmt.Println("ListenAndServeTLS failed:", err.Error())
    }
}
```

### 生成公钥和私钥的方法
https://stackoverflow.com/questions/10175812/how-to-create-a-self-signed-certificate-with-openssl

```
openssl genrsa -out server.key 2048
openssl rsa -in server.key -out server.key
openssl req -sha256 -new -key server.key -out server.csr -subj '/CN=localhost'
openssl x509 -req -sha256 -days 365 -in server.csr -signkey server.key -out server.crt
```

### 访问方式

```
https://localhost:8080
```

**注意是https，否则出现tls: first record does not look like a TLS handshake**


