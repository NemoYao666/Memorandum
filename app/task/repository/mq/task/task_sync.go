package task

import (
	"context"
	"encoding/json"

	"micro-todoList/app/task/repository/mq"
	"micro-todoList/app/task/service"
	"micro-todoList/consts"
	"micro-todoList/idl/pb"
	log "micro-todoList/pkg/logger"
)

type SyncTask struct {
}

func (s *SyncTask) RunTaskCreate(ctx context.Context) error {
	rabbitMqQueue := consts.RabbitMqTaskQueue
	msgs, err := mq.ConsumeMessage(ctx, rabbitMqQueue)
	if err != nil {
		return err
	}
	var forever chan struct{}

	go func() {
		for d := range msgs {
			log.LogrusObj.Infof("Received run Task: %s", d.Body)

			// 落库
			reqRabbitMQ := new(pb.TaskRequest)
			err = json.Unmarshal(d.Body, reqRabbitMQ)
			if err != nil {
				log.LogrusObj.Infof("Received run Task: %s", err)
			}

			err = service.TaskMQ2MySQL(ctx, reqRabbitMQ)
			if err != nil {
				log.LogrusObj.Infof("Received run Task: %s", err)
			}

			d.Ack(false)

		}
	}()

	log.LogrusObj.Infoln(err)
	<-forever

	return nil
}
