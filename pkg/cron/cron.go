package cron

import (
	"github.com/robfig/cron"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"

	"github.com/aeekayy/go-api-base/pkg/models"
)

// Job represents a cron job for the services
type Job struct {
	Config *Config
	Cron   *cron.Cron
	DB     *gorm.DB
}

// PrometheusResponse for prometheus response
type PrometheusResponse struct {
	Status string
	Data   PrometheusResponseData
}

// PrometheusResponseData for Prometheus response
type PrometheusResponseData struct {
	ResultType string
	Result     []PrometheusResponseObject
}

// PrometheusResponseObject for Prometheus response
type PrometheusResponseObject struct {
	Metric Metric
	Value  []interface{}
}

// APIResponse for API respsonses
type APIResponse struct {
	CurrentTime string `json:"currentTime"`
	StartedTime string `json:"startedTime"`
	BuildInfo   string `json:"buildInfo"`
	ClusterID   string `json:"clusterID"`
	NSF         string `json:"nsf"`
	APIBaseURL  string `json:"apiBaseURL"`
}

// Metric represents Prometheus metric
type Metric map[string]string

// NewCron returns new cron job
func NewCron(config *Config, db *gorm.DB) Job {
	log.Info("Returning a new cron")
	cronJob := cron.New()
	cronJob.AddFunc("0 */15 * * *", func() { RetrieveEvents(db) })

	return Job{
		Config: config,
		Cron:   cronJob,
		DB:     db,
	}
}

// RetrieveEvents returns events from the database
func RetrieveEvents(db *gorm.DB) {
	log.Info("[RetrieveEvents] Starting build version retrieval")
	var events []models.Event
	db.Find(&events)
	log.Infof("%+v", events)
}
