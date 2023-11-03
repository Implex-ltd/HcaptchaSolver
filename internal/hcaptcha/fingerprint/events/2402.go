/**
 * Behaviours:
 */

package events

func (B *EventManager) Event_2402() FingerprintEvent {
	return FingerprintEvent{
		2402,
		"[\"Google Inc. (NVIDIA)\",\"ANGLE (NVIDIA, NVIDIA GeForce GTX 1060 6GB (0x00001C03) Direct3D11 vs_5_0 ps_5_0, D3D11)\"]", 
		/*Stringify([]interface{}{
			B.Fingerprint.Webgl.Vendor,
			B.Fingerprint.Webgl.Renderer,
		}),*/
	}
}
