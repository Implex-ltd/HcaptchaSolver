/**
 * Behaviours:
 */

package events

func (B *EventManager) Event_2410() FingerprintEvent {
	return FingerprintEvent{
		2410,
		Stringify(B.Fingerprint.Events["2410"]),
	}
}
