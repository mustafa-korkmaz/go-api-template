package model

//Olive represents olive model
type Olive struct {
	Base
	Kind    string `json:"kind"`
	Country string `json:"country"`
}
