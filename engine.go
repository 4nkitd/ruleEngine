package engine

import (
	"encoding/json"
)

func Run(data string) ([]Event, error) {
	rules := Rules{}
	err := json.Unmarshal([]byte(data), &rules)
	if err != nil {
		return []Event{}, err
	}

	err = rules.Validate()
	if err != nil {
		return []Event{}, err
	}

	egress, err := rules.Run()
	if err != nil {
		return []Event{}, err
	}

	return egress, nil
}

func (r *Rules) Run() ([]Event, error) {

	egress, err := r.RunDecisions()
	if err != nil {
		return []Event{}, err
	}

	return egress, nil
}

func (r *Rules) Validate() error {
	err := ValidateRuleAttributesAndConditions(*r)
	if err != nil {
		return err
	}

	return nil
}
