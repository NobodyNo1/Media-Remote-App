package main

import (
	"log"
	"net/http"
	"github.com/itchyny/volume-go"
)


func main() {
	// start the server
	http.HandleFunc("/volume", handlerVolume)
	http.HandleFunc("/key", handlerKey)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func handlerKey(w http.ResponseWriter, r *http.Request){
	//TODO: pass as json
	log.Default().Printf("KEY")

	event := r.URL.Query().Get("event")
	if event == "play" || event == "pause" {
		keyPress(event == "play")
	} else {
		log.Default().Print("need type")
	}
}

func keyPress(isPlay bool)  {
	playPause()
}

func handlerVolume(w http.ResponseWriter, r *http.Request){
	log.Default().Printf("VOLUME")
	//TODO: pass as json
	status := r.URL.Query().Get("status")
	if status != "" {
		changeVolume(status == "increase")
	} else {
		log.Default().Print("need status")
	}
}

func changeVolume(increase bool) {
	currVol := getVolume()
	if increase{
		currVol++
		log.Default().Print("increase")
	} else {
		currVol--
		log.Default().Print("decrease")
	}
	setVolume(currVol)
	getVolume()
}

func getVolume() int {
	vol, err := volume.GetVolume()
	if err != nil {
		log.Fatalf("get volume failed: %+v", err)
	}
	log.Default().Printf("current volume: %d\n", vol)
	return vol
}

func setVolume(vol int) {
	err := volume.SetVolume(vol)
	if err != nil {
		log.Fatalf("set volume failed: %+v", err)
	}
	log.Default().Printf("set volume success\n")
}
