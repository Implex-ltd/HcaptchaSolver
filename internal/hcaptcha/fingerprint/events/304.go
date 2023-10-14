/**
 * Behaviours:
 */

package events

import "fmt"

func (B *EventManager) Event_304() FingerprintEvent {
	return FingerprintEvent{
		304,
		fmt.Sprintf("%d", B.Fingerprint.Properties.CSS),
	}
}
