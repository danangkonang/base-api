package helper

import (
	"math/rand"
	"strings"
	"time"
)

func unixRandomString(length int) string {
	rand.Seed(time.Now().UnixNano())
	chars := []rune("C2DBEWK6SX8PQ9R3FG4HA1YZLM7NOI5J0TUV" + "luk9vmw6ob1cd2xenf3gp7qr8sah4i5tj0yz" + "0123456789")
	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}
	str := b.String()
	return str
}

func GetTime() string {
	t := time.Now()
	waktu := t.Format("2006-01-02 15:04:05")
	return waktu
}

func random(length int) string {
	chars := []rune("C2DBEWK6SX8PQ9R3FG4HA1YZLM7NOI5J0TUV")
	b := make([]rune, length)
	for i := range b {
		b[i] = chars[rand.Intn(len(chars))]
	}
	return string(b)
}

func RamdomString() string {
	name := random(8) + "_" + unixRandomString(16)
	return name
}
