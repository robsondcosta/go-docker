package models

import (
	"strconv"

	"docker/api/config"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Tabler interface {
	TableName() string
}

func Insert(table string, campos []string, model interface{}) int64 {
	db = config.Connect()
	defer db.Close()

	db.NewRecord(model)
	rows := db.Debug().Table(table).Omit("data_alteracao").Create(model).RowsAffected
	return rows
}

func Update(table string, campos []string, clausula map[string]interface{}, model interface{}) int64 {
	db = config.Connect()
	defer db.Close()

	rows := db.Table(table).Select(campos).Where(clausula).Updates(model).RowsAffected
	return rows
}

func Delete(table string, clausula map[string]interface{}, model interface{}) int64 {
	db = config.Connect()
	defer db.Close()

	rows := db.Table(table).Where(clausula).Delete(model).RowsAffected
	return rows
}

func Select(table string, campos []string, clausula map[string]interface{}, model interface{}) interface{} {
	db = config.Connect()
	defer db.Close()

	db.Table(table).Select(campos).Where(clausula).Find(*&model)
	return model
}

func SelectAll(table string) []interface{} {
	db = config.Connect()
	defer db.Close()

	var array []interface{}
	rows, _ := db.Table(table).Rows()
	columns, _ := rows.Columns()
	tamanho := len(columns)
	value := make([]interface{}, tamanho)
	valuesColumns := make([]interface{}, tamanho)

	for rows.Next() {
		for i := range columns {
			valuesColumns[i] = &value[i]
		}

		rows.Scan(valuesColumns...)
		result := make(map[string]interface{})
		for i, col := range columns {
			val := value[i]
			ret, flag := val.([]byte)
			var valueInterf interface{}

			if flag {
				_, err := strconv.Atoi(string(ret))
				if err == nil {
					valueInterf, _ = strconv.Atoi(string(ret))
				} else {
					valueInterf = string(ret)
				}
			} else {
				valueInterf = val
			}
			result[col] = valueInterf
		}

		array = append(array, result)
	}
	return array
}
