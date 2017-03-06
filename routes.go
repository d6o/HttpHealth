package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/disiqueira/itinerary"
)

func (a *App) serveList(w http.ResponseWriter, r *http.Request) {
	pingList := make(map[string]PingResponse)

	for _, e := range a.URLList {
		p := NewHeadPing(e)
		status, err := p.Ping()
		if err != nil {
			panic(err)
		}
		pingList[p.URL()] = status
	}
	response := NewResponse(w)
	response.Write(http.StatusOK, pingList)
}

func (a *App) serveURLList(w http.ResponseWriter, r *http.Request) {
	res, _ := json.Marshal(&a.URLList)
	w.Write(res)
}

func (a *App) serveURLCreate(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	if !govalidator.IsURL(r.FormValue("url")) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	a.AddURL(r.FormValue("url"))
	w.WriteHeader(http.StatusOK)
}

func (a *App) serveURLDelete(w http.ResponseWriter, r *http.Request) {
	itinerary.RouteToQuery(r)
	fmt.Println(r.URL.Query())
	url := r.URL.Query()["itinerary"][0]

	if !govalidator.IsURL(url) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	for i := 0; i < len(a.URLList); i++ {
		u := a.URLList[i]

		if u == url {
			a.URLList = append(a.URLList[:i], a.URLList[i+1:]...)
		}
	}

	w.WriteHeader(http.StatusNoContent)
}
