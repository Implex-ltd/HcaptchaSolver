/**
 * Behaviours:
 */

package events

func (B *EventManager) Event_407(extraKeys ...[]string) FingerprintEvent {
	defaultKeys := []string{
		"Raven",
		"_sharedLibs",
		"hsw",
		"__wdata",
	}

	keys := append([]string(nil), defaultKeys...)
	if len(extraKeys) > 0 {
		keys = append(keys, extraKeys[0]...)
	}

	return FingerprintEvent{
		407,
		Stringify([]interface{}{
			[]interface{}{
				"loadTimes",
				"csi",
				"app",
			},
			35,
			34,
			nil,
			false,
			false,
			true,
			37,
			true,
			true,
			true,
			true,
			true,
			keys,
			[]interface{}{
				[]interface{}{
					"getElementsByClassName",
					[]interface{}{},
				},
				[]interface{}{
					"getElementById",
					[]interface{}{},
				},
				[]interface{}{
					"querySelector",
					[]interface{}{},
				},
				[]interface{}{
					"querySelectorAll",
					[]interface{}{},
				},
			},
			[]interface{}{},
			true,
		}),
	}
}
