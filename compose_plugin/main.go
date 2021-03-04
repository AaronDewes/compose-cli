/*
   Copyright 2020 Docker Compose CLI authors

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

package main

import (
	"github.com/docker/cli/cli-plugins/manager"
	"github.com/docker/cli/cli-plugins/plugin"
	"github.com/docker/cli/cli/command"
	"github.com/docker/compose-cli/api/context/store"
	"github.com/docker/compose-cli/cli/cmd/compose"
	"github.com/docker/compose-cli/local"
	"github.com/spf13/cobra"
)

func main() {
	local.LoadLocalBackend()
	plugin.Run(func(dockerCli command.Cli) *cobra.Command {
		return compose.Command(store.DefaultContextType)
	},
		manager.Metadata{
			SchemaVersion: "0.1.0",
			Vendor:        "Docker Inc.",
			Version:       "compose 1.0-beta",
		})
}
