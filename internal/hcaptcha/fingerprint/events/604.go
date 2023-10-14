/**
 * Behaviours: https://gist.github.com/nikolahellatrigger/c4d6cf4ddb0ab219c38ddd133dc772eb
 */

package events

import (
	"fmt"
	"strings"
)

func (B *EventManager) Event_604() FingerprintEvent {
	event := B.Fingerprint.Events["702"].(map[string]interface{})

	return FingerprintEvent{
		604,
		Stringify([]interface{}{
			B.Fingerprint.Browser.AppVersion,
			B.UserAgent,
			B.Fingerprint.Browser.DeviceMemory,
			B.Fingerprint.Browser.HardwareConcurrency,
			B.Fingerprint.Browser.Language,
			B.Fingerprint.Browser.Languages,
			B.Fingerprint.Browser.Platform,
			nil,
			[]interface{}{
				fmt.Sprintf("Google Chrome %s", strings.Split(event["BrowserVersion"].(string), ".")[0]),
				"Not;A=Brand 8",
				fmt.Sprintf("Chromium %s", strings.Split(event["BrowserVersion"].(string), "."))[0],
			},
			B.Fingerprint.Browser.Mobile,
			event["OsName"],
			B.Fingerprint.Properties.MIMETypes,
			B.Fingerprint.Properties.Plugins,
			B.Fingerprint.Browser.PDFViewerEnabled,
			false,
			50,
			false,
			false,
			true,
			"[object Keyboard]",
			strings.Contains(strings.ToLower(B.UserAgent), "brave"),
			strings.Contains(strings.ToLower(B.UserAgent), "duckduckgo"),
		}),
	}
}
