/**
 * Behaviours:
 */

package events

func (B *EventManager) Event_2402() FingerprintEvent {
	return FingerprintEvent{
		2402,
		Stringify([]interface{}{
			B.Fingerprint.Webgl.Vendor,
			B.Fingerprint.Webgl.Renderer,
		}),
	}
}
