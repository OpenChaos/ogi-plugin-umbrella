
### HTTP Service Consumer for Ogi

> it is a consumer plug-in for Ogi, it is a simple HTTP Service which consumes Body at a route and passes it to configured Transformer
>
>  `environment variables` & `basic authentication` can be enabled/disabled for consume route

#### Environment Variable Configuration

* `CONSUMER_API_LISTENAT` provides port to bind at, value format is `:8080`.

* `CONSUMER_API_BASICAUTH_ENABLED` allows to enable Basic Auth which is the only way to configure it over API, by default `true`.

* `CONSUMER_API_BASICAUTH_USERNAME` is to be configured with BasicAuth username, default `changeit`.

* `CONSUMER_API_BASICAUTH_PASSWORD` is to be configured with BasicAuth password, default `changeit`.

---

### Functionality

* exposes Prometheus metrices at `/metrics` api path

*  a `/ping` api for health-check

---

### To Build Plug-in

```
mkdir out
go build -o "out/ogi-http-service-consumer.so" -buildmode=plugin . 
```

---
