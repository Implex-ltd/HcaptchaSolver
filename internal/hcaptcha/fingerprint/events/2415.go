/**
 * Behaviours:
 */

package events

func (B *EventManager) Event_2415() FingerprintEvent {
	return FingerprintEvent{
		2415,
		Stringify(B.Fingerprint.Events["2415"]),
	}
}
