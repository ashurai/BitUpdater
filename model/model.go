// Package model is to handle entities / structs
package model

// Symbol struct to manipulate symbol object
type Symbol struct {
	ID                   string `json:"id"`
	BaseCurrency         string `json:"baseCurrency"`
	QuoteCurrency        string `json:"quoteCurrency"`
	QuantityIncrement    string `json:"quantityIncrement"`
	TickSize             string `json:"tickSize"`
	TakeLiquidityRate    string `json:"takeLiquidityRate"`
	ProvideLiquidityRate string `json:"provideLiquidityRate"`
	FeeCurrency          string `json:"feeCurrency"`
}

// Currency struct is to manipuate currency object
type Currency struct {
	ID          string `json:"id"`
	FullName    string `json:"fullName"`
	FeeCurrency string `json:"feeCurrency"`
	Ticker
}

// Ticker model to internal struct
type Ticker struct {
	Ask         string    `json:"ask"`
	Bid         string    `json:"bid"`
	Last        string    `json:"last"`
	Open        string    `json:"open"`
	Low         string    `json:"low"`
	High        string    `json:"high"`
}