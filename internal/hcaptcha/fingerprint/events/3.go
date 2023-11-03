/**
 * Behaviours:
 */

package events

import (
	"fmt"

	"github.com/Implex-ltd/hcsolver/internal/utils"
)

func (B *EventManager) Event_3() FingerprintEvent {
	return FingerprintEvent{
		3,
		 fmt.Sprintf("%.2f", utils.RandomFloat64(30000, 60000)),
	}
}
