/**
 * Behaviours:
 */

package events

func (B *EventManager) Event_2410() FingerprintEvent {
	return FingerprintEvent{
		2410,
		"[16,1024,4096,7,12,120,[23,127,127]]",//Stringify(B.Fingerprint.Events["2410"]),
	}
}
