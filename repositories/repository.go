package repositories

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strings"

	"github.com/ettle/strcase"
	"gorm.io/gorm"
)

// Repository common repository
type Repository struct{}

// NewRepository new repository
func NewRepository() Repository {
	return Repository{}
}

// FindFullHierarchy find full hierarchy
func (r *Repository) FindFullHierarchy(db *gorm.DB, id interface{}, model interface{}, condition map[string]interface{}) bool {
	found := false
	found = !errors.Is(db.First(model, id).Error, gorm.ErrRecordNotFound)
	if found {
		fetchRecord(db, model, condition)
	}

	return found
}

// FetchRecord find data if found Child
// reference: https://stackoverflow.com/questions/29230261/how-preload-a-full-hierarchy-in-go-using-gorm
func fetchRecord(db *gorm.DB, data interface{}, condition map[string]interface{}) {
	var ref reflect.Value
	if reflect.TypeOf(data).Kind() == reflect.Struct {
		ref = reflect.ValueOf(data)
	} else if reflect.TypeOf(data).Kind() == reflect.Ptr {
		ref = reflect.Indirect(reflect.ValueOf(data))
	}

	switch ref.Type().Kind() {
	case reflect.Slice:
		for i := 0; i < ref.Len(); i++ {
			if ref.Index(i).Type().Kind() == reflect.Ptr {
				fetchRecord(db, ref.Index(i).Elem().Addr().Interface(), condition)
			}
		}

	case reflect.Struct:
		for i := 0; i < ref.NumField(); i++ {
			var IDFieldRaw string
			var IDFields []string
			var RefFieldRaw string
			var RefFields []string
			var re *regexp.Regexp
			var matches []string

			if ref.Field(i).CanAddr() && strings.EqualFold(ref.Type().Field(i).Tag.Get("walkrec"), "true") {
				gormflags := ref.Type().Field(i).Tag.Get("gorm")
				if gormflags == "" {
					panic("No gorm flags found!")
				} else {
					re = regexp.MustCompile(`\breferences:([a-zA-Z0-9_,]+)\b`)
					matches = re.FindStringSubmatch(gormflags)
					if len(matches) == 2 {
						IDFieldRaw = matches[1]
						IDFields = strings.Split(IDFieldRaw, ",")
					}
					re = regexp.MustCompile(`\bforeignkey:([a-zA-Z0-9_,]+)\b`)
					matches = re.FindStringSubmatch(gormflags)
					if len(matches) == 2 {
						RefFieldRaw = matches[1]
						RefFields = strings.Split(RefFieldRaw, ",")
					}
				}

				if len(IDFields) == 0 {
					continue
				}

				if len(RefFields) != 0 {
					WhereMap := make(map[string]interface{})
					for fk := 0; fk < len(RefFields); fk++ {
						WhereMap[strcase.ToSnake(RefFields[fk])] = fmt.Sprint(ref.FieldByName(IDFields[fk]))
					}
					object := ref.Field(i).Addr().Interface()
					for k, v := range condition {
						WhereMap[k] = v
					}
					db.Where(WhereMap).Order("id").Find(object)
					if object != nil {
						fetchRecord(db, object, condition)
					}
				} else {
					panic("foreignKey empty!")
				}
			}
		}
	}
}
