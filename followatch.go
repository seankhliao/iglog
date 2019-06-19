package main

import (
	"context"
	"sort"
	"time"

	"github.com/ahmdrz/goinsta"
	"github.com/seankhliao/iglog/iglog"
	log "github.com/sirupsen/logrus"
)

type ProtoDiff struct {
	events    *iglog.Events
	followers *iglog.Users
	following *iglog.Users
}

func (c *Client) EventLog(ctx context.Context, r *iglog.Request) (*iglog.Events, error) {
	log.Infoln("EventLog started")
	return c.pDiff.events, nil
}

func (c *Client) Followers(ctx context.Context, r *iglog.Request) (*iglog.Users, error) {
	log.Infoln("Followers started")
	return c.pDiff.followers, nil
}

func (c *Client) Following(ctx context.Context, r *iglog.Request) (*iglog.Users, error) {
	log.Infoln("Following started")
	return c.pDiff.following, nil
}

func (c *Client) followDiffProto() {
	log.Debugln("FollowDiffProto events")
	c.pDiff = &ProtoDiff{}
	c.pDiff.events = &iglog.Events{
		Events: make([]*iglog.Event, len(c.fDiff.events)),
	}
	for i, e := range c.fDiff.events {
		c.pDiff.events.Events[len(c.fDiff.events)-1-i] = event2proto(e)
	}
	log.Debugln("FollowDiffProto events", len(c.pDiff.events.Events))

	log.Debugln("FollowDiffProto followers")
	c.pDiff.followers = &iglog.Users{}
	for _, u := range c.fDiff.followers {
		c.pDiff.followers.Users = append(c.pDiff.followers.Users, user2proto(u))
	}
	sort.Slice(c.pDiff.followers.Users, func(i, j int) bool {
		return c.pDiff.followers.Users[i].Username < c.pDiff.followers.Users[j].Username
	})
	log.Debugln("FollowDiffProto followers", len(c.pDiff.followers.Users))

	log.Debugln("FollowDiffProto following")
	c.pDiff.following = &iglog.Users{}
	for _, u := range c.fDiff.following {
		c.pDiff.following.Users = append(c.pDiff.following.Users, user2proto(u))
	}
	sort.Slice(c.pDiff.following.Users, func(i, j int) bool {
		return c.pDiff.following.Users[i].Username < c.pDiff.following.Users[j].Username
	})
	log.Debugln("FollowDiffProto following", len(c.pDiff.following.Users))
}

func event2proto(e FollowEvent) *iglog.Event {
	et := iglog.EventType_Unknown
	switch e.Type {
	case FEFollowerGained:
		et = iglog.EventType_FollowerGained
	case FEFollowerLost:
		et = iglog.EventType_FollowerLost
	case FEFollowingGained:
		et = iglog.EventType_FollowingGained
	case FEFollowingLost:
		et = iglog.EventType_FollowingLost
	}
	return &iglog.Event{
		Time: e.TimeStamp.Format(time.RFC3339),
		Type: et,
		User: user2proto(e.User),
	}
}

func user2proto(u goinsta.User) *iglog.User {
	return &iglog.User{
		Id:          u.ID,
		Username:    u.Username,
		Displayname: u.FullName,
	}
}
