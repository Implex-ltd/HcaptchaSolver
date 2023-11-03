/**
 * Behaviours:
 */

package events

func (B *EventManager) Event_2413() FingerprintEvent {
	return FingerprintEvent{
		2413,
		Stringify(B.Fingerprint.Events["2413"]),
	}
}
