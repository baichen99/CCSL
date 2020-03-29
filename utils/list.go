package utils

import (
	"errors"
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/kataras/iris/v12"
	"github.com/lib/pq"
)

// SearchByColumn where clause to execute a pseudo-search on the provided value against the given column
func SearchByColumn(columnName string, searchString string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if searchString != "" {
			// add % to prefix and suffix to match search
			// SQL search: SELECT * FROM column_name WHERE query LIKE value
			value := "%" + searchString + "%"
			query := fmt.Sprintf("%s LIKE ?", columnName)
			return db.Where(query, value)
		}
		return db
	}
}

// FilterByColumn generate function used to create a where clause for string type columns
func FilterByColumn(columnName string, value string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		// SQL select: SELECT * FROM table WHERE column_name = value
		if value != "" {
			query := columnName + " = ?"
			return db.Where(query, value)
		}
		return db

	}
}

// FilterInSubQuery generate function used to select in array
func FilterInSubQuery(columnName string, subQuery string) func(db *gorm.DB) *gorm.DB {
	// SELECT * FROM table WHERE columnName IN(subQuery)
	return func(db *gorm.DB) *gorm.DB {
		if subQuery != "" {
			query := fmt.Sprintf("%s IN(%s)", columnName, subQuery)
			return db.Where(query)
		}
		return db
	}
}

// SearchInList generate function to search value from array
func SearchInList(listName string, value string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		// SQL select in array: SELECT * FROM table WHERE array_to_string(listName, ',') LIKE value
		if value != "" {
			value := "%" + value + "%"
			query := fmt.Sprintf("array_to_string(%s,',') LIKE ?", listName)
			return db.Where(query, value)
		}
		return db
	}
}

// FilterInList generate function to filter value from array
func FilterInList(columnName string, value string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		// SQL select in array: SELECT * FROM table WHERE value = ANY(query)
		if value != "" {
			query := fmt.Sprintf("? = ANY(%s)", columnName)
			return db.Where(query, value)
		}
		return db
	}
}

// FilterMultiInList generate function to filter multiple value from array
func FilterMultiInList(columnName string, value pq.StringArray) func(db *gorm.DB) *gorm.DB {
	// SQL select multiple value in array: SELECT * FROM table WHERE [element_1, element_2] <@ columnName
	return func(db *gorm.DB) *gorm.DB {
		if len(value) != 0 {
			query := fmt.Sprintf("? <@ %s", columnName)
			return db.Where(query, value)
		}
		return db
	}
}

// FilterByArray generate function to search value from a comma escaped string
func FilterByArray(columnName string, value string, escape string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		// SQL select in array: SELECT * FROM table WHERE value = ANY(string_to_array(query,','))
		if value != "" {
			query := fmt.Sprintf("? = ANY(string_to_array(%s,'%s'))", columnName, escape)
			return db.Where(query, value)
		}
		return db
	}
}

// FilterByListParameters generate function to limit get list counts
func FilterByListParameters(parameters GetListParameters) func(db *gorm.DB) *gorm.DB {
	// SQL: ORDER orderQuery LIMIT limit OFFSET offset
	return func(db *gorm.DB) *gorm.DB {
		orderQuery := fmt.Sprintf("%s %s", parameters.OrderBy, parameters.Order)
		if parameters.Limit != 0 {
			return db.
				Order(orderQuery).
				Limit(parameters.Limit).
				Offset(parameters.Limit * (parameters.Page - 1))
		}
		return db.Order(orderQuery)
	}
}

// GetListParamsFromContext gets list params from context
func GetListParamsFromContext(c iris.Context, defaultOrderName string) (listParmas GetListParameters, err error) {
	listParmas.Page = c.URLParamIntDefault("page", 1)
	listParmas.Limit = c.URLParamIntDefault("limit", 10) // 0 means no limit
	listParmas.Order = c.URLParamDefault("order", "asc")
	listParmas.OrderBy = c.URLParamDefault("orderBy", defaultOrderName)
	if !(listParmas.Order == "asc" || listParmas.Order == "desc") {
		err = errors.New("ParamsError")
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
	Username string
	Name     string
	State    string
	Roles    string
}

// GetLexiconListParameters parameters for get word list queries
type GetLexiconListParameters struct {
	GetListParameters
	Initial string
	Chinese string
	English string
	Pos     string
}

// GetPerformerListParameters parameters for get performers list queries
type GetPerformerListParameters struct {
	GetListParameters
	Name     string
	RegionID string
	Gender   string
}

// GetLexicalVideoListParameters parameters for get videos list queries
type GetLexicalVideoListParameters struct {
	GetListParameters
	LexiconID        string
	Initial          string
	Chinese          string
	English          string
	Pos              string
	RegionID         string
	Gender           string
	LeftHandshapeID  string
	RightHandshapeID string
	HandshapeID      string
	WordFormation    string
	Morpheme         string
	PerformerID      string
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

// GetMemberListParameters parameters for get members list queries
type GetMemberListParameters struct {
	GetListParameters
	NameZh string
	NameEn string
}

// GetCarouselListParameters parameters for get carousel list queries
type GetCarouselListParameters struct {
	GetListParameters
	TitleZh string
	TitleEn string
	State   string
}

// GetHandshapeListParameters parameters for get handshape list queries
type GetHandshapeListParameters struct {
	GetListParameters
	Name string
}

// GetJsErrorListParameters parameters for get js error list queries
type GetJsErrorListParameters struct {
	GetListParameters
}

// GetUserLoginListParameters parameters for get login history
type GetUserLoginListParameters struct {
	GetListParameters
	UserID string
	Status string
	IP     string
}

// GetNotificationListParameters parameters for get notification list queries
type GetNotificationListParameters struct {
	GetListParameters
	UserID  string
	Title   string
	Message string
}

// GetClassListParameters parameters for get notification list queries
type GetClassListParameters struct {
	GetListParameters
	Name string
}

// GetCourseListParameters parameters for get notification list queries
type GetCourseListParameters struct {
	GetListParameters
	ClassID string
	Name    string
	Content string
}

// GetAssignmentListParameters parameters for get notification list queries
type GetAssignmentListParameters struct {
	GetListParameters
	CourseID string
	Title    string
	Type     string
	Content  string
}

// GetSubmittedAssignmentsListParameters parameters for get notification list queries
type GetSubmittedAssignmentsListParameters struct {
	GetListParameters
	CreatorID    string
	AssignmentID string
	GraderID     string
	Answer       string
	Comment      string
}

// GetPostListParameters parameters for get Post list queries
type GetPostListParameters struct {
	GetListParameters
	Title     string
	CreatorID string
}

// ReplyListParameters parameters for get Reply list queries
type ReplyListParameters struct {
	GetListParameters
	PostID    string
	CreatorID string
}
