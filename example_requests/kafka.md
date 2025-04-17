### Register user
```shell
res=$(curl -X POST "http://localhost:8080/users?username=alexdat&password=Pass1@_W&email=alex@example.com")
jwt="${res:1:-1}"
```

### Post
```shell
curl -X POST "http://localhost:8080/entry?title=New%20post&description=Post%20content&jwt=$jwt" -i
```

### Get post
```shell
curl -X GET "http://localhost:8080/entry?postId=352604672&&jwt=$jwt" -i
```

### Like
```shell
curl -X POST "http://localhost:8080/like?postId=352604672&&jwt=$jwt" -i
```

### Comment
```shell
curl -X POST "http://localhost:8080/comment?postId=352604672&&text=my%20comment&jwt=$jwt" -i
```
