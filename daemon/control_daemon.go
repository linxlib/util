package daemon

import (
	"github.com/kardianos/service"
	"github.com/linxlib/logs"
	"os"
)

type Service interface {
	Run() error
	GetConfig() *service.Config
}

func Run(svc Service) error {
	if len(os.Args) < 2 {
		daemonctl, err := newDaemonController(svc)
		if err != nil {
			logs.Errorln("Error -> ", err)
			return err
		}
		if err := daemonctl.Run(); err != nil {
			logs.Errorln("Start Error -> ", err)
			return err
		}
		return nil
	}
	switch os.Args[1] {
	case "service":
		if len(os.Args) < 3 {
			logs.Errorln("缺少执行参数。如 service install 或者 service start。")
			break
		}
		daemonctl, err := newDaemonController(svc)
		if err != nil {
			logs.Errorln("Error -> ", err)
			return err
		}
		controlDaemon(daemonctl, os.Args[2])
	}
	return nil
}

// controlDaemon 操作守护进程
func controlDaemon(ctl *DaemonController, action string) {
	switch action {
	case "install":
		if err := ctl.Install(); err != nil {
			logs.Errorln("Install Error -> ", err)
			return
		}
		logs.Println("安装成功!")
	case "uninstall":
		if err := ctl.Uninstall(); err != nil {
			logs.Errorln("Uninstall Error -> ", err)
			return
		}
		logs.Println("卸载成功!")
	case "start":
		if err := ctl.Start(); err != nil {
			logs.Errorln("Start -> ", err)
			return
		}
		logs.Println("开启成功!")
	case "stop":
		if err := ctl.Stop(); err != nil {
			logs.Errorln("Stop -> ", err)
			return
		}
		logs.Println("停止成功!")
	case "restart":
		if err := ctl.Restart(); err != nil {
			logs.Errorln("Restart -> ", err)
			return
		}
		logs.Println("重启成功!")
	default:
		logs.Errorln("参数错误。")
	}
}
