/**
 * Behaviours:
 */

package events

func (B *EventManager) Event_2401() FingerprintEvent {
	return FingerprintEvent{
		2401,
		"4185260598258108199", //Stringify(B.Fingerprint.Hash["2401"]),
	}
}
