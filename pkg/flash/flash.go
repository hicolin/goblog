package flash

import (
	"encoding/gob"
	"goblog/pkg/session"
)

type Flashes map[string]interface{}

var flashKey = "_flashes"

func init() {
	gob.Register(Flashes{})
}

func Info(message string) {
	addFlash("info", message)
}

func Warning(message string) {
	addFlash("warning", message)
}

func Success(message string) {
	addFlash("success", message)
}

func Danger(message string) {
	addFlash("danger", message)
}

func All() Flashes {
	val := session.Get(flashKey)
	flashMessages, ok := val.(Flashes)
	if !ok {
		return nil
	}
	session.Forget(flashKey)
	return flashMessages
}

func addFlash(key string, message string) {
	flashes := Flashes{}
	flashes[key] = message
	session.Put(flashKey, flashes)
	session.Save()
}
