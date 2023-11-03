/**
 * Behaviours:
 */

package events

func (B *EventManager) Event_1302() FingerprintEvent {
	return FingerprintEvent{
		1302,
		"[1,2,3,4]",//Stringify(B.Fingerprint.Events["1302"]),
	}
}
