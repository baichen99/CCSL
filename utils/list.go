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

// FilterByColumn general function used to create a where clause for string type columns
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
func GetListParamsFromContext(c iris.Context, orderName string) (listParmas GetListParameters, err error) {
	listParmas.Page = c.URLParamIntDefault("page", 1)
	listParmas.Limit = c.URLParamIntDefault("limit", 0) // 0 means no limit
	listParmas.Order = c.URLParamDefault("order", "asc")
	listParmas.OrderBy = c.URLParamDefault("orderBy", orderName)
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
	Type    string
	Initial string
	Chinese string
	English string
}

type GetPerformerListParameters struct {
	GetListParameters
	Name   string
	Region string
	Gender string
}

type GetVideoListParameters struct {
	GetListParameters
	Initial        string
	Chinese        string
	English        string
	Type           string
	Name           string
	Region         string
	Gender         string
	LeftSign       string
	RightSign      string
	ConstructType  string
	ConstructWords string
	PerformerID    string
	WordID         string
}
