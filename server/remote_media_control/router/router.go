package router

import (
	"net/http"
	. "remote_media_control/tools"
	. "remote_media_control/router/api"
	. "remote_media_control/router/pages/home"
	. "remote_media_control/data/model"
)

var keyMap = map[string]MediaKey {
	"playPause": PlayPause,
	"forward":   Forward,
	"backward":  Backward,
}

func SetupRouter() {
	emptyHandler := func (w http.ResponseWriter, r *http.Request, _ any){
	}
	http.HandleFunc("/", HandlerHome)
	// Current State Fetch
	GetRequestTemplate[any] {
		Path: "state",
		ParamName: "",
		Values: nil,
		DefaultValue: nil,
	}.Create(
		HandlerStateApi,
		emptyHandler,
	)
	// Volume State Update
	GetRequestTemplate[bool] {
		Path: "volume",
		ParamName: "status",
		Values: map[string]bool{
			"increase": true,
			"decrease": false,
		},
	}.Create(
		HandlerVolumeApi,
		HandlerVolumeHtmx,
	)

	// Media State Update
	GetRequestTemplate[MediaKey] {
			Path: "key",
			ParamName: "event",
			Values: keyMap,
	}.Create(
		HandlerMediaKeysApi,
		HandlerMediaKeysHtmx,
	)
}