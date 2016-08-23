package main

type ExtendedSignalJson struct {
	Param map[string]string `json:"param"`
}
type JsonEvent struct {
	EventName string `json:"eventName"`

	Param map[string]string `json:"param"`

	Signal ExtendedSignalJson `json:"signal"`
}
