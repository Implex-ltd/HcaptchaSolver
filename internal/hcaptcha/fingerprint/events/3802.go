/**
 * Behaviours:
 */

package events

import (
	"fmt"

	"github.com/Implex-ltd/hcsolver/internal/utils"
)

func (B *EventManager) Event_3802() FingerprintEvent {
	return FingerprintEvent{
		3802,
		fmt.Sprintf("%.13f", utils.RandomFloat64(300, 500)),
	}
}
