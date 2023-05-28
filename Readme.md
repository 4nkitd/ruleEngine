Rule Engine Documentation
=========================

Overview
--------

The rule engine is a Go package designed to evaluate and execute rules based on predefined conditions. It provides a flexible and customizable way to make decisions and trigger events based on the specified criteria. The rule engine accepts rule definitions in JSON format.

Installation
------------

To use the rule engine in your Go project, you can install it using the following command:

    
    go get github.com/4nkitd/ruleEngine
      

Usage
-----

To use the rule engine, you need to import the package in your Go code:

    
    import "github.com/4nkitd/ruleEngine"
      

### Creating Rules

The rule engine allows you to define rules consisting of events and conditions. Each rule has one or more conditions that need to be satisfied in order for the associated event to be triggered.

Example rule definition in JSON format:

```json
    {
      "name": "Buy Food",
      "event": {
        "type": "To Buy Food",
        "params": {
          "variable": "marinda",
          "value": "1l"
        }
      },
      "conditions": {
        "all": [
          {
            "variable": "amount",
            "value": "50",
            "operator": "equal"
          },
          {
            "variable": "weather",
            "value": "sunny",
            "operator": "equal"
          },
          {
            "variable": "people",
            "value": "4",
            "operator": "<="
          }
        ]
      }
    }
```

### Evaluating Rules

Once you have defined your rules in JSON format, you can evaluate them by calling the `engine.Run` method:

```go
    results, err := engine.Run(jsonString)
    if err != nil {
      // Handle error
    }
    
```
Supported Operators
-------------------

The rule engine supports the following operators for condition evaluation:

*   equal (=)
*   not equal (!=)
*   greater than (>)
*   less than (<)
*   greater than or equal to (>=)
*   less than or equal to (<=)