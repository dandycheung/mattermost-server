package graphql

import (
	"context"

	"github.com/mattermost/mattermost-server/v5/app"
	"github.com/mattermost/mattermost-server/v5/model"
)

type Resolver struct{
	*app.App
}

func (r *Resolver) Channel() ChannelResolver {
	return &channelResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type channelResolver struct{ *Resolver }

func (r *channelResolver) Members(ctx context.Context, channel *model.Channel) ([]*model.ChannelMember, error) {
	channelMembers, err := r.GetChannelMembersPage(channel.Id, 1, 20)
	if err != nil {
		return nil, err
	}

	members := []*model.ChannelMember{}
	for _, channelMember := range *channelMembers {
		members = append(members, &channelMember)
	}

	return members, nil
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Channels(ctx context.Context, teamId string) ([]*model.Channel, error) {
	channelList, err := r.GetPublicChannelsForTeam(teamId, 0, 20)
	if err != nil {
		return nil, err
	}

	channels := []*model.Channel{}
	for _, c := range *channelList {
		channels = append(channels, c)
	}

	return channels, nil
}
