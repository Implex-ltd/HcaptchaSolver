/**
 * Behaviours:
 */

package events

func (B *EventManager) Event_901() FingerprintEvent {
	return FingerprintEvent{
		901,
		Stringify(B.Fingerprint.Hash["901"]),
	}
}
