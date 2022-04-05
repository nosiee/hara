package convert

import (
	"net/url"
	"reflect"
	"strconv"
	"strings"

	"github.com/xfrr/goffmpeg/models"
)

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

func UrlQueryToOptions(values url.Values) ConversionOptions {
	var options ConversionOptions

	tof := reflect.TypeOf(options)
	vof := reflect.ValueOf(&options).Elem()

	for i := 0; i < vof.NumField(); i++ {
		if v, ok := values[tof.Field(i).Tag.Get("option")]; ok {
			s := strings.Join(v, "")

			switch vof.Field(i).Interface().(type) {
			case string:
				vof.Field(i).SetString(s)
			case uint:
				uv, err := strconv.ParseUint(s, 10, 64)
				if err != nil {
					continue
				}
				vof.Field(i).SetUint(uv)
			case int:
				iv, err := strconv.ParseInt(s, 10, 64)
				if err != nil {
					continue
				}
				vof.Field(i).SetInt(iv)
			case bool:
				bv, err := strconv.ParseBool(s)
				if err != nil {
					continue
				}
				vof.Field(i).SetBool(bv)
			}
		}
	}

	return options
}
