/**
 * Behaviours:
 */

package events

func (B *EventManager) Event_2805() FingerprintEvent {
	return FingerprintEvent{
		2805,
		Stringify(B.Fingerprint.Events["2805"]),
	}
}
