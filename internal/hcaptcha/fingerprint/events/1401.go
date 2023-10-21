/**
 * Behaviours:
 */

package events

import "fmt"

func (B *EventManager) Event_1401() FingerprintEvent {
	return FingerprintEvent{
		1401,
		Stringify(fmt.Sprintf("%s", B.Fingerprint.Timezone[0].(string))),
	}
}
