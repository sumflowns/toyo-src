package book_service

import (
	"fmt"
	"github.com/sumflowns/toyo-src/book-ticket-common/proto/task"
	"sync"
)

type service struct {
}

var (
	s *service
	m sync.RWMutex
)

type Service interface {
	//获取信息
	GetTaskById(req *task.In_GetTaskInfo) (rsp *task.Out_GetTaskInfo)
	//修改信息
	UpdateTaskInfo(req *task.In_UpdateTaskInfo) (rsp *task.Out_UpdateTaskInfo)
	//获取列表
	GetTasks(req *task.In_GetTaskInfoList) (rsp *task.Out_GetTaskInfoList)
	//获取需要抢的列表
	GetNeedTicketList(limit int64, pages int64, status int64) (rsp []task.Task, err error)
	//新建信息
	CreateTask(req *task.In_AddTask) (rsp *task.Out_AddTask)

	StartBathTicket()

	StartBathDoneError()

	//获取用户列表
	GetUserTaskList(req *task.In_GetUserTaskList) (rsp *task.Out_GetTaskInfoList)
	//修改信息
	UpdateTaskStatus(req *task.In_UpdateTaskStatus) (rsp *task.Out_UpdateTaskStatus)
}

//获取服务
func GetService() (Service, error) {
	if s == nil {
		return nil, fmt.Errorf("[GetService] GetService 未初始化")
	}
	return s, nil
}

func Init() {
	m.Lock()
	defer m.Unlock()
	if s != nil {
		return
	}
	s = &service{}
}
