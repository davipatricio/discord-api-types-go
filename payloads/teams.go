package payloads

import "github.com/denkylabs/discord-api-types-go/globals"

// https://discord.com/developers/docs/topics/teams#data-models-team-object
type APITeam struct {
	//A hash of the image of the team's icon
	Icon string
	//The unique id of the team
	Id globals.Snowflake
	//The members of the team
	Members []APITeamMember
	//The name of the team
	Name string
	//The user id of the current team owner
	OwnerUserId globals.Snowflake
}

// https://discord.com/developers/docs/topics/teams#data-models-team-members-object
type APITeamMember struct {
	// The user's membership state on the team
	// See https://discord.com/developers/docs/topics/teams#data-models-membership-state-enum
	MembershipState TeamMemberMembershipState `json:"membership_state"`
	// Will always be `["*"]`
	Permissions [1]string `json:"permissions"`
	// The id of the parent team of which they are a member
	TeamId globals.Snowflake `json:"team_id"`
	// The avatar, discriminator, id, and username of the user
	// See https://discord.com/developers/docs/resources/user#user-object
	User APIUser `json:"user"`
}

// https://discord.com/developers/docs/topics/teams#data-models-membership-state-enum
type TeamMemberMembershipState uint8

// https://discord.com/developers/docs/topics/teams#data-models-membership-state-enum
const (
	TeamMemberMembershipStateInvited  TeamMemberMembershipState = 1
	TeamMemberMembershipStateAccepted TeamMemberMembershipState = 2
)
