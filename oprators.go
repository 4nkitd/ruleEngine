package engine

import (
	"strconv"
	"strings"
)

var OperatorMap = map[string]func(string, string) bool{
	"=":                  equal,
	"!=":                 notEqual,
	">":                  greaterThan,
	"<":                  lessThan,
	">=":                 greaterThanOrEqual,
	"<=":                 lessThanOrEqual,
	"equal":              equal,
	"notEqual":           notEqual,
	"greaterThan":        greaterThan,
	"lessThan":           lessThan,
	"greaterThanOrEqual": greaterThanOrEqual,
	"lessThanOrEqual":    lessThanOrEqual,
}

// evaluateCondition evaluates the condition using the corresponding operator function
func evaluateCondition(value1, value2 string, operator string) bool {
	normalizedOperator := strings.ToLower(operator)
	if op, ok := OperatorMap[normalizedOperator]; ok {
		return op(value1, value2)
	}
	return false
}

// equal compares two string values
func equal(value1, value2 string) bool {
	return value1 == value2
}

// notEqual compares two string values
func notEqual(value1, value2 string) bool {
	return value1 != value2
}

// greaterThan compares two numeric values
func greaterThan(value1, value2 string) bool {
	num1, _ := strconv.ParseFloat(value1, 64)
	num2, _ := strconv.ParseFloat(value2, 64)
	return num1 > num2
}

// lessThan compares two numeric values
func lessThan(value1, value2 string) bool {
	num1, _ := strconv.ParseFloat(value1, 64)
	num2, _ := strconv.ParseFloat(value2, 64)
	return num1 < num2
}

// greaterThanOrEqual compares two numeric values
func greaterThanOrEqual(value1, value2 string) bool {
	num1, _ := strconv.ParseFloat(value1, 64)
	num2, _ := strconv.ParseFloat(value2, 64)
	return num1 >= num2
}

// lessThanOrEqual compares two numeric values
func lessThanOrEqual(value1, value2 string) bool {
	num1, _ := strconv.ParseFloat(value1, 64)
	num2, _ := strconv.ParseFloat(value2, 64)
	return num1 <= num2
}

// ...
