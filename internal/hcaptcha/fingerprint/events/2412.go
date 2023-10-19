/**
 * Behaviours:
 */

package events

func (B *EventManager) Event_2412() FingerprintEvent {
	return FingerprintEvent{
		2412,
		Stringify(B.Fingerprint.Events["2412"]),
	}
}
