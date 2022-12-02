// THIS FILE IS GENERATED. EDITING IS FUTILE.
//
// Generated by:
//     kinds/gen.go
// Using jennies:
//     GoTypesJenny
//     LatestJenny
//
// Run 'make gen-cue' from repository root to regenerate.

package team

// Defines values for Permission.
const (
	PermissionN1 Permission = 1

	PermissionN2 Permission = 2

	PermissionN4 Permission = 4
)

// Team defines model for team.
type Team struct {
	// AccessControl metadata associated with a given resource.
	AccessControl map[string]interface{} `json:"accessControl"`

	// AvatarUrl is the team's avatar URL.
	AvatarUrl *string `json:"avatarUrl,omitempty"`

	// Created indicates when the team was created.
	Created int64 `json:"created"`

	// Email of the team.
	Email *string `json:"email,omitempty"`

	// MemberCount is the number of the team members.
	MemberCount int64 `json:"memberCount"`

	// Name of the team.
	Name string `json:"name"`

	// OrgId is the ID of an organisation the team belongs to.
	OrgId int64 `json:"orgId"`

	// TODO - it seems it's a team_member.permission, unlikely it should belong to the team kind
	Permission Permission `json:"permission"`

	// Updated indicates when the team was updated.
	Updated int64 `json:"updated"`
}

// TODO - it seems it's a team_member.permission, unlikely it should belong to the team kind
type Permission int