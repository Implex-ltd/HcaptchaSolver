/**
 * Behaviours:
 */

package events

func (B *EventManager) Event_3210() FingerprintEvent {
	return FingerprintEvent{
		3210,
		Stringify(B.Fingerprint.Events["3210"]),
	}
}
