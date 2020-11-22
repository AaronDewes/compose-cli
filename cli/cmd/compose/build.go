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

package compose

import (
	"context"

	"github.com/compose-spec/compose-go/cli"
	"github.com/spf13/cobra"

	"github.com/docker/compose-cli/api/client"
)

func buildCommand() *cobra.Command {
	opts := composeOptions{}
	buildCmd := &cobra.Command{
		Use: "build",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runBuild(cmd.Context(), opts)
		},
	}
	buildCmd.Flags().StringVarP(&opts.Name, "project-name", "p", "", "Project name")
	buildCmd.Flags().StringVar(&opts.WorkingDir, "workdir", "", "Work dir")
	buildCmd.Flags().StringArrayVarP(&opts.ConfigPaths, "file", "f", []string{}, "Compose configuration files")

	return buildCmd
}

func runBuild(ctx context.Context, opts composeOptions) error {
	c, err := client.New(ctx)
	if err != nil {
		return err
	}

	options, err := opts.toProjectOptions()
	if err != nil {
		return err
	}

	project, err := cli.ProjectFromOptions(options)
	if err != nil {
		return err
	}

	return c.ComposeService().Build(ctx, project)
}
