package events

import (
	"encoding/json"
)

func Stringify(data any) string {
	jsonString, err := json.Marshal(data)
	if err != nil {
		return ""
	}

	return string(jsonString)
}

func (E *EventManager) Event_407Custom(keys []string) FingerprintEvent {
	return E.Event_407(keys)
}

func (E *EventManager) Event_3401Custom(hash string) FingerprintEvent {
	return E.Event_3401(hash)
}

func (E *EventManager) BuildEvents(customKeys []string, domHash string) (Events [][]interface{}) {
	eventsMethods := []func(*EventManager) FingerprintEvent{
		(*EventManager).Event_3,
		(*EventManager).Event_1902,
		(*EventManager).Event_1901,
		(*EventManager).Event_1101,
		(*EventManager).Event_1103,
		(*EventManager).Event_1105,
		(*EventManager).Event_201,
		(*EventManager).Event_1107,
		(*EventManager).Event_211,
		func(e *EventManager) FingerprintEvent {
			return e.Event_3401Custom(domHash)
		},
		(*EventManager).Event_3403,
		(*EventManager).Event_803,
		(*EventManager).Event_604,
		(*EventManager).Event_2801,
		(*EventManager).Event_2805,
		(*EventManager).Event_107,
		(*EventManager).Event_301,
		(*EventManager).Event_304,
		(*EventManager).Event_1401,
		(*EventManager).Event_1402,
		(*EventManager).Event_1403,
		(*EventManager).Event_3504,
		(*EventManager).Event_3501,
		(*EventManager).Event_3503,
		(*EventManager).Event_3502,
		(*EventManager).Event_401,
		(*EventManager).Event_402,
		func(e *EventManager) FingerprintEvent {
			return e.Event_407Custom(customKeys)
		},
		(*EventManager).Event_412,
		(*EventManager).Event_2408,
		(*EventManager).Event_2402,
		(*EventManager).Event_2420,
		(*EventManager).Event_2401,
		(*EventManager).Event_2407,
		(*EventManager).Event_2409,
		(*EventManager).Event_2410,
		(*EventManager).Event_2411,
		(*EventManager).Event_2412,
		(*EventManager).Event_2413,
		(*EventManager).Event_2414,
		(*EventManager).Event_2415,
		(*EventManager).Event_2416,
		(*EventManager).Event_2417,
		(*EventManager).Event_901,
		(*EventManager).Event_905,
		(*EventManager).Event_1302,
		(*EventManager).Event_1904,
		(*EventManager).Event_702,
		(*EventManager).Event_3210,
		(*EventManager).Event_3211,
		(*EventManager).Event_0,
	}

	for _, eventMethod := range eventsMethods {
		event := eventMethod(E)
		Events = append(Events, []interface{}{
			event.EventID,
			event.Value,
		})
	}

	// utils.ShuffleSlice(Events)
	return Events
}
