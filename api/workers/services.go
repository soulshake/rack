package workers

import (
	"fmt"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ecs"
	"github.com/convox/rack/api/models"
	"github.com/convox/rack/api/provider"
	"github.com/convox/rack/api/structs"
	"github.com/ddollar/logger"
)

// Monitor ECS Cluster.
// Notify on capacity-related and re-convergence events
func StartServicesCapacity() {
	converged, lastEvent := checkConverged()

	for _ = range time.Tick(1 * time.Minute) {
		converged, lastEvent = monitorConverged(converged, *lastEvent.CreatedAt)
	}
}

// Get initial convergence state
func checkConverged() (bool, ecs.ServiceEvent) {
	log := logger.New("ns=services_monitor")

	services, err := models.ClusterServices()
	if err != nil {
		log.Log("fn=checkConverged err=%q", err)

		return true, ecs.ServiceEvent{
			CreatedAt: aws.Time(time.Now()),
		}
	}

	converged := services.IsConverged()
	lastEvent := services.LastEvent()

	log.Log("fn=checkConverged converged=%t lastEventAt=%q", converged, lastEvent.CreatedAt)

	return converged, lastEvent
}

// Get latest convergence state and send notifications
func monitorConverged(lastConverged bool, lastEventAt time.Time) (bool, ecs.ServiceEvent) {
	log := logger.New("ns=services_monitor")

	services, err := models.ClusterServices()

	if err != nil {
		log.Log("fn=monitorConverged err=%q", err)

		return lastConverged, ecs.ServiceEvent{
			CreatedAt: aws.Time(lastEventAt),
		}
	}

	converged := services.IsConverged()
	events := services.EventsSince(lastEventAt)

	log.Log("fn=monitorConverged converged=%t events=%d lastEventAt=%q", converged, len(events), lastEventAt)

	if events.HasCapacityWarning() {
		provider.EventSend(&structs.Event{
			Action: "rack:capacity",
			Status: "error",
			Data: map[string]string{
				"rack": os.Getenv("RACK"),
			},
			Timestamp: time.Now(),
		}, fmt.Errorf(events.CapacityWarning()))
	}

	if converged != lastConverged {
		provider.EventSend(&structs.Event{
			Action: "rack:converge",
			Status: "success",
			Data: map[string]string{
				"rack":      os.Getenv("RACK"),
				"converged": fmt.Sprintf("%t", converged),
			},
			Timestamp: time.Now(),
		}, nil)
	}

	return converged, services.LastEvent()
}
