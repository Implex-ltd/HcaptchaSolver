/**
 * Behaviours:
 */

package events

import "fmt"

func (B *EventManager) Event_402() FingerprintEvent {
	return FingerprintEvent{
		402,
		fmt.Sprintf("%v", B.Fingerprint.Properties.WindowFunctions),
	}
}
