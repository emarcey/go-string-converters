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
