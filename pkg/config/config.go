package config

import (
	"github.com/golang/glog"
	"github.com/kelseyhightower/envconfig"
)

// Config contains overridable configuration options for the Sensor
var Sensor struct {
	ProcFs   string `split_words:"true" default:"/proc"`
	CgroupFs string `split_words:"true" default:"/sys/fs/cgroup"`
	TraceFs  string `split_words:"true" default:"/sys/kernel/debug/tracing"`

	// Node name to use if not the value returned from uname(2)
	NodeName string

	// DockerContainerDir is the path to the directory used for docker
	// container local storage areas (i.e. /var/lib/docker/containers)
	DockerContainerDir string `split_words:"true" default:"/var/lib/docker/containers"`

	// OciContainerDir is the path to the directory used for the
	// container runtime's container state directories
	// (i.e. /var/run/docker/libcontainerd)
	OciContainerDir string `split_words:"true" default:"/var/run/docker/libcontainerd"`

	// Pubsub backend implementation to use
	Pubsub string `default:"stan"`

	// Subscription timeout in seconds
	SubscriptionTimeout int64 `default:"5"`

	TelemetryServiceBindAddress string `default:"127.0.0.1:5051"`
	MonitoringPort              int    `default:"8083"`
}

var ApiServer struct {
	Pubsub         string `default:"stan"`
	Port           int    `default:"8080"`
	ProxyPort      int    `default:"8081"`
	MonitoringPort int    `default:"8082"`
}

var Backplane struct {
	ClusterName       string `default:"c8-backplane"`
	NatsURL           string `default:"nats://localhost:4222"`
	NatsMonitoringURL string `default:"http://localhost:8222"`
	AckWait           int    `default:"1"`
}

var Recorder struct {
	TelemetryServiceURL string `default:"127.0.0.1:5051"`
	DbPath              string `default:"/var/lib/capsule8/recorder"`
	DbFileName          string `default:"recorder.db"`
	DbSizeLimit         string `default:"100mb"`
	MonitoringPort      int    `default:"8084"`
}

func init() {
	err := envconfig.Process("C8_APISERVER", &ApiServer)
	if err != nil {
		glog.Fatal(err)
	}

	err = envconfig.Process("C8_BACKPLANE", &Backplane)
	if err != nil {
		glog.Fatal(err)
	}

	err = envconfig.Process("C8_RECORDER", &Recorder)
	if err != nil {
		glog.Fatal(err)
	}

	err = envconfig.Process("C8_SENSOR", &Sensor)
	if err != nil {
		glog.Fatal(err)
	}
}