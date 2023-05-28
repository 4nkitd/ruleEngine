package engine

import (
	"errors"
	"strings"
)

var supportedTypes = map[string]bool{
	"string": true,
	"number": true,
}

var supportedOperators = map[string]bool{
	"equal":              true,
	"notequal":           true,
	"greaterthan":        true,
	"lessthan":           true,
	"greaterthanorequal": true,
	"lessthanorequal":    true,
	"=":                  true,
	"!=":                 true,
	">":                  true,
	"<":                  true,
	">=":                 true,
	"<=":                 true,
}

func ValidateRuleAttributesAndConditions(rules Rules) error {
	for _, attr := range rules.Attributes {
		if !supportedTypes[attr.Type] {
			return errors.New("unsupported attribute type: " + attr.Type)
		}
	}

	for _, decision := range rules.Decisions {
		for _, condition := range decision.Conditions.AllOf {
			if !supportedOperators[strings.ToLower(condition.Operator)] {
				return errors.New("unsupported condition operator: " + condition.Operator)
			}
		}
	}

	return nil
}
