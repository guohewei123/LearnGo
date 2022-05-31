package zap_log

import (
	"errors"
	"go.uber.org/zap"
	"gopkg.in/alecthomas/kingpin.v2"
	"testing"
)

func TestInitLogger(t *testing.T) {
	if _, err := kingpin.CommandLine.Parse([]string{
		"--log.level", "debug",
		"--log.format", "logfmt",
		"--log.path", "../01_zap_log/log",
		"--log.filename", "test.log",
		"--log.file-max-size", "3",
		"--log.file-max-backups", "2"}); err != nil {
		t.Fatal(err)
	}
	if err := InitLogger(); err != nil {
		t.Fatal(err)
	}

	var logger *zap.SugaredLogger
	logger = Logger
	logger.Infof("测试一下啊：%s", "111")  // logger Infof 用法
	logger.Debugf("测试一下啊：%s", "111") // logger Debugf 用法
	logger.Errorf("测试一下啊：%s", "111") // logger Errorf 用法
	logger.Warnf("测试一下啊：%s", "111")  // logger Warnf 用法
	logger.Infof("测试一下啊：%s, %d, %v, %f", "111", 1111, errors.New("collector returned no data"), 3333.33)
	logger = logger.With("collector", "cpu", "name", "主机") // logger with 用法

}
