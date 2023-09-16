// Copyright © 2018 Raul Sampedro
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

package cli

import (
	"github.com/kkontosis/SaSSHimi/agent"
	"github.com/spf13/cobra"
)

var useHttpProxy bool
var keepBinary bool

// agentCmd represents the agent command
var agentCmd = &cobra.Command{
	Use:   "agent",
	Short: "Run as remote agent process",
	Run: func(cmd *cobra.Command, args []string) {
		agent.Run(useHttpProxy, keepBinary)
	},
}

func init() {
	rootCmd.AddCommand(agentCmd)

	agentCmd.Flags().BoolVar(&useHttpProxy, "use-http", false, "Use HTTP proxy instead of HTTP")
	agentCmd.Flags().BoolVarP(&keepBinary, "keep-binary", "k",  false, "Do not remove binary when closing")
}
