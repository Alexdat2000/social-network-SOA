Register 2 users

```shell
curl -X 'POST' -d '{
"username": "Alex",
"password": "P@ssW0rd",
"email": "alex@example.com"
}' -H 'Accept: application/json' -H 'Content-Type: application/json' 'http://localhost:8080/users' -i
```

```shell
curl -X 'POST' -d '{
"username": "Alex2",
"password": "P@ssW0rd",
"email": "alex2@example.com"
}' -H 'Accept: application/json' -H 'Content-Type: application/json' 'http://localhost:8080/users' -i
```

Create 3 posts

```shell
curl -X 'POST' -d '{
"title": "1",
"content": "1"
}' -H 'Accept: application/json' -H 'Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6IkFsZXgifQ.rZleuSg3WaXOzbdxutIisrEdcCOfMhYKTThcj90Zxa9xjjuy25NXZNfWtzrHB3YFTXhI78pWy5S4AhK1OWk-7c2v_G15GC8Y8PA2T4YBQ8qFbRj3xMiR4VXYISaMhnP3HlcOcLctaTArmZJoan9ml9jbIyvL3PFdfgpxwQu1BlAkefaIxSwPy1QxH9HY0nDfrbxFYNeCX7eAOvzEKS9KHWIOTfTvZacD8TusEbfuvlqvuOf2VQ34rJfDkaGEUi8v7yLZY6QgpvaCbM4YWu46hOMBDtKWtiGcEjj0VDFjBb1OjBOwv523hDcUPdxoSfoavLFmmXRBPdFVtdJ6TJX9KA' \
-H 'Content-Type: application/json' 'http://localhost:8080/posts' -i
```

```shell
curl -X 'POST' -d '{
"title": "2",
"content": "2",
"isPrivate": true
}' -H 'Accept: application/json' -H 'Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6IkFsZXgifQ.rZleuSg3WaXOzbdxutIisrEdcCOfMhYKTThcj90Zxa9xjjuy25NXZNfWtzrHB3YFTXhI78pWy5S4AhK1OWk-7c2v_G15GC8Y8PA2T4YBQ8qFbRj3xMiR4VXYISaMhnP3HlcOcLctaTArmZJoan9ml9jbIyvL3PFdfgpxwQu1BlAkefaIxSwPy1QxH9HY0nDfrbxFYNeCX7eAOvzEKS9KHWIOTfTvZacD8TusEbfuvlqvuOf2VQ34rJfDkaGEUi8v7yLZY6QgpvaCbM4YWu46hOMBDtKWtiGcEjj0VDFjBb1OjBOwv523hDcUPdxoSfoavLFmmXRBPdFVtdJ6TJX9KA' \
-H 'Content-Type: application/json' 'http://localhost:8080/posts' -i
```

```shell
curl -X 'POST' -d '{
"title": "3",
"content": "3"
}' -H 'Accept: application/json' -H 'Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6IkFsZXgyIn0.cHheU27CLjqa6j63kmD-VUStrOFjirEuzV9woe_po3lBdspe4I3H3AU9rgOuLdi_RXuEVChUcPyfTYapBJMlTBBsVEYsjfXkSnNJubzn8CsdiD-q6Ery6dcBb_dPYG7Bp5GLuUM1AgUYaWVNPt1ssxB9cabQHRnjrCqVG9Q7xEfRkSpkhw9hmHVRVO2Co-8MU2Jed8ITDqArp1HYFzJY4FuzG9Vyz-8_lIZhYE_mnEaPTDliko4K_9wyxoAx4G-EORo9V5f-c-cqWuyt1_uceZe0lYEbiRB25iv3XNN4m-8Ka2PFOZLZct2hghESNtX8IU1XCy9tjXQTceG5V55VMA' \
-H 'Content-Type: application/json' 'http://localhost:8080/posts' -i
```

Like post

```shell
curl -X 'POST' -d '' -H 'Accept: application/json' -H 'Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6IkFsZXgifQ.rZleuSg3WaXOzbdxutIisrEdcCOfMhYKTThcj90Zxa9xjjuy25NXZNfWtzrHB3YFTXhI78pWy5S4AhK1OWk-7c2v_G15GC8Y8PA2T4YBQ8qFbRj3xMiR4VXYISaMhnP3HlcOcLctaTArmZJoan9ml9jbIyvL3PFdfgpxwQu1BlAkefaIxSwPy1QxH9HY0nDfrbxFYNeCX7eAOvzEKS9KHWIOTfTvZacD8TusEbfuvlqvuOf2VQ34rJfDkaGEUi8v7yLZY6QgpvaCbM4YWu46hOMBDtKWtiGcEjj0VDFjBb1OjBOwv523hDcUPdxoSfoavLFmmXRBPdFVtdJ6TJX9KA' \
-H 'Content-Type: application/json' 'http://localhost:8080/posts/1/likes' -i
```

Like your own private post

```shell
curl -X 'POST' -d '' -H 'Accept: application/json' -H 'Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6IkFsZXgifQ.rZleuSg3WaXOzbdxutIisrEdcCOfMhYKTThcj90Zxa9xjjuy25NXZNfWtzrHB3YFTXhI78pWy5S4AhK1OWk-7c2v_G15GC8Y8PA2T4YBQ8qFbRj3xMiR4VXYISaMhnP3HlcOcLctaTArmZJoan9ml9jbIyvL3PFdfgpxwQu1BlAkefaIxSwPy1QxH9HY0nDfrbxFYNeCX7eAOvzEKS9KHWIOTfTvZacD8TusEbfuvlqvuOf2VQ34rJfDkaGEUi8v7yLZY6QgpvaCbM4YWu46hOMBDtKWtiGcEjj0VDFjBb1OjBOwv523hDcUPdxoSfoavLFmmXRBPdFVtdJ6TJX9KA' \
-H 'Content-Type: application/json' 'http://localhost:8080/posts/2/likes' -i
```

Like others private post - error

```shell
curl -X 'POST' -d '' -H 'Accept: application/json' -H 'Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6IkFsZXgyIn0.cHheU27CLjqa6j63kmD-VUStrOFjirEuzV9woe_po3lBdspe4I3H3AU9rgOuLdi_RXuEVChUcPyfTYapBJMlTBBsVEYsjfXkSnNJubzn8CsdiD-q6Ery6dcBb_dPYG7Bp5GLuUM1AgUYaWVNPt1ssxB9cabQHRnjrCqVG9Q7xEfRkSpkhw9hmHVRVO2Co-8MU2Jed8ITDqArp1HYFzJY4FuzG9Vyz-8_lIZhYE_mnEaPTDliko4K_9wyxoAx4G-EORo9V5f-c-cqWuyt1_uceZe0lYEbiRB25iv3XNN4m-8Ka2PFOZLZct2hghESNtX8IU1XCy9tjXQTceG5V55VMA' \
-H 'Content-Type: application/json' 'http://localhost:8080/posts/2/likes' -i
```

Comment on posts

```shell
curl -X 'POST' -d '{"text": "Comment 1"}' -H 'Accept: application/json' -H 'Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6IkFsZXgifQ.rZleuSg3WaXOzbdxutIisrEdcCOfMhYKTThcj90Zxa9xjjuy25NXZNfWtzrHB3YFTXhI78pWy5S4AhK1OWk-7c2v_G15GC8Y8PA2T4YBQ8qFbRj3xMiR4VXYISaMhnP3HlcOcLctaTArmZJoan9ml9jbIyvL3PFdfgpxwQu1BlAkefaIxSwPy1QxH9HY0nDfrbxFYNeCX7eAOvzEKS9KHWIOTfTvZacD8TusEbfuvlqvuOf2VQ34rJfDkaGEUi8v7yLZY6QgpvaCbM4YWu46hOMBDtKWtiGcEjj0VDFjBb1OjBOwv523hDcUPdxoSfoavLFmmXRBPdFVtdJ6TJX9KA' \
-H 'Content-Type: application/json' 'http://localhost:8080/posts/1/comments' -i
```

```shell
curl -X 'POST' -d '{"text": "Comment 2"}' -H 'Accept: application/json' -H 'Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6IkFsZXgyIn0.cHheU27CLjqa6j63kmD-VUStrOFjirEuzV9woe_po3lBdspe4I3H3AU9rgOuLdi_RXuEVChUcPyfTYapBJMlTBBsVEYsjfXkSnNJubzn8CsdiD-q6Ery6dcBb_dPYG7Bp5GLuUM1AgUYaWVNPt1ssxB9cabQHRnjrCqVG9Q7xEfRkSpkhw9hmHVRVO2Co-8MU2Jed8ITDqArp1HYFzJY4FuzG9Vyz-8_lIZhYE_mnEaPTDliko4K_9wyxoAx4G-EORo9V5f-c-cqWuyt1_uceZe0lYEbiRB25iv3XNN4m-8Ka2PFOZLZct2hghESNtX8IU1XCy9tjXQTceG5V55VMA' \
-H 'Content-Type: application/json' 'http://localhost:8080/posts/1/comments' -i
```

```shell
curl -X 'POST' -d '{"text": "Comment 3"}' -H 'Accept: application/json' -H 'Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6IkFsZXgifQ.rZleuSg3WaXOzbdxutIisrEdcCOfMhYKTThcj90Zxa9xjjuy25NXZNfWtzrHB3YFTXhI78pWy5S4AhK1OWk-7c2v_G15GC8Y8PA2T4YBQ8qFbRj3xMiR4VXYISaMhnP3HlcOcLctaTArmZJoan9ml9jbIyvL3PFdfgpxwQu1BlAkefaIxSwPy1QxH9HY0nDfrbxFYNeCX7eAOvzEKS9KHWIOTfTvZacD8TusEbfuvlqvuOf2VQ34rJfDkaGEUi8v7yLZY6QgpvaCbM4YWu46hOMBDtKWtiGcEjj0VDFjBb1OjBOwv523hDcUPdxoSfoavLFmmXRBPdFVtdJ6TJX9KA' \
-H 'Content-Type: application/json' 'http://localhost:8080/posts/1/comments' -i
```

Comment on private post - error

```shell
curl -X 'POST' -d '{"text": "Comment 4"}' -H 'Accept: application/json' -H 'Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6IkFsZXgyIn0.cHheU27CLjqa6j63kmD-VUStrOFjirEuzV9woe_po3lBdspe4I3H3AU9rgOuLdi_RXuEVChUcPyfTYapBJMlTBBsVEYsjfXkSnNJubzn8CsdiD-q6Ery6dcBb_dPYG7Bp5GLuUM1AgUYaWVNPt1ssxB9cabQHRnjrCqVG9Q7xEfRkSpkhw9hmHVRVO2Co-8MU2Jed8ITDqArp1HYFzJY4FuzG9Vyz-8_lIZhYE_mnEaPTDliko4K_9wyxoAx4G-EORo9V5f-c-cqWuyt1_uceZe0lYEbiRB25iv3XNN4m-8Ka2PFOZLZct2hghESNtX8IU1XCy9tjXQTceG5V55VMA' \
-H 'Content-Type: application/json' 'http://localhost:8080/posts/2/comments' -i
```

Get comments

```shell
curl -X 'GET' -d '' -H 'Accept: application/json' -H 'Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6IkFsZXgifQ.rZleuSg3WaXOzbdxutIisrEdcCOfMhYKTThcj90Zxa9xjjuy25NXZNfWtzrHB3YFTXhI78pWy5S4AhK1OWk-7c2v_G15GC8Y8PA2T4YBQ8qFbRj3xMiR4VXYISaMhnP3HlcOcLctaTArmZJoan9ml9jbIyvL3PFdfgpxwQu1BlAkefaIxSwPy1QxH9HY0nDfrbxFYNeCX7eAOvzEKS9KHWIOTfTvZacD8TusEbfuvlqvuOf2VQ34rJfDkaGEUi8v7yLZY6QgpvaCbM4YWu46hOMBDtKWtiGcEjj0VDFjBb1OjBOwv523hDcUPdxoSfoavLFmmXRBPdFVtdJ6TJX9KA' \
-H 'Content-Type: application/json' 'http://localhost:8080/posts/1/comments?page=1' -i
```

```shell
curl -X 'GET' -d '' -H 'Accept: application/json' -H 'Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6IkFsZXgifQ.rZleuSg3WaXOzbdxutIisrEdcCOfMhYKTThcj90Zxa9xjjuy25NXZNfWtzrHB3YFTXhI78pWy5S4AhK1OWk-7c2v_G15GC8Y8PA2T4YBQ8qFbRj3xMiR4VXYISaMhnP3HlcOcLctaTArmZJoan9ml9jbIyvL3PFdfgpxwQu1BlAkefaIxSwPy1QxH9HY0nDfrbxFYNeCX7eAOvzEKS9KHWIOTfTvZacD8TusEbfuvlqvuOf2VQ34rJfDkaGEUi8v7yLZY6QgpvaCbM4YWu46hOMBDtKWtiGcEjj0VDFjBb1OjBOwv523hDcUPdxoSfoavLFmmXRBPdFVtdJ6TJX9KA' \
-H 'Content-Type: application/json' 'http://localhost:8080/posts/1/comments?page=2' -i
```
