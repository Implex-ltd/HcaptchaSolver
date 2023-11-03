/**
 * Behaviours:
 */

package events

func (B *EventManager) Event_1401() FingerprintEvent {
	return FingerprintEvent{
		1401,
		Stringify(B.Fingerprint.Timezone[0].(string)),
	}
}
