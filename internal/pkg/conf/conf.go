package conf

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"strconv"
)

var Conf Config

type Config struct {
	Base                        Base          `yaml:"base"`
	Http                        Http          `yaml:"http"`
	GRPC                        GRPC          `yaml:"grpc"`
	MySQL                       MySQL         `yaml:"mysql"`
	JWT                         JWT           `yaml:"jwt"`
	MongoDB                     MongoDB       `yaml:"mongodb"`
	Prometheus                  Prometheus    `yaml:"prometheus"`
	Kafka                       Kafka         `yaml:"kafka"`
	ES                          ES            `yaml:"es"`
	Clients                     Clients       `yaml:"clients"`
	Proof                       Proof         `yaml:"proof"`
	Redis                       Redis         `yaml:"redis"`
	RedisReport                 RedisReport   `yaml:"redisReport"`
	SMTP                        SMTP          `yaml:"smtp"`
	Sms                         Sms           `yaml:"sms"`
	InviteData                  inviteData    `yaml:"inviteData"`
	Log                         Log           `yaml:"log"`
	Pay                         Pay           `yaml:"pay"`
	GeeTest                     GeeTest       `yaml:"geeTest"`
	WechatLogin                 WechatLogin   `yaml:"wechatLogin"`
	CanUsePartitionTotalNum     int           `yaml:"canUsePartitionTotalNum"`
	OneMachineCanConcurrenceNum int           `yaml:"oneMachineCanConcurrenceNum"`
	MachineConfig               MachineConfig `yaml:"machineConfig"`
}

type MachineConfig struct {
	MachineAliveTime      int `yaml:"MachineAliveTime"`
	InitPartitionTotalNum int `yaml:"InitPartitionTotalNum"`
	CpuTopLimit           int `yaml:"CpuTopLimit"`
	MemoryTopLimit        int `yaml:"MemoryTopLimit"`
	DiskTopLimit          int `yaml:"DiskTopLimit"`
}

type Log struct {
	InfoPath string `yaml:"InfoPath"`
	ErrPath  string `yaml:"ErrPath"`
}

type Base struct {
	IsDebug        bool   `mapstructure:"is_debug"`
	Domain         string `mapstructure:"domain"`
	ProjectEnv     string `mapstructure:"project_env"`
	MaxConcurrency int64  `mapstructure:"max_concurrency"`
}

type Http struct {
	Port int `yaml:"port"`
}

type GRPC struct {
	Port int `yaml:"port"`
}

type MySQL struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	DBName   string `yaml:"dbname"`
	Charset  string `yaml:"charset"`
}

type JWT struct {
	Issuer string `yaml:"issuer"`
	Secret string `yaml:"secret"`
}

type MongoDB struct {
	DSN      string `yaml:"dsn"`
	Database string `yaml:"database"`
	PoolSize uint64 `mapstructure:"pool_size"`
}

type Prometheus struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type Kafka struct {
	Host  string `yaml:"host"`
	Topic string `yaml:"topic"`
}

type ES struct {
	Host     string `yaml:"host"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Index    string `yaml:"index"`
	Size     int    `yaml:"size"`
}

type Clients struct {
	Runner Runner
}

type Runner struct {
	RunAPI    string `mapstructure:"run_api"`
	RunScene  string `mapstructure:"run_scene"`
	StopScene string `mapstructure:"stop_scene"`
	RunPlan   string `mapstructure:"run_plan"`
	StopPlan  string `mapstructure:"stop_plan"`
}

type Proof struct {
	InfoLog string `mapstructure:"info_log"`
	ErrLog  string `mapstructure:"err_log"`
}

type Redis struct {
	ClusterAddress string `yaml:"clusterAddress"`
	Password       string `yaml:"password"`
	DB             int    `yaml:"db"`
}
type RedisReport struct {
	ClusterAddress string `yaml:"clusterAddress"`
	Password       string `yaml:"password"`
	DB             int    `yaml:"db"`
}

type SMTP struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Email    string `mapstructure:"email"`
	Password string `mapstructure:"password"`
}

type Sms struct {
	ID     string `mapstructure:"id"`
	Secret string `mapstructure:"secret"`
}

type inviteData struct {
	AesSecretKey string `yaml:"AesSecretKey"`
}

type Pay struct {
	Business                string `yaml:"Business"`
	Theme                   string `yaml:"Theme"`
	PayNotifyApi            string `yaml:"PayNotifyApi"`
	CreateOrderApi          string `yaml:"CreateOrderApi"`
	SelectOrderPayResultApi string `yaml:"SelectOrderPayResultApi"`
	CloseOrderApi           string `yaml:"CloseOrderApi"`
}

type GeeTest struct {
	CaptchaID  string `yaml:"CaptchaID"`
	CaptchaKey string `yaml:"CaptchaKey"`
	ApiServer  string `yaml:"ApiServer"`
}

type WechatLogin struct {
	WechatLoginQrCodeApi string `yaml:"WechatLoginQrCodeApi"`
	WechatScanResultApi  string `yaml:"WechatScanResultApi"`
}

func MustInitConf(configFile string) {
	viper.SetConfigFile(configFile)
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	if err := viper.Unmarshal(&Conf); err != nil {
		panic(fmt.Errorf("unmarshal error config file: %w", err))
	}

	fmt.Println("config initialized")
}

func MustInitConfByEnv() {
	initBase()
	initHttp()
	initMysql()
	initJWT()
	initMongoDB()
	initClients()
	initProof()
	initRedis()
	initRedisReport()
	initSMTP()
	initSms()
	initInviteData()
	initLog()
	initCanUsePartitionTotalNum()
	initOneMachineCanConcurrenceNum()
	initMachineConfig()
}

func initBase() {
	if os.Getenv("RG_IS_DEBUG") == "true" {
		Conf.Base.IsDebug = true
	} else {
		Conf.Base.IsDebug = false
	}

	Conf.Base.Domain = os.Getenv("RG_DOMAIN")

	maxConcurrency, err := strconv.ParseInt(os.Getenv("RG_MAX_CONCURRENCY"), 10, 64)
	if err != nil {
		Conf.Base.MaxConcurrency = 100000
	} else {
		Conf.Base.MaxConcurrency = maxConcurrency
	}

}
func initHttp() {
	httpPort, err := strconv.Atoi(os.Getenv("RG_MANAGEMENT_HTTP_PORT"))
	if err != nil {
		Conf.Http.Port = 30000
	} else {
		Conf.Http.Port = httpPort
	}
}
func initMysql() {
	Conf.MySQL.Host = os.Getenv("RG_MYSQL_HOST")
	if Conf.MySQL.Host == "" {
		Conf.MySQL.Host = "127.0.0.0"
	}

	Conf.MySQL.Username = os.Getenv("RG_MYSQL_USERNAME")
	if Conf.MySQL.Username == "" {
		Conf.MySQL.Username = "root"
	}
	Conf.MySQL.Password = os.Getenv("RG_MYSQL_PASSWORD")
	Conf.MySQL.DBName = os.Getenv("RG_MYSQL_DBNAME")
	if Conf.MySQL.DBName == "" {
		Conf.MySQL.DBName = "runnergo"
	}

	Conf.MySQL.Charset = os.Getenv("RG_MYSQL_CHARSET")
	if Conf.MySQL.Charset == "" {
		Conf.MySQL.Charset = "utf8mb4"
	}

	port, err := strconv.Atoi(os.Getenv("RG_MYSQL_PORT"))
	if err != nil {
		Conf.MySQL.Port = 3306
	} else {
		Conf.MySQL.Port = port
	}
}
func initJWT() {
	Conf.JWT.Issuer = os.Getenv("RG_JWT_ISSUER")
	if Conf.JWT.Issuer == "" {
		Conf.JWT.Issuer = "RunnerGo"
	}
	Conf.JWT.Secret = os.Getenv("RG_JWT_SECRET")
	if Conf.JWT.Secret == "" {
		Conf.JWT.Secret = "RunnerGo#docker"
	}
}
func initMongoDB() {
	mgPassword := os.Getenv("RG_MONGO_PASSWORD")
	Conf.MongoDB.DSN = os.Getenv("RG_MONGO_DSN")
	if Conf.MongoDB.DSN == "" {
		Conf.MongoDB.DSN = fmt.Sprintf("mongodb://runnergo_open:%s@127.0.0.1:27017/runnergo_open", mgPassword)
	}

	Conf.MongoDB.Database = os.Getenv("RG_MONGO_DATABASE")
	if Conf.MongoDB.Database == "" {
		Conf.MongoDB.Database = "runnergo_open"
	}

	Conf.MongoDB.PoolSize = 20
}
func initClients() {
	Conf.Clients.Runner.RunAPI = os.Getenv("RG_CLIENTS_ENGINE_RUN_API")
	if Conf.Clients.Runner.RunAPI == "" {
		Conf.Clients.Runner.RunAPI = "https://127.0.0.0:30000/runner/run_api"
	}
	Conf.Clients.Runner.RunScene = os.Getenv("RG_CLIENTS_ENGINE_RUN_SCENE")
	if Conf.Clients.Runner.RunScene == "" {
		Conf.Clients.Runner.RunScene = "https://127.0.0.0:30000/runner/run_scene"
	}
	Conf.Clients.Runner.StopScene = os.Getenv("RG_CLIENTS_ENGINE_STOP_SCENE")
	if Conf.Clients.Runner.StopScene == "" {
		Conf.Clients.Runner.StopScene = "https://127.0.0.0:30000/runner/stop_scene"
	}
	Conf.Clients.Runner.RunPlan = os.Getenv("RG_CLIENTS_ENGINE_RUN_PLAN")
	if Conf.Clients.Runner.RunPlan == "" {
		Conf.Clients.Runner.RunPlan = "https://127.0.0.0:30000/runner/run_plan"
	}
	Conf.Clients.Runner.StopPlan = os.Getenv("RG_CLIENTS_ENGINE_STOP_PLAN")
	if Conf.Clients.Runner.StopPlan == "" {
		Conf.Clients.Runner.StopPlan = "https://127.0.0.0:30000/runner/stop"
	}
}
func initProof() {
	Conf.Proof.InfoLog = os.Getenv("RG_PROOF_INFO_LOG")
	if Conf.Proof.InfoLog == "" {
		Conf.Proof.InfoLog = "/data/logs/RunnerGo/RunnerGo_management-info.log"
	}
	Conf.Proof.ErrLog = os.Getenv("RG_PROOF_ERR_LOG")
	if Conf.Proof.ErrLog == "" {
		Conf.Proof.ErrLog = "/data/logs/RunnerGo/RunnerGo_management-err.log"
	}
}
func initRedis() {
	Conf.Redis.ClusterAddress = os.Getenv("RG_REDIS_ADDRESS")
	if Conf.Redis.ClusterAddress == "" {
		Conf.Redis.ClusterAddress = "127.0.0.0:6379"
	}
	Conf.Redis.Password = os.Getenv("RG_REDIS_PASSWORD")

	redisDB, err := strconv.Atoi(os.Getenv("RG_REDIS_DB"))
	if err != nil {
		Conf.Redis.DB = 0
	} else {
		Conf.Redis.DB = redisDB
	}
}
func initRedisReport() {
	Conf.RedisReport.ClusterAddress = os.Getenv("RG_REDIS_ADDRESS")
	if Conf.RedisReport.ClusterAddress == "" {
		Conf.RedisReport.ClusterAddress = "127.0.0.0:6379"
	}
	Conf.RedisReport.Password = os.Getenv("RG_REDIS_PASSWORD")

	redisDB, err := strconv.Atoi(os.Getenv("RG_REDIS_DB"))
	if err != nil {
		Conf.RedisReport.DB = 0
	} else {
		Conf.RedisReport.DB = redisDB
	}
}
func initSMTP() {
	Conf.SMTP.Host = os.Getenv("RG_SMTP_HOST")
	port, err := strconv.Atoi(os.Getenv("RG_SMTP_PORT"))
	if err != nil {
		Conf.SMTP.Port = 465
	} else {
		Conf.SMTP.Port = port
	}
	Conf.SMTP.Email = os.Getenv("RG_SMTP_EMAIL")
	Conf.SMTP.Password = os.Getenv("RG_SMTP_PASSWORD")
}
func initSms() {
	Conf.Sms.ID = os.Getenv("RG_SMS_ID")
	Conf.Sms.Secret = os.Getenv("RG_SMS_SECRET")
}
func initInviteData() {
	Conf.InviteData.AesSecretKey = os.Getenv("RG_INVITE_DATA_AES_SECRET_KEY")
	if Conf.InviteData.AesSecretKey == "" {
		Conf.InviteData.AesSecretKey = "RunnerGo"
	}
}
func initLog() {
	Conf.Log.InfoPath = os.Getenv("RG_LOG_INFO_PATH")
	if Conf.Log.InfoPath == "" {
		Conf.Log.InfoPath = "/data/logs/RunnerGo/RunnerGo_management-info.log"
	}
	Conf.Log.ErrPath = os.Getenv("RG_LOG_ERR_PATH")
	if Conf.Log.ErrPath == "" {
		Conf.Log.ErrPath = "/data/logs/RunnerGo/RunnerGo_management-err.log"
	}
}

func initCanUsePartitionTotalNum() {
	canUsePartitionTotalNum, err := strconv.Atoi(os.Getenv("RG_CAN_USE_PARTITION_TOTAL_NUM"))
	if err != nil {
		Conf.CanUsePartitionTotalNum = 2
	} else {
		if canUsePartitionTotalNum == 0 {
			Conf.CanUsePartitionTotalNum = 2
		}
		Conf.CanUsePartitionTotalNum = canUsePartitionTotalNum
	}
}

func initOneMachineCanConcurrenceNum() {
	oneMachineCanConcurrenceNum, err := strconv.Atoi(os.Getenv("RG_ONE_MACHINE_CAN_CONCURRENCE_NUM"))
	if err != nil {
		Conf.OneMachineCanConcurrenceNum = 5000
	} else {
		if oneMachineCanConcurrenceNum == 0 {
			Conf.OneMachineCanConcurrenceNum = 5000
		}
		Conf.OneMachineCanConcurrenceNum = oneMachineCanConcurrenceNum
	}
}

func initMachineConfig() {
	machineAliveTime, err := strconv.Atoi(os.Getenv("RG_MACHINE_ALIVE_TIME"))
	if err != nil {
		Conf.MachineConfig.MachineAliveTime = 10
	} else {
		Conf.MachineConfig.MachineAliveTime = machineAliveTime
	}

	initPartitionTotalNum, err := strconv.Atoi(os.Getenv("RG_INIT_PARTITION_TOTAL_NUM"))
	if err != nil {
		Conf.MachineConfig.InitPartitionTotalNum = 2
	} else {
		Conf.MachineConfig.InitPartitionTotalNum = initPartitionTotalNum
	}

	cpuTopLimit, err := strconv.Atoi(os.Getenv("RG_CPU_TOP_LIMIT"))
	if err != nil {
		Conf.MachineConfig.CpuTopLimit = 65
	} else {
		Conf.MachineConfig.CpuTopLimit = cpuTopLimit
	}

	memoryTopLimit, err := strconv.Atoi(os.Getenv("RG_MEMORY_TOP_LIMIT"))
	if err != nil {
		Conf.MachineConfig.MemoryTopLimit = 65
	} else {
		Conf.MachineConfig.MemoryTopLimit = memoryTopLimit
	}

	diskTopLimit, err := strconv.Atoi(os.Getenv("RG_DISK_TOP_LIMIT"))
	if err != nil {
		Conf.MachineConfig.DiskTopLimit = 55
	} else {
		Conf.MachineConfig.DiskTopLimit = diskTopLimit
	}
}
