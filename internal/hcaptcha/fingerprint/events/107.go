/**
 * Behaviours: https://gist.github.com/nikolahellatrigger/ea00832b010c0db8f0a0d5ca0d467072
 */

package events

func (B *EventManager) Event_107() FingerprintEvent {
	event := B.Fingerprint.Events["107"].(map[string]interface{})

	data := Stringify([]interface{}{
		B.Fingerprint.Screen["Width"],
		B.Fingerprint.Screen["Height"],
		B.Fingerprint.Screen["AvailWidth"],
		B.Fingerprint.Screen["AvailHeight"],
		B.Fingerprint.Screen["ColorDepth"],
		B.Fingerprint.Screen["PixelDepth"],
		event["touchEvent"].(bool),
		B.Fingerprint.Screen["MaxTouchPoints"],
		B.Fingerprint.Screen["DevicePixelRatio"],
		B.Fingerprint.Screen["outerWidth"],
		B.Fingerprint.Screen["outerHeight"],
		event["MediaQueryList"].([]interface{})[0],
		event["MediaQueryList"].([]interface{})[1],
		event["MediaQueryList"].([]interface{})[2],
		event["MediaQueryList"].([]interface{})[3],
	})

	return FingerprintEvent{
		107,
		data,
	}
}
