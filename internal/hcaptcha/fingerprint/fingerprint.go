package fingerprint

import (
	"encoding/base64"
	"encoding/json"
	"fmt"

	"github.com/0xF7A4C6/GoCycle"
	"github.com/Implex-ltd/hcsolver/internal/utils"
)

var (
	CollectFpArray, _ = GoCycle.NewFromFile("../../assets/cleaned.txt")
)

func NewFingerprintBuilder(useragent string) (*Builder, error) {
	fp, err := CollectFpArray.Next()
	if err != nil {
		panic(err)
	}

	decodedData, err := base64.StdEncoding.DecodeString(fp)
	if err != nil {
		fmt.Println("Error decoding Base64:", err)
		return nil, err
	}

	var data NdataCollect
	if err := json.Unmarshal(decodedData, &data); err != nil {
		return nil, err
	}

	data.ParsedEvents = make(map[int]interface{})
	for _, event := range data.FingerprintEvents {
		eventList, ok := event.([]interface{})
		if ok && len(eventList) >= 2 {
			data.ParsedEvents[int(eventList[0].(float64))] = eventList[1]
		}
	}

	return &Builder{
		UserAgent:   useragent,
		CollectedFp: &data,
		HcapVersion: "1b812e2",
		HswVersion:  "1.40.1/7a7fc3d",
	}, nil
}

func (B *Builder) GenerateProfile() (*Profile, error) {
	p := Profile{
		UserAgent: B.UserAgent,
		Screen: Screen{
			ColorDepth:  B.CollectedFp.Components.Screen.ColorDepth,
			PixelDepth:  B.CollectedFp.Components.Screen.PixelDepth,
			Width:       B.CollectedFp.Components.Screen.Width,
			Height:      B.CollectedFp.Components.Screen.Height,
			AvailWidth:  B.CollectedFp.Components.Screen.AvailWidth,
			AvailHeight: B.CollectedFp.Components.Screen.AvailHeight,
		},
		Navigator: Navigator{
			UserAgent:                   "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/116.0.0.0 Safari/537.36",
			Language:                    B.CollectedFp.Components.Navigator.Language,
			Languages:                   B.CollectedFp.Components.Navigator.Languages,
			Platform:                    B.CollectedFp.Components.Navigator.Platform,
			MaxTouchPoints:              B.CollectedFp.Components.Navigator.MaxTouchPoints,
			Webdriver:                   false,
			NotificationQueryPermission: nil,
			PluginsUndefined:            B.CollectedFp.Components.Navigator.PluginsUndefined,
		},
		Hash: Hash{
			Performance:   utils.RandomHash(19), //B.CollectedFp.Components.PerformanceHash,
			Canvas:        utils.RandomHash(19), //B.CollectedFp.Components.CanvasHash,
			WebGL:         B.CollectedFp.Components.WebGlHash,
			WebRTC:        B.CollectedFp.Components.WebrtcHash,
			Audio:         B.CollectedFp.Components.AudioHash,
			ParrentWindow: "17464972382824382480", //utils.RandomHash(20),
		},
		Misc: Misc{
			HasTouch:            false,
			Chrome:              B.CollectedFp.Components.Chrome,
			UniqueKeys:          "0,IntlPolyfill,hcaptcha,__SECRET_EMOTION__,DiscordSentry,grecaptcha,platform,1,__sentry_instrumentation_handlers__,setImmediate,webpackChunkdiscord_app,2,_,GLOBAL_ENV,clearImmediate,__localeData__,__OVERLAY__,__SENTRY__,regeneratorRuntime,hcaptchaOnLoad,__timingFunction,DiscordErrors,__DISCORD_WINDOW_ID,__BILLING_STANDALONE__", //B.CollectedFp.Components.UniqueKeys,
			InvUniqueKeys:       "__wdata,image_label_binary,_sharedLibs,text_free_entry,sessionStorage,hsw,localStorage",                                                                                                                                                                                                                                                  //B.CollectedFp.Components.InvUniqueKeys,
			DeviceMemory:        utils.RandomElementInt([]int{2, 4, 8, 1632, 64, 128}),
			HardwareConcurrency: utils.RandomElementInt([]int{2, 4, 6, 8, 12, 16, 32, 64}),
			ChromeVersion:       "116",
			Os:                  "Windows",
			Arch:                "x86",
			CPU:                 "64",
			BrowserAppVersion:   utils.RandomElementString([]string{"116.0.5845.187", "116.0.5845.179", "116.0.5845.140", "116.0.5845.110", "116.0.5845.96"}),
			Vendor:              "Google Inc. (NVIDIA)",
			Renderer:            "ANGLE (NVIDIA, NVIDIA GeForce RTX 3060 Ti Direct3D11 vs_5_0 ps_5_0, D3D11)",
		},
	}

	B.Profile = &p
	return &p, nil
}

func (B *Builder) Build() (*Ndata, error) {
	if B.Profile == nil {
		return nil, fmt.Errorf("you need to generate profile first")
	}

	N := Ndata{
		Components: Components{
			Version:                   B.HswVersion,
			Navigator:                 B.Profile.Navigator,
			Screen:                    B.Profile.Screen,
			DevicePixelRatio:          B.CollectedFp.Components.DevicePixelRatio,
			HasSessionStorage:         B.CollectedFp.Components.HasSessionStorage,
			HasLocalStorage:           B.CollectedFp.Components.HasLocalStorage,
			HasIndexedDB:              B.CollectedFp.Components.HasIndexedDB,
			WebGlHash:                 B.Profile.Hash.WebGL,
			CanvasHash:                B.Profile.Hash.Canvas,
			HasTouch:                  B.Profile.Misc.HasTouch,
			NotificationAPIPermission: B.CollectedFp.Components.NotificationAPIPermission,
			Chrome:                    B.Profile.Misc.Chrome,
			ToStringLength:            B.CollectedFp.Components.ToStringLength,
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
		FingerprintEvents:           [][]interface{}{},
		FingerprintSuspiciousEvents: []string{},
		Errs: Errs{
			List: []string{},
		},
		Perf: [][]int64{
			{
				1,
				int64(utils.RandomNumber(10, 100)),
			},
			{
				2,
				int64(utils.RandomNumber(50, 300)),
			},
			{
				3,
				int64(utils.RandomNumber(0, 5)),
			},
		},
	}

	eventsMethods := []func(*Builder) FingerprintEvent{
		(*Builder).Event_0,
		(*Builder).Event_3,
		(*Builder).Event_107,
		(*Builder).Event_201,
		(*Builder).Event_211,
		(*Builder).Event_301,
		//(*Builder).Event_302,
		//(*Builder).Event_303,
		(*Builder).Event_304,
		(*Builder).Event_401,
		(*Builder).Event_402,
		(*Builder).Event_407,
		(*Builder).Event_412,
		(*Builder).Event_604,
		(*Builder).Event_702,
		(*Builder).Event_803,
		(*Builder).Event_901,
		(*Builder).Event_905,
		(*Builder).Event_1101,
		(*Builder).Event_1103,
		(*Builder).Event_1105,
		(*Builder).Event_1107,
		(*Builder).Event_1302,
		(*Builder).Event_1401,
		(*Builder).Event_1402,
		(*Builder).Event_1403,
		(*Builder).Event_1901,
		(*Builder).Event_1902,
		(*Builder).Event_1904,
		//(*Builder).Event_2001,
		//(*Builder).Event_2002,
		(*Builder).Event_2401,
		(*Builder).Event_2402,
		//(*Builder).Event_2403,
		(*Builder).Event_2407,
		(*Builder).Event_2408,
		(*Builder).Event_2409,
		(*Builder).Event_2410,
		(*Builder).Event_2411,
		(*Builder).Event_2412,
		(*Builder).Event_2413,
		(*Builder).Event_2414,
		(*Builder).Event_2415,
		(*Builder).Event_2416,
		(*Builder).Event_2417,
		(*Builder).Event_2420,
		(*Builder).Event_2801,
		(*Builder).Event_2805,
		(*Builder).Event_3210,
		(*Builder).Event_3211,
		(*Builder).Event_3401,
		(*Builder).Event_3403,
		(*Builder).Event_3501,
		(*Builder).Event_3502,
		(*Builder).Event_3503,
		(*Builder).Event_3504,
		//(*Builder).Event_3505,
	}

	for _, eventMethod := range eventsMethods {
		event := eventMethod(B)
		N.FingerprintEvents = append(N.FingerprintEvents, []interface{}{
			event.EventID,
			event.Value,
		})
	}

	return &N, nil
}
