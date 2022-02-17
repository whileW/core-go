package pkg

import (
	"fmt"
	"github.com/whileW/core-go/pkg/conf"
	"github.com/whileW/core-go/pkg/system_variable"
	"github.com/whileW/core-go/pkg/util/xcolor"
	"github.com/whileW/core-go/pkg/util/xdebug"
	"github.com/whileW/core-go/pkg/util/xgo"
	"sync"
	"time"
)

// 程序总控制器（入口、程序的生命周期）
type Application struct {
	StartTime 			time.Time

	DisableParserFlag	bool
	initOnce     		sync.Once
}

var (
	default_application		= 	&Application{}
)

func StartUp() {
	default_application.initialize(
		//default_application.parseFlags,
		conf.Initialize,
		//loki.InitByConf,
		default_application.printBanner,
	)
}
// initialize application
func (app *Application) initialize(initfns ...func()error) {
	app.initOnce.Do(func() {
		app.StartTime = time.Now()

		err := xgo.SerialUntilError(initfns...)()
		if err != nil {
			panic(err)
		}
	})
}

////parseFlags init
//func (app *Application) parseFlags() error {
//	//if app.isDisable(DisableParserFlag) {
//	//	app.logger.Info("parseFlags disable", xlog.FieldMod(ecode.ModApp))
//	//	return nil
//	//}
//	if app.DisableParserFlag {
//		fmt.Println("parseFlags disable")
//		return nil
//	}
//	return flag.Parse()
//}
//printBanner init
func (app *Application) printBanner() error {
	if xdebug.IsTestingMode() {
		return nil
	}

	const banner = `
  ____  ___________   ____              ____   ____  
_/ ___\/  _ \_  __ \_/ __ \   ______   / ___\ /  _ \ 
\  \__(  <_> )  | \/\  ___/  /_____/  / /_/  >  <_> )
 \___  >____/|__|    \___  >          \___  / \____/ 
     \/                  \/          /_____/
`
	fmt.Println(xcolor.Green(banner))
	fmt.Println("Welcome to core-go, starting application ...")
	fmt.Println()
	fmt.Println(fmt.Sprintf("启动时间：%v",app.StartTime))
	fmt.Println(fmt.Sprintf("系统名称：%s",system_variable.SystemName))
	fmt.Println(fmt.Sprintf("GoVersion：%s",system_variable.GoVersion))
	fmt.Println(fmt.Sprintf("Env：%s",system_variable.Env))
	return nil
}

// 禁用解析flag
func DisableParserFlag()  {
	default_application.DisableParserFlag = true
}