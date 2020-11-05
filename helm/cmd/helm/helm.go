package main // import "helm.sh/helm/v3/cmd/helm"

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"sigs.k8s.io/yaml"

	// Import to initialize client auth plugins.
	_ "k8s.io/client-go/plugin/pkg/client/auth"

	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/cli"
	"helm.sh/helm/v3/pkg/gates"
	kubefake "helm.sh/helm/v3/pkg/kube/fake"
	"helm.sh/helm/v3/pkg/release"
	"helm.sh/helm/v3/pkg/storage/driver"
)

// FeatureGateOCI是检查“helm chart”和“helm registry”命令是否可以工作的feature gate
const FeatureGateOCI = gates.Gate("HELM_EXPERIMENTAL_OCI")

var settings = cli.New()

func init() {
	log.SetFlags(log.Lshortfile)
}

func debug(format string, v ...interface{}) {
	if settings.Debug {
		format = fmt.Sprintf("[debug] %s\n", format)
		log.Output(2, fmt.Sprintf(format, v...))
	}
}

func warning(format string, v ...interface{}) {
	format = fmt.Sprintf("WARNING: %s\n", format)
	fmt.Fprintf(os.Stderr, format, v...)
}

func main() {
	actionConfig := new(action.Configuration)
	cmd, err := newRootCmd(actionConfig, os.Stdout, os.Args[1:])
	if err != nil {
		debug("%+v", err)
		os.Exit(1)
	}

	// run when each command's execute method is called
	cobra.OnInitialize(func() {
		helmDriver := os.Getenv("HELM_DRIVER")
		if err := actionConfig.Init(settings.RESTClientGetter(), settings.Namespace(), helmDriver, debug); err != nil {
			log.Fatal(err)
		}
		if helmDriver == "memory" {
			loadReleasesInMemory(actionConfig)
		}
	})

	if err := cmd.Execute(); err != nil {
		debug("%+v", err)
		switch e := err.(type) {
		case pluginError:
			os.Exit(e.code)
		default:
			os.Exit(1)
		}
	}
}

func checkOCIFeatureGate() func(_ *cobra.Command, _ []string) error {
	return func(_ *cobra.Command, _ []string) error {
		if !FeatureGateOCI.IsEnabled() {
			return FeatureGateOCI.Error()
		}
		return nil
	}
}

// This function loads releases into the memory storage if the
// environment variable is properly set.
func loadReleasesInMemory(actionConfig *action.Configuration) {
	filePaths := strings.Split(os.Getenv("HELM_MEMORY_DRIVER_DATA"), ":")
	if len(filePaths) == 0 {
		return
	}

	store := actionConfig.Releases
	mem, ok := store.Driver.(*driver.Memory)
	if !ok {
		// For an unexpected reason we are not dealing with the memory storage driver.
		return
	}

	actionConfig.KubeClient = &kubefake.PrintingKubeClient{Out: ioutil.Discard}

	for _, path := range filePaths {
		b, err := ioutil.ReadFile(path)
		if err != nil {
			log.Fatal("Unable to read memory driver data", err)
		}

		releases := []*release.Release{}
		if err := yaml.Unmarshal(b, &releases); err != nil {
			log.Fatal("Unable to unmarshal memory driver data: ", err)
		}

		for _, rel := range releases {
			if err := store.Create(rel); err != nil {
				log.Fatal(err)
			}
		}
	}
	// Must reset namespace to the proper one
	mem.SetNamespace(settings.Namespace())
}
