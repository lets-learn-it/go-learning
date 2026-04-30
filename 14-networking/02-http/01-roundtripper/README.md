# RoundTripper

## Program Execution

```sh
$ go run .
[headerRoundTripper] before request sent
[basicAuthRoundTripper] before request sent
[basicAuthRoundTripper] after request sent
[headerRoundTripper] after request sent
200 OK
{
  "authenticated": true, 
  "user": "pskp"
}
```

## Sequence Diagram

```mermaid
sequenceDiagram
    participant Client
    participant HTTPClient as http.Client
    participant HeaderRT as headerRoundTripper
    participant BasicAuthRT as basicAuthRoundTripper
    participant Transport as http.Transport
    participant Server

    Client->>HTTPClient: Do(req)
    HTTPClient->>HeaderRT: RoundTrip(req)

    HeaderRT->>HeaderRT: Add custom headers
    HeaderRT->>BasicAuthRT: RoundTrip(req)

    BasicAuthRT->>BasicAuthRT: Add Authorization header (Basic)
    BasicAuthRT->>Transport: RoundTrip(req)

    Transport->>Server: Send HTTP request
    Server-->>Transport: HTTP response

    Transport-->>BasicAuthRT: Response
    BasicAuthRT-->>HeaderRT: Response
    HeaderRT-->>HTTPClient: Response
    HTTPClient-->>Client: Response
```
