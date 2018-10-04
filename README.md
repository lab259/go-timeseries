[![CircleCI](https://circleci.com/gh/lab259/go-timeseries.svg?style=shield)](https://circleci.com/gh/lab259/go-timeseries)

# go-timeseries

## Architecture

**Pipeline** are a series of aggregations performed in sequence over an
Event that generates a set of data that will be stored using a `Storage`
abstraction.

**Event** is an occurrence that happens and need to be measured.

**Aggregations** are `Operation`s that will aggregate information about
a specific event and will be kept in the system.

**Operation** is a transformation that will be performed in the database
by the `Storage`.

**Granularity** is the measure of time that specifies how precise a
datetime is.

**Storage** is the abstraction that will receive the data aggregated and
spread through an database to