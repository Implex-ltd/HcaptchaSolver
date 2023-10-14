/**
 * Behaviours:
 */

package events

func (B *EventManager) Event_2414() FingerprintEvent {
	return FingerprintEvent{
		2414,
		Stringify(B.Fingerprint.Events["2414"]),
	}
}
