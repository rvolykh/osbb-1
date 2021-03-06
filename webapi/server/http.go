package server

import (
	"context"
	"fmt"
	"log"

	pb "github.com/i3odja/osbb/contracts/notifications"
	"github.com/i3odja/osbb/webapi/client"
)

type HTTP struct {
}

func AllNotifications(ctx context.Context, c *client.Notifications) error {
	r, err := c.SendNotification(ctx, &pb.SendRequest{UserId: "0673419017", Notification: nil})
	if err != nil {
		return fmt.Errorf("could not send notification: %w", err)
	}
	log.Printf("Sending...: %s", r.GetSResponse())

	rb, err := c.BroadcastNotification(ctx, &pb.BroadcastRequest{Notification: nil})
	if err != nil {
		return fmt.Errorf("could not broadcast notification: %w", err)
	}
	log.Printf("Broadcasting...: %s", rb.GetBResponse())

	rm, err := c.MyNotification(ctx, &pb.MyRequest{
		Notification: []*pb.Notification{
			{
				Title: "My testing",
			},
		},
	})
	if err != nil {
		return fmt.Errorf("could not my notification: %w", err)
	}
	log.Printf("my...: %s", rm.GetMResponse())

	rs, err := c.SendNotification(ctx, &pb.SendRequest{
		UserId: "0731674016",
	})
	if err != nil {
		return fmt.Errorf("could not send notification: %w", err)
	}
	log.Printf("Sending...: %s", rs.GetSResponse())

	return nil
}
