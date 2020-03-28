// Package main to run base/boot application
package main

import (
	//"fmt"
	"github.com/ashurai/BitUpdater/handler"
	"github.com/ashurai/BitUpdater/model"
	"log"
	"net/http"
	"sync"
	"github.com/gorilla/mux"
)

// PreSymbolsList configure your list of symbols
var PreSymbolsList = []string{"BTCUSD", "ETHBTC"}//, "FGVTCJ"}//, "FGVTCJ"} 
// innvalid for test = FGVTCJ, another valid value for test = FGVTCJ

// Use the symbole type from handler 
// which is importing struct from models
var sy *handler.Symbol

func main() {
	var wg sync.WaitGroup
	lenOfSy := len(PreSymbolsList)
	valChan := make(chan model.Symbol, lenOfSy)

	// Validate all the valid symbols
	wg.Add(lenOfSy)
	sy.ValidateSymbols(PreSymbolsList, &wg, valChan)
	wg.Wait()
	close(valChan)

	result := []model.Symbol{}
	for ch := range valChan {
		result = append(result, ch)
	}
	log.Println(result)
	log.Println("start localhost")
	handleRequests()
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/currency/all", handler.GetAllCurrency)
	router.HandleFunc("/currency/{currency}", handler.GetCurrencyByID)
	
	log.Fatal(http.ListenAndServe(":8099", router))
}
