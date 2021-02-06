package main

import (
	"github.com/simplemoon/golog/conf"
	"github.com/simplemoon/golog/utils"

	"github.com/fsnotify/fsnotify"
	//"github.com/heirko/go-contrib/logrusHelper"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// 具体的定义
type logger struct {
	*logrus.Entry
}

var (
	gLogger  *logger
	viperCfg = viper.New()
	cfg      conf.Config // 配置文件
)

func initViper() {
	filename := utils.GetFileName(cfg.Path)
	viperCfg.SetConfigName(filename)
	// 如果 dir 不为空
	if cfg.Dir != "" {
		viperCfg.AddConfigPath(cfg.Dir)
	} else {
		viperCfg.AddConfigPath(".")
	}
	if cfg.ReloadWhenChange {
		viperCfg.WatchConfig()
		viperCfg.OnConfigChange(onConfigChange)
	}
	// 检查是否需要远程配置
	if cfg.EndPoints == "" {
		return
	}
	for _, p := range cfg.RemoteProviders {
		viperCfg.AddRemoteProvider(p, cfg.EndPoints, filename)
	}
	if cfg.ReloadWhenChange {
		viperCfg.WatchRemoteConfig()
	}
}

func loadConfig() error {
	err := viperCfg.ReadInConfig()
	if err == nil {
		return nil
	}
	err = viperCfg.ReadRemoteConfig()
	if err == nil {
		return nil
	}
	return err
}

// 初始化logger配置
func initLogger() {
	// 重新设置logger
	// 从配置之中解析结构体
	//var c = logrusHelper.UnmarshalConfiguration(viperCfg)
	//err := logrusHelper.SetConfig(logrus.StandardLogger(), c)
	//if err != nil {
	//	log.Printf("load from viper err %v", err)
	//}
	// TODO: 设置对应的配置路径
}

// 配置改变
func onConfigChange(in fsnotify.Event) {
	// log的配置改变时调用
	switch in.Op {
	case fsnotify.Write, fsnotify.Create:
	default:
		return
	}
	// 判断文件的名称
	filename := utils.GetFileName(cfg.Path)
	if filename != in.Name {
		return
	}
	// 重新设置
	initLogger()
}

// 设置配置路径
func SetConfig(path, dir, endpoint string, watch bool, provider ...string) {
	// 其他的配置
	if path == "" {
		// 当前目录下的config.json文件的配置
		cfg = conf.Config{
			Path:             "config.json",
			Dir:              ".",
			ReloadWhenChange: true,
		}
		return
	}
	// 配置文件的名字
	cfg = conf.Config{
		Path:             path,
		Dir:              dir,
		EndPoints:        endpoint,
		RemoteProviders:  provider,
		ReloadWhenChange: watch,
	}
}

// 初始化
func Init() {
	// 初始化viper
	initViper()
	// 加载配置
	err := loadConfig()
	if err != nil {
		panic(err)
	}
	// 设置logger的配置了,从文件之中读取配置文件
	initLogger()
}

// 输出debug的日志
func Debug(format string, args ...interface{}) {
	if gLogger == nil {
		return
	}
	gLogger.Debugf(format, args...)
}

func Warn(format string, args ...interface{}) {
	if gLogger == nil {
		return
	}
	gLogger.Warnf(format, args...)
}

func Info(format string, args ...interface{}) {
	if gLogger == nil {
		return
	}
	gLogger.Infof(format, args...)
}

func Error(format string, args ...interface{}) {
	if gLogger == nil {
		return
	}
	gLogger.Errorf(format, args...)
}

func Fatal(format string, args ...interface{}) {
	if gLogger == nil {
		return
	}
	gLogger.Fatalf(format, args...)
}

func WithField(key string, value interface{}) *logrus.Entry {
	if gLogger == nil {
		return nil
	}
	return gLogger.WithField(key, value)
}

func WithFields(fields logrus.Fields) *logrus.Entry {
	if gLogger == nil {
		return nil
	}
	return gLogger.WithFields(fields)
}

func WithError(err error) *logrus.Entry {
	if gLogger == nil {
		return nil
	}
	return gLogger.WithError(err)
}
