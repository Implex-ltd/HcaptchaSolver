/**
 * Behaviours:
 */

package events

func (B *EventManager) Event_803() FingerprintEvent {
	return FingerprintEvent{
		803,
		Stringify(B.Fingerprint.Events["803"]),
	}
}
