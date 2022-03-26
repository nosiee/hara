package convert

import "github.com/xfrr/goffmpeg/models"

type VideoOptionsMap map[string]any

func (opt VideoOptionsMap) AddFunc(key string, f any) {
	opt[key] = f
}

func (opt VideoOptionsMap) CallFunc(key string, m *models.Mediafile, value any) {
	if _, ok := opt[key]; !ok {
		return
	}

	switch value.(type) {
	case string:
		if len(value.(string)) == 0 {
			break
		}
		opt[key].(func(*models.Mediafile, string))(m, value.(string))
	case int:
		if value.(int) == 0 {
			break
		}
		opt[key].(func(*models.Mediafile, int))(m, value.(int))
	case uint:
		if value.(uint) == 0 {
			break
		}
		opt[key].(func(*models.Mediafile, uint))(m, value.(uint))
	case bool:
		if !value.(bool) {
			break
		}
		opt[key].(func(*models.Mediafile, bool))(m, value.(bool))
	}
}
