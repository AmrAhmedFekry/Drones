package Utils

import (
	"GoGin/Application"
	"GoGin/Models"
	"context"
	"log"
	"time"

	"github.com/procyon-projects/chrono"
)

// set your own schedule functions here
func RunScheduleTasks() {
	// Battery Log Task every 10 minutes
	batteryLogAction := func() {
		var drones []Models.Drone
		application := Application.NewApp()
		application.DB.Find(&drones)
		for _, drone := range drones {
			var DroneBatteryLog Models.DroneBatteryLogs
			DroneBatteryLog.DroneId = drone.ID
			DroneBatteryLog.DroneBatteryCapacity = drone.BatteryCapacity
			application.DB.Create(&DroneBatteryLog)
		}
	}
	ScheduleTask(batteryLogAction, 600)
}

// RunScheduleTasks is a function to run schedule tasks
func ScheduleTask(action func(), rate int) {

	taskScheduler := chrono.NewDefaultTaskScheduler()

	_, err := taskScheduler.ScheduleAtFixedRate(func(ctx context.Context) {
		action()
	}, time.Duration(rate)*time.Second)

	if err == nil {
		log.Print("Task has been scheduled successfully.")
	}
}
