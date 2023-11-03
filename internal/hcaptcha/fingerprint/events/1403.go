/**
 * Behaviours:
 */

package events

func (B *EventManager) Event_1403() FingerprintEvent {
	return FingerprintEvent{
		1403,
		Stringify(EncStr(B.Fingerprint.Timezone[0].(string))),
	}
}
