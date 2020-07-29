// Copyright 2017 Baidu, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file
// except in compliance with the License. You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software distributed under the
// License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
// either express or implied. See the License for the specific language governing permissions
// and limitations under the License.

package net_tools

import (
	"os/exec"
	"runtime"
	"unicode/utf8"
)

import (
	"github.com/axgle/mahonia"
)

func Ping(address *string, count string) (string, error) {
	var (
		pingCountArg []string
	)

	if runtime.GOOS == "windows" {
		pingCountArg = []string{"-n", count, *address}
	} else {
		pingCountArg = []string{"-c", count, *address}
	}

	pingOut, err := exec.Command("ping", pingCountArg...).Output()

	ret := string(pingOut)
	// TODO in windows the output of command ping maybe is gbk, we need write a ping instead using
	// ping of system
	if !utf8.ValidString(ret) {
		ret = mahonia.NewDecoder("gbk").ConvertString(ret)
	}
	return ret, err
}
