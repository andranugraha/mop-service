package logger

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/metadata"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

var (
	infoLog  *logrus.Entry
	errorLog *logrus.Entry
	dataLog  *logrus.Entry
)

func init() {
	logDir := "log"
	infoLogger := logrus.New()
	infoFile, err := os.OpenFile(filepath.Join(logDir, "info.log"), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		logrus.Fatalf("Failed to log to info file, using default stderr: %v", err)
	}
	infoLogger.SetOutput(io.MultiWriter(os.Stdout, infoFile))
	infoLogger.SetFormatter(&logrus.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
	})
	infoLogger.SetLevel(logrus.InfoLevel)
	infoLog = infoLogger.WithField("level", "info")

	errorLogger := logrus.New()
	errorFile, err := os.OpenFile(filepath.Join(logDir, "error.log"), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		logrus.Fatalf("Failed to log to error file, using default stderr: %v", err)
	}
	errorLogger.SetOutput(io.MultiWriter(os.Stderr, errorFile))
	errorLogger.SetFormatter(&logrus.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
	})
	errorLogger.SetLevel(logrus.WarnLevel)
	errorLog = errorLogger.WithField("level", "error")

	dataLogger := logrus.New()
	dataFile, err := os.OpenFile(filepath.Join(logDir, "data.log"), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		logrus.Fatalf("Failed to log to data file, using default stderr: %v", err)
	}
	dataLogger.SetOutput(io.MultiWriter(os.Stdout, dataFile))
	dataLogger.SetFormatter(&logrus.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
	})
	dataLogger.SetLevel(logrus.InfoLevel)
	dataLog = dataLogger.WithField("level", "data")
}

func Data(ctx context.Context, logGroup string, req interface{}, res interface{}) {
	md, _ := metadata.FromIncomingContext(ctx)
	logEntry := dataLog.WithFields(logrus.Fields{
		"group": logGroup,
		"request_id": func() string {
			if len(md.Get("request_id")) > 0 {
				return md.Get("request_id")[0]
			}
			return "0"
		}(),
		"user_id": func() string {
			if len(md.Get("user_id")) > 0 {
				return md.Get("user_id")[0]
			}
			return "0"
		}(),
	})

	reqJSON, _ := json.Marshal(req)
	resJSON, _ := json.Marshal(res)
	formattedMsg := fmt.Sprintf("response_data_record | request: %v, response: %v", string(reqJSON), string(resJSON))

	logEntry.Info(formattedMsg)
}

func Info(ctx context.Context, logGroup, msg string, args ...interface{}) {
	formattedMsg := msg
	if len(args) > 0 {
		formattedMsg = fmt.Sprintf(msg, args...)
	}

	md, _ := metadata.FromIncomingContext(ctx)
	logEntry := infoLog.WithFields(logrus.Fields{
		"group": logGroup,
		"request_id": func() string {
			if len(md.Get("request_id")) > 0 {
				return md.Get("request_id")[0]
			}
			return "0"
		}(),
		"user_id": func() string {
			if len(md.Get("user_id")) > 0 {
				return md.Get("user_id")[0]
			}
			return "0"
		}(),
	})
	logEntry.Info(formattedMsg)
}

func Error(ctx context.Context, logGroup, msg string, args ...interface{}) {
	formattedMsg := msg
	if len(args) > 0 {
		formattedMsg = fmt.Sprintf(msg, args...)
	}

	md, _ := metadata.FromIncomingContext(ctx)
	logEntry := errorLog.WithFields(logrus.Fields{
		"group": logGroup,
		"request_id": func() string {
			if len(md.Get("request_id")) > 0 {
				return md.Get("request_id")[0]
			}
			return "0"
		}(),
		"user_id": func() string {
			if len(md.Get("user_id")) > 0 {
				return md.Get("user_id")[0]
			}
			return "0"
		}(),
	})
	logEntry.Error(formattedMsg)
}

func Warn(ctx context.Context, logGroup, msg string, args ...interface{}) {
	formattedMsg := msg
	if len(args) > 0 {
		formattedMsg = fmt.Sprintf(msg, args...)
	}

	md, _ := metadata.FromIncomingContext(ctx)
	logEntry := errorLog.WithFields(logrus.Fields{
		"group": logGroup,
		"request_id": func() string {
			if len(md.Get("request_id")) > 0 {
				return md.Get("request_id")[0]
			}
			return "0"
		}(),
		"user_id": func() string {
			if len(md.Get("user_id")) > 0 {
				return md.Get("user_id")[0]
			}
			return "0"
		}(),
	})
	logEntry.Warn(formattedMsg)
}

func Debug(ctx context.Context, logGroup, msg string, args ...interface{}) {
	formattedMsg := msg
	if len(args) > 0 {
		formattedMsg = fmt.Sprintf(msg, args...)
	}

	md, _ := metadata.FromIncomingContext(ctx)
	logEntry := infoLog.WithFields(logrus.Fields{
		"group": logGroup,
		"request_id": func() string {
			if len(md.Get("request_id")) > 0 {
				return md.Get("request_id")[0]
			}
			return "0"
		}(),
		"user_id": func() string {
			if len(md.Get("user_id")) > 0 {
				return md.Get("user_id")[0]
			}
			return "0"
		}(),
	})
	logEntry.Debug(formattedMsg)
}

func Fatal(ctx context.Context, logGroup, msg string, args ...interface{}) {
	formattedMsg := msg
	if len(args) > 0 {
		formattedMsg = fmt.Sprintf(msg, args...)
	}

	md, _ := metadata.FromIncomingContext(ctx)
	logEntry := errorLog.WithFields(logrus.Fields{
		"group": logGroup,
		"request_id": func() string {
			if len(md.Get("request_id")) > 0 {
				return md.Get("request_id")[0]
			}
			return "0"
		}(),
		"user_id": func() string {
			if len(md.Get("user_id")) > 0 {
				return md.Get("user_id")[0]
			}
			return "0"
		}(),
	})
	logEntry.Fatal(formattedMsg)
}

func Panic(ctx context.Context, msg string, args ...interface{}) {
	formattedMsg := msg
	if len(args) > 0 {
		formattedMsg = fmt.Sprintf(msg, args...)
	}

	var logGroup string
	stack := make([]byte, 4096)
	length := runtime.Stack(stack, false)

	stackStr := string(stack[:length])
	lines := strings.Split(stackStr, "\n")
	if len(lines) > 3 {
		// Extracting the file name and line number
		fileLineParts := strings.Fields(lines[3])
		if len(fileLineParts) >= 3 {
			fileName := fileLineParts[1]
			lineNumber := fileLineParts[2]
			logGroup = fmt.Sprintf("%s:%s", fileName, lineNumber)
		} else {
			errorLog.Info(lines)
			logGroup = lines[3]
		}
	} else {
		logGroup = lines[0]
	}

	md, _ := metadata.FromIncomingContext(ctx)
	logEntry := errorLog.WithFields(logrus.Fields{
		"group": logGroup,
		"request_id": func() string {
			if len(md.Get("request_id")) > 0 {
				return md.Get("request_id")[0]
			}
			return "0"
		}(),
		"user_id": func() string {
			if len(md.Get("user_id")) > 0 {
				return md.Get("user_id")[0]
			}
			return "0"
		}(),
	})
	logEntry.Error(formattedMsg)
}
