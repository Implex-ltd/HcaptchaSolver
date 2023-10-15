/**
 * Behaviours:
 */

package events

import (
	"fmt"
	"time"
)

func (B *EventManager) Event_3504() FingerprintEvent {
	return FingerprintEvent{
		3504,
		fmt.Sprintf("%.1f", float64(time.Now().UnixNano())/1e6),
	}
}
