### Non-existing user (request fails)
```shell
curl -X GET "http://localhost:8080/users?username=alex" -i
```

### Register user without email (request fails)
```shell
curl -X POST "http://localhost:8080/users?username=alex&password=Pass1@_W" -i
```

### Register user with weak password
```shell
curl -X POST "http://localhost:8080/users?username=alex&password=PassW0rd&email=alex@example.com" -i
```

### Register user
```shell
curl -X POST "http://localhost:8080/users?username=alex10&password=Pass1@_W&email=alex@example.com" -i
```

### Get user info (null fields are skipped)
```shell
curl -X GET "http://localhost:8080/users?username=alex" -i
```

### Login
```shell
curl -X GET "http://localhost:8080/login?username=alex&password=Pass1@_W" -i
```

### Login (incorrect password)
```shell
curl -X GET "http://localhost:8080/login?username=alex&password=Pass1@_Q" -i
```

### Add info
```shell
curl -X PATCH "http://localhost:8080/users?fieldName=firstName&newValue=Alex&jwt=eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFsZXgifQ.uRpnS0KeuVtnLfDHFfi41rAunuN3YU3lBonx7a8e4ikj3nmAmhWo-oArZJP5pjJgSGQ319lmibESjXAzwH4Cxz0qN1giZbS3SLNxvHx1dVbVOxk8twXCL71NhaMz2WHSQ_K32gbCXkalqUHmeiUAKLPO1vXPVWPSzr9XmmfG9QA9a-Trka47XiM8xlsDmCTYZ-K7FkbaccNyGqPWR0IbYmhyanSTE7PECTsUnFBrjwfDOsua6aixpuU2fkCMPT3R6TWfMZauPZVvU4InGwEuC1Ye0nAf1CkqD5kszJhFGVo78iT1a7OBm3pjbFBN6E0eadsREg_g5QUqsH9WKNd9qw" -i
```

### More info
```shell
curl -X PATCH "http://localhost:8080/users?fieldName=dateOfBirth&newValue=1970-11-15&jwt=eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFsZXgifQ.uRpnS0KeuVtnLfDHFfi41rAunuN3YU3lBonx7a8e4ikj3nmAmhWo-oArZJP5pjJgSGQ319lmibESjXAzwH4Cxz0qN1giZbS3SLNxvHx1dVbVOxk8twXCL71NhaMz2WHSQ_K32gbCXkalqUHmeiUAKLPO1vXPVWPSzr9XmmfG9QA9a-Trka47XiM8xlsDmCTYZ-K7FkbaccNyGqPWR0IbYmhyanSTE7PECTsUnFBrjwfDOsua6aixpuU2fkCMPT3R6TWfMZauPZVvU4InGwEuC1Ye0nAf1CkqD5kszJhFGVo78iT1a7OBm3pjbFBN6E0eadsREg_g5QUqsH9WKNd9qw" -i
```

### Get again
```shell
curl -X GET "http://localhost:8080/users?username=alex" -i
```

### Overwrite
```shell
curl -X PATCH "http://localhost:8080/users?fieldName=dateOfBirth&newValue=1984-04-23&jwt=eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFsZXgifQ.uRpnS0KeuVtnLfDHFfi41rAunuN3YU3lBonx7a8e4ikj3nmAmhWo-oArZJP5pjJgSGQ319lmibESjXAzwH4Cxz0qN1giZbS3SLNxvHx1dVbVOxk8twXCL71NhaMz2WHSQ_K32gbCXkalqUHmeiUAKLPO1vXPVWPSzr9XmmfG9QA9a-Trka47XiM8xlsDmCTYZ-K7FkbaccNyGqPWR0IbYmhyanSTE7PECTsUnFBrjwfDOsua6aixpuU2fkCMPT3R6TWfMZauPZVvU4InGwEuC1Ye0nAf1CkqD5kszJhFGVo78iT1a7OBm3pjbFBN6E0eadsREg_g5QUqsH9WKNd9qw" -i
```

### Check
```shell
curl -X GET "http://localhost:8080/users?username=alex" -i
```

### Some examples of validations
#### Wrong date of birth format
```shell
curl -X PATCH "http://localhost:8080/users?fieldName=dateOfBirth&newValue=02-01-1970&jwt=eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFsZXgifQ.uRpnS0KeuVtnLfDHFfi41rAunuN3YU3lBonx7a8e4ikj3nmAmhWo-oArZJP5pjJgSGQ319lmibESjXAzwH4Cxz0qN1giZbS3SLNxvHx1dVbVOxk8twXCL71NhaMz2WHSQ_K32gbCXkalqUHmeiUAKLPO1vXPVWPSzr9XmmfG9QA9a-Trka47XiM8xlsDmCTYZ-K7FkbaccNyGqPWR0IbYmhyanSTE7PECTsUnFBrjwfDOsua6aixpuU2fkCMPT3R6TWfMZauPZVvU4InGwEuC1Ye0nAf1CkqD5kszJhFGVo78iT1a7OBm3pjbFBN6E0eadsREg_g5QUqsH9WKNd9qw" -i
```

#### Incorrect email
```shell
curl -X PATCH "http://localhost:8080/users?fieldName=email&newValue=alex@dat@gmail.com&jwt=eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFsZXgifQ.uRpnS0KeuVtnLfDHFfi41rAunuN3YU3lBonx7a8e4ikj3nmAmhWo-oArZJP5pjJgSGQ319lmibESjXAzwH4Cxz0qN1giZbS3SLNxvHx1dVbVOxk8twXCL71NhaMz2WHSQ_K32gbCXkalqUHmeiUAKLPO1vXPVWPSzr9XmmfG9QA9a-Trka47XiM8xlsDmCTYZ-K7FkbaccNyGqPWR0IbYmhyanSTE7PECTsUnFBrjwfDOsua6aixpuU2fkCMPT3R6TWfMZauPZVvU4InGwEuC1Ye0nAf1CkqD5kszJhFGVo78iT1a7OBm3pjbFBN6E0eadsREg_g5QUqsH9WKNd9qw" -i
```

#### Incorrect phone number
```shell
curl -X PATCH "http://localhost:8080/users?fieldName=phone&newValue=betterCallSaul&jwt=eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFsZXgifQ.uRpnS0KeuVtnLfDHFfi41rAunuN3YU3lBonx7a8e4ikj3nmAmhWo-oArZJP5pjJgSGQ319lmibESjXAzwH4Cxz0qN1giZbS3SLNxvHx1dVbVOxk8twXCL71NhaMz2WHSQ_K32gbCXkalqUHmeiUAKLPO1vXPVWPSzr9XmmfG9QA9a-Trka47XiM8xlsDmCTYZ-K7FkbaccNyGqPWR0IbYmhyanSTE7PECTsUnFBrjwfDOsua6aixpuU2fkCMPT3R6TWfMZauPZVvU4InGwEuC1Ye0nAf1CkqD5kszJhFGVo78iT1a7OBm3pjbFBN6E0eadsREg_g5QUqsH9WKNd9qw" -i
```

#### Incorrect field
```shell
curl -X PATCH "http://localhost:8080/users?fieldName=bio&newValue=betterCallSaul&jwt=eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFsZXgifQ.uRpnS0KeuVtnLfDHFfi41rAunuN3YU3lBonx7a8e4ikj3nmAmhWo-oArZJP5pjJgSGQ319lmibESjXAzwH4Cxz0qN1giZbS3SLNxvHx1dVbVOxk8twXCL71NhaMz2WHSQ_K32gbCXkalqUHmeiUAKLPO1vXPVWPSzr9XmmfG9QA9a-Trka47XiM8xlsDmCTYZ-K7FkbaccNyGqPWR0IbYmhyanSTE7PECTsUnFBrjwfDOsua6aixpuU2fkCMPT3R6TWfMZauPZVvU4InGwEuC1Ye0nAf1CkqD5kszJhFGVo78iT1a7OBm3pjbFBN6E0eadsREg_g5QUqsH9WKNd9qw" -i
```

### All of that didn't break profile
```shell
curl -X GET "http://localhost:8080/users?username=alex" -i
```
