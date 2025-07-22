package task3

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

// Employee 员工结构体
type Employee struct {
	ID         int    `db:"id"`
	Name       string `db:"name"`
	Department string `db:"department"`
	Salary     int    `db:"salary"`
}

func GetTechEmployees(db *sqlx.DB) ([]Employee, error) {
	const query = `SELECT id, name, department, salary FROM employees WHERE department = ?`

	var employees []Employee
	err := db.Select(&employees, query, "技术部")
	if err != nil {
		return nil, fmt.Errorf("查询技术部员工失败: %v", err)
	}

	return employees, nil
}

func GetHighestPaidEmployee(db *sqlx.DB) (Employee, error) {
	const query = `SELECT id, name, department, salary 
                  FROM employees 
                  ORDER BY salary DESC 
                  LIMIT 1`

	var employee Employee
	err := db.Get(&employee, query)
	if err != nil {
		return Employee{}, fmt.Errorf("查询最高薪资员工失败: %v", err)
	}

	return employee, nil
}

func Run_case3(db *sqlx.DB) {
	//初始化表
	db.MustExec(`CREATE TABLE IF NOT EXISTS employees (
                  id INT AUTO_INCREMENT PRIMARY KEY,
                  name VARCHAR(255) NOT NULL,
                  department VARCHAR(255) NOT NULL,
                  salary INT NOT NULL
                )`)

	//模拟数据
	db.MustExec(`INSERT INTO employees (name, department, salary) VALUES 
                ('张三', '技术部', 5000),
                ('李四', '销售部', 4000),
                ('王五', '技术部', 6000),
                ('赵六', '销售部', 5500),
                ('孙七', '技术部', 7000)`)

	// 查询技术部员工
	techEmployees, err := GetTechEmployees(db)
	if err != nil {
		fmt.Printf("错误: %v\n", err)
	} else {
		fmt.Println("技术部员工:")
		for _, emp := range techEmployees {
			fmt.Printf("ID: %d, 姓名: %s, 薪资: %d\n", emp.ID, emp.Name, emp.Salary)
		}
	}

	// 查询最高薪资员工
	topEmployee, err := GetHighestPaidEmployee(db)
	if err != nil {
		fmt.Printf("错误: %v\n", err)
	} else {
		fmt.Printf("\n最高薪资员工: %s, 部门: %s, 薪资: %d\n",
			topEmployee.Name, topEmployee.Department, topEmployee.Salary)
	}
}
