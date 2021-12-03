/**
 * @project ads-serving
 * @Author 27
 * @Description //TODO
 * @Date 2021/11/20 11:40 11月
 **/
package logic

import (
	"fmt"
	"kcli/pkg"
	"os/exec"

	"github.com/urfave/cli/v2"
)

func GenProtoFile(c *cli.Context) error {
	protoFilePath := c.String("grpc-gen-path")
	protoFileName := c.String("proto-file-name")

	// use protoc --proto_path ./common/meta --go_out=plugins=grpc:./common/meta advert.proto
	autoGenProtoCmd := exec.Command("protoc", "--proto_path", protoFilePath,
		"--go_out=plugins=grpc:"+protoFilePath, protoFileName)

	pkg.LogFatal(autoGenProtoCmd)
	fmt.Println("【kcli gen protofile  success】")
	return nil
}
