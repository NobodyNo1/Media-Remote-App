package home

import (
	"fmt"
	"log"
	"net/http"
	. "remote_media_control/controller"
	. "remote_media_control/data/model"
	. "remote_media_control/tools"
	"text/template"
	"strings"
)

func commandToMessage(command CommandLog) string {
	return fmt.Sprintf(
		"id: %d, name: %s, param: %s, status: %s, occuredAt: %s ",
		command.ID, command.Name, command.Param, command.Status, command.OccuredAt,
	)
}

func getLogMessages() map[string][]string {
	commandChain := GetCommandChain()
	result := make([]string, len(commandChain))
	for idX, chain := range commandChain{
		var messageBuilder strings.Builder
		for j, item := range chain {
			messageBuilder.WriteString(commandToMessage(item))
			if j != len(chain) - 1 {
				messageBuilder.WriteString("\n")
			}
		}
		result[idX] = messageBuilder.String()
	}
	return map[string][]string {
		// plain text?
		"Commands": result,
	}
}

func HandlerHome(w http.ResponseWriter, r *http.Request) {
	tmplt := template.Must(template.ParseFiles("router/pages/home/index2.html"))
	tmplt.Execute(w, getLogMessages())
}

func HandlerVolumeHtmx(w http.ResponseWriter, r *http.Request, increase bool) {
	volume, err := ChangeVolume(increase)
	if err != nil {

	}
	log.Default().Print(volume)
}

func HandlerMediaKeysHtmx(
	w http.ResponseWriter,
	r *http.Request,
	mediaKey MediaKey,
) {
	err := MediaKeyPress(mediaKey)
	if err != nil {

	}
}
