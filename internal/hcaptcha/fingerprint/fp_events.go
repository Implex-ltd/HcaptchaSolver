package fingerprint

import (
	"encoding/json"
	"fmt"
	"strings"
)

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
		//"73.19999999552965",
	}
}

func (B *Builder) Event_107() FingerprintEvent {
	/*data := []interface{}{
		B.Profile.Screen.Width,
		B.Profile.Screen.Height,
		B.Profile.Screen.AvailWidth,
		B.Profile.Screen.AvailHeight,
		B.Profile.Screen.ColorDepth,
		B.Profile.Screen.PixelDepth,
		false,
		0,
		1,
		1979, // AvailLeft ?
		1399, // AvailTop  ?
		true,
		true,
		true,
		false,
	}*/

	data := B.CollectedFp.ParsedEvents[107].(string)

	return FingerprintEvent{
		107,
		data, //B.Stringify(data),
	}
}

func (B *Builder) Event_201() FingerprintEvent {
	return FingerprintEvent{
		201,
		B.CollectedFp.ParsedEvents[201].(string),
		//"13122422878918113034",
	}
}

func (B *Builder) Event_211() FingerprintEvent {
	/*data := []interface{}{
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
	}*/

	data := B.CollectedFp.ParsedEvents[211]

	return FingerprintEvent{
		211,
		B.Stringify(data),
	}
}

func (B *Builder) Event_301() FingerprintEvent {
	return FingerprintEvent{
		301,
		"8383473043360077444",
	}
}

func (B *Builder) Event_302() FingerprintEvent {
	data := []interface{}{
		0,
		1,
		2,
		3,
		4,
	}

	return FingerprintEvent{
		302,
		B.Stringify(data),
	}
}

func (B *Builder) Event_303() FingerprintEvent {
	data := []interface{}{
		"Arial",
		"\"Segoe UI\"",
	}

	return FingerprintEvent{
		303,
		B.Stringify(data),
	}
}

func (B *Builder) Event_304() FingerprintEvent {
	return FingerprintEvent{
		304,
		"623",
	}
}

func (B *Builder) Event_401() FingerprintEvent {
	return FingerprintEvent{
		401,
		"\"Europe/Paris\"",
	}
}

func (B *Builder) Event_402() FingerprintEvent {
	return FingerprintEvent{
		402,
		B.CollectedFp.ParsedEvents[402].(string), //"1116",
	}
}

func (B *Builder) Event_407() FingerprintEvent {
	/*data := []interface{}{
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
	}*/

	data := B.CollectedFp.ParsedEvents[407]

	return FingerprintEvent{
		407,
		string(data.(string)), //B.Stringify(data),
	}
}

func (B *Builder) Event_412() FingerprintEvent {
	return FingerprintEvent{
		412,
		B.CollectedFp.ParsedEvents[412].(string), //"15584660433093862032",
	}
}

func (B *Builder) Event_604() FingerprintEvent {
	data := []interface{}{
		strings.Split(B.UserAgent, "Mozilla/")[1],
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
		false,
		B.Profile.Misc.Os,
		2,
		5,
		true,
		false,
		0,
		false,
		false,
		true,
		"[object Keyboard]",
		false,
		false,
	}

	return FingerprintEvent{
		604,
		B.Stringify(data),
	}
}

func (B *Builder) Event_702() FingerprintEvent {
	data := []interface{}{
		B.Profile.Misc.Os,
		"15.0.0",
		nil,
		B.Profile.Misc.CPU,
		B.Profile.Misc.Arch,
		B.Profile.Misc.BrowserAppVersion,
	}

	return FingerprintEvent{
		702,
		B.Stringify(data),
	}
}

func (B *Builder) Event_803() FingerprintEvent {
	data := []interface{}{
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
	}

	return FingerprintEvent{
		803,
		B.Stringify(data),
	}
}

func (B *Builder) Event_901() FingerprintEvent {
	return FingerprintEvent{
		901,
		B.CollectedFp.ParsedEvents[901].(string), //"16132118391739044799",
	}
}

func (B *Builder) Event_905() FingerprintEvent {
	data := []interface{}{
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
	}

	return FingerprintEvent{
		905,
		B.Stringify(data),
	}
}

func (B *Builder) Event_1101() FingerprintEvent {
	return FingerprintEvent{
		1101,
		B.CollectedFp.ParsedEvents[1101].(string), //"1308681860871815407",
	}
}

func (B *Builder) Event_1103() FingerprintEvent {
	return FingerprintEvent{
		1103,
		B.CollectedFp.ParsedEvents[1103].(string), //"4932383211497360507",
	}
}

func (B *Builder) Event_1105() FingerprintEvent {
	return FingerprintEvent{
		1105,
		B.CollectedFp.ParsedEvents[1105].(string), //"17157476241021694346",
	}
}

func (B *Builder) Event_1107() FingerprintEvent {
	/*data := []interface{}{
		[]interface{}{
			29,
			[]interface{}{
				29,
				29,
				29,
				255,
				29,
				29,
				29,
				255,
				29,
				29,
				29,
				255,
				29,
				29,
				29,
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
	}*/

	data := B.CollectedFp.ParsedEvents[1107]

	return FingerprintEvent{
		1107,
		string(data.(string)), //B.Stringify(data),
	}
}

func (B *Builder) Event_1401() FingerprintEvent {
	return FingerprintEvent{
		1401,
		"\"Europe/Paris\"",
	}
}

func (B *Builder) Event_1402() FingerprintEvent {
	data := []interface{}{
		"Europe/Paris",
		-60,
		-60,
		-3203647761000,
		"heure d’été d’Europe centrale",
		"fr",
	}

	return FingerprintEvent{
		1402,
		B.Stringify(data),
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
		string(data.(string)), //B.Stringify(data),
	}
}

func (B *Builder) Event_1901() FingerprintEvent {
	return FingerprintEvent{
		1901,
		B.CollectedFp.ParsedEvents[1901].(string), //"15307345790125003576",
	}
}

func (B *Builder) Event_1902() FingerprintEvent {
	return FingerprintEvent{
		1902,
		B.CollectedFp.ParsedEvents[1902].(string), //"57",
	}
}

func (B *Builder) Event_2001() FingerprintEvent {
	return FingerprintEvent{
		2001,
		"13414760775080815217",
	}
}

func (B *Builder) Event_2002() FingerprintEvent {
	data := []interface{}{
		"denied",
		"denied",
	}

	return FingerprintEvent{
		2002,
		B.Stringify(data),
	}
}

func (B *Builder) Event_2401() FingerprintEvent {
	return FingerprintEvent{
		2401,
		"17670611538850778206",
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
		string(data.(string)), //B.Stringify(data),
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
		string(data.(string)), //B.Stringify(data),
	}
}

func (B *Builder) Event_2407() FingerprintEvent {
	return FingerprintEvent{
		2407,
		B.CollectedFp.ParsedEvents[2407].(string), //"13177607191192652685",
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
	/*data := []interface{}{
		4,
		120,
		4,
	}*/
	data := B.CollectedFp.ParsedEvents[2415]

	return FingerprintEvent{
		2415,
		string(data.(string)), //B.Stringify(data),
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

func (B *Builder) Event_2801() FingerprintEvent {
	return FingerprintEvent{
		2801,
		B.CollectedFp.ParsedEvents[2801].(string), //"4631229088072584217",
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
	/*data := []interface{}{
		143254600089,
		143254600089,
		nil,
		nil,
		4294705152,
		true,
		true,
		true,
		nil,
	}*/
	data := B.CollectedFp.ParsedEvents[3210]

	return FingerprintEvent{
		3210,
		string(data.(string)), //B.Stringify(data),
	}
}

func (B *Builder) Event_3211() FingerprintEvent {
	/*data := []interface{}{
		"mYudn2AdmWgTOxQZ",
		"11",
		"3",
		"ZKRKYBDWNDUMW",
	}*/
	data := B.CollectedFp.ParsedEvents[3211]

	return FingerprintEvent{
		3211,
		string(data.(string)), //B.Stringify(data),
	}
}

func (B *Builder) Event_3401() FingerprintEvent {
	return FingerprintEvent{
		3401,
		B.CollectedFp.ParsedEvents[3401].(string), //"4226317358175830201",
	}
}

func (B *Builder) Event_3403() FingerprintEvent {
	/*data := []interface{}{
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
	}*/
	data := B.CollectedFp.ParsedEvents[3403]

	return FingerprintEvent{
		3403,
		string(data.(string)), //B.Stringify(data),
	}
}

func (B *Builder) Event_3501() FingerprintEvent {
	/*data := []interface{}{
		[]interface{}{
			"navigation:newassets.hcaptcha.com",
			0,
			8.5,
		},
		[]interface{}{
			"script:newassets.hcaptcha.com",
			0.15000000223517418,
			3.649999998509884,
		},
		[]interface{}{
			"xmlhttprequest:hcaptcha.com",
			0,
			40.399999998509884,
		},
	}*/
	data := B.CollectedFp.ParsedEvents[3501]

	return FingerprintEvent{
		3501,
		string(data.(string)), //B.Stringify(data),
	}
}

func (B *Builder) Event_3502() FingerprintEvent {
	return FingerprintEvent{
		3502,
		B.CollectedFp.ParsedEvents[3502].(string), //"6.199999999254942",
	}
}

func (B *Builder) Event_3503() FingerprintEvent {
	return FingerprintEvent{
		3503,
		B.CollectedFp.ParsedEvents[3503].(string), //"0.15000000223517418",
	}
}

func (B *Builder) Event_3504() FingerprintEvent {
	return FingerprintEvent{
		3504,
		B.CollectedFp.ParsedEvents[3504].(string), //"1696117557863.2",
	}
}

func (B *Builder) Event_3505() FingerprintEvent {
	/*data := []interface{}{
		0.09999999403953552,
		27,
	}*/
	data := B.CollectedFp.ParsedEvents[3505]

	return FingerprintEvent{
		3505,
		string(data.(string)), //B.Stringify(data),
	}
}
