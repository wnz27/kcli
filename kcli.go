/**
 * @project ads-serving
 * @Author 27
 * @Description //TODO
 * @Date 2021/11/16 17:51 11月
 **/
package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"

	"kcli/logic"
)

func ParseFlag() {
	cmd := &cli.App{
		Name:    "kcli command tool",
		Usage:   "kcli 通用命令行工具",
		Version: "0.0.1",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "service_name",
				Aliases: []string{"sn"},
				Usage:   "服务名称可以为: \n \t1、service1 \n \t2、service2 \n \t3、service3\n",
			},

			// TODO 本地使用cli 可以根据服务写config 的文件位置
			//  如果该工具迭代的比较好的话要放在服务器跑的话里面的逻辑可以根据env来搞
			&cli.StringFlag{
				Name:     "config",
				Aliases:  []string{"c"},
				FilePath: "./config.json",
				Usage: "当前仅支持json文件, 默认为当前目录下的config.json文件\n" +
					"\t没有指定的话会 默认为 在 internal/服务名称/ 的路径下拼接config.json 文件\n",
			},

			&cli.StringFlag{
				Name:  "env",
				Value: "dev",
				Usage: "app运行环境, 当前仅支持: dev、test、prod",
			},
			&cli.StringFlag{
				Name:    "mode",
				Aliases: []string{"m"},
				Value:   "debug",
				Usage:   "app运行模式, 仅支持debug, release",
			},
			&cli.StringFlag{
				Name:    "grpc-gen-path",
				Aliases: []string{"gpath"},
				Value:   "./common/meta",
				Usage: "grpc 自动生成的代码文件的输出位置(建议绝对路径, 如果相对路径则根据本命令行执行位置来判断即可)\n" +
					"\t（暂默认 proto 文件也在这个位置）默认是: ./common/meta\n",

			},
			&cli.StringFlag{
				Name:    "proto-file-name",
				Aliases: []string{"pfile"},
				Value:   "",
				Usage:   "目标 proto 文件名称, \"service 和 client 在同一个文件\n",
			},
			&cli.StringFlag{
				Name:    "lint-file",
				Aliases: []string{"lfmt"},
				Value:   "",
				Usage:   "目标 需要执行 gofmt 的文件 路径",
			},
			//&cli.StringFlag{
			//	Name:    "encrypt",
			//	Aliases: []string{"e"},
			//	Value:   "",
			//	Usage:   "加密, 值为要加密的文本",
			//},
		},
		Commands: []*cli.Command{
			{
				Name:   "install",
				Usage:  "安装 kcli 命令行工具所有依赖 （Todo）",
				Action: logic.InstallDependencies,
			},
			{
				Name:   "run",
				Usage:  "启动微服务",
				Action: logic.ExecAdsServingRun,
			},
			{
				Name:   "build",
				Usage:  "编译微服务",
				Action: logic.ExecAdsServingBuild,
			},
			{
				Name: "gen-grpc",
				Usage: "根据 proto 文件 自动生成代码文件\n" +
					"\tuse: protoc --proto_path xxx --go_out=plugins=grpc: xxxx xxxx.proto\n",
				Action: logic.GenProtoFile,
			},
			{
				Name: "lint-fmt",
				Usage: "修复这类 lint 报错: File is not `gofmt`-ed with `-s` (gofmt)\n" +
					"\tuse: gofmt -s -w {{ lint-file }}\n",
				Action: logic.FixLintFmtErr,
			},
		},
	}

	err := cmd.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	ParseFlag()
}
