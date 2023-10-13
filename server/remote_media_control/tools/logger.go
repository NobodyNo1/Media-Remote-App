package tools

import (
	"time"
	"fmt"
)

type Status string

const (
	Started	Status = "started"
	Fail 	Status = "fail"
	Success Status = "success"
)

type CommandLog struct {
	//TODO: FROM
	ID int
	Name string
	Param any
	LogStatus Status
	OccuredAt time.Time
}

// TODO: Linked List
var idCounter int = 0
var chainIdCounter int = 0
var requests []CommandLog
var commands []CommandLog
var commandChain [][]CommandLog

func GetCommands() []CommandLog{
	return commands
}

func GetCommandChain() [][]CommandLog{
	return commandChain
}

func LogRequest(){
}


func CreateLogCommand(id int, name string, status Status, param any, message string) CommandLog {
	return CommandLog { 
		ID: id,
		Name: name,
		Param: param,
		LogStatus: status,
		OccuredAt: time.Now(),
	}
}

func CreateChainLogCommand(
	name string,
	status Status,
	param any,
	message string,
) CommandLog {
	return CreateLogCommand(chainIdCounter, name, status, param, message)
}

func CreateChainLogCommandF(
	name string,
	status Status,
	param any,
	messageFormat string, 
	a ...any,
) CommandLog {
	message := fmt.Sprintf(messageFormat, a)
	return CreateChainLogCommand(name, status, param, message)
}

func LogCommandChain(chain []CommandLog){
	if len(commandChain) > 10 {
		_, commandChain = dequeueChain(commandChain)
	}
	commandChain = enqueueChain(commandChain, chain)
	chainIdCounter++
}

func LogCommand(name string, status Status, param any, message string) {
	fmt.Printf("name:%s, status: %s, \n message:%s", name, status, message)
	if len(commands) > 10 {
		_, commands = dequeue(commands)
	}
	newCommand := CreateLogCommand(chainIdCounter, name, status, param, message)
	commands = enqueue(commands, newCommand)
	idCounter++
}

func LogCommandF(
	name string,
	status Status,
	param any,
	messageFormat string, 
	a ...any) {
	message := fmt.Sprintf(messageFormat, a)
	LogCommand(name, status, param, message)
}


func enqueue(queue []CommandLog, element CommandLog) []CommandLog { 
	queue = append(queue, element) // Simply append to enqueue. 
	return queue 
} 
	  
func dequeue(queue []CommandLog) (CommandLog, []CommandLog) { 
	element := queue[0]
	if len(queue) == 1 { 
	 var tmp = []CommandLog{} 
	 return element, tmp 
	  
	} 
	  
	return element, queue[1:]
} 


func enqueueChain(queue [][]CommandLog, element []CommandLog) [][]CommandLog { 
	queue = append(queue, element) // Simply append to enqueue. 
	return queue 
} 
	  
func dequeueChain(queue [][]CommandLog) ([]CommandLog, [][]CommandLog) { 
	element := queue[0]
	if len(queue) == 1 { 
		var tmp = [][]CommandLog{} 
		return element, tmp 
	} 
	  
	return element, queue[1:]
} 