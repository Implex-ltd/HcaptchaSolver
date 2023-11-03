/**
 * Behaviours:
 */

package events

import "fmt"

func (B *EventManager) Event_1302() FingerprintEvent {
	fmt.Println(B.Fingerprint.Events["1302"])
	return FingerprintEvent{
		1302,
		Stringify(B.Fingerprint.Events["1302"]),
	}
}
