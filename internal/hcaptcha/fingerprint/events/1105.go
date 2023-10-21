/**
 * Behaviours:
 */

package events

func (B *EventManager) Event_1105() FingerprintEvent {
	return FingerprintEvent{
		1105,
		Stringify(B.Fingerprint.Hash["1105"]),
	}
}
