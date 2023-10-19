/**
 * Behaviours:
 */

package events

func (B *EventManager) Event_905() FingerprintEvent {
	return FingerprintEvent{
		905,
		Stringify(B.Fingerprint.Events["905"]),
	}
}
