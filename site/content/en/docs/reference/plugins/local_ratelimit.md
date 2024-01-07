---
title: Local Ratelimit
---

## Description

This plugin limits the number of requests per second, by leveraging Envoy's `local_ratelimit` filter. The limiter is run before authentication.

## Attribute

|       |         |
|-------|---------|
| Type  | Traffic |
| Order | Pre     |

## Configuration

See the corresponding [Envoy documentation](https://www.envoyproxy.io/docs/envoy/v1.28.0/configuration/http/http_filters/local_rate_limit_filter).

## Usage

Assumed we have the HTTPRoute below attached to `localhost:10000`, and a backend server listening to port `8080`:

```yaml
apiVersion: gateway.networking.k8s.io/v1
kind: HTTPRoute
metadata:
  name: default
spec:
  parentRefs:
  - name: default
    namespace: default
  rules:
  - matches:
    - path:
        type: PathPrefix
        value: /
    backendRefs:
    - name: backend
      port: 8080
```

By applying the configuration below, the rate of requests to `http://localhost:10000/` is limited to 1 request per second:

```yaml
apiVersion: mosn.io/v1
kind: HTTPFilterPolicy
metadata:
  name: policy
spec:
  targetRef:
    group: gateway.networking.k8s.io
    kind: HTTPRoute
    name: default
  filters:
    local_ratelimit:
      config:
        stat_prefix: http_policy_local_rate_limiter
        token_bucket:
          max_tokens: 1
          tokens_per_fill: 1
          fill_interval: 1s
        filter_enabled:
          default_value:
            numerator: 100
            denominator: HUNDRED
        filter_enforced:
          default_value:
            numerator: 100
            denominator: HUNDRED
```

We can test it out:

```
$ while true; do curl -I http://localhost:10000/ 2>/dev/null | head -1 ;done
HTTP/1.1 200 OK
HTTP/1.1 429 Too Many Requests
HTTP/1.1 429 Too Many Requests
```

```
$ while true; do curl -I http://localhost:10000/ 2>/dev/null | head -1 ; sleep 1; done
HTTP/1.1 200 OK
HTTP/1.1 200 OK
```