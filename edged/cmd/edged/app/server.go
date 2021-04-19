package app

import (
	"errors"
	"fmt"
	"github.com/kubeedge/beehive/pkg/core"
	"github.com/kubeedge/kubeedge/common/constants"
	"github.com/kubeedge/kubeedge/edged/cmd/edged/app/options"
	"github.com/kubeedge/kubeedge/edged/pkg/client"
	"github.com/kubeedge/kubeedge/edged/pkg/edged"
	"github.com/kubeedge/kubeedge/pkg/apis/componentconfig/edged/v1alpha1"
	"github.com/kubeedge/kubeedge/pkg/apis/componentconfig/edged/v1alpha1/validation"
	"github.com/kubeedge/kubeedge/pkg/util"
	"github.com/kubeedge/kubeedge/pkg/util/flag"
	"github.com/kubeedge/kubeedge/pkg/version"
	"github.com/kubeedge/kubeedge/pkg/version/verflag"
	"github.com/mitchellh/go-ps"
	"github.com/spf13/cobra"
	cliflag "k8s.io/component-base/cli/flag"
	"k8s.io/component-base/cli/globalflag"
	"k8s.io/component-base/term"
	"k8s.io/klog/v2"
	"os"
)

// NewEdgedCommand create edged cmd
func NewEdgedCommand() *cobra.Command {
	opts := options.NewEdgedOptions()
	cmd := &cobra.Command{
		Use: "edged",
		Long: `what?`,
		Run: func(cmd *cobra.Command, args []string) {
			verflag.PrintAndExitIfRequested()
			flag.PrintMinConfigAndExitIfRequested(v1alpha1.NewMinEdgedConfig())
			flag.PrintDefaultConfigAndExitIfRequested(v1alpha1.NewDefaultEdgedConfig())
			flag.PrintFlags(cmd.Flags())

			if errs := opts.Validate(); len(errs) > 0 {
				klog.Fatal(util.SpliceErrors(errs))
			}

			config, err := opts.Config()
			if err != nil {
				klog.Fatal(err)
			}

			if errs := validation.ValidateEdgedConfiguration(config); len(errs) > 0 {
				klog.Fatal(util.SpliceErrors(errs.ToAggregate().Errors()))
			}

			// To help debugging, immediately log version
			klog.Infof("Version: %+v", version.Get())

			// Check the running environment by default
			checkEnv := os.Getenv("CHECK_EDGECORE_ENVIRONMENT")
			if checkEnv != "false" {
				// Check running environment before run edge core
				if err := environmentCheck(); err != nil {
					klog.Fatal(fmt.Errorf("Failed to check the running environment: %v", err))
				}
			}

			// get edge node local ip
			if config.Modules.Edged.NodeIP == "" {
				hostnameOverride, err := os.Hostname()
				if err != nil {
					hostnameOverride = constants.DefaultHostnameOverride
				}
				localIP, _ := util.GetLocalIP(hostnameOverride)
				config.Modules.Edged.NodeIP = localIP
			}

			client.InitEdgedClient(config.KubeAPIConfig)


			registerModules(config)
			// start all modules
			core.Run()
		},
	}
	fs := cmd.Flags()
	namedFs := opts.Flags()
	flag.AddFlags(namedFs.FlagSet("global"))
	verflag.AddFlags(namedFs.FlagSet("global"))
	globalflag.AddGlobalFlags(namedFs.FlagSet("global"), cmd.Name())
	for _, f := range namedFs.FlagSets {
		fs.AddFlagSet(f)
	}

	usageFmt := "Usage:\n  %s\n"
	cols, _, _ := term.TerminalSize(cmd.OutOrStdout())
	cmd.SetUsageFunc(func(cmd *cobra.Command) error {
		fmt.Fprintf(cmd.OutOrStderr(), usageFmt, cmd.UseLine())
		cliflag.PrintSections(cmd.OutOrStderr(), namedFs, cols)
		return nil
	})
	cmd.SetHelpFunc(func(cmd *cobra.Command, args []string) {
		fmt.Fprintf(cmd.OutOrStdout(), "%s\n\n"+usageFmt, cmd.Long, cmd.UseLine())
		cliflag.PrintSections(cmd.OutOrStdout(), namedFs, cols)
	})

	return cmd
}


// findProcess find a running process by name
func findProcess(name string) (bool, error) {
	processes, err := ps.Processes()
	if err != nil {
		return false, err
	}

	for _, process := range processes {
		if process.Executable() == name {
			return true, nil
		}
	}

	return false, nil
}



// environmentCheck check the environment before edged start
// if Check failed,  return errors
func environmentCheck() error {
	// if kubelet is running, return error
	if find, err := findProcess("kubelet"); err != nil {
		return err
	} else if find {
		return errors.New("Kubelet should not running on edge node when running edgecore")
	}

	// if kube-proxy is running, return error
	if find, err := findProcess("kube-proxy"); err != nil {
		return err
	} else if find {
		return errors.New("Kube-proxy should not running on edge node when running edgecore")
	}

	return nil
}

// registerModules register all the modules started in edged
func registerModules(c *v1alpha1.EdgedConfig) {
	edged.Register(c.Modules.Edged)
}