/**
 * Behaviours:
 */

package events

func (B *EventManager) Event_3401() FingerprintEvent {
	return FingerprintEvent{
		3401,
		"4226317358175830201", //Stringify(B.Fingerprint.Hash["3401"]),
	}
}
