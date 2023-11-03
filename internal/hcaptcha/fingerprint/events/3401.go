/**
 * Behaviours:
 */

package events

//import "github.com/Implex-ltd/hcsolver/internal/utils"

func (B *EventManager) Event_3401() FingerprintEvent {
	return FingerprintEvent{
		3401,
		"14582940537062521554", //utils.RandomHash(19),
	}
}
