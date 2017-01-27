//  Copyright 2016 Red Hat, Inc.
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/funktionio/funktion/pkg/version"
)

func init() {
	var versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Print the version of funktion.",
		Long:  `Print the version of funktion.`,
		Run: func(command *cobra.Command, args []string) {
			fmt.Println("funktion version:", version.GetVersion())
		},
	}
	RootCmd.AddCommand(versionCmd)
}
