package job

import (
	"log"
	"time"

	"caige/schedule-service"
	"caige/utils"
)

func ScheduleJobList() []schedule_service.ScheduleJob {
	jobs := make([]schedule_service.ScheduleJob, 0)
	//jobs = append(jobs, schedule_service.NewScheduleJob(test()))
	return jobs
}

func test() (string, time.Time, schedule_service.JobFunc, string) {
	name := "计算代理商收入"
	t := time.Date(0, 0, 0, 0, 0, 20, 0, utils.Location())
	f := func(logger *log.Logger) {

	}
	p := schedule_service.SCHEDULE_JOB_PERIOD_MINUTES
	return name, t, f, p
}
