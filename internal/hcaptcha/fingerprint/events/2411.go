/**
 * Behaviours:
 */

package events

func (B *EventManager) Event_2411() FingerprintEvent {
	return FingerprintEvent{
		2411,
		Stringify(B.Fingerprint.Events["2411"]),
	}
}
