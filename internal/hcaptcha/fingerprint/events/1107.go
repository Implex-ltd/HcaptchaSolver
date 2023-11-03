/**
 * Behaviours:
 */

package events


func (B *EventManager) Event_1107() FingerprintEvent {
	return FingerprintEvent{
		1107,
		Stringify(B.Fingerprint.Events["201"]),
	}
}
