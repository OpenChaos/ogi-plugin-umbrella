
### Tail File Consumer for Ogi

> it is a consumer plug-in for Ogi, it is a simple tail file consumer that reads data line-by-line and passes it to configured Transformer

#### Environment Variable Configuration

* `OGI_FILE_TO_CONSUME` provides port to bind at, value format is `:8080`.

---

### To Build Plug-in

```
mkdir -p ../../out
go build -o "../../out/ogi-tail-file-consumer.so" -buildmode=plugin . 
```

---
