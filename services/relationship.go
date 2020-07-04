package services

var IRelationshipService RelationshipService

type RelationshipService interface {
	CreateRelationship(userA, userB string) error
	DeleteRelationship(userA, userB string) error
	GetAllFriends(username string) ([]string, error)
	IsFriend(userA, userB string) (bool, error)
}
