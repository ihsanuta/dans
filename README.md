# DANS APPLICATION Backend

## How to run application

### setup env
```
DB_HOST=127.0.0.1
DB_PORT=3306
DB_USER=root
DB_PASS=mauFJcuf5dhRMQrjj
DB_NAME=dans

SERVER_PORT=8880
SERVER_ADDRESS=0.0.0.0

APPLICATION_NAME=dans
LOGIN_EXPIRATION_DURATION=60
JWT_SIGNATURE_KEY=dans
```

### run docker compose
```
docker-compose up -d
```

### run migration
```
migrate -database 'mysql://root:mauFJcuf5dhRMQrjj@tcp(localhost:3306)/dans?query' -path ./db/migrations up
```

### run application
```
go run .
```

### Curl REST API
#### Login
```
curl --location --request POST 'http://localhost:8880/api/v1/login' \
--header 'Content-Type: application/json' \
--data-raw '{
    "username":"admin",
    "password":"admin"
}'
```
#### Get List
```
curl --location --request GET 'http://localhost:8880/api/v1/job?type=Full Time&location=ehningen' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkYXRhIjp7ImlkIjoiIiwidXNlcm5hbWUiOiJhZG1pbiIsInBhc3N3b3JkIjoiYWRtaW4iLCJjcmVhdGVkX2F0IjoiIn0sImV4cCI6MTcxMTg5NTIxMywiaWF0IjoxNjgwMjcyODEzfQ.RyfU71StwZ7ymJXZyPsPa9TRHE059DPv6mvBLsYAHeA'
```
#### Get By ID
```
curl --location --request GET 'http://localhost:8880/api/v1/job/ecbe528e-ae60-45ad-9706-76819ae07c85' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkYXRhIjp7ImlkIjoiIiwidXNlcm5hbWUiOiJhZG1pbiIsInBhc3N3b3JkIjoiYWRtaW4iLCJjcmVhdGVkX2F0IjoiIn0sImV4cCI6MTcxMTg5NTIxMywiaWF0IjoxNjgwMjcyODEzfQ.RyfU71StwZ7ymJXZyPsPa9TRHE059DPv6mvBLsYAHeA'
```