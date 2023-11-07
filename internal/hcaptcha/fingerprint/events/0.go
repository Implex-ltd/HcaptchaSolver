/**
 * Behaviours:
 */

package events


import (
	"fmt"

	"github.com/Implex-ltd/hcsolver/internal/utils"
)


func (B *EventManager) Event_0() FingerprintEvent {
	return FingerprintEvent{
		0,
		fmt.Sprintf("%.14f", utils.RandomFloat64(15, 40)),
	}
}
