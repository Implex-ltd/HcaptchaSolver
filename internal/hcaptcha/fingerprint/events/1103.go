/**
 * Behaviours:
 */

package events

func (B *EventManager) Event_1103() FingerprintEvent {
	return FingerprintEvent{
		1103,
		Stringify(B.Fingerprint.Hash["1103"]),
	}
}
