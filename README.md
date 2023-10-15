# **Crapsolver backend API**

## **Running SurrealDB**

### Docker

```
docker run --rm --pull always -p 8000:8000 -v /mydata:/mydata surrealdb/surrealdb:1.0.0-beta.9-20230402 start --log trace --user root --pass rootnikoontoplmao5245
```

### Windows
```
surreal.exe start --user root --pass root --bind 0.0.0.0:8080 file:mydatabase.db
```