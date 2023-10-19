package events

type EventManager struct {
	UserAgent, HcapVersion, Href string
	Fingerprint                  *FpModel
}

type FingerprintEvent struct {
	EventID int64
	Value   string
}

type FpModel struct {
	Hash       map[string]any         `json:"hash"`
	Properties Properties             `json:"properties"`
	Browser    Browser                `json:"browser"`
	Connection Connection             `json:"connection"`
	Timezone   []interface{}          `json:"timezone"`
	Screen     map[string]interface{} `json:"screen"`
	Webgl      Webgl                  `json:"webgl"`
	Events     map[string]interface{} `json:"events"`
}

type Connection struct {
	DownlinkMax bool    `json:"downlinkMax"`
	Rtt         float64 `json:"rtt"`
}

type Properties struct {
	CSS             int64 `json:"css"`
	MIMETypes       int64 `json:"MimeTypes"`
	Plugins         int64 `json:"Plugins"`
	WindowFunctions int64 `json:"WindowFunctions"`
}

type Browser struct {
	UserAgent           string   `json:"UserAgent"`
	AppVersion          string   `json:"AppVersion"`
	DeviceMemory        int64    `json:"DeviceMemory"`
	HardwareConcurrency int64    `json:"HardwareConcurrency"`
	Language            string   `json:"Language"`
	Languages           []string `json:"Languages"`
	PDFViewerEnabled    bool     `json:"PDFViewerEnabled"`
	Platform            string   `json:"platform"`
	Mobile              bool     `json:"Mobile"`
}

type Webgl struct {
	Vendor   string `json:"vendor"`
	Renderer string `json:"renderer"`
}
