/**
 * Behaviours:
 */

package events

func (B *EventManager) Event_2416() FingerprintEvent {
	return FingerprintEvent{
		2416,
		Stringify(B.Fingerprint.Events["2416"]),
	}
}
