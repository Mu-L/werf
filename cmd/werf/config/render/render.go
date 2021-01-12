package render

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/werf/werf/cmd/werf/common"
	"github.com/werf/werf/pkg/config"
	"github.com/werf/werf/pkg/git_repo"
	"github.com/werf/werf/pkg/true_git"
	"github.com/werf/werf/pkg/werf"
)

var commonCmdData common.CmdData

func NewCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                   "render [IMAGE_NAME...]",
		DisableFlagsInUseLine: true,
		Short:                 "Render werf.yaml",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := common.ProcessLogOptions(&commonCmdData); err != nil {
				common.PrintHelp(cmd)
				return err
			}

			if err := werf.Init(*commonCmdData.TmpDir, *commonCmdData.HomeDir); err != nil {
				return fmt.Errorf("initialization error: %s", err)
			}

			if err := common.InitGiterminismInspector(&commonCmdData); err != nil {
				return err
			}

			if err := git_repo.Init(); err != nil {
				return err
			}

			if err := true_git.Init(true_git.Options{LiveGitOutput: *commonCmdData.LogVerbose || *commonCmdData.LogDebug}); err != nil {
				return err
			}

			projectDir, err := common.GetProjectDir(&commonCmdData)
			if err != nil {
				return fmt.Errorf("getting project dir failed: %s", err)
			}

			giterminismManager, err := common.GetGiterminismManager(&commonCmdData)
			if err != nil {
				return err
			}

			configOpts := common.GetWerfConfigOptions(&commonCmdData, false)

			customWerfConfigRelPath, err := common.GetCustomWerfConfigRelPath(projectDir, &commonCmdData)
			if err != nil {
				return err
			}

			customWerfConfigTemplatesDirRelPath, err := common.GetCustomWerfConfigTemplatesDirRelPath(projectDir, &commonCmdData)
			if err != nil {
				return err
			}

			return config.RenderWerfConfig(common.BackgroundContext(), customWerfConfigRelPath, customWerfConfigTemplatesDirRelPath, args, giterminismManager, configOpts)
		},
	}

	common.SetupDir(&commonCmdData, cmd)
	common.SetupConfigTemplatesDir(&commonCmdData, cmd)
	common.SetupConfigPath(&commonCmdData, cmd)
	common.SetupEnvironment(&commonCmdData, cmd)

	common.SetupGiterminismInspectorOptions(&commonCmdData, cmd)

	common.SetupTmpDir(&commonCmdData, cmd)
	common.SetupHomeDir(&commonCmdData, cmd)

	common.SetupLogOptions(&commonCmdData, cmd)

	return cmd
}
