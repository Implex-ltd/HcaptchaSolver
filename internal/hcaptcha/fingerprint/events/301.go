/**
 * Behaviours:
 */

package events

func (B *EventManager) Event_301() FingerprintEvent {
	return FingerprintEvent{
		301,
		Stringify(B.Fingerprint.Hash["301"]),
	}
}
