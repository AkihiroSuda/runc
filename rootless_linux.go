// +build linux

package main

import (
	"os"

	"github.com/opencontainers/runc/libcontainer/system"
	"github.com/urfave/cli"
)

func shouldUseLenientCgroupManager(context *cli.Context) (bool, error) {
	// for backward compatibility, "--rootless" is used for explicitly specifying
	// whether we should use the lenient cgroup manager, which ignores permission errors.
	if context != nil {
		b, err := parseBoolOrAuto(context.GlobalString("rootless"))
		if err != nil {
			return false, err
		}
		if b != nil {
			return *b, nil
		}
		// nil b stands for "auto detect"
	}
	if context.GlobalBool("systemd-cgroup") {
		return false, nil
	}
	if os.Geteuid() != 0 {
		return true, nil
	}
	if !system.RunningInUserNS() {
		// uid == 0 , in the initial ns (i.e. the real root)
		return false, nil
	}
	// euid = 0, in a userns.
	// As we are unaware of cgroups path, we can't determine whether we have the full
	// access to the cgroups path.
	// Either way, we can safely decide to use the leninent cgroups manager.
	return true, nil
}

func shouldHonorXDGRuntimeDir() bool {
	if os.Geteuid() != 0 {
		return true
	}
	if !system.RunningInUserNS() {
		// uid == 0 , in the initial ns (i.e. the real root)
		return false
	}
	// euid = 0, in a userns.
	u := os.Getenv("USER")
	return u != "" && u != "root" && os.Getenv("XDG_RUNTIME_DIR") != ""
}
