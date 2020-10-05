# Mustache Proxy
This service handles mustache templates based on the passed query parameters.

## Usage
```
docker build --tag mustache-proxy .
docker run --env ALLOWED_TARGETS=www.target1.com,target2.com,traget3.co.uk --p 5555:5555 mustache-proxy --host 0.0.0.0
```
