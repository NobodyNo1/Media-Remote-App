package volume_control

import "github.com/itchyny/volume-go"

func GetVolume() (int, error) {
	return volume.GetVolume()
}

func SetVolume(vol int) error {
	return volume.SetVolume(vol)
}