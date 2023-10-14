/**
 * Behaviours: https://gist.github.com/nikolahellatrigger/c4d6cf4ddb0ab219c38ddd133dc772eb
 */

package events

import (
	"fmt"
	"log"
	"strings"
)

func (B *EventManager) Event_604() FingerprintEvent {
	event := B.Fingerprint.Events["702"].(map[string]interface{})
	log.Println(B.Fingerprint.Browser.UserAgent)

	return FingerprintEvent{
		604,
		Stringify([]interface{}{
			B.Fingerprint.Browser.AppVersion,
			B.Fingerprint.Browser.UserAgent,
			B.Fingerprint.Browser.DeviceMemory,
			B.Fingerprint.Browser.HardwareConcurrency,
			B.Fingerprint.Browser.Language,
			B.Fingerprint.Browser.Languages,
			B.Fingerprint.Browser.Platform,
			nil,
			[]interface{}{
				fmt.Sprintf("Google Chrome %s", strings.Split(event["BrowserVersion"].(string), ".")[0]),
				"Not=A?Brand 99",
				fmt.Sprintf("Chromium %s", strings.Split(event["BrowserVersion"].(string), "."))[0],
			},
			B.Fingerprint.Browser.Mobile,
			event["OsName"],
			B.Fingerprint.Properties.MIMETypes,
			B.Fingerprint.Properties.Plugins,
			B.Fingerprint.Browser.PDFViewerEnabled,

			B.Fingerprint.Connection.DownlinkMax, // "downlinkMax" in navigator.connection,
			B.Fingerprint.Connection.Rtt,         // null == navigator.connection ? void 0 : navigator.connection.rtt,
			false,                                // navigator.webdriver,
			false,                                // null === (result = window.clientInformation) || void 0 === result ? void 0 : result.webdriver,
			true,                                 // "share" in navigator,

			"[object Keyboard]",
			strings.Contains(strings.ToLower(B.Fingerprint.Browser.UserAgent), "brave"),
			strings.Contains(strings.ToLower(B.Fingerprint.Browser.UserAgent), "duckduckgo"),
		}),
	}
}
