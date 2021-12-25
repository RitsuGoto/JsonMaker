package model

type Elements struct {
	Value []string
}

type Holder struct {
	Value map[string]string
}

type ToJson struct {
	Value map[string][]Holder
}
