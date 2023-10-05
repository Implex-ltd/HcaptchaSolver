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

func (B *Builder) Stringify(data any) string {
	jsonString, err := json.Marshal(data)
	if err != nil {
		return ""
	}

	return string(jsonString)
}

func (B *Builder) Event_0() FingerprintEvent {
	return FingerprintEvent{
		0,
		B.CollectedFp.ParsedEvents[0].(string),
	}
}

func (B *Builder) Event_3() FingerprintEvent {
	return FingerprintEvent{
		3,
		//"32575",
		B.CollectedFp.ParsedEvents[3].(string),
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

	/*if fp := B.CollectedFp.ParsedEvents[107]; fp != nil {
		data = B.CollectedFp.ParsedEvents[107].(string)
	}*/

	return FingerprintEvent{
		107,
		data,
	}
}

func (B *Builder) Event_201() FingerprintEvent {
	/*data := "13122422878918113034"

	if fp := B.CollectedFp.ParsedEvents[201]; fp != nil {
		data = fp.(string)
	}*/

	return FingerprintEvent{
		201,
		"4226317358175830201",
		//data,
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

	if fp := B.CollectedFp.ParsedEvents[211]; fp != nil {
		data = fp.(string)
	}

	return FingerprintEvent{
		211,
		data,
	}
}

func (B *Builder) Event_301() FingerprintEvent {
	/*data := "8383473043360077444"

	if fp := B.CollectedFp.ParsedEvents[301]; fp != nil {
		data = fp.(string)
	}*/

	return FingerprintEvent{
		301,
		"8383473043360077444",
		//data,
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

	if fp := B.CollectedFp.ParsedEvents[302]; fp != nil {
		data = fp.(string)
	}

	return FingerprintEvent{
		302,
		data,
	}
}

func (B *Builder) Event_303() FingerprintEvent {
	data := B.Stringify([]interface{}{
		"Arial",
		"\"Segoe UI\"",
	})

	if fp := B.CollectedFp.ParsedEvents[303]; fp != nil {
		data = fp.(string)
	}

	return FingerprintEvent{
		303,
		data,
	}
}

func (B *Builder) Event_304() FingerprintEvent {
	/*data := "623"

	if fp := B.CollectedFp.ParsedEvents[304]; fp != nil {
		data = fp.(string)
	}*/

	return FingerprintEvent{
		304,
		"623",
	}
}

func (B *Builder) Event_401() FingerprintEvent {
	/*data := "2400869836852708862"

	if fp := B.CollectedFp.ParsedEvents[401]; fp != nil {
		data = fp.(string)
	}*/

	return FingerprintEvent{
		401,
		"2400869836852708862",
		//data,
	}
}

func (B *Builder) Event_402() FingerprintEvent {
	data := "1109"

	/*if fp := B.CollectedFp.ParsedEvents[402]; fp != nil {
		data = fp.(string)
	}*/

	return FingerprintEvent{
		402,
		data,
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
	/*data := "15584660433093862032"

	if fp := B.CollectedFp.ParsedEvents[412]; fp != nil {
		data = fp.(string)
	}*/

	return FingerprintEvent{
		412,
		"15584660433093862032",
		//data,
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

	/*if fp := B.CollectedFp.ParsedEvents[604]; fp != nil {
		data = fp.(string)
	}*/

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

	/*if fp := B.CollectedFp.ParsedEvents[702]; fp != nil {
		data = fp.(string)
	}*/

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

	if fp := B.CollectedFp.ParsedEvents[803]; fp != nil {
		data = fp.(string)
	}

	return FingerprintEvent{
		803,
		data,
	}
}

func (B *Builder) Event_901() FingerprintEvent {
	/*data := "16132118391739044799"

	if fp := B.CollectedFp.ParsedEvents[901]; fp != nil {
		data = fp.(string)
	}*/

	return FingerprintEvent{
		901,
		//data,
		"135869055876678538",
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

	/*if fp := B.CollectedFp.ParsedEvents[905]; fp != nil {
		data = fp.(string)
	}*/

	return FingerprintEvent{
		905,
		data,
	}
}

func (B *Builder) Event_1101() FingerprintEvent {
	return FingerprintEvent{
		1101,
		"2028058230851665858",
		//B.CollectedFp.ParsedEvents[1101].(string), //"1308681860871815407",
	}
}

func (B *Builder) Event_1103() FingerprintEvent {
	return FingerprintEvent{
		1103,
		"4932383211497360507",
		//B.CollectedFp.ParsedEvents[1103].(string), //"4932383211497360507",
	}
}

func (B *Builder) Event_1105() FingerprintEvent {
	return FingerprintEvent{
		1105,
		"17157476241021694346",
		//B.CollectedFp.ParsedEvents[1105].(string), //"17157476241021694346",
	}
}

func (B *Builder) Event_1107() FingerprintEvent {
	x := utils.RandomNumber(15, 250)

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

	if fp := B.CollectedFp.ParsedEvents[1107]; fp != nil {
		data = fp.(string)
	}

	return FingerprintEvent{
		1107,
		data,
	}
}

func (B *Builder) Event_1302() FingerprintEvent {
	data := B.Stringify([]interface{}{
		0, 1, 2, 3, 4,
	})

	if fp := B.CollectedFp.ParsedEvents[1302]; fp != nil {
		data = fp.(string)
	}

	return FingerprintEvent{
		1302,
		data,
	}
}

func (B *Builder) Event_1401() FingerprintEvent {
	data := "\"Europe/Paris\""

	if fp := B.CollectedFp.ParsedEvents[1401]; fp != nil {
		data = fp.(string)
	}

	return FingerprintEvent{
		1401,
		data,
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

	if fp := B.CollectedFp.ParsedEvents[1402]; fp != nil {
		data = fp.(string)
	}

	return FingerprintEvent{
		1402,
		data,
	}
}

func (B *Builder) Event_1403() FingerprintEvent {
	/*data := []interface{}{
		"DzBXcMZkmlelYzPGD=Yk",
		"1",
		"3",
		"TDEDEVJMPMMFH",
	}*/

	data := B.CollectedFp.ParsedEvents[1403]

	return FingerprintEvent{
		1403,
		string(data.(string)),
	}
}

func (B *Builder) Event_1901() FingerprintEvent {
	return FingerprintEvent{
		1901,
		"15307345790125003576",
		//B.CollectedFp.ParsedEvents[1901].(string), //"15307345790125003576",
	}
}

func (B *Builder) Event_1902() FingerprintEvent {
	return FingerprintEvent{
		1902,
		B.CollectedFp.ParsedEvents[1902].(string), //"57",
	}
}

func (B *Builder) Event_1904() FingerprintEvent {
	return FingerprintEvent{
		1904,
		B.CollectedFp.ParsedEvents[1904].(string),
	}
}

func (B *Builder) Event_2001() FingerprintEvent {
	data := "13414760775080815217"

	if fp := B.CollectedFp.ParsedEvents[2001]; fp != nil {
		data = fp.(string)
	}

	return FingerprintEvent{
		2001,
		data,
	}
}

func (B *Builder) Event_2002() FingerprintEvent {
	data := B.Stringify([]interface{}{
		"denied",
		"denied",
	})

	if fp := B.CollectedFp.ParsedEvents[2002]; fp != nil {
		data = fp.(string)
	}

	return FingerprintEvent{
		2002,
		data,
	}
}

func (B *Builder) Event_2401() FingerprintEvent {
	/*data := utils.RandomHash(18)

	if fp := B.CollectedFp.ParsedEvents[2401]; fp != nil {
		data = fp.(string)
	}*/

	return FingerprintEvent{
		2401,
		//data,
		"13854066076378004045",
	}
}

func (B *Builder) Event_2402() FingerprintEvent {
	/*data := []interface{}{
		B.Profile.Misc.Vendor,
		B.Profile.Misc.Renderer,
	}*/

	data := B.CollectedFp.ParsedEvents[2402]

	return FingerprintEvent{
		2402,
		string(data.(string)),
	}
}

func (B *Builder) Event_2403() FingerprintEvent {
	/*data := []interface{}{
		B.Profile.Misc.Vendor,
		B.Profile.Misc.Renderer,
	}*/

	data := B.CollectedFp.ParsedEvents[2403]

	return FingerprintEvent{
		2403,
		string(data.(string)),
	}
}

func (B *Builder) Event_2407() FingerprintEvent {
	return FingerprintEvent{
		2407,
		"13177607191192652685",
		//B.CollectedFp.ParsedEvents[2407].(string), //"13177607191192652685",
	}
}

func (B *Builder) Event_2408() FingerprintEvent {
	return FingerprintEvent{
		2408,
		"true",
	}
}

func (B *Builder) Event_2409() FingerprintEvent {
	/*data := []interface{}{
		2147483647,
		2147483647,
		4294967294,
	}*/
	data := B.CollectedFp.ParsedEvents[2409]

	return FingerprintEvent{
		2409,
		string(data.(string)), //B.Stringify(data),
	}
}

func (B *Builder) Event_2410() FingerprintEvent {
	/*data := []interface{}{
		16,
		1024,
		4096,
		7,
		12,
		120,
		[]interface{}{23, 127, 127},
	}*/
	data := B.CollectedFp.ParsedEvents[2410]

	return FingerprintEvent{
		2410,
		string(data.(string)), //B.Stringify(data),
	}
}

func (B *Builder) Event_2411() FingerprintEvent {
	/*data := []interface{}{
		32767,
		32767,
		16384,
		8,
		8,
		8,
	}*/
	data := B.CollectedFp.ParsedEvents[2411]

	return FingerprintEvent{
		2411,
		string(data.(string)), //B.Stringify(data),
	}
}

func (B *Builder) Event_2412() FingerprintEvent {
	/*data := []interface{}{
		1,
		1024,
		1,
		1,
		4,
	}*/
	data := B.CollectedFp.ParsedEvents[2412]

	return FingerprintEvent{
		2412,
		string(data.(string)), //B.Stringify(data),
	}
}

func (B *Builder) Event_2413() FingerprintEvent {
	/*data := []interface{}{
		2147483647,
		2147483647,
		2147483647,
		2147483647,
	}*/

	data := B.CollectedFp.ParsedEvents[2413]

	return FingerprintEvent{
		2413,
		string(data.(string)), //B.Stringify(data),
	}
}

func (B *Builder) Event_2414() FingerprintEvent {
	/*data := []interface{}{
		16384,
		32,
		16384,
		2048,
		2,
		2048,
	}*/
	data := B.CollectedFp.ParsedEvents[2414]

	return FingerprintEvent{
		2414,
		string(data.(string)), //B.Stringify(data),
	}
}

func (B *Builder) Event_2415() FingerprintEvent {
	data := B.Stringify([]interface{}{
		4,
		120,
		4,
	})

	if fp := B.CollectedFp.ParsedEvents[2415]; fp == nil {
		data = fp.(string)
	}

	return FingerprintEvent{
		2415,
		data,
	}
}

func (B *Builder) Event_2416() FingerprintEvent {
	/*data := []interface{}{
		24,
		24,
		65536,
		212988,
		200704,
	}*/
	data := B.CollectedFp.ParsedEvents[2416]

	return FingerprintEvent{
		2416,
		string(data.(string)), //B.Stringify(data),
	}
}

func (B *Builder) Event_2417() FingerprintEvent {
	/*data := []interface{}{
		16,
		4095,
		30,
		16,
		16380,
		120,
		12,
		120,
		[]interface{}{23, 127, 127},
	}*/
	data := B.CollectedFp.ParsedEvents[2417]

	return FingerprintEvent{
		2417,
		string(data.(string)), //B.Stringify(data),
	}
}

func (B *Builder) Event_2420() FingerprintEvent {
	data := B.CollectedFp.ParsedEvents[2420]

	return FingerprintEvent{
		2420,
		string(data.(string)),
	}
}

func (B *Builder) Event_2801() FingerprintEvent {
	return FingerprintEvent{
		2801,
		"4631229088072584217",
		//B.CollectedFp.ParsedEvents[2801].(string), //"4631229088072584217",
	}
}

func (B *Builder) Event_2805() FingerprintEvent {
	/*data := []interface{}{
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
	}*/
	data := B.CollectedFp.ParsedEvents[2805]

	return FingerprintEvent{
		2805,
		string(data.(string)), //B.Stringify(data),
	}
}

func (B *Builder) Event_3210() FingerprintEvent {
	data := B.Stringify([]interface{}{
		143254600089,
		143254600089,
		nil,
		nil,
		4294705152,
		true,
		true,
		true,
		nil,
	})

	if fp := B.CollectedFp.ParsedEvents[3210]; fp != nil {
		data = fp.(string)
	}

	return FingerprintEvent{
		3210,
		data,
	}
}

func (B *Builder) Event_3211() FingerprintEvent {
	data := B.Stringify([]interface{}{
		"mYudn2AdmWgTOxQZ",
		"11",
		"3",
		"ZKRKYBDWNDUMW",
	})

	if fp := B.CollectedFp.ParsedEvents[3211]; fp != nil {
		data = fp.(string)
	}

	return FingerprintEvent{
		3211,
		data,
	}
}

func (B *Builder) Event_3401() FingerprintEvent {
	return FingerprintEvent{
		3401,
		B.CollectedFp.ParsedEvents[3401].(string),
		//utils.RandomHash(20), //"4226317358175830201",
	}
}

func (B *Builder) Event_3403() FingerprintEvent {
	data := []interface{}{
		[]interface{}{
			[]interface{}{
				fmt.Sprintf("https://newassets.hcaptcha.com/captcha/v1/%s/hcaptcha.js", B.HcapVersion),
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
	}

	return FingerprintEvent{
		3403,
		B.Stringify(data),
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

	/*if fp := B.CollectedFp.ParsedEvents[3501]; fp != nil {
		data = fp.(string)
	}*/

	return FingerprintEvent{
		3501,
		data,
	}
}

func (B *Builder) Event_3502() FingerprintEvent {
	return FingerprintEvent{
		3502,
		fmt.Sprintf("%.14f", utils.RandomFloat64(20, 50)),
		//B.CollectedFp.ParsedEvents[3502].(string),
		//"6.199999999254942",
	}
}

func (B *Builder) Event_3503() FingerprintEvent {
	return FingerprintEvent{
		3503,
		fmt.Sprintf("%.14f", utils.RandomFloat64(10, 20)),
		//B.CollectedFp.ParsedEvents[3503].(string),
		//"0.15000000223517418",
	}
}

func (B *Builder) Event_3504() FingerprintEvent {
	return FingerprintEvent{
		3504,
		fmt.Sprintf("%.1f", float64(time.Now().UnixNano())/1e6),
	}
}

// not used
func (B *Builder) Event_3505() FingerprintEvent {
	data := B.Stringify([]interface{}{
		0.09999999403953552,
		27,
	})

	if fp := B.CollectedFp.ParsedEvents[3505]; fp != nil {
		data = fp.(string)
	}

	return FingerprintEvent{
		3505,
		data,
	}
}
