/*
Copyright 2016 Skippbox, Ltd.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package toast

import (
	"fmt"
	"os"

	"github.com/bitnami-labs/kubewatch/config"
	kbEvent "github.com/bitnami-labs/kubewatch/pkg/event"
	"github.com/gen2brain/beeep"
)

// Toast handler implements handler.Handler interface,
// Notify event as Toast
type Toast struct {
	Mode string
}

// Init prepares Toast configuration
func (t *Toast) Init(c *config.Config) error {
	mode := c.Handler.Toast.Mode

	if mode == "" {
		mode = os.Getenv("KW_TOAST_MODE")
	}

	t.Mode = mode
	if mode == "alert" || mode == "notify" {
		return nil
	} else {
		fmt.Println("Invalid Mode")
		return nil
	}

}

func (t *Toast) ObjectCreated(obj interface{}) {
	notifyToast(t, obj, "created")
}

func (t *Toast) ObjectDeleted(obj interface{}) {
	notifyToast(t, obj, "deleted")
}

func (t *Toast) ObjectUpdated(oldObj, newObj interface{}) {
	notifyToast(t, newObj, "updated")
}

func notifyToast(t *Toast, obj interface{}, action string) {
	e := kbEvent.New(obj, action)
	toastMessage := prepareToastMessage(e, t)
	switch t.Mode {
	case "notify":
		err := beeep.Notify("Kubewatch", toastMessage, "./docs/kubernetes.png")
		if err != nil {
			panic(err)
		}
	case "alert":
		err := beeep.Alert("Kubewatch", toastMessage, "./docs/kubernetes.png")
		if err != nil {
			panic(err)
		}
	}
}

func prepareToastMessage(e kbEvent.Event, t *Toast) string {
	var message string
	if e.Reason == "deleted" {
		message = e.Kind + " : " + e.Name + " : " + e.Reason
	} else {
		message = e.Kind + " : " + e.Namespace + "/" + e.Name + " : " + e.Reason
	}
	return message
}
