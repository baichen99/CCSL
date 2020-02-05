package models

import "github.com/lib/pq"

type Class struct {
    Base
    Name     string         `json:"name"`
    Teachers pq.StringArray `gorm:"uuid[];NOT NULL;DEFAULT:array[]::uuid[]" json:"teachers"`
    Students pq.StringArray `gorm:"uuid[];NOT NULL;DEFAULT:array[]::uuid[]" json:"students"`
    Details  string         `json:"details"`
}
