package initialize

import (
	"flag"
	"fmt"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/fsnotify/fsnotify"
	"github.com/pkg/errors"
	"github.com/sevenzx/eztodo/conf"
	"github.com/spf13/viper"
)

// LoadConfigFile 使用viper加载配置文件
// 优先级: 命令行 > 默认值 > 自定义路径
func LoadConfigFile(paths ...string) {
	var path string

	if len(paths) == 0 {
		flag.StringVar(&path, "c", "", "choose path file.")
		flag.Parse()
		if path != "" {
			// 1. 使用命令行
			hlog.Infof("Using cmd line file: %s\n", path)
		} else {
			// 2. 使用默认值
			path = "config.yaml"
			hlog.Infof("Using default file: %s\n", path)
		}
	} else {
		// 3. 使用自定义路径
		path = paths[0]
		hlog.Infof("Using custom file: %s\n", path)
	}

	v := viper.New()
	// 指定配置文件路径
	v.SetConfigFile(path)
	// 指定配置文件类型
	v.SetConfigType("yaml")
	// 读取配置信息
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error path file: %s \n", err))
	}
	// 监控配置文件变化
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		hlog.Infof("viper: config file(%s) changed", e.Name)
		if err = v.Unmarshal(&conf.Config); err != nil {
			hlog.Errorf("%+v", errors.Wrap(err, "unmarshal config file error"))
		}
	})
	// 将配置文件加载到conf.Config
	if err = v.Unmarshal(&conf.Config); err != nil {
		panic(err)
	}
}
