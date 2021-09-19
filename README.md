# Fare Estimation Script

## Overview

The application estimates ride fares given rides formatted as csv as follows:
```
id_ride,lat,lng,timestamp
```

## Usage

CSV input is read from stdin, output is written to stdout (inspired by the Unix `cpio` command).

Example:

```
zcat sample/paths.csv.gz | go run ./cmd
```

## Testing

```
go test ./...
```

## Project structure and design

### Structure

- Project root contains the domain rules and entities, as well as some helper methods and structs
- `./cmd` contains the main executable code and a simple integration test
- `./rawridefetcher` contains the component responsible for reading the source and ride aggregation by id
- `./fareestimator` contains the core logic, agnostic to input and output formats. This is where the fare calculation is performed
- `./fareestimatewriter` contains the output formatting

### Design

The application represents a very simple pipeline of 3 stages: input reading and aggregation, input processing, output formatting and writing.

## Assumptions

- assume rides take place on planet Earth
- assume the planet is a perfect sphere
- assume we can disregard leap seconds-related quirks in Unix time
- assume float64 is precise enough to perform the calculations of speed, distance and ride fare
- assume segment timestamp can be safely represented by the first of its two timestamps
- rides having less than 2 points are considered invalid and result in an error
