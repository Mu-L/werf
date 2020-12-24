package giterminism_inspector

import (
	"context"
	"errors"
	"fmt"

	"github.com/werf/logboek"
	"github.com/werf/werf/pkg/giterminism_inspector/config"
)

const giterminismDocPageURL = "https://werf.io/v1.2-alpha/documentation/advanced/configuration/giterminism.html"

var (
	LooseGiterminism         bool
	NonStrict                bool
	DevMode                  bool
	ReportedUncommittedPaths []string
	ReportedUntrackedPaths   []string

	giterminismConfig config.GiterminismConfig
)

type InspectionOptions struct {
	LooseGiterminism bool
	NonStrict        bool
	DevMode          bool
}

func Init(projectPath string, opts InspectionOptions) error {
	LooseGiterminism = opts.LooseGiterminism
	NonStrict = opts.NonStrict
	DevMode = opts.DevMode

	if c, err := config.PrepareConfig(projectPath); err != nil {
		return err
	} else {
		giterminismConfig = c
	}

	return nil
}

func IsUncommittedConfigAccepted() bool {
	return giterminismConfig.Config.AllowUncommitted
}

func IsUncommittedDockerfileAccepted(path string) (bool, error) {
	return giterminismConfig.Config.Dockerfile.IsUncommittedAccepted(path)
}

func IsUncommittedDockerignoreAccepted(path string) (bool, error) {
	return giterminismConfig.Config.Dockerfile.IsUncommittedDockerignoreAccepted(path)
}

func ReportUntrackedFile(ctx context.Context, path string) error {
	for _, p := range ReportedUntrackedPaths {
		if p == path {
			return nil
		}
	}
	ReportedUntrackedPaths = append(ReportedUntrackedPaths, path)

	if NonStrict {
		logboek.Context(ctx).Warn().LogF("WARNING: Untracked file %s was not taken into account (more info %s)\n", path, giterminismDocPageURL)
		return nil
	} else {
		return fmt.Errorf("restricted usage of untracked file %s (more info %s)", path, giterminismDocPageURL)
	}
}

func ReportUncommittedFile(ctx context.Context, path string) error {
	for _, p := range ReportedUncommittedPaths {
		if p == path {
			return nil
		}
	}
	ReportedUncommittedPaths = append(ReportedUncommittedPaths, path)

	if NonStrict {
		logboek.Context(ctx).Warn().LogF("WARNING: Uncommitted file %s was not taken into account (more info %s)\n", path, giterminismDocPageURL)
		return nil
	} else {
		return fmt.Errorf("restricted usage of uncommitted file %s (more info %s)", path, giterminismDocPageURL)
	}
}

func ReportConfigStapelGitBranch(_ context.Context) error {
	if giterminismConfig.Config.Stapel.Git.AllowBranch {
		return nil
	}

	return errors.New(`Pay attention, werf uses git repository history to calculate stages digests. Thus, the usage of remote git mapping with branch (by default, it is master branch) might break the reproducibility of previous builds. New commits in the branch will make previously built stages not usable.

* Previous pipeline jobs (e.g. deploy) might not be retried without the image rebuild after git remote branch is modified.
* If git remote branch is modified unexpectedly it might lead to the inexplicably failed pipeline. For instance, the modification occurs after successful build and the following pipeline jobs will be failed due to changing of stages digests alongside the branch HEAD.

If you want to use the branch for remote git mapping instead of commit or tag, then loose some restrictions of the werf giterminism mode with option --loose-giterminism (or WERF_LOOSE_GITERMINISM=1 env var). However it is NOT RECOMMENDED to use the remote git mapping in a such way. Use a particular unchangeable reference, tag, or commit to provide controllable and predictable lifecycle of software.`)
}

func ReportConfigStapelMountBuildDir(_ context.Context) error {
	if giterminismConfig.Config.Stapel.Mount.AllowBuildDir {
		return nil
	}

	return fmt.Errorf("'mount { from: build_dir, ... }' is forbidden due to enabled giterminism mode (more info %s), it is recommended to avoid this directive", giterminismDocPageURL)
}

func ReportConfigStapelMountFromPath(_ context.Context, fromPath string) error {
	if isAccepted, err := giterminismConfig.Config.Stapel.Mount.IsFromPathAccepted(fromPath); err != nil {
		return err
	} else if isAccepted {
		return nil
	}

	return fmt.Errorf("'mount { fromPath: %s, ... }' is forbidden due to enabled giterminism mode (more info %s), it is recommended to avoid this directive", fromPath, giterminismDocPageURL)
}

func ReportConfigDockerfileContextAddFile(_ context.Context, contextAddFile string) error {
	if isAccepted, err := giterminismConfig.Config.Dockerfile.IsContextAddFileAccepted(contextAddFile); err != nil {
		return err
	} else if isAccepted {
		return nil
	}

	return fmt.Errorf("'contextAddFile %s' is forbidden due to enabled giterminism mode (more info %s), it is recommended to avoid this directive", contextAddFile, giterminismDocPageURL)
}

func ReportConfigGoTemplateRenderingEnv(_ context.Context, envName string) error {
	if isAccepted, err := giterminismConfig.Config.GoTemplateRendering.IsEnvNameAccepted(envName); err != nil {
		return err
	} else if isAccepted {
		return nil
	}

	return fmt.Errorf("env name %s is forbidden due to enabled giterminism mode (more info %s)", envName, giterminismDocPageURL)
}

func PrintInspectionDebrief(ctx context.Context) {
	if NonStrict {
		if len(ReportedUncommittedPaths) > 0 || len(ReportedUntrackedPaths) > 0 {
			logboek.Context(ctx).Warn().LogLn()
			logboek.Context(ctx).Warn().LogF("### Giterminism inspection debrief ###\n")
			logboek.Context(ctx).Warn().LogLn()

			if len(ReportedUncommittedPaths) > 0 {
				logboek.Context(ctx).Warn().LogF("Following uncommitted files were not taken into account:\n")
				for _, path := range ReportedUncommittedPaths {
					logboek.Context(ctx).Warn().LogF(" - %s\n", path)
				}
				logboek.Context(ctx).Warn().LogLn()
			}

			if len(ReportedUntrackedPaths) > 0 {
				logboek.Context(ctx).Warn().LogF("Following untracked files were not taken into account:\n")
				for _, path := range ReportedUntrackedPaths {
					logboek.Context(ctx).Warn().LogF(" - %s\n", path)
				}
				logboek.Context(ctx).Warn().LogLn()
			}

			logboek.Context(ctx).Warn().LogF("More info about giterminism in the werf avaiable on the page: %s\n", giterminismDocPageURL)
			logboek.Context(ctx).Warn().LogLn()
		}
	}
}
