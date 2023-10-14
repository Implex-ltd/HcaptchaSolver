/**
 * Behaviours:
 */

package events

func (B *EventManager) Event_3211() FingerprintEvent {
	return FingerprintEvent{
		3211,
		Stringify(EncStr("143254600089")),
	}
}
