package main

import (
	"os"

	"go.uber.org/zap"
)

func main() {
	log_file = os.OpenFile("Log.json",os.O_APPEND | os.O_CREATE | os.O_WRONLY,0644)
	sugar_logger := logInit(false, log_file)

	//Pass sugar to other objects
	sugar_logger.Info("Starting Creator.")
}

func logInit(d bool, f *os.File) *zap.SugaredLogger {

    pe := zap.NewProductionEncoderConfig()

    fileEncoder := zapcore.NewJSONEncoder(pe)

    pe.EncodeTime = zapcore.ISO8601TimeEncoder # The encoder can be customized for each output
    consoleEncoder := zapcore.NewConsoleEncoder(pe)

    level := zap.ErrorLevel
    if d {
        level = zap.DebugLevel
    }

    core := zapcore.NewTee(
        zapcore.NewCore(fileEncoder, zapcore.AddSync(f), level),
        zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), level),
    )

    l := zap.New(core) // Creating the logger

    return l.Sugar()
}
