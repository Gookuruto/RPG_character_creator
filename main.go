package main

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	log_file, err := os.OpenFile("Log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic("Error when opening log file.")
	}
	defer log_file.Close()
	sugar_logger := logInit(false, log_file)

	//Pass sugar to other objects
	sugar_logger.Info("Starting Creator.")
}

func logInit(d bool, f *os.File) *zap.SugaredLogger {

	pe := zap.NewProductionEncoderConfig()

	fileEncoder := zapcore.NewJSONEncoder(pe)

	pe.EncodeTime = zapcore.ISO8601TimeEncoder // The encoder can be customized for each output
	//consoleEncoder := zapcore.NewConsoleEncoder(pe)

	level := zap.InfoLevel
	if d {
		level = zap.DebugLevel
	}

	core := zapcore.NewTee(
		zapcore.NewCore(fileEncoder, zapcore.AddSync(f), level),
		//zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), level), //Uncomment if you want logs in console
	)

	l := zap.New(core) // Creating the logger

	return l.Sugar()
}
