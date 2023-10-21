/**
 * Behaviours:
 */

package events

func (B *EventManager) Event_401() FingerprintEvent {
	return FingerprintEvent{
		401,
		Stringify(B.Fingerprint.Hash["401"]),
	}
}
