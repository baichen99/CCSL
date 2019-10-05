package utils

import (
	"errors"

	"github.com/jinzhu/gorm"
	"github.com/kataras/iris"
)

// SearchByColumn where clause to execute a pseudo-search on the provided value against the given column
func SearchByColumn(columnName string, searchString string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if searchString != "" {
			// add % to prefix and suffix to match search
			// SQL search: SELECT * FROM column_name WHERE query LIKE value
			value := "%" + searchString + "%"
			query := columnName + " LIKE ? "
			return db.Where(query, value)
		}
		return db
	}
}

// FilterByColumn generate function used to create a where clause for string type columns
func FilterByColumn(columnName string, value string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		// SQL select: SELECT * FROM column_name WHERE query = value
		if value != "" {
			query := columnName + " = ?"
			return db.Where(query, value)
		}
		return db

	}
}

// FilterInSubQuery generate function used to select in array
func FilterInSubQuery(columnName string, subQuery string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if subQuery != "" {
			query := columnName + " IN (" + subQuery + ")"
			return db.Where(query)
		}
		return db
	}
}

// FilterByArray generare function to search value from a comma escaped string
func FilterByArray(columnName string, value string, escape string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		// SQL select in array: SELCET * FROM column_name WHERE value = ANY(string_to_array(query,','))
		if value != "" {
			query := "? = ANY(string_to_array(" + columnName + ", '" + escape + "'))"
			return db.Where(query, value)
		}
		return db
	}
}

// GetListParamsFromContext gets list params from context
func GetListParamsFromContext(c iris.Context, defaultOrderName string) (listParmas GetListParameters, err error) {
	listParmas.Page = c.URLParamIntDefault("page", 1)
	listParmas.Limit = c.URLParamIntDefault("limit", 10) // 0 means no limit
	listParmas.Order = c.URLParamDefault("order", "asc")
	listParmas.OrderBy = c.URLParamDefault("orderBy", defaultOrderName)
	if !(listParmas.Order == "asc" || listParmas.Order == "desc") {
		err = errors.New("PramsError")
	}
	return
}

// GetListParameters base parameters needed to define the get list queries
type GetListParameters struct {
	Page    int
	Limit   int
	OrderBy string
	Order   string
}

// GetUserListParameters parameters for get users list queries
type GetUserListParameters struct {
	GetListParameters
	UserType   string
	SearchName string
}

// GetWordListParameters parameters for get word list queries
type GetWordListParameters struct {
	GetListParameters
	Initial string
	Chinese string
	English string
	Pos     string
}

// GetPerformerListParameters parameters for get performers list queries
type GetPerformerListParameters struct {
	GetListParameters
	Name   string
	Region string
	Gender string
}

// GetVideoListParameters parameters for get videos list queries
type GetVideoListParameters struct {
	GetListParameters
	WordID         string
	Initial        string
	Chinese        string
	English        string
	Pos            string
	RegionID       string
	Gender         string
	LeftSignID     string
	RightSignID    string
	SignID         string
	ConstructType  string
	ConstructWords string
	PerformerID    string
}

// GetNewsListParameters parameters for get list queries
type GetNewsListParameters struct {
	GetListParameters
	Column   string
	Title    string
	Type     string
	Text     string
	Language string
	State    string
}

// GetMemberListParameters parameters for get members list queris
type GetMemberListParameters struct {
	GetListParameters
	NameZh string
	NameEn string
}

// GetCarouselListParameters parameters for get carousel list queris
type GetCarouselListParameters struct {
	GetListParameters
	Title string
	State string
}

// GetSignListParameters parameters for get sign list queries
type GetSignListParameters struct {
	GetListParameters
	Name string
}
