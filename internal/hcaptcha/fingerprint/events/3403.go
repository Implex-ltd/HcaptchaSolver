/**
 * Behaviours:
 */

package events

import "fmt"

func (B *EventManager) Event_3403() FingerprintEvent {
	return FingerprintEvent{
		3403,
		Stringify([]interface{}{
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
		}),
	}
}
