package controller

import (
	"errors"
	"remote_media_control/controller/media"
	"remote_media_control/controller/volume_control"
	. "remote_media_control/data/model"
	. "remote_media_control/tools"
)

// TODO: make command logging simpler 

func getVolumeCommand(chain []CommandLog) (int, error) {
	process := "get-volume"
	chain = append(chain, CreateChainLogCommand(process, Started, "", "" ))
	
	curVol , err := volume_control.GetVolume()
	if err != nil {
		chain = append(chain, CreateChainLogCommandF(
			process,
			Fail,
			"",
			"can't get current state of volume: %v",
			err,
		))

		return curVol, err
	}
	chain = append(chain, CreateChainLogCommand(process, Success, "", ""))
	return curVol , err
}

func GetVolume() (int, error) {
	var logChain []CommandLog
	volume, err :=  getVolumeCommand(logChain)
	LogCommandChain(logChain)
	return volume,err
}


func ChangeVolume(increase bool) (int, error) {
	var logChain []CommandLog
	process := "change-volume"
	logChain = append(logChain, CreateChainLogCommand(process, Started, increase, ""))
	curVol , err := volume_control.GetVolume()
	if err != nil {
		LogCommandChain(logChain)
		return curVol, err
	}
	newVol := getUpdatedVolume(curVol, increase)
	err = volume_control.SetVolume(newVol)
	if err != nil {
		logChain = append(
			logChain,
			CreateChainLogCommandF(
				process,
				Fail,
				increase,
				"can't update volume %v", 
				err,
			),
		)
		LogCommandChain(logChain)
		return curVol, err
	}
	
	currVol, err := getVolumeCommand(logChain)
	if err != nil {
		// even if getting volume failed we know that volume updated
		return newVol, nil
	}
	logChain = append(
		logChain,
		 CreateChainLogCommand(process, Success,increase, ""),
	)
	LogCommandChain(logChain)
	return currVol, err
}

func getUpdatedVolume(curVolume int, increase bool) int {
	newVol := curVolume
	if increase {
		newVol += 5
	} else {
		newVol -= 5
	}
	if newVol > 100{
		newVol = 100
	} else if newVol < 0 {
		newVol = 0
	}
	return newVol
}

func MediaKeyPress(key MediaKey) error  {
	var logChain []CommandLog
	//todo: NEXT, PREVIOUS, etc..
	process := "media-key"
	logChain = append(
		logChain,
		CreateChainLogCommand(process, Success, string(key), ""),
	)
	var err error
	if key == PlayPause {
		err = media.PlayPause()
	} else { 
		err = errors.New("unsupported action")
	}
	if err != nil {
		logChain = append(
			logChain,
			CreateChainLogCommandF(
				process,
				Fail,
				key,
				"can't get updated volume %v", 
				err,
			),
		)
	} else {
		logChain = append(
			logChain,
			CreateChainLogCommand(
				process,
				Success,
				key,
				"",
			),
		)
	}
	LogCommandChain(logChain)
	return err
}