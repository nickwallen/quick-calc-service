
## Quick Calc Service

A GraphQL-based service for [Quick Calc](https://github.com/nickwallen/quick-calc); a simple calculator that operates on units of measure and allows for quick conversions.

### Quick Start

1. Launch the service.
    ```
    make run
    ```
   
1. Explore the service and documentation using the GraphiQL interface at `http://localhost:8080`. Use the query examples below to get started.

1. Query the service endpoint at `http://localhost:8080/query`.

    ```
    curl \
        --header "Content-Type: application/json" \
        --data '{ "query": "{ units { name measureOf } }" }' \
        http://localhost:8080/query \
   | jq
    ```


### Queries

1. Evaluate an expression.
    ```
    { 
      evaluate(expr: "2 stones + 0.5 tons in pounds") {
        value
        units {
          pluralName
        }
      } 
    }
   ```
   ```
       "evaluate": {
         "value": 1148,
         "units": {
           "pluralName": "pounds"
         }
       }
   ```

1. List all available units of measure.
    ```
    {
      units {
        name
        pluralName
        measureOf
        partOf
      }
    }
    ```

1. Find unit by name.
    ```
    {
      unitByName(name: "stones") {
        name
        pluralName
        measureOf
        partOf
      }
    }
    ```


