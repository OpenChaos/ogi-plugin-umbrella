
## Ogi Plugin Umbrella

> This is home to primary selected plug-ins for [Ogi](https://github.com/OpenChaos/ogi) which might be of higher relevance, so there is a single place to find them.
>
> For people unware of [Ogi](https://github.com/OpenChaos/ogi), it's a simple pluggable task pipeline creator in format of old-school ETL but with easy & better combination possibilities among its three steps of `Consumer`, `Transformer`, and `Producer`. Each can load any plug-in giving many possibilities along-with mixing in custom private plug-ins as well.

![ogi plugins umbrella](docs/ogi-plugins.png "ogi plugins umbrella")

---

### Plugin List

#### Consumer

* [File](./consumers/ogi-file-consumer): reads a file line-by-line and uses each line as one entity

* [HTTP Service](./consumers/ogi-http-service-consumer): simply passes request body to transformer, supports basic-auth


#### Transformer

* [Yahoo! Finance! Get Ticker Data](./transformers/ogi-yfinance-transformer): fetchs YFin ticker data for prodivided Symbol and passthrough to Producer


#### Producer

* [File](./producers/ogi-file-producer): appends provided data as string to a file

* [HTTP Request](./producers/ogi-http-producer): make an HTTP Request based on provided data

* [Slack Notifier](./producers/ogi-slack-producer): to send Slack Channel Notifications

* [Yahoo! Finance! Is It 52Week Low?](producers/ogi-yfin52low-producer): checks if YFin chart data last closed at 52 week low

---
