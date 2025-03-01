# Gateway Service

Accepts requests from users (web interface) with REST API. Confirms request correctness with Users service (checks that
user has valid session and has rights to execute the request). Then if the request is related to content, it is sent to
the Content service.
