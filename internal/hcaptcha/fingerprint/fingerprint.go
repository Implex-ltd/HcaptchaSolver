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
	WASM              = "1.40.10"
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

	//hash, _ := RandHash([]byte("Raven,alert,atob,blur,btoa,caches,cancelAnimationFrame,cancelIdleCallback,captureEvents,chrome,clearInterval,clearTimeout,clientInformation,close,closed,confirm,cookieStore,createImageBitmap,credentialless,crossOriginIsolated,crypto,customElements,devicePixelRatio,document,documentPictureInPicture,external,fence,fetch,find,focus,frameElement,frames,getComputedStyle,getScreenDetails,getSelection,history,indexedDB,innerHeight,innerWidth,isSecureContext,launchQueue,length,localStorage,location,locationbar,matchMedia,menubar,moveBy,moveTo,name,navigation,navigator,onabort,onafterprint,onanimationend,onanimationiteration,onanimationstart,onappinstalled,onauxclick,onbeforeinput,onbeforeinstallprompt,onbeforematch,onbeforeprint,onbeforetoggle,onbeforeunload,onbeforexrselect,onblur,oncancel,oncanplay,oncanplaythrough,onchange,onclick,onclose,oncontentvisibilityautostatechange,oncontextlost,oncontextmenu,oncontextrestored,oncuechange,ondblclick,ondevicemotion,ondeviceorientation,ondeviceorientationabsolute,ondrag,ondragend,ondragenter,ondragleave,ondragover,ondragstart,ondrop,ondurationchange,onemptied,onended,onerror,onfocus,onformdata,ongotpointercapture,onhashchange,oninput,oninvalid,onkeydown,onkeypress,onkeyup,onlanguagechange,onload,onloadeddata,onloadedmetadata,onloadstart,onlostpointercapture,onmessage,onmessageerror,onmousedown,onmouseenter,onmouseleave,onmousemove,onmouseout,onmouseover,onmouseup,onmousewheel,onoffline,ononline,onpagehide,onpageshow,onpause,onplay,onplaying,onpointercancel,onpointerdown,onpointerenter,onpointerleave,onpointermove,onpointerout,onpointerover,onpointerrawupdate,onpointerup,onpopstate,onprogress,onratechange,onrejectionhandled,onreset,onresize,onscroll,onscrollend,onsearch,onsecuritypolicyviolation,onseeked,onseeking,onselect,onselectionchange,onselectstart,onslotchange,onstalled,onstorage,onsubmit,onsuspend,ontimeupdate,ontoggle,ontransitioncancel,ontransitionend,ontransitionrun,ontransitionstart,onunhandledrejection,onunload,onvolumechange,onwaiting,onwebkitanimationend,onwebkitanimationiteration,onwebkitanimationstart,onwebkittransitionend,onwheel,open,openDatabase,opener,origin,originAgentCluster,outerHeight,outerWidth,pageXOffset,pageYOffset,parent,performance,personalbar,postMessage,print,prompt,queryLocalFonts,queueMicrotask,releaseEvents,reportError,requestAnimationFrame,requestIdleCallback,resizeBy,resizeTo,scheduler,screen,screenLeft,screenTop,screenX,screenY,scroll,scrollBy,scrollTo,scrollX,scrollY,scrollbars,self,sessionStorage,setInterval,setTimeout,sharedStorage,showDirectoryPicker,showOpenFilePicker,showSaveFilePicker,speechSynthesis,status,statusbar,stop,structuredClone,styleMedia,toolbar,top,trustedTypes,visualViewport,webkitCancelAnimationFrame,webkitRequestAnimationFrame,webkitRequestFileSystem,webkitResolveLocalFileSystemURL,window"))
	//hash, _ := RandHash([]byte(utils.RandomString(80)))

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
			PluginsUndefined:            false,
		},
		Hash: Hash{
			Performance:   "2372271609278715010",//utils.RandomHash(19), //HashString([]byte("navigation:newassets.hcaptcha.comscript:newassets.hcaptcha.comxmlhttprequest:hcaptcha.com")),
			Canvas:        utils.RandomHash(19),
			WebGL:         "-1",
			WebRTC:        "-1",
			Audio:         "-1",
			ParrentWindow: "13419404057851147340", //utils.RandomHash(19),
			CommonKeys:    2125906006,
		},
		Misc: Misc{
			UniqueKeys:      "__localeData__,regeneratorRuntime,0,__BILLING_STANDALONE__,webpackChunkdiscord_app,platform,__SECRET_EMOTION__,__SENTRY__,hcaptcha,hcaptchaOnLoad,__timingFunction,DiscordErrors,clearImmediate,__OVERLAY__,grecaptcha,GLOBAL_ENV,setImmediate,1,IntlPolyfill,__DISCORD_WINDOW_ID",
			InvUniqueKeys:   "__wdata,sessionStorage,localStorage,hsw,_sharedLibs",
			CommonKeysTails: "chrome,fence,caches,cookieStore,ondevicemotion,ondeviceorientation,ondeviceorientationabsolute,launchQueue,sharedStorage,documentPictureInPicture,onbeforematch,getScreenDetails,openDatabase,queryLocalFonts,showDirectoryPicker,showOpenFilePicker,showSaveFilePicker,originAgentCluster,credentialless,speechSynthesis,oncontentvisibilityautostatechange,onscrollend,webkitRequestFileSystem,webkitResolveLocalFileSystemURL,Raven", //"__wdata,image_label_binary,_sharedLibs,text_free_entry,sessionStorage,hsw,localStorage",
		},
	}

	B.Profile = &p
	return &p, nil
}

func (B *Builder) Build(jwt string, isSubmit bool) (*Ndata, error) {
	/*Profile, err := B.GenerateProfile()
	if err != nil {
		return nil, err
	}*/

	token, err := ParseJWT(jwt)
	if err != nil {
		return nil, err
	}

	stamp, err := GetStamp(uint(token.Difficuly), token.PowData)
	if err != nil {
		return nil, fmt.Errorf("pow error")
	}

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
			// 0.13396570613838654
			utils.RandomFloat64Precission(0, 1, 10000000000000000.0),
			//0.0004674382952947198,
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
			Chrome:                    strings.Contains(B.Manager.Fingerprint.Browser.UserAgent, "Chrome"),
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
			CommonKeysHash:  B.Profile.Hash.CommonKeys,
			CommonKeysTail:  B.Profile.Misc.CommonKeysTails,
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
		Stamp:                       stamp,
		Href:                        B.Manager.Href,
		Errs: Errs{
			List: []string{},
		},
		StackData: []string{},
		Perf: [][]any{
			{
				1,
				float64(utils.RandomNumber(5, 100)),
			},
			{
				2,
				float64(utils.RandomNumber(20, 300)),
			},
			{
				3,
				0.0, //int64(utils.RandomNumber(0, 5)),
			},
		},
	}

	if isSubmit {
		N.StackData = []string{
			"new Promise (<anonymous>)",
		}
		b, err := json.Marshal(N)
		if err != nil {
			panic(err)
		}

		_, rand_int := RandHash(b)
		N.Rand = append(N.Rand, rand_int)
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
