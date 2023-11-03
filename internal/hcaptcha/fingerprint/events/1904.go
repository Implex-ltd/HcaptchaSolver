/**
 * Behaviours:
 */

package events

func (B *EventManager) Event_1904() FingerprintEvent {
	return FingerprintEvent{
		1904,
		Stringify(B.Fingerprint.Events["1904"]),
	}
}
