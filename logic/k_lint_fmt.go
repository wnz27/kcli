/**
 * @project ads-serving
 * @Author 27
 * @Description //TODO
 * @Date 2021/12/2 14:13 12月
 **/
package logic

import (
	"fmt"
	"os/exec"

	"github.com/urfave/cli/v2"

	"kcli/pkg"
)

func FixLintFmtErr(c *cli.Context) error {
	targetFixFile := c.String("lint-file")
	autoFixFmtCmd := exec.Command("gofmt", "-s", "-w", targetFixFile)
	pkg.LogFatal(autoFixFmtCmd)
	fmt.Println("【kcli fix lint fmt success】")
	return nil
}
