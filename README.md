# Introduction

I am trying to learn go by writing small programs. This one is
inspired by `httpbin.org`.

Here are the features:

### Headers

Requesting `http://localhost:8799/headers/` will yield a response
like below:

```json
{
  "headers": {
    "Accept": "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8",
    "Accept-Encoding": "gzip, deflate",
    "Accept-Language": "en-US,en;q=0.5",
    "Connection": "keep-alive",
    "Cookie": "csrftoken=vMAAEmcfMPSjToiFqDtq7O6Nhoack1pQTZdbcpjFqieASJi9whARd6CqvOa57uvY; _ga=GA1.1.1395202726.1515485876; hblid=GnGY8MdTRPm5BQjO3m39N0J02OF0Ao2x; olfsk=olfsk824434969431056; sessionid=xx6ft2gxb5t7hm2dkm60rl1sw1wh9mmn",
    "Upgrade-Insecure-Requests": "1",
    "User-Agent": "Mozilla/5.0 (X11; Linux x86_64; rv:60.0) Gecko/20100101 Firefox/60.0"
  },
  "ip": "[::1]"
}
// GET http://localhost:8799/headers/
// HTTP/1.1 200 OK
// Content-Type: application/json
```


### User agent

Requesting `http://localhost:8799/user-agent/` will yield a response
like below:

```json
{
  "user_agent": "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/66.0.3359.181 Safari/537.36",
  "ip": "[::1]"
}
// GET http://localhost:8799/user-agent/
// HTTP/1.1 200 OK
// Content-Type: application/json
```


### Time

Return current time in different formats:

Sending a request to `http://localhost:8799/now/` will yield a
response like below:


```json
{
  "rfc3339": "2018-06-07T13:24:07+05:45",
  "ansi_c": "Thu Jun  7 13:24:07 2018",
  "unix_date": "Thu Jun  7 13:24:07 +0545 2018",
  "unix_seconds": 1528357147,
  "unix_nano_seconds": 1528357147091860422,
  "ip": "[::1]"
}
// POST http://localhost:8799/now/
// HTTP/1.1 200 OK
// Content-Type: application/json
```


### Status code

Return a response with the same HTTP status code specified in request. Example:

A request on `http://localhost:8799/status/400/` will return a
response like this:

```json
{
  "message": "Bad Request",
  "status_code": 400,
  "ip": "[::1]",
  "method": "GET"
}
// GET http://localhost:8799/status/400/
// HTTP/1.1 400 Bad Request
// Content-Type: application/json
```

Other HTTP methods such as `PATCH`, `PUT` etc. are supported.
