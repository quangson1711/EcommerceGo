package config

import (
	"io"
	"log"
	"os"
	"time"

	"gopkg.in/natefinch/lumberjack.v2"
)

func SetupLogger() {
	// Lấy ngày hiện tại để sử dụng trong tên file backup
	now := time.Now()
	date := now.Format("2006-01-02") // Định dạng YYYY-MM-DD

	// Tạo một logger mới với cấu hình quay vòng log
	logFile := &lumberjack.Logger{
		Filename:   "logs/" + "app-" + date + ".log",
		MaxSize:    10, // MB
		MaxBackups: 5,
		MaxAge:     30,   // days
		LocalTime:  true, // Sử dụng múi giờ cục bộ
	}

	// Tạo một writer đa kênh kết hợp stdout và file log
	multiWriter := io.MultiWriter(os.Stdout, logFile)

	// Đặt output của log package là writer đa kênh này
	log.SetOutput(multiWriter)
}
