/**
 * Behaviours:
 */

package events

func (B *EventManager) Event_2411() FingerprintEvent {
	return FingerprintEvent{
		2411,
		"[32767,32767,16384,8,8,8]",//	Stringify(B.Fingerprint.Events["2411"]),
	}
}
