/**
 * Behaviours:
 */

package events

func (B *EventManager) Event_2407() FingerprintEvent {
	return FingerprintEvent{
		2407,
		Stringify(B.Fingerprint.Hash["2407"]),
	}
}
