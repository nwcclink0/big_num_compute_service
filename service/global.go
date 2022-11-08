package service

import (
	"big_num_compute_service/config"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var (
	BigNumComputeConf     config.ConfYaml
	QueueComputeWorkerNum chan uint64
	LogAccess             *logrus.Logger
	LogError              *logrus.Logger
	db                    *gorm.DB
)
