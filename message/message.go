package fileupload

import (
	"encoding/json"
	"log"
)

const (
	LOGOUTPUT_ON  logout = 1
	LOGOUTPUT_OFF logout = 0
)

type logout int

type Message struct {
	Name   string `json:name`
	Status string `json:"status"`
	Code   int    `json:"code"`
}

// MessageのJSON変換
func (m *Message) Output() string {
	msg := Message{Status: m.Status, Code: m.Code}
	bytes, _ := json.Marshal(msg)
	return string(bytes)

}

//メッセージの設定
func (m *Message) InputMessage(msg string, flag logout) {
	m.Status = msg
	if flag == LOGOUTPUT_ON {
		log.Println(m.Status)
	}
}
