/**
 * Behaviours:
 */

package events

func (B *EventManager) Event_1402() FingerprintEvent {
	return FingerprintEvent{
		1402,
		"[\"Europe/Paris\",-60,-60,-3203647761000,\"Central European Standard Time\",\"en-US\"]",//Stringify(B.Fingerprint.Timezone),
	}
}
