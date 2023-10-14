package config

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/Implex-ltd/hcsolver/internal/hcaptcha/fingerprint"
	"github.com/Implex-ltd/hcsolver/internal/recognizer"
	"github.com/mattn/go-colorable"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	Logger *zap.Logger
	Config = Cfg{}
)

type Cfg struct {
	API struct {
		Port int `toml:"port"`
	} `toml:"api"`
	Ratelimit struct {
		APIMax        int `toml:"api_max"`
		APIExpiration int `toml:"api_expiration"`
	} `toml:"ratelimit"`
	Database struct {
		Username string `toml:"username"`
		Password string `toml:"password"`
		IP       string `toml:"ip"`
		Port     int    `toml:"port"`
	} `toml:"database"`
}

func CreateLogFile() *os.File {
	logFileName := fmt.Sprintf("../../assets/logs/%s", time.Now().Format("2006-01-02_15-04-05")+".json")

	err := os.MkdirAll("../../assets/logs/", os.ModePerm)
	if err != nil {
		panic(err)
	}

	file, err := os.OpenFile(logFileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	return file
}

func LoadSettings() {
	rand.New(rand.NewSource(time.Now().UnixNano()))

	if _, err := toml.DecodeFile("../../scripts/config.toml", &Config); err != nil {
		panic(err)
	}

	encoder := zap.NewDevelopmentEncoderConfig()
	encoder.EncodeLevel = zapcore.CapitalColorLevelEncoder

	fileEncoder := zap.NewProductionEncoderConfig()
	fileEncoder.EncodeTime = zapcore.ISO8601TimeEncoder

	core := zapcore.NewTee(
		zapcore.NewCore(
			zapcore.NewConsoleEncoder(encoder),
			zapcore.AddSync(colorable.NewColorableStdout()),
			zapcore.DebugLevel,
		),
		zapcore.NewCore(
			zapcore.NewJSONEncoder(fileEncoder),
			zapcore.AddSync(CreateLogFile()),
			zapcore.DebugLevel,
		),
	)

	Logger = zap.New(core)
	count, err := recognizer.LoadHash("../../assets/hash.csv")
	if err != nil {
		panic(err)
	}

	Logger.Info("Loaded hash csv",
		zap.Int("count", count),
	)

	recognizer.Hashlist.Range(func(key, value interface{}) bool {
		prompt := key.(string)
		hashes := value.([]string)

		Logger.Info("Loaded hash",
			zap.String("prompt", prompt),
			zap.Int("count", len(hashes)),
		)

		return true
	})

	selectCount, err := recognizer.LoadHashSelect("../../assets/area_hash.csv")
	if err != nil {
		panic(err)
	}

	Logger.Info("Loaded select hash csv",
		zap.Int("count", selectCount),
	)

	recognizer.Selectlist.Range(func(key, value interface{}) bool {
		prompt := key.(string)
		hashDataList := value.([]recognizer.HashData)

		Logger.Info("Loaded hash data",
			zap.String("prompt", prompt),
			zap.Int("count", len(hashDataList)),
		)

		return true
	})

	recognizer.LoadAnswer("../../assets/questions.txt")

	// recognizer http client
	t := http.DefaultTransport.(*http.Transport).Clone()
	t.MaxIdleConns = 500
	t.MaxConnsPerHost = 500
	t.MaxIdleConnsPerHost = 500

	recognizer.Client = &http.Client{
		Timeout:   15 * time.Second,
		Transport: t,
	}

	fingerprint.CollectFpArray.RandomiseIndex()
}
