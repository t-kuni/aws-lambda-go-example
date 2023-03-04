# aws-lambda-go-example

```
docker build -t aws-lambda-go-example .
```

```
docker-compose up
```

```
curl -XPOST "http://localhost:10000/2015-03-31/functions/function/invocations" -d '{"name": "lambda"}' 
```