package service

import (
	"big_num_compute_service/config"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"net"
)

var (
	BigNumComputeConf  config.ConfYaml
	QueueComputeWorker chan net.Conn
	LogAccess          *logrus.Logger
	LogError           *logrus.Logger
	db                 *gorm.DB
)
