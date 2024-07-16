package data

type AttributeDefinition struct {
	Id     int64            `json:"id"`
	Name   string           `json:"name"`
	Type   string           `json:"type"`
	Values []AttributeValue `json:"values"`
}

type AttributeValue struct {
	Id    int64  `json:"id"`
	Value string `json:"value"`
}
