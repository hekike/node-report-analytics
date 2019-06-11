# node-report-analytics

[![CircleCI](https://circleci.com/gh/hekike/node-report-analytics.svg?style=svg&circle-token=7fc6be4be1ab10cc0029a42ef6e349fb5f9eae06)](https://circleci.com/gh/hekike/node-report-analytics)

Processes Node.js [Diagnostic Report](https://nodejs.org/api/report.html)(s).

Requires `Node 12.x` or greater.

## Usage

```
Usage:
  node-report-analytics [command]

Available Commands:
  elastic     Loads Diagnostics Report(s) to Elasticsearch
  help        Help about any command
  stats       Prints the aggregated statistics of Diagnostics Report(s)

Flags:
  -h, --help   help for node-report-analytics
```

### Stats

Print aggregated results on terminal.

```
make generate-report
make stats
```

Output:
```
+---------+-------+-------------------+------------------+------------------+
| HASH    | COUNT | MESSAGE           | FIRST OCCURENCE  | LAST OCCURENCE   |
+---------+-------+-------------------+------------------+------------------+
| 7b14b96 |     2 | Error: My Error   | 2019-06-10 20:14 | 2019-06-10 20:15 |
| 2aac2df |     1 | Error: My Error 2 | 2019-06-10 20:15 | 2019-06-10 20:15 |
+---------+-------+-------------------+------------------+------------------+
```

### Elastic

Insert reports by hash to Elasticsearch.

```
make docker-compose
make elastic
```

Output:
```
7b14b968202c3218514e54f627c20de43323a631, 2019-06-10 20:14:52 +0000 UTC, Error: My Error
```

## Test

```
make test
```
