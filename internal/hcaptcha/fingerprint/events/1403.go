/**
 * Behaviours:
 */

package events

import "fmt"

func (B *EventManager) Event_1403() FingerprintEvent {
	return FingerprintEvent{
		1403,
		Stringify(EncStr(fmt.Sprintf("%s", B.Fingerprint.Timezone[0].(string)))),
	}
}
