// Package handler to handle all your logics
package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"github.com/ashurai/BitUpdater/model"
	"sync"
	"github.com/gorilla/mux"
)

type Symbol struct {
	*model.Symbol
}

var mp = map[string]model.Symbol{}

// ValidateSymbols validate all configured symbols
// and only have those are valid in a slice
func (Sy *Symbol) ValidateSymbols(list []string, wg *sync.WaitGroup, valChan chan model.Symbol) {
	sym := model.Symbol{}
	//wg.Add(len(list))
	for _, v := range list {
		go func(v string) {
			defer wg.Done()
			u, err := url.Parse("https://api.hitbtc.com")
			u.Path = path.Join("/api/2/public/symbol/", v)
			if err != nil {
				fmt.Printf("while parsing url facing an issue %s", err)
			}

			resp, err := http.Get(u.String())
			if err != nil {
				fmt.Printf("getting error from api %s", err)
			}
			b, _ := ioutil.ReadAll(resp.Body)
			defer resp.Body.Close()
			
			err = json.Unmarshal(b, &sym)
			if err != nil {
				fmt.Printf("getting error from api read %s", err)
			}
			valChan <- sym 
			mp[sym.ID] = sym
		}(v)	
	}
}

// GetCurrencyById function to get by symbol id symbole details
func GetCurrencyByID(w http.ResponseWriter, r *http.Request) {
	crDetail := model.Currency{}
	rVars := mux.Vars(r)
	cr := rVars["currency"]
	
	symbol := mp[cr]
	if symbol.BaseCurrency == "" {
		log.Printf("Not a valid symbol type")
		
	}
	u, err := url.Parse("https://api.hitbtc.com")
	u.Path = path.Join("/api/2/public/currency/", symbol.BaseCurrency)
	resp, err := http.Get(u.String())
	if err != nil {
		fmt.Printf("getting error from api %s", err)
	}
	
	b, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	err = json.Unmarshal(b, &crDetail)
	crDetail.FeeCurrency = symbol.FeeCurrency
	if err != nil {
		fmt.Printf("getting error from api read %s", err)
	}

	tick := getTicker(symbol.ID)// Get latest currency details from bit api
	crDetail.Ticker = tick
	json.NewEncoder(w).Encode(crDetail)
}

// getTicker a function to get all latest values in background
func getTicker(cur string) model.Ticker {
	ticker := model.Ticker{}

	u, err := url.Parse("https://api.hitbtc.com")
	u.Path = path.Join("/api/2/public/ticker/", cur)
	resp, err := http.Get(u.String())
	if err != nil {
		fmt.Printf("getting error from ticker api %s", err)
	}
	b, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	err = json.Unmarshal(b, &ticker)
	return ticker
}

// GetAllCurrency to list out all the currency
func GetAllCurrency(w http.ResponseWriter, r *http.Request) {
	crs := []model.Currency{}
	var wag sync.WaitGroup
	cr := make(chan model.Currency, len(mp))
	for _, v := range mp {
		if v.QuoteCurrency != "" {
			wag.Add(1)
			go func(ID string, feeCurrency string, baseCurrency string){
				crDetail := model.Currency{}
				u, err := url.Parse("https://api.hitbtc.com")
				u.Path = path.Join("/api/2/public/currency/", baseCurrency)
				
				resp, err := http.Get(u.String())
				if err != nil {
					fmt.Printf("getting error from api %s", err)
				}
				fmt.Println(u.String())
				b, _ := ioutil.ReadAll(resp.Body)
				defer resp.Body.Close()

				err = json.Unmarshal(b, &crDetail)
				crDetail.FeeCurrency = feeCurrency
				if err != nil {
					fmt.Printf("getting error from api read %s", err)
				}

				tick := getTicker(ID)// Get latest currency details from bit api
				crDetail.Ticker = tick
				log.Println(crDetail)
				cr <- crDetail
				wag.Done()
			}(v.ID, v.FeeCurrency, v.BaseCurrency)
		} else {
			crDetail := model.Currency{}
			cr <- crDetail
		} 
	}
	wag.Wait()
	close(cr)

	for cur := range cr {
		if cur.ID != "" {
			crs = append(crs, cur)
		}
	}

	json.NewEncoder(w).Encode(crs)
}
