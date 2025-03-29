### Register user
```shell
curl -X POST "http://localhost:8080/users?username=alex&password=Pass1@_W&email=alex@example.com" -i
```

### Create post
```shell
curl -X POST "http://localhost:8080/entry?title=New%20post&description=Post%20content&jwt=eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFsZXgifQ.uRpnS0KeuVtnLfDHFfi41rAunuN3YU3lBonx7a8e4ikj3nmAmhWo-oArZJP5pjJgSGQ319lmibESjXAzwH4Cxz0qN1giZbS3SLNxvHx1dVbVOxk8twXCL71NhaMz2WHSQ_K32gbCXkalqUHmeiUAKLPO1vXPVWPSzr9XmmfG9QA9a-Trka47XiM8xlsDmCTYZ-K7FkbaccNyGqPWR0IbYmhyanSTE7PECTsUnFBrjwfDOsua6aixpuU2fkCMPT3R6TWfMZauPZVvU4InGwEuC1Ye0nAf1CkqD5kszJhFGVo78iT1a7OBm3pjbFBN6E0eadsREg_g5QUqsH9WKNd9qw" -i
```
