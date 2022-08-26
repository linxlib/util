package gorm

import "gorm.io/gorm/schema"

type SameNamer struct {
}

func (s SameNamer) TableName(table string) string {
	return table
}

func (s SameNamer) SchemaName(table string) string {
	return table
}

func (s SameNamer) ColumnName(table, column string) string {
	return column
}

func (s SameNamer) JoinTableName(joinTable string) string {
	return joinTable
}

func (s SameNamer) RelationshipFKName(relationship schema.Relationship) string {
	return "REL_" + relationship.Name
}

func (s SameNamer) CheckerName(table, column string) string {
	return "CHK_" + table + "_" + column
}

func (s SameNamer) IndexName(table, column string) string {
	return "IDX_" + table + "_" + column
}
