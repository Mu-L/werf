package config

import (
	"context"
	"fmt"

	"github.com/werf/werf/pkg/giterminism_inspector"
)

type Mount struct {
	To   string
	From string
	Type string

	raw *rawMount
}

func (c *Mount) validate() error {
	if c.raw.FromPath != "" {
		if err := giterminism_inspector.ReportConfigStapelMountFromPath(context.Background(), c.raw.FromPath); err != nil {
			return err
		}
	} else if c.Type == "build_dir" {
		if err := giterminism_inspector.ReportConfigStapelMountBuildDir(context.Background()); err != nil {
			return err
		}
	}

	if c.raw.From != "" && c.raw.FromPath != "" {
		return newDetailedConfigError(fmt.Sprintf("cannot use `from: %s` and `fromPath: %s` at the same time for mount!", c.raw.From, c.raw.FromPath), c, c.raw.rawStapelImage.doc)
	}

	if c.To == "" || !isAbsolutePath(c.To) {
		return newDetailedConfigError("`to: PATH` absolute path required for mount!", c.raw, c.raw.rawStapelImage.doc)
	} else if c.Type == "custom_dir" {
		if c.From == "" {
			return newDetailedConfigError("`fromPath: PATH` absolute or relative path required for mount!", c.raw, c.raw.rawStapelImage.doc)
		}
	} else if c.Type != "tmp_dir" && c.Type != "build_dir" {
		return newDetailedConfigError(fmt.Sprintf("invalid `from: %s` for mount: expected `tmp_dir` or `build_dir`!", c.Type), c.raw, c.raw.rawStapelImage.doc)
	}
	return nil
}
