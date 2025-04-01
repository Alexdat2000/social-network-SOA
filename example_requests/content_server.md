### Register user
```shell
curl -X POST "http://localhost:8080/users?username=alex&password=Pass1@_W&email=alex@example.com" -i
```

### Create post
```shell
curl -X POST "http://localhost:8080/entry?title=New%20post&description=Post%20content&jwt=eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFsZXgifQ.uRpnS0KeuVtnLfDHFfi41rAunuN3YU3lBonx7a8e4ikj3nmAmhWo-oArZJP5pjJgSGQ319lmibESjXAzwH4Cxz0qN1giZbS3SLNxvHx1dVbVOxk8twXCL71NhaMz2WHSQ_K32gbCXkalqUHmeiUAKLPO1vXPVWPSzr9XmmfG9QA9a-Trka47XiM8xlsDmCTYZ-K7FkbaccNyGqPWR0IbYmhyanSTE7PECTsUnFBrjwfDOsua6aixpuU2fkCMPT3R6TWfMZauPZVvU4InGwEuC1Ye0nAf1CkqD5kszJhFGVo78iT1a7OBm3pjbFBN6E0eadsREg_g5QUqsH9WKNd9qw" -i
```

### Get post
```shell
curl -X GET "http://localhost:8080/entry?postId=693151983&jwt=eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFsZXgifQ.uRpnS0KeuVtnLfDHFfi41rAunuN3YU3lBonx7a8e4ikj3nmAmhWo-oArZJP5pjJgSGQ319lmibESjXAzwH4Cxz0qN1giZbS3SLNxvHx1dVbVOxk8twXCL71NhaMz2WHSQ_K32gbCXkalqUHmeiUAKLPO1vXPVWPSzr9XmmfG9QA9a-Trka47XiM8xlsDmCTYZ-K7FkbaccNyGqPWR0IbYmhyanSTE7PECTsUnFBrjwfDOsua6aixpuU2fkCMPT3R6TWfMZauPZVvU4InGwEuC1Ye0nAf1CkqD5kszJhFGVo78iT1a7OBm3pjbFBN6E0eadsREg_g5QUqsH9WKNd9qw" -i
```

### Create private post
```shell
curl -X POST "http://localhost:8080/entry?title=Post%20two&description=Post%20content&isPrivate=true&jwt=eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFsZXgifQ.uRpnS0KeuVtnLfDHFfi41rAunuN3YU3lBonx7a8e4ikj3nmAmhWo-oArZJP5pjJgSGQ319lmibESjXAzwH4Cxz0qN1giZbS3SLNxvHx1dVbVOxk8twXCL71NhaMz2WHSQ_K32gbCXkalqUHmeiUAKLPO1vXPVWPSzr9XmmfG9QA9a-Trka47XiM8xlsDmCTYZ-K7FkbaccNyGqPWR0IbYmhyanSTE7PECTsUnFBrjwfDOsua6aixpuU2fkCMPT3R6TWfMZauPZVvU4InGwEuC1Ye0nAf1CkqD5kszJhFGVo78iT1a7OBm3pjbFBN6E0eadsREg_g5QUqsH9WKNd9qw" -i
```

### Get private post
```shell
curl -X GET "http://localhost:8080/entry?postId=644114018&jwt=eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFsZXgifQ.uRpnS0KeuVtnLfDHFfi41rAunuN3YU3lBonx7a8e4ikj3nmAmhWo-oArZJP5pjJgSGQ319lmibESjXAzwH4Cxz0qN1giZbS3SLNxvHx1dVbVOxk8twXCL71NhaMz2WHSQ_K32gbCXkalqUHmeiUAKLPO1vXPVWPSzr9XmmfG9QA9a-Trka47XiM8xlsDmCTYZ-K7FkbaccNyGqPWR0IbYmhyanSTE7PECTsUnFBrjwfDOsua6aixpuU2fkCMPT3R6TWfMZauPZVvU4InGwEuC1Ye0nAf1CkqD5kszJhFGVo78iT1a7OBm3pjbFBN6E0eadsREg_g5QUqsH9WKNd9qw" -i
```

### Register another user
```shell
curl -X POST "http://localhost:8080/users?username=alex2&password=Pass1@_W&email=alex@example.com" -i
```

### View previous post
```shell
curl -X GET "http://localhost:8080/entry?postId=644114018&jwt=eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFsZXgyIn0.ZutugQoZhoJsGmt8CF-rkNjQs-fMySGi8WHiAPKlL9cXlWocaLTYzJSuL7roCfiIqwBkHR-K_eO0Hx5wI_WFIlKdfYjIl28aZhZWRfae8BZgmY4KP_zPcl887EcjztHG5l2TWMzlvos7bK-YGHUdfaQDmrtNmtbhH-5_yXonuleJPjk5IPPYa1_MBlpSANZsK-O2e1MLd_PKYw4qi2IMXiVToCv3dBOBHXr4kERNbIV38xHBaawJw05Eq37fkYOC-jgiUsEvT9ZXEAI_nW8TVqvONy-PvgQ_ZLGy0aGL_4saj8By3woJEEgEKGucIOpqm8zrN8XeYtajOY5Qs-mT2A" -i
```

### List posts
```shell
curl -X GET "http://localhost:8080/list?page=1" -i
```

### Create third post
```shell
curl -X POST "http://localhost:8080/entry?title=Post%203&description=Post%20content&jwt=eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFsZXgifQ.uRpnS0KeuVtnLfDHFfi41rAunuN3YU3lBonx7a8e4ikj3nmAmhWo-oArZJP5pjJgSGQ319lmibESjXAzwH4Cxz0qN1giZbS3SLNxvHx1dVbVOxk8twXCL71NhaMz2WHSQ_K32gbCXkalqUHmeiUAKLPO1vXPVWPSzr9XmmfG9QA9a-Trka47XiM8xlsDmCTYZ-K7FkbaccNyGqPWR0IbYmhyanSTE7PECTsUnFBrjwfDOsua6aixpuU2fkCMPT3R6TWfMZauPZVvU4InGwEuC1Ye0nAf1CkqD5kszJhFGVo78iT1a7OBm3pjbFBN6E0eadsREg_g5QUqsH9WKNd9qw" -i
```

### Patch post
```shell
curl -X PUT "http://localhost:8080/entry?postId=1034037483&title=Updated&description=UpdatedDesc&tags=a,b&jwt=eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFsZXgifQ.uRpnS0KeuVtnLfDHFfi41rAunuN3YU3lBonx7a8e4ikj3nmAmhWo-oArZJP5pjJgSGQ319lmibESjXAzwH4Cxz0qN1giZbS3SLNxvHx1dVbVOxk8twXCL71NhaMz2WHSQ_K32gbCXkalqUHmeiUAKLPO1vXPVWPSzr9XmmfG9QA9a-Trka47XiM8xlsDmCTYZ-K7FkbaccNyGqPWR0IbYmhyanSTE7PECTsUnFBrjwfDOsua6aixpuU2fkCMPT3R6TWfMZauPZVvU4InGwEuC1Ye0nAf1CkqD5kszJhFGVo78iT1a7OBm3pjbFBN6E0eadsREg_g5QUqsH9WKNd9qw" -i
```

### List posts
```shell
curl -X GET "http://localhost:8080/list?page=1" -i
```

### Delete post
```shell
curl -X DELETE "http://localhost:8080/entry?postId=529831335&jwt=eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFsZXgifQ.uRpnS0KeuVtnLfDHFfi41rAunuN3YU3lBonx7a8e4ikj3nmAmhWo-oArZJP5pjJgSGQ319lmibESjXAzwH4Cxz0qN1giZbS3SLNxvHx1dVbVOxk8twXCL71NhaMz2WHSQ_K32gbCXkalqUHmeiUAKLPO1vXPVWPSzr9XmmfG9QA9a-Trka47XiM8xlsDmCTYZ-K7FkbaccNyGqPWR0IbYmhyanSTE7PECTsUnFBrjwfDOsua6aixpuU2fkCMPT3R6TWfMZauPZVvU4InGwEuC1Ye0nAf1CkqD5kszJhFGVo78iT1a7OBm3pjbFBN6E0eadsREg_g5QUqsH9WKNd9qw" -i
```

### List posts
```shell
curl -X GET "http://localhost:8080/list?page=1" -i
```

### Create post on alex
```shell
curl -X POST "http://localhost:8080/entry?title=New%20post&description=Post%20content&jwt=eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFsZXgifQ.uRpnS0KeuVtnLfDHFfi41rAunuN3YU3lBonx7a8e4ikj3nmAmhWo-oArZJP5pjJgSGQ319lmibESjXAzwH4Cxz0qN1giZbS3SLNxvHx1dVbVOxk8twXCL71NhaMz2WHSQ_K32gbCXkalqUHmeiUAKLPO1vXPVWPSzr9XmmfG9QA9a-Trka47XiM8xlsDmCTYZ-K7FkbaccNyGqPWR0IbYmhyanSTE7PECTsUnFBrjwfDOsua6aixpuU2fkCMPT3R6TWfMZauPZVvU4InGwEuC1Ye0nAf1CkqD5kszJhFGVo78iT1a7OBm3pjbFBN6E0eadsREg_g5QUqsH9WKNd9qw" -i
```

### Delete on alex2
```shell
curl -X DELETE "http://localhost:8080/entry?postId=1441433335&jwt=eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFsZXgyIn0.ZutugQoZhoJsGmt8CF-rkNjQs-fMySGi8WHiAPKlL9cXlWocaLTYzJSuL7roCfiIqwBkHR-K_eO0Hx5wI_WFIlKdfYjIl28aZhZWRfae8BZgmY4KP_zPcl887EcjztHG5l2TWMzlvos7bK-YGHUdfaQDmrtNmtbhH-5_yXonuleJPjk5IPPYa1_MBlpSANZsK-O2e1MLd_PKYw4qi2IMXiVToCv3dBOBHXr4kERNbIV38xHBaawJw05Eq37fkYOC-jgiUsEvT9ZXEAI_nW8TVqvONy-PvgQ_ZLGy0aGL_4saj8By3woJEEgEKGucIOpqm8zrN8XeYtajOY5Qs-mT2A" -i
```
