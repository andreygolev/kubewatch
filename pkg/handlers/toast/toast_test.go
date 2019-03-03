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
	"reflect"
	"testing"

	"github.com/bitnami-labs/kubewatch/config"
)

var toastErrMsg = `
%s

You need to set Toast mode for Toast notify,
using "--mode/-m" or using environment variables:

export KW_TOAST_MODE=toast_mode

Command line flags will override environment variables

`

func TestToastInit(t *testing.T) {
	s := &Toast{}
	expectedError := fmt.Errorf(toastErrMsg, "Missing Toast mode")

	var Tests = []struct {
		toast config.Toast
		err   error
	}{
		{config.Toast{Mode: "foo"}, nil},
		{config.Toast{}, expectedError},
	}

	for _, tt := range Tests {
		c := &config.Config{}
		c.Handler.Toast = tt.toast
		if err := s.Init(c); !reflect.DeepEqual(err, tt.err) {
			t.Fatalf("Init(): %v", err)
		}
	}
}
