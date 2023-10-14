/**
 * Behaviours:
 */

package events

func (B *EventManager) Event_201() FingerprintEvent {
	return FingerprintEvent{
		201,
		Stringify(B.Fingerprint.Hash["201"]),
	}
}
