package main

import (
	"flag"
	"os"
	"strings"
	"time"

	"github.com/fabric8io/jadvisor/sinks"
	"github.com/fabric8io/jadvisor/sources"
	"github.com/golang/glog"
)

var argPollDuration = flag.Duration("poll_duration", 10*time.Second, "Polling duration")

func main() {
	flag.Parse()
	glog.Infof(strings.Join(os.Args, " "))
	glog.Infof("jAdvisor version %v", jadvisorVersion)
	err := doWork()
	if err != nil {
		glog.Error(err)
		os.Exit(1)
	}
	os.Exit(0)
}

func doWork() error {
	source, err := sources.NewSource()
	if err != nil {
		return err
	}
	sink, err := sinks.NewSink()
	if err != nil {
		return err
	}
	ticker := time.NewTicker(*argPollDuration)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			data, err := source.GetData()
			if err != nil {
				return err
			}

			glog.V(2).Infof("Got stats: %v", data)

			if err := sink.StoreData(data); err != nil {
				return err
			}
		}
	}
	return nil
}
