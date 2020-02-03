# List all dates in a month

This command line tool returns a list of all dates in a given year/month, each on its own line.

## Usage

List all dates in February 2020 with the default sprintf pattern of __* [[date]] __

```
all_dates_in_month 2020 2
```


List all dates in February 2020 with the given sprintf pattern.

```
all_dates_in_month 2020 2 "- %s"
```

## Installation

```
go get -u github.com/jakoubek/all_dates_in_month
```
