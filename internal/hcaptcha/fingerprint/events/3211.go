/**
 * Behaviours:
 */

package events

import (
	"strconv"
)

func (B *EventManager) Event_3211() FingerprintEvent {
	event := B.Fingerprint.Events["3210"].([]any)
	numberAsString := strconv.FormatFloat(event[0].(float64), 'f', -1, 64)

	return FingerprintEvent{
		3211,
		Stringify(EncStr(numberAsString)),
	}
}
