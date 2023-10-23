package fingerprint

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"regexp"
	"strings"

	"github.com/0xF7A4C6/GoCycle"
	"github.com/Implex-ltd/hcsolver/internal/hcaptcha/fingerprint/events"
	"github.com/Implex-ltd/hcsolver/internal/utils"
)

var (
	CollectFpArray, _ = GoCycle.NewFromFile("../../assets/cleaned.txt")
	hcRegex           = regexp.MustCompile(`/captcha/v1/[A-Za-z0-9]+/static/images`)
	VERSION, _        = checkForUpdate()
	WASM              = "1.40.7"
)

func NewFingerprintBuilder(useragent, href string) (*Builder, error) {
	fp, err := CollectFpArray.Next()
	if err != nil {
		panic(err)
	}

	decodedData, err := base64.StdEncoding.DecodeString(fp)
	if err != nil {
		fmt.Println("Error decoding Base64:", err)
		return nil, err
	}

	var data events.FpModel
	if err := json.Unmarshal(decodedData, &data); err != nil {
		return nil, err
	}

	return &Builder{
		Manager: &events.EventManager{
			Href:        href,
			UserAgent:   useragent,
			HcapVersion: VERSION,
			Fingerprint: &data,
		},
	}, nil
}

func (B *Builder) GenerateProfile() (*Profile, error) {
	uaSplit := strings.Split(B.Manager.UserAgent, "Mozilla/")
	if len(uaSplit) != 2 {
		return nil, fmt.Errorf("invalid UA")
	}

	B.Manager.UserAgent = B.Manager.Fingerprint.Browser.UserAgent

	p := Profile{
		UserAgent: B.Manager.UserAgent,
		Screen: Screen{
			ColorDepth:  B.Manager.Fingerprint.Screen["ColorDepth"].(float64),
			PixelDepth:  B.Manager.Fingerprint.Screen["PixelDepth"].(float64),
			Width:       B.Manager.Fingerprint.Screen["Width"].(float64),
			Height:      B.Manager.Fingerprint.Screen["Height"].(float64),
			AvailWidth:  B.Manager.Fingerprint.Screen["AvailWidth"].(float64),
			AvailHeight: B.Manager.Fingerprint.Screen["AvailHeight"].(float64),
		},
		Navigator: Navigator{
			UserAgent:                   B.Manager.Fingerprint.Browser.UserAgent,
			Language:                    B.Manager.Fingerprint.Browser.Language,
			Languages:                   B.Manager.Fingerprint.Browser.Languages,
			Platform:                    B.Manager.Fingerprint.Browser.Platform,
			MaxTouchPoints:              B.Manager.Fingerprint.Screen["MaxTouchPoints"].(float64),
			NotificationQueryPermission: nil,
			PluginsUndefined:            true,
		},
		Hash: Hash{
			Performance:   "2047758435847122209", //utils.RandomHash(19),
			Canvas:        utils.RandomHash(19),
			WebGL:         "-1",
			WebRTC:        "-1",
			Audio:         "-1",
			ParrentWindow: "2556339636007144308", //utils.RandomHash(19),
		},
		Misc: Misc{
			UniqueKeys:    "0,IntlPolyfill,hcaptcha,__SECRET_EMOTION__,DiscordSentry,grecaptcha,platform,1,__sentry_instrumentation_handlers__,setImmediate,webpackChunkdiscord_app,_,GLOBAL_ENV,clearImmediate,__localeData__,__OVERLAY__,__SENTRY__,regeneratorRuntime,hcaptchaOnLoad,__timingFunction,DiscordErrors,__DISCORD_WINDOW_ID,__BILLING_STANDALONE__",
			InvUniqueKeys: "__wdata,image_label_binary,_sharedLibs,text_free_entry,sessionStorage,hsw,localStorage",
		},
	}

	B.Profile = &p
	return &p, nil
}

func (B *Builder) Build(jwt string) (*Ndata, error) {
	/*Profile, err := B.GenerateProfile()
	if err != nil {
		return nil, err
	}*/

	token, err := ParseJWT(jwt)
	if err != nil {
		return nil, err
	}

	/*stamp, err := GetStamp(token.PowData)
	if err != nil {
		return nil, fmt.Errorf("pow error")
	}*/

	V := strings.Split(token.Location, "https://newassets.hcaptcha.com/c/")
	if len(V) == 1 {
		return nil, fmt.Errorf("cant parse jwt location")
	}

	N := Ndata{
		ProofSpec: ProofSpec{
			Difficulty:      int64(token.Difficuly),
			FingerprintType: int64(token.FingerprintType),
			Type:            token.VmType,
			Data:            token.PowData,
			Location:        token.Location,
			TimeoutValue:    int64(token.TimeoutValue),
		},
		Rand: []float64{
			utils.RandomFloat64Precission(0, 1, 100000000000000000.0),
			utils.RandomFloat64Precission(0, 1, 100000000000000000.0),
		},
		Components: Components{
			Version:                   fmt.Sprintf("%v/%v", WASM, V[1]),
			Navigator:                 B.Profile.Navigator,
			Screen:                    B.Profile.Screen,
			DevicePixelRatio:          B.Manager.Fingerprint.Screen["DevicePixelRatio"].(float64),
			HasSessionStorage:         true,
			HasLocalStorage:           true,
			HasIndexedDB:              true,
			WebGlHash:                 B.Profile.Hash.WebGL,
			CanvasHash:                B.Profile.Hash.Canvas,
			HasTouch:                  B.Manager.Fingerprint.Events["107"].(map[string]interface{})["touchEvent"].(bool),
			NotificationAPIPermission: "Denied",
			Chrome:                    true, //strings.Contains(B.Manager.Fingerprint.Browser.UserAgent, "Chrome"),
			ToStringLength:            33,
			ErrFirefox:                nil,
			RBotScore:                 0,
			RBotScoreSuspiciousKeys:   []string{},
			RBotScore2:                0,
			AudioHash:                 B.Profile.Hash.Audio,
			Extensions: []bool{
				false,
			},
			ParentWinHash:   B.Profile.Hash.ParrentWindow,
			WebrtcHash:      B.Profile.Hash.WebRTC,
			PerformanceHash: B.Profile.Hash.Performance,
			UniqueKeys:      B.Profile.Misc.UniqueKeys,
			InvUniqueKeys:   B.Profile.Misc.InvUniqueKeys,
			Features: Features{
				PerformanceEntries: true,
				WebAudio:           true,
				WebRTC:             true,
				Canvas2D:           true,
				Fetch:              true,
			},
		},
		FingerprintEvents:           B.Manager.BuildEvents(),
		FingerprintSuspiciousEvents: []string{},
		//Stamp:                       stamp,
		Href: B.Manager.Href,
		Errs: Errs{
			List: []string{},
		},
		StackData: []string{},
		Perf: [][]int64{
			{
				1,
				int64(utils.RandomNumber(5, 100)),
			},
			{
				2,
				int64(utils.RandomNumber(20, 300)),
			},
			/*{
				3,
				0.0, //int64(utils.RandomNumber(0, 5)),
			},*/
		},
	}

	return &N, nil
}

func checkForUpdate() (string, error) {
	response, err := http.Get("https://hcaptcha.com/1/api.js?render=explicit&onload=hcaptchaOnLoad")
	if err != nil {
		return "", err
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	var found string
	for _, match := range hcRegex.FindAllStringSubmatch(string(body), -1) {
		found = match[0]
		break
	}

	if found == "" {
		return "", fmt.Errorf("cant find version")
	}

	version := strings.Split(strings.Split(found, "v1/")[1], "/static")[0]

	log.Println("load version", version)

	return version, nil
}
