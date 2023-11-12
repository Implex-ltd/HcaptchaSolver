/**
 * Behaviours:
 */

package events

import _ "github.com/Implex-ltd/hcsolver/internal/utils"

func (B *EventManager) Event_3401(hash string) FingerprintEvent {
	return FingerprintEvent{
		3401,
		hash, //utils.RandomHash(19),
	}
}
