package engine

import "sort"

type Rules struct {
	Name       string      `json:"name"`
	Decisions  []Decision  `json:"decisions"`
	Attributes []Attribute `json:"attributes"`
}

type Decision struct {
	Event      Event      `json:"event"`
	Conditions Conditions `json:"conditions"`
}

type Event struct {
	Type   string  `json:"type"`
	Params Options `json:"params"`
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

	// Calculate scores for each decision
	scores := make(map[int]int)
	for i, d := range c.Decisions {
		scores[i] = len(d.Conditions.AllOf)
	}

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
	}

	return results, nil
}
