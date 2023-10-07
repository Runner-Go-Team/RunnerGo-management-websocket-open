package autoPlan

import (
	"context"
	"github.com/Runner-Go-Team/RunnerGo-management-websocket-open/internal/pkg/dal/mao"
	"github.com/Runner-Go-Team/RunnerGo-management-websocket-open/internal/pkg/dal/model"
	"github.com/Runner-Go-Team/RunnerGo-management-websocket-open/internal/pkg/dal/run_plan"
	"github.com/shirou/gopsutil/load"
	"sync"
)

type Baton struct {
	Ctx      context.Context
	PlanID   string
	TeamID   string
	UserID   string
	SceneIDs []string
	RunType  int

	reportID    string
	plan        *model.AutoPlan
	scenes      []*model.Target
	testCase    []*model.Target
	testCaseIDs []string
	//task            map[int64]*mao.Task // sceneID 对应任务配置
	ConfigTask      ConfigTask // 任务配置
	globalVariables []*model.Variable
	sceneFlows      map[string]*mao.Flow
	sceneCaseFlows  map[string]*mao.SceneCaseFlow
	sceneVariables  map[string][]*model.Variable
	importVariables map[string][]*model.VariableImport
	balance         *DispatchMachineBalance
	stress          []*run_plan.Stress
	MachineList     []*HeartBeat
	RealRunParam    RealRunParam
}

type RealRunParam struct {
	PlanId        string        `json:"plan_id" bson:"plan_id"`             // 计划id
	PlanName      string        `json:"plan_name" bson:"plan_name"`         // 计划名称
	ReportId      string        `json:"report_id" bson:"report_id"`         // 报告名称
	TeamId        string        `json:"team_id" bson:"team_id"`             // 团队id
	ReportName    string        `json:"report_name" bson:"report_name"`     // 报告名称
	MachineNum    int64         `json:"machine_num" bson:"machine_num"`     // 使用的机器数量
	ConfigTask    ConfigTask    `json:"config_task" bson:"config_task"`     // 任务配置
	Variable      []PlanKv      `json:"variable" bson:"variable"`           // 全局变量
	Scenes        []Scene       `json:"scenes" bson:"scenes"`               // 场景
	Configuration Configuration `json:"configuration" bson:"configuration"` // 场景配置
}

// ConfigTask 任务配置
type ConfigTask struct {
	TaskType     int64  `json:"task_type" bson:"task_type"`           // 任务类型：0. 普通任务； 1. 定时任务； 2. cicd任务
	TaskMode     int64  `json:"task_mode" bson:"task_mode"`           // 1. 按用例执行
	SceneRunMode int64  `json:"scene_run_mode" bson:"scene_run_mode"` // 2. 同时执行； 1. 顺序执行
	CaseRunMode  int64  `json:"case_run_mode" bson:"case_run_mode"`   // 2. 同时执行； 1. 顺序执行
	Remark       string `json:"remark" bson:"remark"`                 // 备注
}

type PlanKv struct {
	Var string `json:"Var"`
	Val string `json:"Val"`
}

type Scene struct {
	PlanId                  string           `json:"plan_id" bson:"plan_id"`
	SceneId                 string           `json:"scene_id" bson:"scene_id"`     // 场景Id
	IsChecked               int32            `json:"is_checked" bson:"is_checked"` // 是否启用
	ParentId                string           `json:"parentId" bson:"parent_id"`
	CaseId                  string           `json:"case_id" bson:"case_id"`
	Partition               int32            `json:"partition"`
	MachineNum              int64            `json:"machine_num" bson:"machine_num"` // 使用的机器数量
	ReportId                string           `json:"report_id" bson:"report_id"`
	TeamId                  string           `json:"team_id" bson:"team_id"`
	SceneName               string           `json:"scene_name" bson:"scene_name"` // 场景名称
	Version                 int64            `json:"version" bson:"version"`
	Debug                   string           `json:"debug" bson:"debug"`
	EnablePlanConfiguration bool             `json:"enable_plan_configuration" bson:"enable_plan_configuration"` // 是否启用计划的任务配置，默认为true，
	Nodes                   []*run_plan.Node `json:"nodes" bson:"nodes"`                                         // 事件列表
	ConfigTask              ConfigTask       `json:"config_task" bson:"config_task"`                             // 任务配置
	Configuration           Configuration    `json:"configuration" bson:"configuration"`                         // 场景配置
	Variable                []KV             `json:"variable" bson:"variable"`                                   // 场景配置
	Cases                   []Scene          `json:"cases" bson:"cases"`
}

type Configuration struct {
	ParameterizedFile ParameterizedFile `json:"parameterizedFile" bson:"parameterizedFile"`
	Variable          []KV              `json:"variable" bson:"variable"`
}

// ParameterizedFile 参数化文件
type ParameterizedFile struct {
	Paths         []FileList     `json:"paths"` // 文件地址
	RealPaths     []string       `json:"real_paths"`
	VariableNames *VariableNames `json:"variable_names"` // 存储变量及数据的map
}

type FileList struct {
	IsChecked int64  `json:"is_checked"` // 1 开， 2： 关
	Path      string `json:"path"`
}

type VariableNames struct {
	VarMapList map[string][]string `json:"var_map_list"`
	Index      int                 `json:"index"`
	Mu         sync.Mutex          `json:"mu"`
}

type KV struct {
	Key   string      `json:"key" bson:"key"`
	Value interface{} `json:"value" bson:"value"`
}

type Event struct {
	Id                string   `json:"id" bson:"id"`
	ReportId          int64    `json:"report_id" bson:"report_id"`
	TeamId            int64    `json:"team_id" bson:"team_id"`
	IsCheck           bool     `json:"is_check" bson:"is_check"`
	Type              string   `json:"type" bson:"type"` //   事件类型 "request" "controller"
	PreList           []string `json:"pre_list" bson:"pre_list"`
	NextList          []string `json:"next_list"   bson:"next_list"`
	Tag               bool     `json:"tag" bson:"tag"` // Tps模式下，该标签代表以该接口为准
	Debug             string   `json:"debug" bson:"debug"`
	Mode              int64    `json:"mode"`                 // 模式类型
	RequestThreshold  int64    `json:"request_threshold"`    // Rps（每秒请求数）阈值
	ResponseThreshold int64    `json:"response_threshold"`   // 响应时间阈值
	ErrorThreshold    float32  `json:"error_threshold"`      // 错误率阈值
	PercentAge        int64    `json:"percent_age"`          // 响应时间线
	Weight            int64    `json:"weight" bson:"weight"` // 权重，并发分配的比例
	Api               Api      `json:"api"`
	Var               string   `json:"var"`     // if控制器key，值某个变量
	Compare           string   `json:"compare"` // 逻辑运算符
	Val               string   `json:"val"`     // key对应的值
	Name              string   `json:"name"`    // 控制器名称
	WaitTime          int      `json:"wait_ms"` // 等待时长，ms
}

// Api 请求数据
type Api struct {
	TargetId   int64   `json:"target_id" bson:"target_id"`
	Name       string  `json:"name" bson:"name"`
	TeamId     int64   `json:"team_id" bson:"team_id"`
	TargetType string  `json:"target_type" bson:"target_type"` // api/webSocket/tcp/grpc
	Method     string  `json:"method" bson:"method"`           // 方法 GET/POST/PUT
	Request    Request `json:"request" bson:"request"`
	//Parameters    *sync.Map            `json:"parameters" bson:"parameters"`
	Assert        []*AssertionText     `json:"assert" bson:"assert"`         // 验证的方法(断言)
	Timeout       int64                `json:"timeout" bson:"timeout"`       // 请求超时时间
	Regex         []*RegularExpression `json:"regex" bson:"regex"`           // 正则表达式
	Debug         string               `json:"debug" bson:"debug"`           // 是否开启Debug模式
	Connection    int64                `json:"connection" bson:"connection"` // 0:websocket长连接
	Configuration *Configuration       `json:"configuration" bson:"configuration"`
	Variable      []*KV                `json:"variable" bson:"variable"` // 全局变量
	HttpApiSetup  HttpApiSetup         `json:"http_api_setup" bson:"http_api_setup"`
}

type HttpApiSetup struct {
	IsRedirects         int    `json:"is_redirects"`  // 是否跟随重定向 0: 是   1：否
	RedirectsNum        int    `json:"redirects_num"` // 重定向次数>= 1; 默认为3
	ReadTimeOut         int    `json:"read_time_out"` // 请求超时时间
	WriteTimeOut        int    `json:"write_time_out"`
	ClientName          string `json:"client_name"`
	KeepAlive           bool   `json:"keep_alive"`
	MaxIdleConnDuration int32  `json:"max_idle_conn_duration"`
	MaxConnPerHost      int32  `json:"max_conn_per_host"`
	UserAgent           bool   `json:"user_agent"`
	MaxConnWaitTimeout  int64  `json:"max_conn_wait_timeout"`
}

type Request struct {
	PreUrl    string     `json:"pre_url" bson:"pre_url"`
	URL       string     `json:"url" bson:"url"`
	Parameter []*VarForm `json:"parameter" bson:"parameter"`
	Header    *Header    `json:"header" bson:"header"` // Headers
	Query     *Query     `json:"query" bson:"query"`
	Body      *Body      `json:"body" bson:"body"`
	Auth      *Auth      `json:"auth" bson:"auth"`
	Cookie    *Cookie    `json:"cookie" bson:"cookie"`
}
type Header struct {
	Parameter []*VarForm `json:"parameter" bson:"parameter"`
}
type Query struct {
	Parameter []*VarForm `json:"parameter" bson:"parameter"`
}

type Cookie struct {
	Parameter []*VarForm
}

type Auth struct {
	Type          string    `json:"type" bson:"type"`
	KV            *KV       `json:"kv" bson:"kv"`
	Bearer        *Bearer   `json:"bearer" bson:"bearer"`
	Basic         *Basic    `json:"basic" bson:"basic"`
	Digest        *Digest   `json:"digest"`
	Hawk          *Hawk     `json:"hawk"`
	Awsv4         *AwsV4    `json:"awsv4"`
	Ntlm          *Ntlm     `json:"ntlm"`
	Edgegrid      *Edgegrid `json:"edgegrid"`
	Oauth1        *Oauth1   `json:"oauth1"`
	Bidirectional TLS       `json:"bidirectional"`
}

type TLS struct {
	CaCert     string `json:"ca_cert"`
	CaCertName string `json:"ca_cert_name"`
}

type Bearer struct {
	Key string `json:"key" bson:"key"`
}

type Basic struct {
	UserName string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
}

type Digest struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	Realm     string `json:"realm"`
	Nonce     string `json:"nonce"`
	Algorithm string `json:"algorithm"`
	Qop       string `json:"qop"`
	Nc        string `json:"nc"`
	Cnonce    string `json:"cnonce"`
	Opaque    string `json:"opaque"`
}

type Hawk struct {
	AuthID             string `json:"authId"`
	AuthKey            string `json:"authKey"`
	Algorithm          string `json:"algorithm"`
	User               string `json:"user"`
	Nonce              string `json:"nonce"`
	ExtraData          string `json:"extraData"`
	App                string `json:"app"`
	Delegation         string `json:"delegation"`
	Timestamp          string `json:"timestamp"`
	IncludePayloadHash int    `json:"includePayloadHash"`
}

type AwsV4 struct {
	AccessKey          string `json:"accessKey"`
	SecretKey          string `json:"secretKey"`
	Region             string `json:"region"`
	Service            string `json:"service"`
	SessionToken       string `json:"sessionToken"`
	AddAuthDataToQuery int    `json:"addAuthDataToQuery"`
}

type Ntlm struct {
	Username            string `json:"username"`
	Password            string `json:"password"`
	Domain              string `json:"domain"`
	Workstation         string `json:"workstation"`
	DisableRetryRequest int    `json:"disableRetryRequest"`
}

type Edgegrid struct {
	AccessToken   string `json:"accessToken"`
	ClientToken   string `json:"clientToken"`
	ClientSecret  string `json:"clientSecret"`
	Nonce         string `json:"nonce"`
	Timestamp     string `json:"timestamp"`
	BaseURi       string `json:"baseURi"`
	HeadersToSign string `json:"headersToSign"`
}

type Oauth1 struct {
	ConsumerKey          string `json:"consumerKey"`
	ConsumerSecret       string `json:"consumerSecret"`
	SignatureMethod      string `json:"signatureMethod"`
	AddEmptyParamsToSign int    `json:"addEmptyParamsToSign"`
	IncludeBodyHash      int    `json:"includeBodyHash"`
	AddParamsToHeader    int    `json:"addParamsToHeader"`
	Realm                string `json:"realm"`
	Version              string `json:"version"`
	Nonce                string `json:"nonce"`
	Timestamp            string `json:"timestamp"`
	Verifier             string `json:"verifier"`
	Callback             string `json:"callback"`
	TokenSecret          string `json:"tokenSecret"`
	Token                string `json:"token"`
}

type Body struct {
	Mode      string     `json:"mode" bson:"mode"`
	Raw       string     `json:"raw" bson:"raw"`
	Parameter []*VarForm `json:"parameter" bson:"parameter"`
}

type RegularExpression struct {
	IsChecked int         `json:"is_checked"` // 1 选中, -1未选
	Type      int         `json:"type"`       // 0 正则  1 json
	Var       string      `json:"var"`        // 变量
	Express   string      `json:"express"`    // 表达式
	Val       interface{} `json:"val"`        // 值
}

// AssertionText 文本断言 0
type AssertionText struct {
	IsChecked    int    `json:"is_checked"`    // 1 选中  -1 未选
	ResponseType int8   `json:"response_type"` //  1:ResponseHeaders; 2:ResponseData; 3: ResponseCode;
	Compare      string `json:"compare"`       // Includes、UNIncludes、Equal、UNEqual、GreaterThan、GreaterThanOrEqual、LessThan、LessThanOrEqual、Includes、UNIncludes、NULL、NotNULL、OriginatingFrom、EndIn
	Var          string `json:"var"`
	Val          string `json:"val"`
}

// VarForm 参数表
type VarForm struct {
	IsChecked   int64       `json:"is_checked" bson:"is_checked"`
	Type        string      `json:"type" bson:"type"`
	FileBase64  []string    `json:"fileBase64"`
	Key         string      `json:"key" bson:"key"`
	Value       interface{} `json:"value" bson:"value"`
	NotNull     int64       `json:"not_null" bson:"not_null"`
	Description string      `json:"description" bson:"description"`
	FieldType   string      `json:"field_type" bson:"field_type"`
}

type UsableMachineMap struct {
	IP               string // IP地址(包含端口号)
	Region           string // 机器所属区域
	Weight           int64  // 权重
	UsableGoroutines int64  // 可用协程数
}

// 压力机心跳上报数据
type HeartBeat struct {
	Name              string        `json:"name"`               // 机器名称
	CpuUsage          float64       `json:"cpu_usage"`          // CPU使用率
	CpuLoad           *load.AvgStat `json:"cpu_load"`           // CPU负载信息
	MemInfo           []MemInfo     `json:"mem_info"`           // 内存使用情况
	Networks          []Network     `json:"networks"`           // 网络连接情况
	DiskInfos         []DiskInfo    `json:"disk_infos"`         // 磁盘IO情况
	MaxGoroutines     int64         `json:"max_goroutines"`     // 当前机器支持最大协程数
	CurrentGoroutines int64         `json:"current_goroutines"` // 当前已用协程数
	ServerType        int64         `json:"server_type"`        // 压力机类型：0-主力机器，1-备用机器
	CreateTime        int64         `json:"create_time"`        // 数据上报时间（时间戳）
}

type MemInfo struct {
	Total       uint64  `json:"total"`
	Used        uint64  `json:"used"`
	Free        uint64  `json:"free"`
	UsedPercent float64 `json:"usedPercent"`
}

type DiskInfo struct {
	Total       uint64  `json:"total"`
	Free        uint64  `json:"free"`
	Used        uint64  `json:"used"`
	UsedPercent float64 `json:"usedPercent"`
}

type Network struct {
	Name        string `json:"name"`
	BytesSent   uint64 `json:"bytesSent"`
	BytesRecv   uint64 `json:"bytesRecv"`
	PacketsSent uint64 `json:"packetsSent"`
	PacketsRecv uint64 `json:"packetsRecv"`
}
