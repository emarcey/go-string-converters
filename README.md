[![Go Report Card](https://goreportcard.com/badge/github.com/emarcey/go-string-converters)](https://goreportcard.com/report/github.com/emarcey/go-string-converters) [![Build Status](https://travis-ci.org/emarcey/go-string-converters.svg?branch=master)](https://travis-ci.org/emarcey/go-string-converters)

# go-string-converters
Some little helpers to make convert go strings around

## Supported Conversions

* PascalCase (also called UpperCamelCase)
    * Does not convert camelCase to PascalCase
* camelCase (also called LowerCamelCase)
    * Does not convert PascalCase to CamelCase
* l337
    * Basic implementation, only does a handful of number-letter conversions
    * Has functions to encode and decode
* snake_case:
    * Aggressive converter that converts any punctuation or PascalCase/camelCase to snake_case
    * Both regular (lowercase) and SCREAMING variants
* kebab-case:
    * Aggressive converter that converts any punctuation or PascalCase/camelCase to kebab-case
    * Both regular (lowercase) and SCREAMING variants
