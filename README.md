### start server
```
go run main.go
```

### example
```
curl -g 'http://localhost:8080/graphql?query={user(id:"2"){name}}'
>>>
{
  "data":{
    "user":{
      "name": "Lee"
    }
  }
}
```
