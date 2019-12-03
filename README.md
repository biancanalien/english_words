# English vocabularies

Just a personal application to reminder the new vocabularies of english class and study GO API

## Usage

- Create vocabulary

``` bash
curl -X POST \
  http://localhost:8080/ \
  -H 'Accept: */*' \
  -H 'Accept-Encoding: gzip, deflate' \
  -H 'Cache-Control: no-cache' \
  -H 'Connection: keep-alive' \
  -H 'Content-Length: 131' \
  -H 'Content-Type: application/json' \
  -H 'Host: localhost:8080' \
  -d '{
	"Expression": "",
	"Meaning": "",
	"Translation": ""
}'
```

- Update vocabulary by {id}

``` bash
curl -X PUT \
  http://localhost:8080/{id} \
  -H 'Accept: */*' \
  -H 'Accept-Encoding: gzip, deflate' \
  -H 'Cache-Control: no-cache' \
  -H 'Connection: keep-alive' \
  -H 'Content-Length: 30' \
  -H 'Content-Type: text/plain' \
  -H 'Host: localhost:8080' \
  -H 'cache-control: no-cache' \
  -d '{
	"Translation": ""
}'
```

- Get all vocabularies

``` bash
curl -X GET \
  http://localhost:8080/ \
  -H 'Accept: */*' \
  -H 'Accept-Encoding: gzip, deflate' \
  -H 'Cache-Control: no-cache' \
  -H 'Connection: keep-alive' \
  -H 'Host: localhost:8080' \
  -H 'cache-control: no-cache'
```