package log

import (
	"os"
	"path/filepath"
	"time"

	"github.com/aliyun/aliyun_assist_client/common/pathutil"
	rotatelogs "github.com/aliyun/aliyun_assist_client/thirdparty/file-rotatelogs"
	log "github.com/aliyun/aliyun_assist_client/thirdparty/sirupsen/logrus"
	"github.com/pkg/errors"
)

var Log *log.Logger
var defaultLevel log.Level = log.InfoLevel

type Fields = log.Fields

// DefaultCommonFields returns preset fields for each log message. All default
// common fields MUST be prefixed by an underscore to avoid conflict with fields
// added later.
func DefaultCommonFields() log.Fields {
	return log.Fields{
		"_pid": os.Getpid(),
	}
}

func InitLog(filename string, logpath string, ignoreRotationError bool) {
	if logpath != "" {
		pathutil.SetLogPath(filepath.Join(logpath, "log"))
	}
	logdir, _ := pathutil.GetLogPath()

	writer, err := rotatelogs.New(
		filepath.Join(logdir, filename+".%Y%m%d"),
		rotatelogs.WithMaxAge(time.Duration(24*30)*time.Hour),    //最长保留30天
		rotatelogs.WithRotationTime(time.Duration(24)*time.Hour), //每天进行一次日志切割
		rotatelogs.WithLinkName(filepath.Join(logdir, filename)),         // 为日志文件创建一个名字不变的链接
		rotatelogs.IgnoreRotationError(ignoreRotationError),
	)
	if err != nil {
		log.Errorf("config local file system logger error. %+v", errors.WithStack(err))
	}

	Log = log.New()
	Log.SetFormatter(&CustomLogrusTextFormatter{
		CommonFields: DefaultCommonFields(),
	})
	Log.SetOutput(writer)
	Log.SetLevel(defaultLevel)
}

func GetLogger() *log.Logger {
	if Log == nil {
		InitLog("aliyun_assist_test", "", false)
	}
	return Log
}
