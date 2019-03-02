package controllers

// Follower is a class not shared between bounded contexts to meet requirement of no source code dependencies
type Follower struct {
	ID         string
	Name       string
	SocialTags []string
}
