package core

import (
	"WebIM/service/ws"
)

func InitSocket() {
	go ws.Manager.Start()
}
