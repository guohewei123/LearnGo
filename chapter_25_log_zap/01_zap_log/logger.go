package zap_log

// go get -u go.uber.org/zap
// go get -u github.com/natefinch/lumberjack
// go get gopkg.in/alecthomas/kingpin.v2
import (
    "github.com/natefinch/lumberjack"
    "go.uber.org/zap"
    "go.uber.org/zap/zapcore"
    "gopkg.in/alecthomas/kingpin.v2"
    "os"
    "path/filepath"
    "strconv"
)

const (
    DefaultLogPath           = "/var/log/test" // 【默认】日志文件路径
    DefaultLogfileName       = "test.log"      // 【默认】日志文件名称
    DefaultLogLevel          = "info"          // 【默认】日志打印级别 debug  info  warning  error
    DefaultLogfileMaxSize    = 5               // 【日志分割】  【默认】单个日志文件最多存储量 单位(MB)
    DefaultLogfileMaxBackups = 10              // 【日志分割】  【默认】日志备份文件最多数量
    LogMaxAge                = 1000            // 【默认】日志保留时间，单位: 天 (day)
    LogCompress              = false           // 【默认】是否压缩日志
    LogStdout                = false           // 【默认】是否输出到控制台
)

var Logger *zap.SugaredLogger // 定义日志打印全局变量

// kingpin 可以在启动时通过输入参数，来修改日志参数
var (
    Level             = kingpin.Flag("log.level", "Only log messages with the given severity or above. One of: [debug, info, warn, error]").Default(DefaultLogLevel).String()
    Format            = kingpin.Flag("log.format", "Output format of log messages. One of: [logfmt, json]").Default("logfmt").String()
    LogPath           = kingpin.Flag("log.path", "Output log path").Default(DefaultLogPath).String()
    LogFilename       = kingpin.Flag("log.filename", "Output log filename").Default(DefaultLogfileName).String()
    LogfileMaxSize    = kingpin.Flag("log.file-max-size", "Output logfile max size, unit MB").Default(strconv.Itoa(DefaultLogfileMaxSize)).Int()
    LogfileMaxBackups = kingpin.Flag("log.file-max-backups", "Output logfile max backups").Default(strconv.Itoa(DefaultLogfileMaxBackups)).Int()
)

// InitLogger 初始化 logger
func InitLogger() error {
    logLevel := map[string]zapcore.Level{
        "debug": zapcore.DebugLevel,
        "info":  zapcore.InfoLevel,
        "warn":  zapcore.WarnLevel,
        "error": zapcore.ErrorLevel,
    }
    writeSyncer, err := getLogWriter() // 日志文件配置 文件位置和切割
    if err != nil {
        return err
    }
    encoder := getEncoder()       // 获取日志输出编码
    level, ok := logLevel[*Level] // 日志打印级别
    if !ok {
        level = logLevel["info"]
    }
    core := zapcore.NewCore(encoder, writeSyncer, level)
    logger := zap.New(core, zap.AddCaller()) //  zap.AddCaller() 输出日志打印文件和行数如： logger/logger_test.go:33
    Logger = logger.Sugar()
    return nil
}

// 编码器(如何写入日志)
func getEncoder() zapcore.Encoder {
	zap.NewProductionConfig()

    encoderConfig := zap.NewProductionEncoderConfig()
    encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder   // looger 时间格式 例如: 2021-09-11T20:05:54.852+0800
    encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder // 输出Level序列化为全大写字符串，如 INFO DEBUG ERROR
    //encoderConfig.EncodeCaller = zapcore.FullCallerEncoder
    //encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
    if *Format == "json" {
        return zapcore.NewJSONEncoder(encoderConfig) // 以json格式写入
    }
    return zapcore.NewConsoleEncoder(encoderConfig) // 以logfmt格式写入
}

// 获取日志输出方式  日志文件 控制台
func getLogWriter() (zapcore.WriteSyncer, error) {
    // 判断日志路径是否存在，如果不存在就创建
    if exist := IsExist(*LogPath); !exist {
        if *LogPath == "" {
            *LogPath = DefaultLogPath
        }
        if err := os.MkdirAll(*LogPath, os.ModePerm); err != nil {
            *LogPath = DefaultLogPath
            if err := os.MkdirAll(*LogPath, os.ModePerm); err != nil {
                return nil, err
            }
        }
    }
    // 日志文件 与 日志切割 配置
    lumberJackLogger := &lumberjack.Logger{
        Filename:   filepath.Join(*LogPath, *LogFilename), // 日志文件路径
        MaxSize:    *LogfileMaxSize,                       // 单个日志文件最大多少 MB
        MaxBackups: *LogfileMaxBackups,                    // 日志备份数量
        MaxAge:     LogMaxAge,                             // 日志最长保留时间
        Compress:   LogCompress,                           // 是否压缩日志
    }
    if LogStdout {
        // 日志同时输出到控制台和日志文件中
        return zapcore.NewMultiWriteSyncer(zapcore.AddSync(lumberJackLogger), zapcore.AddSync(os.Stdout)), nil
    } else {
        // 日志只输出到控制台
        return zapcore.AddSync(lumberJackLogger), nil
    }
}

// IsExist 判断文件或者目录是否存在
func IsExist(path string) bool {
    _, err := os.Stat(path)
    return err == nil || os.IsExist(err)
}
