/**
 * Behaviours:
 */

package events

func (B *EventManager) Event_1402() FingerprintEvent {
	return FingerprintEvent{
		1402,
		Stringify(B.Fingerprint.Timezone),
	}
}
