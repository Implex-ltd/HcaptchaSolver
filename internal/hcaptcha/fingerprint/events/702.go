/**
 * Behaviours:
 */

package events

func (B *EventManager) Event_702() FingerprintEvent {
	event := B.Fingerprint.Events["702"].(map[string]interface{})

	return FingerprintEvent{
		702,
		Stringify([]interface{}{
			event["OsName"],
			event["OsVersion"],
			event["AndroidModel"],
			event["Arch"],
			event["CPU"],
			event["BrowserVersion"],
		}),
	}
}
