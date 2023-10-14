/**
 * Behaviours:
 */

package events

func (B *EventManager) Event_2420() FingerprintEvent {
	return FingerprintEvent{
		2420,
		Stringify([]interface{}{
			EncStr(B.Fingerprint.Webgl.Vendor),
			EncStr(B.Fingerprint.Webgl.Renderer),
		}),
	}
}
