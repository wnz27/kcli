/**
 * @project ads-serving
 * @Author 27
 * @Description //TODO
 * @Date 2021/11/20 11:14 11月
 **/
package logic

import (
	"fmt"
	"kcli/pkg"
	"os/exec"

	"github.com/urfave/cli/v2"
)

// InstallDependencies 安装依赖
func InstallDependencies(c *cli.Context) error {
	fmt.Println("Unimplement!!!!")
	return nil
}

// defaultConfigPath 填充config 路径, 从根目录配置
func defaultConfigPath(c *cli.Context, serviceName string) string {
	cliConfigPath := c.String("config")
	if len(cliConfigPath) != 0 {
		return cliConfigPath
	}
	return "./config/local/" + serviceName + "_config_local.json"
}

// ExecAdsServingRun 运行服务 需执行编译
func ExecAdsServingRun(c *cli.Context) error {
	serviceName := c.String("service_name")
	pkg.ValidServiceName(serviceName)
	// 运行微服务
	configPath := defaultConfigPath(c, serviceName)

	envStr := c.String("env")
	appMode := c.String("mode")

	fmt.Printf("kcli starting service: %s \n \tuse config: %s \n \tenv: %s \n \tmode: %s\n",
		serviceName, configPath, envStr, appMode)

	execPath := "./internal/" + serviceName
	cmd := exec.Command("./"+serviceName, "-m", appMode, "-c", configPath, "-e", envStr)
	// 这里运行情况会持续输出到命令行
	pkg.LogFatalChangeDirForRun(cmd, execPath)
	return nil
}


// ExecAdsServingBuild 编译服务
func ExecAdsServingBuild(c *cli.Context) error {
	// 编译微服务
	serviceName := c.String("service_name")
	pkg.ValidServiceName(serviceName)

	servicePath := "./internal/" + serviceName
	outTargetPath := "./internal/" + serviceName + "/"

	buildCmd := exec.Command("go", "build", "-o", outTargetPath, servicePath)
	pkg.LogFatal(buildCmd)
	fmt.Println("【kcli build success】")
	return nil
}
