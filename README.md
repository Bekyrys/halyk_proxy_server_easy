# Server for proxying requests

## Introduction

The project was created for educational purposes and is an HTTP proxy server written in Golang.  It accepts requests, sends them to the specified external services, and returns a response in JSON format.
## Requirements
- Docker
- Docker Compose
- Golang
- 
## How to install 
1. Clone the repository:
   ```sh
   git clone https://github.com/Bekyrys/halyk_proxy_server_easy/
    ```
  ```sh
     cd halyk_proxy_server_easy
```
2. Build and run using Docker:
```sh
make build
```
```sh
make up
```
3. Run the server:

```sh
make run
```
## Deployed on RENDER:
### Base URL
https://halyk-proxy-server-easy.onrender.com
### Request Body:
```sh
{
  "method": "GET",           
  "url": "http://example.com", 
  "headers": { "Authorization": "Bearer token" }
}
```
### Response:
```sh
{
  "id": "requestId",         
  "status": 200,            
  "headers": { <array of headers from a third-party service response> },
  "length": 1234             
}
```

### Response Example:
```sh
{
"method": "GET",
"url": "http://google.com",
"headers": { "Authorization": "Bearer token" }
}

{
    "id": "52d5966f-b2dd-4cd3-8949-7949c5cb3d18",
    "status": 200,
    "headers": {
        "Cache-Control": "private, max-age=0",
        "Content-Security-Policy-Report-Only": "object-src 'none';base-uri 'self';script-src 'nonce-0qpsjpe1eaXGNAd545gnpQ' 'strict-dynamic' 'report-sample' 'unsafe-eval' 'unsafe-inline' https: http:;report-uri https://csp.withgoogle.com/csp/gws/other-hp",
        "Content-Type": "text/html; charset=ISO-8859-1",
        "Date": "Tue, 09 Jul 2024 08:25:33 GMT",
        "Expires": "-1",
        "P3p": "CP=\"This is not a P3P policy! See g.co/p3phelp for more info.\"",
        "Server": "gws",
        "Set-Cookie": "AEC=AVYB7crN-RFuiKsAueY3wqwawBhzXgR5AzoWkmHva-EIdBWcjsidirIbmfU; expires=Sun, 05-Jan-2025 08:25:33 GMT; path=/; domain=.google.com; Secure; HttpOnly; SameSite=lax",
        "X-Frame-Options": "SAMEORIGIN",
        "X-Xss-Protection": "0"
    },
    "length": 21189
}
```
