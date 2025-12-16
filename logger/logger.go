package logger

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Config struct {
	Level  string `mapstructure:"level"`
	Format string `mapstructure:"format"`
}

var Log *logrus.Logger

func Init(config Config, component string) error {
	Log = logrus.New()

	level, err := logrus.ParseLevel(config.Level)
	if err != nil {
		return err
	}
	Log.SetLevel(level)

	if config.Format == "json" {
		Log.SetFormatter(&logrus.JSONFormatter{
			TimestampFormat: "2006-01-02T15:04:05.000Z07:00",
		})
	} else {
		Log.SetFormatter(&logrus.TextFormatter{
			FullTimestamp:   true,
			TimestampFormat: "2006-01-02 15:04:05",
			DisableColors:   true,
		})
	}

	Log.SetOutput(os.Stdout)

	Log = Log.WithFields(logrus.Fields{
		"component": component,
		"version":   GetVersion(),
	}).Logger

	return nil
}

func FromGinContext(c *gin.Context) *logrus.Entry {
	fields := logrus.Fields{}

	if requestID, exists := c.Get("request_id"); exists {
		if id, ok := requestID.(string); ok && id != "" {
			fields["request_id"] = id
		}
	}
	if userID, exists := c.Get("user_id"); exists {
		if id, ok := userID.(string); ok && id != "" {
			fields["user_id"] = id
		}
	}
	if teamID, exists := c.Get("team_id"); exists {
		if id, ok := teamID.(string); ok && id != "" {
			fields["team_id"] = id
		}
	}

	fields["method"] = c.Request.Method
	fields["path"] = c.Request.URL.Path

	return Log.WithFields(fields)
}

func GetLogger(c *gin.Context) *logrus.Entry {
	if logger, exists := c.Get("logger"); exists {
		if log, ok := logger.(*logrus.Entry); ok {
			return log
		}
	}
	return FromGinContext(c)
}

func L(c *gin.Context) *logrus.Entry {
	return GetLogger(c)
}

func GetVersion() string {
	return "1.0.0"
}
