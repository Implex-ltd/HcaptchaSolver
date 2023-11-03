/**
 * Behaviours:
 */

package events

//import "github.com/Implex-ltd/hcsolver/internal/utils"

func (B *EventManager) Event_3401() FingerprintEvent {
	return FingerprintEvent{
		3401,
		"4226317358175830201",//utils.RandomHash(19), //Stringify(B.Fingerprint.Hash["3401"]),
	}
}
