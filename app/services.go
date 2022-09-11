package main

// Tasks

func GetTasks() ([]Task, error) {
	var tasks []Task

	if err := Db.Model(&Task{}).Preload("Assignees").
		Preload("Tags").Find(&tasks).Error; err != nil {
		return tasks, err
	}
	return tasks, nil
}

func GetTasksWithTag(tagName string) ([]Task, error) {
	var tasks []Task

	if err := Db.Model(&Task{}).Preload("Assignees").Preload("Tags").
		Joins("JOIN task_tags ON tasks.id = task_tags.task_id").
		Joins("JOIN tags ON tags.id = task_tags.tag_id").
		Where("tags.name = ?", tagName).Find(&tasks).Error; err != nil {
		return tasks, err
	}

	return tasks, nil
}

func GetTasksWithUser(userName string) ([]Task, error) {
	var tasks []Task

	if err := Db.Model(&Task{}).Preload("Assignees").Preload("Tags").
		Joins("JOIN task_users ON tasks.id = task_users.task_id").
		Joins("JOIN users ON users.id = task_users.user_id").
		Where("users.name = ?", userName).Find(&tasks).Error; err != nil {
		return tasks, err
	}

	return tasks, nil
}

func GetTasksWithUserAndTag(userName string, tagName string) ([]Task, error) {
	var tasks []Task

	if err := Db.Model(&Task{}).Preload("Assignees").Preload("Tags").
		Joins("JOIN task_tags ON tasks.id = task_tags.task_id").
		Joins("JOIN tags ON tags.id = task_tags.tag_id").
		Joins("JOIN task_users ON tasks.id = task_users.task_id").
		Joins("JOIN users ON users.id = task_users.user_id").
		Where("users.name = ? AND tags.name = ?", userName, tagName).
		Find(&tasks).Error; err != nil {
		return tasks, err
	}

	return tasks, nil
}

func GetTask(id string) (Task, error) {
	var task Task

	if err := Db.Model(&Task{}).Preload("Assignees").Preload("Tags").First(&task, id).Error; err != nil {
		return task, err
	}

	return task, nil
}

func PostTask(task Task) error {
	if err := Db.Create(&task).Error; err != nil {
		return err
	}
	return nil
}

func UpdateTask(task Task) error {
	if err := Db.Save(&task).Error; err != nil {
		return err
	}
	return nil
}

// Users

func GetUser(id string) (User, error) {
	var user User

	if err := Db.Model(&User{}).Preload("Tasks").First(&user, id).Error; err != nil {
		return user, err
	}

	return user, nil
}

func GetUsers() ([]User, error) {
	var users []User

	if err := Db.Model(&User{}).Preload("Tasks").Find(&users).Error; err != nil {
		return users, err
	}
	return users, nil
}

func PostUser(user User) error {
	if err := Db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func UpdateUser(user User) error {
	if err := Db.Save(&user).Error; err != nil {
		return err
	}
	return nil
}
