/**
 * Behaviours:
 */

package events

func (B *EventManager) Event_1403() FingerprintEvent {
	return FingerprintEvent{
		1403,
		"[\"KY5Z3dSZKmlCFa5Bne=w\",\"7\",\"2\",\"QMXJIQCXOSKVH\"]", //Stringify(EncStr(B.Fingerprint.Timezone[0].(string))),
	}
}
