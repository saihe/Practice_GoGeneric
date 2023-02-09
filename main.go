package main

import "practice_gogeneric/model"

func TableName[T model.TableModel](tableModel T) T {
	return tableModel
}
