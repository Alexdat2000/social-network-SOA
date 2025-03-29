### Register user
```shell
curl -X POST "http://localhost:8080/users?username=alex&password=Pass1@_W&email=alex@example.com" -i
```

### Login
```shell
curl -X GET "http://localhost:8080/login?username=alex&password=Pass1@_W" -i
```
