/**
 * Behaviours:
 */

package events

import (
	"fmt"

	"github.com/Implex-ltd/hcsolver/internal/utils"
)

func (B *EventManager) Event_3503() FingerprintEvent {
	return FingerprintEvent{
		3503,
		fmt.Sprintf("%.14f", utils.RandomFloat64(0, 1)),
	}
}
