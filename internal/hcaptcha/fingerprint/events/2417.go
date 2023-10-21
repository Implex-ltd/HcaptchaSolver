/**
 * Behaviours:
 */

package events

func (B *EventManager) Event_2417() FingerprintEvent {
	return FingerprintEvent{
		2417,
		Stringify(B.Fingerprint.Events["2417"]),
	}
}
