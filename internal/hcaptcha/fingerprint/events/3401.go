/**
 * Behaviours:
 */

package events

import _ "github.com/Implex-ltd/hcsolver/internal/utils"

func (B *EventManager) Event_3401() FingerprintEvent {
	return FingerprintEvent{
		3401,
		"2942030374453181095", //utils.RandomHash(19),
	}
}
