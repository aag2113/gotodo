package db

import "time"

type Task struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"createdAt"`
	Status    string    `json:"status"`
}

func CreateTask(t Task) (Task, error) {
	sql := `
	INSERT INTO tasks (id, title, created_at, status)
	VALUES ($1, $2, $3, $4)
	RETURNING id, title, created_at, status`
	var task Task
	err := DataBase.QueryRow(sql, t.ID, t.Title, t.CreatedAt, t.Status).Scan(&task.ID, &task.Title, &task.CreatedAt, &task.Status)
	if err != nil {
		panic(err)
	}
	return task, err
}

func GetTask(id string) (Task, error) {
	sql := `
	SELECT id, title, created_at, status
	FROM tasks
	WHERE id = $1`
	var err error
	var task Task
	err = DataBase.QueryRow(sql, id).Scan(&task.ID, &task.Title, &task.CreatedAt, &task.Status)
	// TODO: Handle 404
	if err != nil {
		panic(err)
	}
	return task, err
}

func GetAllTasks() ([]*Task, error) {
	sql := `SELECT id, title, created_at, status from tasks;`
	rows, err := DataBase.Query(sql)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	tasks := make([]*Task, 0)

	for rows.Next() {
		t := new(Task)
		err := rows.Scan(&t.ID, &t.Title, &t.CreatedAt, &t.Status)
		if err != nil {
			panic(err)
		}
		tasks = append(tasks, t)
	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}
	return tasks, err
}

func UpdateTask(t Task) (Task, error) {
	sql := `
	UPDATE tasks
	SET title=$2, status=$3
	WHERE id=$1
	RETURNING id, title, created_at, status`
	var err error
	err = DataBase.QueryRow(sql, t.ID, t.Title, t.Status).Scan(&t.ID, &t.Title, &t.CreatedAt, &t.Status)
	// TODO: Handle 404
	if err != nil {
		panic(err)
	}
	return t, err
}
