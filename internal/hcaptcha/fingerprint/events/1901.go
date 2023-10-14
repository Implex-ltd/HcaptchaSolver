/**
 * Behaviours:
 */

package events

func (B *EventManager) Event_1901() FingerprintEvent {
	return FingerprintEvent{
		1901,
		Stringify(B.Fingerprint.Hash["1901"]),
	}
}
