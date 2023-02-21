[![Coverage](https://sonarcloud.io/api/project_badges/measure?project=dnitsch_simplelog&metric=coverage)](https://sonarcloud.io/summary/new_code?id=dnitsch_simplelog)
[![Vulnerabilities](https://sonarcloud.io/api/project_badges/measure?project=dnitsch_simplelog&metric=vulnerabilities)](https://sonarcloud.io/summary/new_code?id=dnitsch_simplelog)
[![Bugs](https://sonarcloud.io/api/project_badges/measure?project=dnitsch_simplelog&metric=bugs)](https://sonarcloud.io/summary/new_code?id=dnitsch_simplelog)
[![Maintainability Rating](https://sonarcloud.io/api/project_badges/measure?project=dnitsch_simplelog&metric=sqale_rating)](https://sonarcloud.io/summary/new_code?id=dnitsch_simplelog)

# Simple Log

Simple wrapper around zeroLog logger with simple interface for most common use cases 

Additionally a [logr]https://github.com/go-logr/logr compliant interface can be instantiated with the same "constructor" signature as the standard implementation.

Exposes 3 log levels and returns an instance of a logger with set-able level and writer.

- `DEBUG`
- `INFO`
- `ERROR`
