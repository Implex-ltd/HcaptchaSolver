package fingerprint

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/Implex-ltd/hcsolver/internal/utils"
)

/*
 Hash that change from chrome version:
 	- 1101
	- 3401
	- 2401
	- 2001
	- 301
	- 401
	- 901
*/

/**
 * hash to be undetect:
	* 1101 h
	* 201  h
	* 3401 h
	* 3501
	* 3503
	* 2401 h
	* 2407 h
	* 1904
	* 3210
	* 702
*/

var (
	Hardcode = true
	hashs    = map[int]string{
		//1101: "5976439626078859368",
		201:  "17864504342478246445", // never change 13122422878918113034
		3401: "12039597181259123542",
		//2401: "3828429617252524255",
		//2407: "13177607191192652685",
		301: "8383473043360077444",
	}
)

func (B *Builder) Stringify(data any) string {
	jsonString, err := json.Marshal(data)
	if err != nil {
		return ""
	}

	return string(jsonString)
}

func (B *Builder) GetByID(id int, hardcoded string) string {
	if fp := B.CollectedFp.ParsedEvents[id]; fp != nil {
		if !Hardcode {
			hardcoded = fp.(string)
		}
	}

	return hardcoded
}

func (B *Builder) Event_0() FingerprintEvent {
	return FingerprintEvent{
		0,
		fmt.Sprintf("%.14f", utils.RandomFloat64(15, 40)),
	}
}

func (B *Builder) Event_3() FingerprintEvent {
	return FingerprintEvent{
		3,
		fmt.Sprintf("%.2f", utils.RandomFloat64(30000, 60000)),
	}
}

// https://gist.github.com/nikolahellatrigger/ea00832b010c0db8f0a0d5ca0d467072
func (B *Builder) Event_107() FingerprintEvent {
	data := B.Stringify([]interface{}{
		B.Profile.Screen.Width,
		B.Profile.Screen.Height,
		B.Profile.Screen.AvailWidth,
		B.Profile.Screen.AvailHeight,
		B.Profile.Screen.ColorDepth,
		B.Profile.Screen.PixelDepth,
		false,
		B.Profile.Navigator.MaxTouchPoints,
		B.CollectedFp.Components.DevicePixelRatio,
		B.Profile.Screen.Width - int64(utils.RandomNumber(100, 300)),  // outerWidth
		B.Profile.Screen.Height - int64(utils.RandomNumber(100, 300)), // outerHeight
		true,
		true,
		true,
		false,
	})

	return FingerprintEvent{
		107,
		data,
	}
}

func (B *Builder) Event_201() FingerprintEvent {
	return FingerprintEvent{
		201,
		B.GetByID(201, hashs[201]),
	}
}

func (B *Builder) Event_211() FingerprintEvent {
	data := B.Stringify([]interface{}{
		-6.172840118408203,
		-20.710678100585938,
		120.71067810058594,
		-20.710678100585938,
		141.42135620117188,
		120.71067810058594,
		-20.710678100585938,
		141.42135620117188,
		-20.710678100585938,
		-20.710678100585938,
		0,
		0,
		300,
		150,
		false,
		[]interface{}{
			0,
			15,
			33,
			34,
			35,
			37,
			39,
			75,
		},
	})

	return FingerprintEvent{
		211,
		B.GetByID(211, data),
	}
}

// Value: CSS Properties list (hash is different accross browsers but shared between devices)
func (B *Builder) Event_301() FingerprintEvent {
	return FingerprintEvent{
		301,
		B.GetByID(301, hashs[301]),
	}
}

func (B *Builder) Event_302() FingerprintEvent {
	data := B.Stringify([]interface{}{
		0,
		1,
		2,
		3,
		4,
	})

	return FingerprintEvent{
		302,
		B.GetByID(302, data),
	}
}

func (B *Builder) Event_303() FingerprintEvent {
	data := B.Stringify([]interface{}{
		"Arial",
		"\"Segoe UI\"",
	})

	return FingerprintEvent{
		303,
		B.GetByID(303, data),
	}
}

func (B *Builder) Event_304() FingerprintEvent {
	return FingerprintEvent{
		304,
		B.GetByID(304, "623"),
	}
}

// Object.getOwnPropertyNames(window)
func (B *Builder) Event_401() FingerprintEvent {
	return FingerprintEvent{
		401,
		B.GetByID(401, B.CollectedFp.ParsedEvents[401].(string)),
	}
}

func (B *Builder) Event_402() FingerprintEvent {
	return FingerprintEvent{
		402,
		B.GetByID(402, "1109"),
	}
}

func (B *Builder) Event_407() FingerprintEvent {
	data := B.Stringify([]interface{}{
		[]interface{}{
			"loadTimes",
			"csi",
			"app",
		},
		35,
		34,
		nil,
		false,
		false,
		true,
		37,
		true,
		true,
		true,
		true,
		true,
		[]interface{}{
			"Raven",
			"_sharedLibs",
			"hsw",
			"__wdata",
		},
		[]interface{}{
			[]interface{}{
				"getElementsByClassName",
				[]interface{}{},
			},
			[]interface{}{
				"getElementById",
				[]interface{}{},
			},
			[]interface{}{
				"querySelector",
				[]interface{}{},
			},
			[]interface{}{
				"querySelectorAll",
				[]interface{}{},
			},
		},
		[]interface{}{},
		true,
	})

	return FingerprintEvent{
		407,
		data,
	}
}

func (B *Builder) Event_412() FingerprintEvent {
	return FingerprintEvent{
		412,
		B.GetByID(412, B.CollectedFp.ParsedEvents[412].(string)),
	}
}

// https://gist.github.com/nikolahellatrigger/c4d6cf4ddb0ab219c38ddd133dc772eb
func (B *Builder) Event_604() FingerprintEvent {
	data := B.Stringify([]interface{}{
		B.Profile.Misc.AppVersion,
		B.UserAgent,
		B.Profile.Misc.DeviceMemory,
		B.Profile.Misc.HardwareConcurrency,
		B.Profile.Navigator.Language,
		B.Profile.Navigator.Languages,
		B.Profile.Navigator.Platform,
		nil,
		[]interface{}{
			fmt.Sprintf("Google Chrome %s", B.Profile.Misc.ChromeVersion),
			"Not;A=Brand 8",
			fmt.Sprintf("Chromium %s", B.Profile.Misc.ChromeVersion),
		},
		B.Profile.Misc.Mobile,
		B.Profile.Misc.Os,
		2,                               // mimeTypes len
		5,                               // plugins len
		B.Profile.Misc.PDFViewerEnabled, // pdf
		false,
		50,
		false,
		false,
		true,
		"[object Keyboard]",
		strings.Contains(strings.ToLower(B.UserAgent), "brave"),
		strings.Contains(strings.ToLower(B.UserAgent), "duckduckgo"),
	})

	return FingerprintEvent{
		604,
		data,
	}
}

func (B *Builder) Event_702() FingerprintEvent {
	data := B.Stringify([]interface{}{
		B.Profile.Misc.Os,
		"14.0.0",
		nil,
		B.Profile.Misc.CPU,
		B.Profile.Misc.Arch,
		B.Profile.Misc.BrowserVersion,
	})

	return FingerprintEvent{
		702,
		data,
	}
}

func (B *Builder) Event_803() FingerprintEvent {
	data := B.Stringify([]interface{}{
		1,
		4,
		5,
		7,
		9,
		12,
		20,
		21,
		24,
		25,
		29,
	})

	return FingerprintEvent{
		803,
		B.GetByID(803, data),
	}
}

// Some voices fingerprinting data
func (B *Builder) Event_901() FingerprintEvent {
	return FingerprintEvent{
		901,
		B.GetByID(901, B.CollectedFp.ParsedEvents[901].(string)),
	}
}

func (B *Builder) Event_905() FingerprintEvent {
	data := B.Stringify([]interface{}{
		[]interface{}{
			true,
			"fr-FR",
			true,
			"Microsoft Hortense - French (France)",
			"Microsoft Hortense - French (France)",
		},
		[]interface{}{
			false,
			"fr-FR",
			true,
			"Microsoft Julie - French (France)",
			"Microsoft Julie - French (France)",
		},
		[]interface{}{
			false,
			"fr-FR",
			true,
			"Microsoft Paul - French (France)",
			"Microsoft Paul - French (France)",
		},
	})

	/*data := B.Stringify([]interface{}{
		[]interface{}{
			true,
			"en-US",
			true,
			"Microsoft David - English (United States)",
			"Microsoft David - English (United States)",
		},
		[]interface{}{
			false,
			"en-US",
			true,
			"Microsoft Mark - English (United States)",
			"Microsoft Mark - English (United States)",
		},
		[]interface{}{
			false,
			"en-US",
			true,
			"Microsoft Zira - English (United States)",
			"Microsoft Zira - English (United States)",
		},
	})*/

	return FingerprintEvent{
		905,
		B.GetByID(905, data),
	}
}

// Canvas fingerprinting
func (B *Builder) Event_1101() FingerprintEvent {
	return FingerprintEvent{
		1101,
		//B.GetByID(1101, B.CollectedFp.Components.CanvasHash),
		//B.GetByID(1101, B.GetByID(1101, hashs[1101])),
		B.GetByID(1101, B.CollectedFp.ParsedEvents[1101].(string)),
	}
}

func (B *Builder) Event_1103() FingerprintEvent {
	return FingerprintEvent{
		1103,
		B.GetByID(1103, B.CollectedFp.ParsedEvents[1103].(string)),
	}
}

func (B *Builder) Event_1105() FingerprintEvent {
	return FingerprintEvent{
		1105,
		B.GetByID(1105, B.CollectedFp.ParsedEvents[1105].(string)),
	}
}

func (B *Builder) Event_1107() FingerprintEvent {
	x := utils.RandomNumber(240, 250)

	data := B.Stringify([]interface{}{
		[]interface{}{
			x,
			[]interface{}{
				x,
				x,
				x,
				255,
				x,
				x,
				x,
				255,
				x,
				x,
				x,
				255,
				x,
				x,
				x,
				255,
			},
		},
		[]interface{}{
			[]interface{}{
				11,
				0,
				0,
				95.96875,
				15,
				4,
				96.765625,
			},
			[]interface{}{
				[]interface{}{
					12,
					0,
					-1,
					113.125,
					17,
					4,
					113,
				},
				[]interface{}{
					11,
					0,
					0,
					111,
					12,
					4,
					111,
				},
				[]interface{}{
					11,
					0,
					0,
					95.96875,
					15,
					4,
					96.765625,
				},
				[]interface{}{
					11,
					0,
					0,
					95.96875,
					15,
					4,
					96.765625,
				},
				[]interface{}{
					11,
					0,
					0,
					95.96875,
					15,
					4,
					96.765625,
				},
				[]interface{}{
					11,
					0,
					0,
					95.96875,
					15,
					4,
					96.765625,
				},
				[]interface{}{
					11,
					0,
					0,
					95.96875,
					15,
					4,
					96.765625,
				},
				[]interface{}{
					11,
					0,
					0,
					95.96875,
					15,
					4,
					96.765625,
				},
				[]interface{}{
					12,
					0,
					0,
					109.640625,
					14,
					3,
					110.1953125,
				},
			},
		},
	})

	return FingerprintEvent{
		1107,
		B.GetByID(1107, data),
	}
}

func (B *Builder) Event_1302() FingerprintEvent {
	data := B.Stringify([]interface{}{
		0, 1, 2, 3, 4,
	})

	return FingerprintEvent{
		1302,
		B.GetByID(1302, data),
	}
}

func (B *Builder) Event_1401() FingerprintEvent {
	data := "\"Europe/Paris\""

	return FingerprintEvent{
		1401,
		B.GetByID(1401, data),
	}
}

func (B *Builder) Event_1402() FingerprintEvent {
	data := B.Stringify([]interface{}{
		"Europe/Paris",
		-60,
		-60,
		-3203647761000,
		"heure d’été d’Europe centrale",
		"fr",
	})

	return FingerprintEvent{
		1402,
		B.GetByID(1402, data),
	}
}

func (B *Builder) Event_1403() FingerprintEvent {
	return FingerprintEvent{
		1403,
		B.Stringify(EncStr(B.CollectedFp.ParsedEvents[1403].(string))),
	}
}

// Math Fingerprinting, Hash is different across browser types
func (B *Builder) Event_1901() FingerprintEvent {
	return FingerprintEvent{
		1901,
		B.GetByID(1901, B.CollectedFp.ParsedEvents[1901].(string)),
	}
}

func (B *Builder) Event_1902() FingerprintEvent {
	return FingerprintEvent{
		1902,
		B.GetByID(1902, "57"),
	}
}

func (B *Builder) Event_1904() FingerprintEvent {
	return FingerprintEvent{
		1904,
		B.GetByID(1904, "[0,11411,11411]"),
	}
}

// disabled (Navigator Permissions)
func (B *Builder) Event_2001() FingerprintEvent {
	return FingerprintEvent{
		2001,
		B.GetByID(2001, "13414760775080815217"),
	}
}

// disabled
func (B *Builder) Event_2002() FingerprintEvent {
	data := B.Stringify([]interface{}{
		"denied",
		"denied",
	})

	return FingerprintEvent{
		2002,
		B.GetByID(2002, data),
	}
}

// WebGL properties hash
func (B *Builder) Event_2401() FingerprintEvent {
	return FingerprintEvent{
		2401,
		//B.GetByID(2401, hashs[2401]),
		B.GetByID(2401, B.CollectedFp.ParsedEvents[2401].(string)),
	}
}

func (B *Builder) Event_2402() FingerprintEvent {
	/*data := B.Stringify([]interface{}{
		B.Profile.Misc.Vendor,
		B.Profile.Misc.Renderer,
	})*/

	data := B.CollectedFp.ParsedEvents[2402].(string)

	return FingerprintEvent{
		2402,
		B.GetByID(2402, data),
	}
}

// disabled
func (B *Builder) Event_2403() FingerprintEvent {
	/*data := []interface{}{
		B.Profile.Misc.Vendor,
		B.Profile.Misc.Renderer,
	}*/

	data := B.CollectedFp.ParsedEvents[2403].(string)

	return FingerprintEvent{
		2403,
		B.GetByID(2403, data),
	}
}

func (B *Builder) Event_2407() FingerprintEvent {
	return FingerprintEvent{
		2407,
		//B.GetByID(2407, hashs[2407]),
		B.GetByID(2407, B.CollectedFp.ParsedEvents[2407].(string)),
	}
}

func (B *Builder) Event_2408() FingerprintEvent {
	return FingerprintEvent{
		2408,
		"true",
	}
}

func (B *Builder) Event_2409() FingerprintEvent {
	data := B.Stringify([]interface{}{
		2147483647,
		2147483647,
		4294967294,
	})

	return FingerprintEvent{
		2409,
		data,
	}
}

func (B *Builder) Event_2410() FingerprintEvent {
	data := B.Stringify([]interface{}{
		16,
		1024,
		4096,
		7,
		12,
		120,
		[]interface{}{23, 127, 127},
	})

	return FingerprintEvent{
		2410,
		data,
	}
}

func (B *Builder) Event_2411() FingerprintEvent {
	data := B.Stringify([]interface{}{
		32767,
		32767,
		16384,
		8,
		8,
		8,
	})

	return FingerprintEvent{
		2411,
		data,
	}
}

func (B *Builder) Event_2412() FingerprintEvent {
	data := B.Stringify([]interface{}{
		1,
		1024,
		1,
		1,
		4,
	})

	return FingerprintEvent{
		2412,
		data,
	}
}

func (B *Builder) Event_2413() FingerprintEvent {
	data := B.Stringify([]interface{}{
		2147483647,
		2147483647,
		2147483647,
		2147483647,
	})

	return FingerprintEvent{
		2413,
		data,
	}
}

func (B *Builder) Event_2414() FingerprintEvent {
	data := B.Stringify([]interface{}{
		16384,
		32,
		16384,
		2048,
		2,
		2048,
	})

	return FingerprintEvent{
		2414,
		data,
	}
}

func (B *Builder) Event_2415() FingerprintEvent {
	data := B.Stringify([]interface{}{
		4,
		120,
		4,
	})

	return FingerprintEvent{
		2415,
		data,
	}
}

func (B *Builder) Event_2416() FingerprintEvent {
	data := B.Stringify([]interface{}{
		24,
		24,
		65536,
		212988,
		200704,
	})

	return FingerprintEvent{
		2416,
		data,
	}
}

func (B *Builder) Event_2417() FingerprintEvent {
	data := B.Stringify([]interface{}{
		16,
		4095,
		30,
		16,
		16380,
		120,
		12,
		120,
		[]interface{}{23, 127, 127},
	})

	return FingerprintEvent{
		2417,
		data,
	}
}

func (B *Builder) Event_2420() FingerprintEvent {
	var out []string
	json.Unmarshal([]byte(B.CollectedFp.ParsedEvents[2402].(string)), &out)

	data := B.Stringify([]interface{}{
		EncStr(out[0]),
		EncStr(out[1]),
	})

	return FingerprintEvent{
		2420,
		data,
	}
}

// audio fp hash
func (B *Builder) Event_2801() FingerprintEvent {
	return FingerprintEvent{
		2801,
		B.GetByID(2801, B.CollectedFp.ParsedEvents[2801].(string)),
	}
}

func (B *Builder) Event_2805() FingerprintEvent {
	data := B.Stringify([]interface{}{
		[]interface{}{
			277114314453,
			277114314460,
			277114314451,
			357114314456,
			277114314452,
			554228628898,
			57114314443,
			717114314371391,
			554228628897,
			277114314456,
			1108457257862,
			277114314450,
			554228628919,
			277114314460,
			277114314451,
		},
		false,
	})

	return FingerprintEvent{
		2805,
		data,
	}
}

func (B *Builder) Event_3210() FingerprintEvent {
	data := B.Stringify([]interface{}{
		104480525378, // change
		104480525378, // change
		nil,
		nil,
		4294705152,
		true,
		true,
		true,
		nil,
	})

	return FingerprintEvent{
		3210,
		B.GetByID(3210, data),
	}
}

func (B *Builder) Event_3211() FingerprintEvent {
	return FingerprintEvent{
		3211,
		B.Stringify(EncStr("143254600089")),
	}
}

// win func  hash ?
func (B *Builder) Event_3401() FingerprintEvent {
	return FingerprintEvent{
		3401,
		//B.GetByID(3401, B.CollectedFp.ParsedEvents[3401].(string)),
		B.GetByID(3401, hashs[3401]),
		//B.GetByID(3401, utils.RandomHash(20)),
	}
}

func (B *Builder) Event_3403() FingerprintEvent {
	data := B.Stringify([]interface{}{
		[]interface{}{
			[]interface{}{
				fmt.Sprintf("https://newassets.hcaptcha.com/captcha/v1/%s/hcaptcha.js", VERSION),
				0,
				5,
			},
		},
		[]interface{}{
			[]interface{}{
				"*",
				84,
				9,
			},
		},
	})

	return FingerprintEvent{
		3403,
		data,
	}
}

func (B *Builder) Event_3501() FingerprintEvent {
	data := B.Stringify([]interface{}{
		[]interface{}{
			"img:imgs.hcaptcha.com",
			0,
			utils.RandomFloat64(20, 60),
		},
		[]interface{}{
			"navigation:newassets.hcaptcha.com",
			utils.RandomFloat64(10, 50),
			utils.RandomFloat64(10, 50),
			utils.RandomFloat64(10, 50),
		},
		[]interface{}{
			"script:newassets.hcaptcha.com",
			utils.RandomFloat64(5, 50),
			utils.RandomFloat64(10, 50),
		},
		[]interface{}{
			"xmlhttprequest:hcaptcha.com",
			0,
			utils.RandomFloat64(150, 250),
		},
	})

	return FingerprintEvent{
		3501,
		data,
	}
}

func (B *Builder) Event_3502() FingerprintEvent {
	return FingerprintEvent{
		3502,
		fmt.Sprintf("%.14f", utils.RandomFloat64(1, 20)),
	}
}

func (B *Builder) Event_3503() FingerprintEvent {
	return FingerprintEvent{
		3503,
		fmt.Sprintf("%.14f", utils.RandomFloat64(0, 1)),
	}
}

func (B *Builder) Event_3504() FingerprintEvent {
	return FingerprintEvent{
		3504,
		fmt.Sprintf("%.1f", float64(time.Now().UnixNano())/1e6),
	}
}

// disabled
func (B *Builder) Event_3505() FingerprintEvent {
	data := B.Stringify([]interface{}{
		0.09999999403953552,
		27,
	})

	return FingerprintEvent{
		3505,
		B.GetByID(3505, data),
	}
}
