package api

import (
	"net/http"
	"encoding/json"
	. "remote_media_control/controller"
	. "remote_media_control/data/model"
)

// TODO: gRPC
type State struct {
	Volume	int `json:"volume"`
}

type DefaultResponse struct {
	Status 	bool `json:"status"`
	Message string `json:"message"`
}

func createSuccessResponse() DefaultResponse {
	return DefaultResponse {
		Status: true,
	}
}

func createFailResponse(message string) DefaultResponse {
	return DefaultResponse {
		Status: false,
		Message: message,
	}
}

func HandlerStateApi(w http.ResponseWriter, r *http.Request, _ any) {
	volume, err := GetVolume()
	if err != nil{
		resp:=createFailResponse(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(resp)
		return
	}
	stateResponse := State{
		Volume: volume,
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(stateResponse)
}

func HandlerVolumeApi(w http.ResponseWriter, r *http.Request, increase bool) {
	volume, err := ChangeVolume(increase)
	if err != nil {
		resp:=createFailResponse(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(resp)
		return
	}
	stateResponse := State{
		Volume: volume,
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(stateResponse)
}


func HandlerMediaKeysApi(
	w http.ResponseWriter,
	r *http.Request,
	mediaKey MediaKey,
) {
	err:= MediaKeyPress(mediaKey)
	if err != nil {
		resp:=createFailResponse(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(resp)
		return
	}
	stateResponse := createSuccessResponse()
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(stateResponse)
}