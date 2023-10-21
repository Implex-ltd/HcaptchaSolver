/**
 * Behaviours:
 */

package events

import (
	"fmt"

	"github.com/Implex-ltd/hcsolver/internal/utils"
)

func (B *EventManager) Event_3502() FingerprintEvent {
	return FingerprintEvent{
		3502,
		fmt.Sprintf("%.14f", utils.RandomFloat64(1, 20)),
	}
}
