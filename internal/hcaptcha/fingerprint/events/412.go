/**
 * Behaviours:
 */

package events

func (B *EventManager) Event_412() FingerprintEvent {
	return FingerprintEvent{
		412,
		Stringify(B.Fingerprint.Hash["412"]),
	}
}
