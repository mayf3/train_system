package models

import (
	"fmt"
	"testing"
)

func TestDbInit(t *testing.T) {
	InitDb()
}

func TestDbInsert(t *testing.T) {
	da := &DbAgent{}

	info := map[string]interface{}{
		"username": "mayf3",
		"password": "123456",
		"email":    "mayf3@qq.com",
	}
	result, err := da.From("user").Insert(info)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(result.LastInsertId())
}

func TestDbCount(t *testing.T) {
	da := &DbAgent{}

	rows, err := da.From("user").Where("username = ?", "mayf3").Count()
	defer rows.Close()

	if err != nil {
		fmt.Println(err)
	}

	num := 0
	rows.Next()
	err = rows.Scan(&num)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(num)
}

func TestDbUpdate(t *testing.T) {
	da := &DbAgent{}

	result, err := da.From("user").Where("username = ?", "mayf3").Update(map[string]interface{}{
		"username": "hello",
	})

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(result.LastInsertId())
}

func TestDbDelete(t *testing.T) {
	da := &DbAgent{}

	_, err := da.From("user").Delete()
	if err != nil {
		fmt.Println(err)
	}

	da.Clear()
	result, err := da.From("user").Where("username = ?", "hello").Delete()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result.LastInsertId())
}
