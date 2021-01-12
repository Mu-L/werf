package helm

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/werf/werf/cmd/werf/common"
	"github.com/werf/werf/pkg/git_repo"
	"github.com/werf/werf/pkg/true_git"
	"github.com/werf/werf/pkg/werf"
)

var getReleaseCmdData common.CmdData

func NewGetReleaseCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                   "get-release",
		DisableFlagsInUseLine: true,
		Short:                 "Print Helm Release name that will be used in current configuration with specified params",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := common.ProcessLogOptions(&getReleaseCmdData); err != nil {
				common.PrintHelp(cmd)
				return err
			}

			return runGetRelease()
		},
	}

	common.SetupDir(&getReleaseCmdData, cmd)
	common.SetupConfigTemplatesDir(&getReleaseCmdData, cmd)
	common.SetupConfigPath(&getReleaseCmdData, cmd)
	common.SetupEnvironment(&getReleaseCmdData, cmd)

	common.SetupGiterminismInspectorOptions(&getReleaseCmdData, cmd)

	common.SetupTmpDir(&getReleaseCmdData, cmd)
	common.SetupHomeDir(&getReleaseCmdData, cmd)
	common.SetupDockerConfig(&getReleaseCmdData, cmd, "")

	common.SetupLogOptions(&getReleaseCmdData, cmd)

	return cmd
}

func runGetRelease() error {
	if err := werf.Init(*getReleaseCmdData.TmpDir, *getReleaseCmdData.HomeDir); err != nil {
		return fmt.Errorf("initialization error: %s", err)
	}

	if err := common.InitGiterminismInspector(&getReleaseCmdData); err != nil {
		return err
	}

	if err := git_repo.Init(); err != nil {
		return err
	}

	if err := true_git.Init(true_git.Options{LiveGitOutput: *getReleaseCmdData.LogVerbose || *getReleaseCmdData.LogDebug}); err != nil {
		return err
	}

	projectDir, err := common.GetProjectDir(&getReleaseCmdData)
	if err != nil {
		return fmt.Errorf("getting project dir failed: %s", err)
	}

	giterminismManager, err := common.GetGiterminismManager(&getReleaseCmdData)
	if err != nil {
		return err
	}

	werfConfig, err := common.GetRequiredWerfConfig(common.BackgroundContext(), projectDir, &getReleaseCmdData, giterminismManager, common.GetWerfConfigOptions(&getReleaseCmdData, false))
	if err != nil {
		return fmt.Errorf("unable to load werf config: %s", err)
	}

	release, err := common.GetHelmRelease("", *getReleaseCmdData.Environment, werfConfig)
	if err != nil {
		return err
	}

	fmt.Println(release)

	return nil
}
