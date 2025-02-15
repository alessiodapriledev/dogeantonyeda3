# dogeantonyeda3
spot vaddias
* ### chiamate per ottenere il JWT TOKEN
```bash
curl -X POST http://localhost:5000/login -H "Content-Type: application/json" -d '{"username": "admin", "password": "password"}'
```
* ### chiamata per ottenere la pagina iniziale

```bash
curl -X GET "http://localhost:5000/" -H "Authorization: Bearer MY_BEARER_TOKEN"
```

* ### chiamata generica al layer API
```bash
curl -X GET "http://localhost:5000/api/v0/*" -H "Authorization: Bearer MY_BEARER_TOKEN"
```

* ### chiamata POST
```bash
curl -X POST http://localhost:5000/api/v0/alerts \
  -H "Authorization: Bearer MY_BEARER_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "lat": 40.7128,
    "lng": -74.0060,
    "message": "This is a test alert message"
  }'
```

* ### chiamata GET
```bash
curl -X GET http://localhost:5000/api/v0/alerts \
  -H "Authorization: Bearer MY_BEARER_TOKEN"
```
