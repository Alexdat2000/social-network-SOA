openssl genrsa -out credentials/private.pem 2048
openssl rsa -in credentials/private.pem -outform PEM -pubout -out credentials/public.pem
