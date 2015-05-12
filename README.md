# badhttp

An HTTP server that can be configured, via simple HTTP PATCH requests, to do bad things.
The server associates each configuration with a request name and internally
holds a `map[string]config` to determine what to do when a `/{request_name}`
comes in.

# Admin requests you can make

## `PATCH /admin/{request_name}/delay/{delay_in_seconds}`
Make all requests to `/{request_name}` take `{delay_in_seconds}` seconds before they return

## `PATCH /admin/{request_name}/code/{http_code}`
Make all requests to `/{request_name}` return with `{http_code}`
