
### YFin HTTP Client

> It's a transformer plug-in for Ogi.
>
> This makes HTTP Requests for provided YFinance Chart Name to `https://query1.finance.yahoo.com/v8/finance/chart/{stock}`, and passthrough body to Producer.

#### Environment Variable Configuration

* SSL Verify in HTTP Requests, set env var `OGI_HTTP_REQUEST_SKIP_SSL`, default: `true`.

* YFin API ticker range `OGI_YFIN_RANGE`; default: `max`.

* YFin API ticker interval `OGI_YFIN_INTERVAL`; default: `1d`.

> Valid Ranges: 1d, 5d, 1mo, 3mo, 6mo, 1y, 2y, 5y, 10y, ytd, max

---
