# Twirp example: Haberdasher

## Usage

### Run the server

```
go run cmd/haberdasher-server/main.go
```

### cURL

```
curl -s --header "Content-Type: application/json" --data '{"inches": 24}' http://localhost:8080/twirp/athega.Haberdasher/MakeHat | jq .
```

### Run the client

```
go run cmd/haberdasher-client/main.go
```

## Links

  - https://twitchtv.github.io/twirp/docs/intro.html
  - https://twitchtv.github.io/twirp/docs/example.html
  - https://blog.twitch.tv/en/2018/01/16/twirp-a-sweet-new-rpc-framework-for-go-5f2febbf35f/#b403
