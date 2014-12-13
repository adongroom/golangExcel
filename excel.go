// create by hanzhendong
package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/tealeg/xlsx"
)

func main() {
	db, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/test?charset=utf8")
	checkErr(err)
	excelFileName := "./ccc.xlsx"
	xlFile, error := xlsx.OpenFile(excelFileName)
	if error != nil {
		fmt.Println("出错了,小伙")
	}

	//插入数据
	stmt, err := db.Prepare("INSERT excelInfo SET id=?,name=?")
	checkErr(err)
	for _, sheet := range xlFile.Sheets {
		for _, row := range sheet.Rows {

			fmt.Printf("%s\n", sheet.Rows)
			if row.Cells[0].String() == "id" || row.Cells[1].String() == "name" {
				continue
			}
			stmt.Exec(row.Cells[0].String(), row.Cells[1].String())
			checkErr(err)

		}
	}
}

//检查错误
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
