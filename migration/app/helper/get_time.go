package helper

import "time"

func GetTime() string {
	t := time.Now()
	waktu := t.Format("2006-01-02 15:04:05")
	return waktu
}
