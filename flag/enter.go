package flag

import (
	"errors"
	"fmt"
	"os"
	"server/global"

	"github.com/urfave/cli"
	"go.uber.org/zap"
)

// 定义 CLI 标志，用于不同操作的命令行选项
var (
	sqlFlag = &cli.BoolFlag{
		Name:  "sql",
		Usage: "Initializes the srtucture of the MySQL database table.",
	}
	sqlExportFlag = &cli.BoolFlag{
		Name:  "sql-export",
		Usage: "Exports the structure of the MySQL database table.",
	}
	sqlImportFlag = &cli.StringFlag{
		Name:  "sql-import",
		Usage: "Imports SQL data from a specified file.",
	}
)

// Run 执行基于命令行标志的相应操作
// 它处理不同的标志，执行相应操作，并记录成功或错误的消息
func Run(c *cli.Context) {
	// 检查是否设置了多个标志
	if c.NumFlags() > 1 {
		err := cli.NewExitError("Only one command can be specified", 1)
		global.Log.Error("Invalid command usage:", zap.Error(err))
		os.Exit(1)
	}

	// 根据不同的标志选择执行的操作
	switch {
	case c.Bool(sqlFlag.Name):
		if err := SQL(); err != nil {
			global.Log.Error("Failed to create table structure:", zap.Error(err))
			return
		} else {
			global.Log.Info("Successfully created table structure")
		}
	case c.Bool(sqlExportFlag.Name):
		if err := SQLExport(); err != nil {
			global.Log.Error("Failed to export table structure:", zap.Error(err))
			return
		} else {
			global.Log.Info("Successfully exported table structure")
		}
	case c.IsSet(sqlImportFlag.Name):
		if errs := SQLImport(c.String(sqlImportFlag.Name)); len(errs) > 0 {
			var combinedErrors string
			for _, err := range errs {
				combinedErrors += err.Error() + "\n"
			}
			err := errors.New(combinedErrors)
			global.Log.Error("Failed to import SQL data:", zap.Error(err))
		} else {
			global.Log.Info("Successfully imported SQL data")
		}
	default:
		err := cli.NewExitError("unknown command", 1)
		global.Log.Error(err.Error(), zap.Error(err))
	}
}

// NewApp 创建并配置一个新的 CLI 应用程序，设置标志和默认操作
func NewApp() *cli.App {
	app := cli.NewApp()
	app.Name = "Go Blog"
	app.Flags = []cli.Flag{
		sqlFlag,
		sqlExportFlag,
		sqlImportFlag,
	}
	app.Action = Run
	return app
}

// InitFlag 初始化并运行 CLI 应用程序
func InitFlag() {
	if len(os.Args) > 1 {
		app := NewApp()
		err := app.Run(os.Args)
		if err != nil {
			global.Log.Error("Application execution encountered an error:", zap.Error(err))
			os.Exit(1)
		}
		if os.Args[1] == "-h" || os.Args[1] == "-help" {
			fmt.Println("Displaying help message...")
		}
		os.Exit(0)
	}
}
