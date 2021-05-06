// Copyright 2021 Ray Harris
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package fork creates a fork of your application and executes it.
// Standard input, output, and error are redirected to /dev/null.
// This detaches your application from the command line and runs it
// in the background.
//
// To use, import it into your application:
//
//     import (
//         _ "github.com/ray-harris/fork"
//     )
package fork

import (
	"os"
	"syscall"
)

const (
	FORK_FLAG     = "FORKED_PROCESS_INDICATOR"
	FORK_FLAG_SET = FORK_FLAG + "=true"
)

func fork() {
	if os.Getenv(FORK_FLAG) == "" {
		null, err := os.Create(os.DevNull)
		if err != nil {
			panic(err)
		}
		fd := null.Fd()
		syscall.ForkExec(os.Args[0], os.Args,
			&syscall.ProcAttr{
				Env: append(os.Environ(), FORK_FLAG_SET),
				Sys: &syscall.SysProcAttr{
					Setsid: true,
				},
				Files: []uintptr{fd, fd, fd},
			},
		)
		os.Exit(0)
	}
}

func init() {
	fork()
}
