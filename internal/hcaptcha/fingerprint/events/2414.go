/**
 * Behaviours:
 */

package events

func (B *EventManager) Event_2414() FingerprintEvent {
	return FingerprintEvent{
		2414,
		"[16384,32,16384,2048,2,2048]",//Stringify(B.Fingerprint.Events["2414"]),
	}
}
