
### TCP Check Consumer

> It is a consumer plug-in for Ogi, checks is TCP Connection is working to provided set of connection strings.. on failure calls a Transform for whatever required.

#### Environment Variable Configuration

* Checks & config loaded from `OGI_TCP_CHECK_YAML`; of below structure & default path "/tmp/ogi-tcp-check.yaml"

```
---
checks:
  - localhost:8080
  - example.com:443
interval_seconds: 15
backoff: 1
max_backoff: 20
```

---

### To Build Plug-in

```
mkdir -p ../../out
go build -o "../../out/ogi-tcp-check-consumer.so" -buildmode=plugin .
```

---
