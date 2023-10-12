package tools

import (
	"fmt"
	"net/http"
	"strings"
	"errors"
)

func isHtmx(r *http.Request) bool {
	value := r.Header.Get("HX-Requested-With"); 
	return strings.ToLower(value) == "xmlhttprequest"
}

type GetRequestTemplate[T any] struct {
	Path string
	ParamName string
	Values map[string]T //here should be generic value
	DefaultValue T
}

func (grt GetRequestTemplate[T]) Create(
	handlerApi func(w http.ResponseWriter, r *http.Request, query T),
	handlerHtmx func(w http.ResponseWriter, r *http.Request, query T),
) {
	CreateSingleParamRequestHandler(
		grt, handlerApi, handlerHtmx,
	)
}

func handleBasic[T any](
	grt GetRequestTemplate[T],
	w http.ResponseWriter, r *http.Request,
	success func(T), fail func(error),
){
	// no query and validation is needed
	if grt.Values == nil {
		success(grt.DefaultValue)
		return
	}

	query := r.URL.Query().Get(grt.ParamName) 
	if query == "" {
		fail(errors.New("Empty Parameters"))
	}
	
	for key, value := range grt.Values {
		if key == query {
			success(value)
			return
		}
	}
	fail(errors.New("wrong parameters"))
}

func CreateSingleParamRequestHandler[T any](
	grt GetRequestTemplate[T],
	handlerApi func(w http.ResponseWriter, r *http.Request, query T),
	handlerHtmx func(w http.ResponseWriter, r *http.Request, query T),
) {
	apiPath:= fmt.Sprintf("/api/%s", grt.Path)
	htmxPath:= fmt.Sprintf("/%s", grt.Path)
	http.HandleFunc(apiPath, func(w http.ResponseWriter, r *http.Request) {
		success := func(query T)  {
			handlerApi(w,r, query)
		}
		fail := func(err error){
				// TODO: FIX THIS, cause it is handled by htmx and the api
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("{ }"))
		}

		handleBasic(grt,w,r,  success, fail)
	})

	http.HandleFunc(htmxPath, func(w http.ResponseWriter, r *http.Request) {
		success := func(query T)  {
			handlerHtmx(w,r, query)
		}
		fail := func(err error){
			panic("HTMX ERROR")
		}
		handleBasic(grt,w,r,success, fail)
	})

}