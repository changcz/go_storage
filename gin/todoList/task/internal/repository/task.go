package repository

import "task/internal/service"

type Task struct {
	TaskID    uint `gorm:"primarykey"` // id
	UserID    uint `gorm:"index"`      // 用户id
	Status    int  `gorm:"default:0"`
	Title     string
	Content   string `gorm:"type:longtext"`
	StartTime int64
	EndTime   int64
}

func (t *Task) TaskCreat(req *service.TaskRequest) (err error){
	task := &Task{
		UserID:     uint(req.UserID),
		Title:     req.Title,
		Content:   req.Content,
		StartTime: int64(req.StartTime),
		EndTime:   int64(req.EndTime),
	}

	return DB.Model(&Task{}).Create(&task).Error
}

func (t *Task) TaskUpdate(req *service.TaskRequest)(err error)   {
	err = DB.Where("task_id=?", req.TaskID).First(&t).Error
	if err != nil {
		return err
	}

	t.Title = req.Title
	t.Content = req.Content
	t.Status = int(req.Status)
	t.StartTime = int64(req.StartTime)
	t.EndTime = int64(req.EndTime)
	return DB.Save(&t).Error
}

func (t *Task) TaskShow(req *service.TaskRequest)(taskList []*Task,err error)   {
	return taskList, DB.Model(&Task{}).Where("user_id = ?",req.UserID).Find(&taskList).Error
}

func (t Task) TaskDelete(req *service.TaskRequest)(err error)   {
	err = DB.Where("task_id=?", req.TaskID).First(&t).Error
	if err != nil {
		return err
	}

	return DB.Model(&Task{}).Delete("task_id=?",req.TaskID).Error
}

func BuildTasks(item []*Task) (tList []*service.TaskModel) {
	for _, v := range item {
		f := BuildTask(v)
		tList = append(tList, f)
	}
	return tList
}


func BuildTask(item *Task) *service.TaskModel {
	return &service.TaskModel{
		TaskID:    uint32(item.TaskID),
		UserID:    uint32(item.UserID),
		Status:    uint32(item.Status),
		Title:     item.Title,
		Content:   item.Content,
		StartTime: uint32(item.StartTime),
		EndTime:   uint32(item.EndTime),
	}
}
