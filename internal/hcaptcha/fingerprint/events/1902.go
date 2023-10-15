/**
 * Behaviours:
 */

package events

func (B *EventManager) Event_1902() FingerprintEvent {
	return FingerprintEvent{
		1902,
		Stringify(B.Fingerprint.Events["1902"]),
	}
}
