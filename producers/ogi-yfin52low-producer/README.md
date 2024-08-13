
### YFin 52Low

> it is a producer plug-in for Ogi, that checks if last Closing Price was 52week Low for given YFin Chart Data

* It logs the result; also return JSON response bytes.. so flow-back to Consumer could be used as well.

* If using `ogi-yfinance-transformer`; set env `OGI_YFIN_RANGE=1y` for 52week data to be provided only. As it primarily picks lowest price from all Closed Prices as of now.

---
