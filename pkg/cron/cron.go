package cron

import (
	"github.com/robfig/cron"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"

	"github.com/aeekayy/go-api-base/pkg/models"
)

type CronJob struct {
	Config		*Config
	Cron		*cron.Cron
	DB			*gorm.DB
}

type PrometheusResponse struct {
	Status			string
	Data 			PrometheusResponseData
}

type PrometheusResponseData struct {
	ResultType		string
	Result			[]PrometheusResponseObject
}

type PrometheusResponseObject struct {
	Metric			Metric
	Value			[]interface{} 
}

type ApiResponse struct {
	CurrentTime			string			`json:"currentTime"`
	StartedTime			string			`json:"startedTime"`
	BuildInfo			string			`json:"buildInfo"`
	ClusterID			string			`json:"clusterID"`
	NSF					string			`json:"nsf"`
	ApiBaseURL			string			`json:"apiBaseURL"`
	EdgeUIAppID			string			`json:"edgeUIAppID"`
}

type Metric map[string]string

func NewCron(config *Config, db *gorm.DB) CronJob {
	log.Info("Returning a new cron")
	cronJob := cron.New()
	cronJob.AddFunc("0 */15 * * *", func() { RetrieveEvents(db) })

	return CronJob{ 
		Config: config,
		Cron:	cronJob,
		DB:		db,
	}
}

func RetrieveEvents(db *gorm.DB) {
	log.Info("[RetrieveEvents] Starting build version retrieval")
	var events []models.Event
	db.Find(&events)
	log.Infof("%+v", events)
}
