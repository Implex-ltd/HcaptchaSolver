/**
 * Behaviours:
 */

package events

func (B *EventManager) Event_2409() FingerprintEvent {
	return FingerprintEvent{
		2409,
		Stringify(B.Fingerprint.Events["2409"]),
	}
}
