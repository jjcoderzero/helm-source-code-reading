package main

import (
	"fmt"
	"io"
	"path/filepath"

	"github.com/spf13/cobra"

	"helm.sh/helm/v3/cmd/helm/require"
	"helm.sh/helm/v3/pkg/chart"
	"helm.sh/helm/v3/pkg/chartutil"
	"helm.sh/helm/v3/pkg/helmpath"
)

const createDesc = `
This command creates a chart directory along with the common files and
directories used in a chart.

For example, 'helm create foo' will create a directory structure that looks
something like this:

    foo/
    ├── .helmignore   # Contains patterns to ignore when packaging Helm charts.
    ├── Chart.yaml    # Information about your chart
    ├── values.yaml   # The default values for your templates
    ├── charts/       # Charts that this chart depends on
    └── templates/    # The template files
        └── tests/    # The test files

'helm create' takes a path for an argument. If directories in the given path
do not exist, Helm will attempt to create them as it goes. If the given
destination exists and there are files in that directory, conflicting files
will be overwritten, but other files will be left alone.
`

type createOptions struct { // create结构体
	starter    string // --starter
	name       string
	starterDir string
}

func newCreateCmd(out io.Writer) *cobra.Command {
	o := &createOptions{} // 初始化结构体

	cmd := &cobra.Command{
		Use:   "create NAME",
		Short: "create a new chart with the given name",
		Long:  createDesc,
		Args:  require.ExactArgs(1),
		ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			if len(args) == 0 {
				// 当完成可能是路径的名称的参数时，允许文件完成
				return nil, cobra.ShellCompDirectiveDefault
			}
			// 没有更多的完成，因此禁用文件完成
			return nil, cobra.ShellCompDirectiveNoFileComp
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			o.name = args[0] // 设置name
			o.starterDir = helmpath.DataPath("starters") // 设置starter目录
			return o.run(out)
		},
	}

	cmd.Flags().StringVarP(&o.starter, "starter", "p", "", "the name or absolute path to Helm starter scaffold") // starter选项
	return cmd
}

func (o *createOptions) run(out io.Writer) error {
	fmt.Fprintf(out, "Creating %s\n", o.name)

	chartname := filepath.Base(o.name) // 获取chart名称
	cfile := &chart.Metadata{ // 构造chart元数据
		Name:        chartname,
		Description: "A Helm chart for Kubernetes",
		Type:        "application",
		Version:     "0.1.0",
		AppVersion:  "0.1.0",
		APIVersion:  chart.APIVersionV2,
	}

	if o.starter != "" {
		lstarter := filepath.Join(o.starterDir, o.starter) // 从启动器创建
		if filepath.IsAbs(o.starter) { // 如果路径是绝对的，我们不想在它前面加上启动器文件夹
			lstarter = o.starter
		}
		return chartutil.CreateFrom(cfile, filepath.Dir(o.name), lstarter)
	}

	chartutil.Stderr = out
	_, err := chartutil.Create(chartname, filepath.Dir(o.name))
	return err
}
