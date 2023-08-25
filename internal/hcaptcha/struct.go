package hcaptcha

import (
	"time"

	"github.com/Implex-ltd/cleanhttp/cleanhttp"
	"github.com/Implex-ltd/fingerprint-client/fpclient"
	"github.com/Implex-ltd/hcsolver/internal/solver"
	"go.uber.org/zap"
)

type Config struct {
	UserAgent string
	SiteKey   string
	Domain    string
	Proxy     string

	Logger *zap.Logger
}

type Hcap struct {
	Config      *Config
	Http        *cleanhttp.CleanHttp
	Fingerprint *fpclient.Fingerprint

	// metrics
	AnswerProcessing time.Duration
	HswProcessing    time.Duration

	Logger *zap.Logger
}

type Motion struct {
	IsCheck bool
	Answers map[string]string
}

/*
	Payloads
*/

type Pdc struct {
	S   int64 `json:"s"`
	N   int   `json:"n"`
	P   int   `json:"p"`
	Gcs int   `json:"gcs"`
}

/*
	general
*/

type C struct {
	Type string `json:"type"`
	Req  string `json:"req"`
}

/*
	checksiteconfig
*/

type SiteConfig struct {
	Features Features `json:"features"`
	C        C        `json:"c"`
	Pass     bool     `json:"pass"`
}

type Features struct {
	A11YChallenge bool `json:"a11y_challenge"`
}

/*
	getcaptcha
*/

type Captcha struct {
	C                        C                 `json:"c"`
	ChallengeURI             string            `json:"challenge_uri"`
	Key                      string            `json:"key"`
	RequestConfig            RequestConfig     `json:"request_config"`
	RequestType              string            `json:"request_type"`
	RequesterQuestion        RequesterQuestion `json:"requester_question"`
	RequesterQuestionExample []string          `json:"requester_question_example"`
	Tasklist                 []solver.TaskList `json:"tasklist"`
	BypassMessage            string            `json:"bypass-message"`

	RequesterRestrictedAnswerSet map[string]map[string]string `json:"requester_restricted_answer_set"`
}

type RequestConfig struct {
	Version                      int    `json:"version"`
	ShapeType                    string `json:"shape_type"`
	MinPoints                    int    `json:"min_points"`
	MaxPoints                    int    `json:"max_points"`
	MinShapesPerImage            int    `json:"min_shapes_per_image"`
	MaxShapesPerImage            int    `json:"max_shapes_per_image"`
	RestrictToCoords             string `json:"restrict_to_coords"`
	MinimumSelectionAreaPerShape string `json:"minimum_selection_area_per_shape"`
	MultipleChoiceMaxChoices     int    `json:"multiple_choice_max_choices"`
	MultipleChoiceMinChoices     int    `json:"multiple_choice_min_choices"`
	OverlapThreshold             string `json:"overlap_threshold"`
	AnswerType                   string `json:"answer_type"`
	MaxValue                     string `json:"max_value"`
	MinValue                     string `json:"min_value"`
	MaxLength                    string `json:"max_length"`
	MinLength                    string `json:"min_length"`
	SigFigs                      string `json:"sig_figs"`
	KeepAnswersOrder             bool   `json:"keep_answers_order"`
}

type RequesterQuestion struct {
	Fr string `json:"fr"`
	En string `json:"en"`
}

type Tasklist struct {
	DatapointURI string `json:"datapoint_uri"`
	TaskKey      string `json:"task_key"`
}

/*
	motion data
*/

type Empty struct {
}

type CheckData struct {
	St  int64 `json:"st"`
	Dct int64 `json:"dct"`

	Mm   [][]int64 `json:"mm"`
	MmMp float64   `json:"mm-mp"`

	Md   [][]int64 `json:"md"`
	MdMp float64   `json:"md-mp"`

	Mu   [][]int64 `json:"mu"`
	MuMp float64   `json:"mu-mp"`

	TopLevel TopLevel `json:"topLevel"`
	V        int64    `json:"v"`
}

type GetData struct {
	St int64 `json:"st"`

	Mm   [][]int64 `json:"mm"`
	MmMp float64   `json:"mm-mp"`

	Md   [][]int64 `json:"md"`
	MdMp float64   `json:"md-mp"`

	Mu   [][]int64 `json:"mu"`
	MuMp float64   `json:"mu-mp"`

	V        int64    `json:"v"`
	TopLevel TopLevel `json:"topLevel"`

	Session    []string `json:"session"`
	WidgetList []string `json:"widgetList"`
	WidgetID   string   `json:"widgetId"`
	Href       string   `json:"href"`
	Prev       Prev     `json:"prev"`
}

type TopLevel struct {
	Inv  bool   `json:"inv"`
	St   int64  `json:"st"`
	Sc   Sc     `json:"sc"`
	Nv   Nv     `json:"nv"`
	DR   string `json:"dr"`
	Exec bool   `json:"exec"`

	Wn   [][]int64 `json:"wn"`
	WnMp float64   `json:"wn-mp"`

	Xy   [][]int64 `json:"xy"`
	XyMp float64   `json:"xy-mp"`

	Mm   [][]int64 `json:"mm"`
	MmMp float64   `json:"mm-mp"`

	//Md   [][]int64 `json:"md"`
	//MdMp int64     `json:"md-mp"`
	//Mu   [][]int64 `json:"mu"`
	//MuMp int64     `json:"mu-mp"`
}

type Nv struct {
	/*
		Iphone params:
			- Standalone bool `json:"standalone"`
	*/

	Clipboard              Empty         `json:"clipboard"`
	VendorSub              string        `json:"vendorSub"`
	ProductSub             string        `json:"productSub"`
	Vendor                 string        `json:"vendor"`
	MaxTouchPoints         int64         `json:"maxTouchPoints"`
	Scheduling             Empty         `json:"scheduling"`
	UserActivation         Empty         `json:"userActivation"`
	DoNotTrack             interface{}   `json:"doNotTrack"`
	Geolocation            Empty         `json:"geolocation"`
	Connection             Empty         `json:"connection"`
	PDFViewerEnabled       bool          `json:"pdfViewerEnabled"`
	WebkitTemporaryStorage Empty         `json:"webkitTemporaryStorage"`
	HardwareConcurrency    int64         `json:"hardwareConcurrency"`
	CookieEnabled          bool          `json:"cookieEnabled"`
	AppCodeName            string        `json:"appCodeName"`
	AppName                string        `json:"appName"`
	AppVersion             string        `json:"appVersion"`
	Platform               string        `json:"platform"`
	Product                string        `json:"product"`
	UserAgent              string        `json:"userAgent"`
	Language               string        `json:"language"`
	Languages              []string      `json:"languages"`
	OnLine                 bool          `json:"onLine"`
	Webdriver              bool          `json:"webdriver"`
	Bluetooth              Empty         `json:"bluetooth"`
	Credentials            Empty         `json:"credentials"`
	Keyboard               Empty         `json:"keyboard"`
	Managed                Empty         `json:"managed"`
	MediaDevices           Empty         `json:"mediaDevices"`
	Storage                Empty         `json:"storage"`
	ServiceWorker          Empty         `json:"serviceWorker"`
	VirtualKeyboard        Empty         `json:"virtualKeyboard"`
	WakeLock               Empty         `json:"wakeLock"`
	DeviceMemory           int64         `json:"deviceMemory"`
	Ink                    Empty         `json:"ink"`
	HID                    Empty         `json:"hid"`
	Locks                  Empty         `json:"locks"`
	MediaCapabilities      Empty         `json:"mediaCapabilities"`
	MediaSession           Empty         `json:"mediaSession"`
	Permissions            Empty         `json:"permissions"`
	Presentation           Empty         `json:"presentation"`
	Serial                 Empty         `json:"serial"`
	USB                    Empty         `json:"usb"`
	WindowControlsOverlay  Empty         `json:"windowControlsOverlay"`
	Xr                     Empty         `json:"xr"`
	UserAgentData          UserAgentData `json:"userAgentData"`
	Plugins                []string      `json:"plugins"`
}

type Brand struct {
	Brand   string `json:"brand"`
	Version string `json:"version"`
}

type UserAgentData struct {
	Brands   []Brand `json:"brands"`
	Mobile   bool    `json:"mobile"`
	Platform string  `json:"platform"`
}

type Sc struct {
	AvailWidth  int64       `json:"availWidth"`
	AvailHeight int64       `json:"availHeight"`
	Width       int64       `json:"width"`
	Height      int64       `json:"height"`
	ColorDepth  int64       `json:"colorDepth"`
	PixelDepth  int64       `json:"pixelDepth"`
	AvailLeft   int64       `json:"availLeft"`
	AvailTop    int64       `json:"availTop"`
	Onchange    interface{} `json:"onchange"`
	IsExtended  bool        `json:"isExtended"`
}

type Prev struct {
	Escaped          bool `json:"escaped"`
	Passed           bool `json:"passed"`
	ExpiredChallenge bool `json:"expiredChallenge"`
	ExpiredResponse  bool `json:"expiredResponse"`
}

/*
	Mouse curves
*/

type Box struct {
	Start, End Point
}

type Point struct {
	X, Y, T int64
}

/*
	Anwsers
*/

type LabelAreaSelect struct {
	TaskType   string     `json:"task_type"`
	Question   string     `json:"question"`
	EntityType string     `json:"entity_type"`
	Tasklist   []Tasklist `json:"tasklist"`
}

type LabelBinaryPayload struct {
	TaskType string     `json:"task_type"`
	Question string     `json:"question"`
	Tasklist []Tasklist `json:"tasklist"`
}

type AiSolverResponse struct {
	Success bool           `json:"success"`
	Data    map[string]any `json:"data"`
}

/*
	checkcaptcha
*/

type PayloadCheckChallenge struct {
	V            string            `json:"v"`
	JobMode      string            `json:"job_mode"`
	Answers      map[string]string `json:"answers"`
	Serverdomain string            `json:"serverdomain"`
	Sitekey      string            `json:"sitekey"`
	MotionData   string            `json:"motionData"`
	N            string            `json:"n"`
	C            string            `json:"c"`
}

type ResponseCheckCaptcha struct {
	C                 C      `json:"c"`
	Pass              bool   `json:"pass"`
	GeneratedPassUUID string `json:"generated_pass_UUID"`
	Expiration        int    `json:"expiration"`
}
