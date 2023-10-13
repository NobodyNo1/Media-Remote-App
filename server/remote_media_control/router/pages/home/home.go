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

func commandToMessage(command CommandLog, short bool) string {
	if(short){
		// TODO: spacing is not kept in the html (tailwind)
		return fmt.Sprintf(
			"%-10s %-10s | %v",
			fmt.Sprintf(
				"[%s]",
				strings.ToUpper(string(command.LogStatus)),
			),
			command.Name,
			command.Param,
		)
	}
	return fmt.Sprintf(
		"id: %d, name: %s, param: %s, status: %s, occuredAt: %s ",
		command.ID, command.Name, command.Param, command.LogStatus, command.OccuredAt,
	)
}

func getLogMessages(onlyFinalStatuses bool, short bool) map[string][]string {
	commandChain := GetCommandChain()
	result := make([]string, len(commandChain))
	for idX, chain := range commandChain {
		var messageBuilder strings.Builder
		
		if(onlyFinalStatuses) {
			messageBuilder.WriteString(commandToMessage(chain[len(chain)-1],short))
		} else {
			for j, item := range chain {
				messageBuilder.WriteString(commandToMessage(item,short))
				if j != len(chain) - 1 {
					messageBuilder.WriteString("\n")
				}
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
	tmplt := template.Must(template.ParseFiles("router/pages/home/index.html"))
	logs := getLogMessages(true, true)
	for _,i := range logs {
		log.Default().Print(i)
	}
	tmplt.Execute(w, logs)
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
