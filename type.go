package engine

import (
	"sort"
)

type Rules struct {
	Name       string      `json:"name"`
	Limit      int         `json:"limit"`
	Decisions  []Decision  `json:"decisions"`
	Attributes []Attribute `json:"attributes"`
}

type Decision struct {
	Event      Event      `json:"event"`
	Conditions Conditions `json:"conditions"`
}

type Event struct {
	Type   string    `json:"type"`
	Params []Options `json:"params"`
}

type Options struct {
	Key   string `json:"variable"`
	Value string `json:"value"`
}

type Conditions struct {
	AllOf []Condition `json:"all"`
	OneOf []Condition `json:"any"`
}

type Condition struct {
	Variable string `json:"variable"`
	Operator string `json:"operator"`
	Value    string `json:"value"`
}

type Attribute struct {
	Variable string `json:"variable"`
	Value    string `json:"value"`
	Type     string `json:"type"`
}

func (c *Rules) RunDecisions() ([]Event, error) {
	results := []Event{}

	limit := c.Limit
	eventsCaptured := 0

	if limit == 0 {
		limit = len(c.Decisions)
	}

	// Calculate scores for each decision
	scores := make(map[int]int)
	anyScores := make(map[int]int)

	hasAny := false

	for i, d := range c.Decisions {
		scores[i] = len(d.Conditions.AllOf)
		anyScores[i] = len(d.Conditions.OneOf)

		if len(d.Conditions.OneOf) > 0 {
			hasAny = true
		}

	}

	// If there are no decisions with any conditions, return the first decision
	if !hasAny {

		for _, d := range c.Decisions {
			processThis := false

			if eventsCaptured >= limit {
				break
			}

			for _, condition := range d.Conditions.OneOf {

				for _, attribute := range c.Attributes {
					if attribute.Variable == condition.Variable {
						if attribute.Value == condition.Value {
							if evaluateCondition(attribute.Value, condition.Value, condition.Operator) {
								processThis = true
								// skip this loop
								break
							}
						}
					}
				}

			}
			if processThis {
				results = append(results, d.Event)
			}

			eventsCaptured++

		}

		return results, nil
	}

	if len(c.Decisions) == 0 {
		return results, nil
	}
	eventsCaptured = 0
	// Sort the decision indexes based on scores in descending order
	sortedIndexes := make([]int, 0, len(scores))
	for i := range scores {
		sortedIndexes = append(sortedIndexes, i)
	}
	sort.Slice(sortedIndexes, func(i, j int) bool {
		return scores[sortedIndexes[i]] > scores[sortedIndexes[j]]
	})

	// Iterate over decisions based on sorted indexes and evaluate conditions
	for _, index := range sortedIndexes {
		d := c.Decisions[index]
		totalConditions := len(d.Conditions.AllOf)
		matchedConditions := 0
		if eventsCaptured >= limit {
			break
		}
		for _, condition := range d.Conditions.AllOf {
			for _, attribute := range c.Attributes {
				if attribute.Variable == condition.Variable {
					if attribute.Value == condition.Value {
						if evaluateCondition(attribute.Value, condition.Value, condition.Operator) {
							matchedConditions++
						}
					}
				}
			}
		}

		if matchedConditions == totalConditions {
			results = append(results, d.Event)
		}
		eventsCaptured++
	}

	return results, nil
}
