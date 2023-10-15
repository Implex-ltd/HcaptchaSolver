/**
 * Behaviours:
 */

package events

import "github.com/Implex-ltd/hcsolver/internal/utils"

func (B *EventManager) Event_3501() FingerprintEvent {
	return FingerprintEvent{
		3501,
		Stringify([]interface{}{
			/*[]interface{}{
				"img:imgs.hcaptcha.com",
				0,
				utils.RandomFloat64(20, 60),
			},*/
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
		}),
	}
}
