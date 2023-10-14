/**
 * Behaviours:
 */

package events

func (B *EventManager) Event_1101() FingerprintEvent {
	return FingerprintEvent{
		1101,
		Stringify(B.Fingerprint.Hash["1101"]),
	}
}
