package photo

import (
	"a6-api/utils/helper"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type Space struct{}

type SpaceListFields struct {
	Id      uint8
	Name    string
	Summary string
}

type SpaceListParams struct {
	IsDefault  string `form:"-" json:"is_default,omitempty" map:"field:is_default:1"`
	IsShow     string `form:"-" json:"is_show,omitempty" map:"field:is_show;default:1"`
	OrderField string `form:"orderField" json:"orderField" map:"field:orderField;default:id"`
}

func (*Space) TableName() string {
	prefix := photo.TablePrefix()
	table := strings.ToLower(reflect.TypeOf(Space{}).Name())
	return fmt.Sprintf("%s_%s", prefix, table)
}

func (*SpaceListParams) Error(err interface{}) string {
	if e, ok := err.(*strconv.NumError); ok {
		return fmt.Sprintf("This value %s is illegal", e.Num)
	}
	//Other errors
	return "Invalid params"
}

func (s *Space) List(baseParamsMaps map[string]interface{}, listParamsMaps map[string]interface{}) (fields []SpaceListFields, pagin map[string]interface{}, notFound bool) {
	var totalNum uint32

	db = db.Table(s.TableName())
	db.Where(listParamsMaps).Count(&totalNum)

	//get pagin data
	totalPage, offset := helper.Paginator(totalNum, baseParamsMaps["perPageNum"].(uint16), baseParamsMaps["page"].(uint16))

	if err := db.Where(listParamsMaps).Offset(offset).Order(baseParamsMaps["orderField"].(string) + " " + baseParamsMaps["orderType"].(string)).Limit(baseParamsMaps["perPageNum"].(uint16)).Scan(&fields).Error; err != nil {
		return []SpaceListFields{}, pagin, true
	}

	//get pagin info
	pagin = helper.GeneratePaginInfo(totalNum, totalPage, baseParamsMaps["page"].(uint16), baseParamsMaps["perPageNum"].(uint16), offset)
	return
}
