/**
 * Behaviours:
 */

package events

func (B *EventManager) Event_3401() FingerprintEvent {
	return FingerprintEvent{
		3401,
		Stringify(B.Fingerprint.Hash["3401"]),
	}
}
