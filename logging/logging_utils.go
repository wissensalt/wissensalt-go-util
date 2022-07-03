package logging

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/natefinch/lumberjack"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

var AppLogger *log.Logger

func Init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to load .env file : ", err)
	} else {
		loggerFileName := os.Getenv("logger.filename")
		e, err := os.OpenFile(loggerFileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)

		if err != nil {
			fmt.Printf("error opening file: %v", err)
			os.Exit(1)
		}

		AppLogger = log.New(e, "", log.Ldate|log.Ltime)
		maxFileSize, _ := strconv.Atoi(os.Getenv("logger.maxsize"))
		maxBackups, _ := strconv.Atoi(os.Getenv("logger.maxbackups"))
		maxAge, _ := strconv.Atoi(os.Getenv("logger.maxage"))
		compress, _ := strconv.ParseBool(os.Getenv("logger.compress"))
		AppLogger.SetOutput(&lumberjack.Logger{
			Filename:   loggerFileName,
			MaxSize:    maxFileSize, // megabytes
			MaxBackups: maxBackups,
			MaxAge:     maxAge,   //days
			Compress:   compress, // disabled by default
		})
	}
}

func CheckFatal(err error) {
	if err != nil {
		log.Fatal("Fatal Exception occurred : ", err)
	}
}


func FindMissingNumbers(p_Array []int) []int {
	var result []int
	for a := 1; a < len(p_Array); a++ {
		for b := 1 + p_Array[a-1]; b < p_Array[a]; b++ {
			result = append(result, b)
		}
	}

	return result
}

func FindMissingDates(p_Array []time.Time) []time.Time {
	var result []time.Time

	return result
}


var src = rand.NewSource(time.Now().UnixNano())

const letterBytes = "1234567890ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

func RandStringBytesMaskImprSrc(n int) string {
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}
