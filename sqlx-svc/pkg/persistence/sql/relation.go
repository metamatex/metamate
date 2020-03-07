package sql

import "fmt"

func GetRelation(tableName, relationName string, isActive bool, id string) (q string, values []interface{}) {
	var whereField string
	var onField string
	if isActive {
		onField = "passive_id_value"
		whereField = "active_id_value"
	} else {
		onField = "active_id_value"
		whereField = "passive_id_value"
	}

	q = fmt.Sprintf("SELECT * FROM %v INNER JOIN %v ON %v.id_value=%v.%v WHERE %v.%v=?", tableName, relationName, tableName, relationName, onField, relationName, whereField)

	values = append(values, id)

	return
}
