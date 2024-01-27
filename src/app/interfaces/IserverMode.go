package interfaces

import "app/api"

type Iserver interface {
	RunServer() (string, error)
}

func IserverModeMain() {
	RunServer()
}

func RunServer() error {
	api.RunListener()
	return nil
}