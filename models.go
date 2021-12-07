package main

type OrdSubmit struct {
	Element string `json:"element"`
	Module  string `json:"module"`
}

type FiniteFieldSubmit struct {
	PX string `json:"px"`
}

type FiniteFieldResult struct {
	Sum string `json:"sum"`
	Mul string `json:"mul"`
}
