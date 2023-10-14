package fingerprint

import "github.com/Implex-ltd/hcsolver/internal/hcaptcha/fingerprint/events"

type Profile struct {
	UserAgent string
	Screen    Screen
	Navigator Navigator
	Hash      Hash
	Misc      Misc
}

type Builder struct {
	Manager *events.EventManager
	Profile *Profile
}

type Misc struct {
	UniqueKeys, InvUniqueKeys string
}

type Hash struct {
	Performance, Canvas, WebGL, WebRTC, Audio, ParrentWindow string
}

type Ndata struct {
	ProofSpec                   ProofSpec       `json:"proof_spec"`
	Rand                        []float64       `json:"rand"`
	Components                  Components      `json:"components"`
	FingerprintEvents           [][]interface{} `json:"fingerprint_events"`
	Messages                    interface{}     `json:"messages"`
	StackData                   interface{}     `json:"stack_data"`
	FingerprintSuspiciousEvents []string        `json:"fingerprint_suspicious_events"`
	Stamp                       string          `json:"stamp"`
	Errs                        Errs            `json:"errs"`
	Href                        string          `json:"href"`
	Perf                        [][]int64       `json:"perf"`
}

type Components struct {
	Version                   string      `json:"version"`
	Navigator                 Navigator   `json:"navigator"`
	Screen                    Screen      `json:"screen"`
	DevicePixelRatio          float64     `json:"device_pixel_ratio"`
	HasSessionStorage         bool        `json:"has_session_storage"`
	HasLocalStorage           bool        `json:"has_local_storage"`
	HasIndexedDB              bool        `json:"has_indexed_db"`
	WebGlHash                 string      `json:"web_gl_hash"`
	CanvasHash                string      `json:"canvas_hash"`
	HasTouch                  bool        `json:"has_touch"`
	NotificationAPIPermission string      `json:"notification_api_permission"`
	Chrome                    bool        `json:"chrome"`
	ToStringLength            int64       `json:"to_string_length"`
	ErrFirefox                interface{} `json:"err_firefox"`
	RBotScore                 int64       `json:"r_bot_score"`
	RBotScoreSuspiciousKeys   []string    `json:"r_bot_score_suspicious_keys"`
	RBotScore2                int64       `json:"r_bot_scorevents"`
	AudioHash                 string      `json:"audio_hash"`
	Extensions                []bool      `json:"extensions"`
	ParentWinHash             string      `json:"parent_win_hash"`
	WebrtcHash                string      `json:"webrtc_hash"`
	PerformanceHash           string      `json:"performance_hash"`
	UniqueKeys                string      `json:"unique_keys"`
	InvUniqueKeys             string      `json:"inv_unique_keys"`
	Features                  Features    `json:"features"`
}

type Features struct {
	PerformanceEntries bool `json:"performance_entries"`
	WebAudio           bool `json:"web_audio"`
	WebRTC             bool `json:"web_rtc"`
	Canvas2D           bool `json:"canvas_2d"`
	Fetch              bool `json:"fetch"`
}

type Navigator struct {
	UserAgent                   string      `json:"user_agent"`
	Language                    string      `json:"language"`
	Languages                   []string    `json:"languages"`
	Platform                    string      `json:"platform"`
	MaxTouchPoints              float64     `json:"max_touch_points"`
	Webdriver                   bool        `json:"webdriver"`
	NotificationQueryPermission interface{} `json:"notification_query_permission"`
	PluginsUndefined            bool        `json:"plugins_undefined"`
}

type Screen struct {
	ColorDepth  float64 `json:"color_depth"`
	PixelDepth  float64 `json:"pixel_depth"`
	Width       float64 `json:"width"`
	Height      float64 `json:"height"`
	AvailWidth  float64 `json:"avail_width"`
	AvailHeight float64 `json:"avail_height"`
}

type Errs struct {
	List []string `json:"list"`
}

type ProofSpec struct {
	Difficulty      int64  `json:"difficulty"`
	FingerprintType int64  `json:"fingerprint_type"`
	Type            string `json:"_type"`
	Data            string `json:"data"`
	Location        string `json:"_location"`
	TimeoutValue    int64  `json:"timeout_value"`
}
