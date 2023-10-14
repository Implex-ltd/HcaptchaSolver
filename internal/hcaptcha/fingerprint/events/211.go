/**
 * Behaviours:
 */

package events

func (B *EventManager) Event_211() FingerprintEvent {
	return FingerprintEvent{
		211,
		Stringify(B.Fingerprint.Events["211"]),
	}
}
