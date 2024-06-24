package data

type Attribute struct {
	Id     int64
	Name   string
	Type   string
	Values []AttributeValue
}

type AttributeValue struct {
	Id    int64
	Value string
}
