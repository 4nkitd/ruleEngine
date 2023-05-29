package engine

import (
	"encoding/json"
	"fmt"
	"testing"
)

var data = `
{
    "name": "discount",
    "decisions": [
        {
            "event": {
                "type": "apply discount",
                "params": [{
                    "variable": "discountCode",
                    "value": "winter50"
                },
                {
                    "variable": "blame",
                    "value": "summer"
                }]
            },
            "conditions": {
                "all": [
                    {
                        "variable": "amount",
                        "value": "200",
                        "operator": "equal"
                    },
                    {
                        "variable": "category",
                        "value": "food",
                        "operator": "equal"
                    },
                    {
                        "variable": "discountCode",
                        "value": "winter50",
                        "operator": "equal"
                    }
                ]
            }
        },
        {
            "event": {
                "type": "apply fee",
                "params": [{
                    "variable": "feeCode",
                    "value": "free"
                },
                {
                    "variable": "blame",
                    "value": "free"
                }]
            },
            "conditions": {
                "any": [
                    {
                        "variable": "amount",
                        "value": "200",
                        "operator": "equal"
                    },
                    {
                        "variable": "category",
                        "value": "food",
                        "operator": "equal"
                    },
                    {
                        "variable": "discountCode",
                        "value": "winter50",
                        "operator": "equal"
                    }
                ]
            }
        }
    ],
    "attributes": [
        {
            "variable": "amount",
            "value": "200",
            "type": "number"
        },
        {
            "variable": "category",
            "value": "food",
            "type": "string"
        },
        {
            "variable": "discountCode",
            "value": "winter50",
            "type": "string"
        }
    ]
}`

var dataWithIncorrectConitions = `
{
    "name": "discount",
    "decisions": [
        {
            "event": {
                "type": "apply discount",
                "params": {
                    "variable": "discountCode",
                    "value": "winter50"
                }
            },
            "conditions": {
                "all": [
                    {
                        "variable": "amount",
                        "value": "200",
                        "operator": "equal"
                    },
                    {
                        "variable": "category",
                        "value": "food",
                        "operator": "equal"
                    },
                    {
                        "variable": "discountCode",
                        "value": "winter50",
                        "operator": "==="
                    }
                ]
            }
        }
    ],
    "attributes": [
        {
            "variable": "amount",
            "value": "200",
            "type": "number"
        },
        {
            "variable": "category",
            "value": "food",
            "type": "string"
        },
        {
            "variable": "discountCode",
            "value": "winter50",
            "type": "string"
        }
    ]
}`

func TestRun(t *testing.T) {

	resp, err := Run(data)

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	jsonResp, _ := json.Marshal(resp)
	fmt.Println(string(jsonResp))
	t.Log(jsonResp)

}

func TestConditionChecking(t *testing.T) {
	resp, err := Run(dataWithIncorrectConitions)

	if resp == nil {
		t.Errorf("Unexpected response: %v", resp)
	}

	t.Logf("Success error: %v", err)
}
