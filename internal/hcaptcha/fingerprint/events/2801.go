/**
 * Behaviours:
 */

package events

func (B *EventManager) Event_2801() FingerprintEvent {
	return FingerprintEvent{
		2801,
		Stringify(B.Fingerprint.Hash["2801"]),
	}
}
