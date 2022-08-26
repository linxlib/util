package daemon

import (
	"fmt"
	"github.com/kardianos/service"
	"os"
)

type Daemon struct {
	config *service.Config
	svc    Service
	errs   chan error
}

func NewDaemon(svc Service) *Daemon {
	c := svc.GetConfig()
	if c == nil {
		return nil
	}
	return &Daemon{
		config: c,
		errs:   make(chan error, 100),
		svc:    svc,
	}
}

func (d *Daemon) Start(s service.Service) error {
	go d.run()
	return nil
}

func (d *Daemon) run() {
	// 运行逻辑
	err := d.svc.Run()
	if err != nil {
		fmt.Println("Start Server Error -> ", err)
		os.Exit(1)
		return
	}
}

func (d *Daemon) Stop(s service.Service) error {
	if service.Interactive() {
		os.Exit(0)
	}
	return nil
}

// DaemonController 守护进程控制器
type DaemonController struct {
	service service.Service
}

// newDaemonController 实例化守护进程控制器
func newDaemonController(svc Service) (*DaemonController, error) {
	d := NewDaemon(svc)
	if d == nil {
		return nil, fmt.Errorf("not service config")
	}
	if service.Platform() == "linux-systemd" {
		d.config.Option = service.KeyValue{
			"LimitNOFILE": 40960,
		}
	}

	s, err := service.New(d, d.config)
	if err != nil {
		return nil, err
	}
	return &DaemonController{
		service: s,
	}, nil
}

// Install 安装守护进程
func (d *DaemonController) Install() error {
	return d.service.Install()
}

// Uninstall 卸载守护进程
func (d *DaemonController) Uninstall() error {
	d.service.Stop()
	return d.service.Uninstall()
}

// Start 开始守护进程
func (d *DaemonController) Start() error {
	return d.service.Start()
}

// Stop 停止守护进程
func (d *DaemonController) Stop() error {
	return d.service.Stop()
}

// Restart 重启守护进程
func (d *DaemonController) Restart() error {
	return d.service.Restart()
}

// Run 命令行运行守护进程
func (d *DaemonController) Run() error {
	return d.service.Run()
}
