// generated by metactl sdk gen 
package sdk

const (
	FeedRelationsName = "FeedRelations"
)

type FeedRelations struct {
    ContainsStatuses *StatusesCollection `json:"containsStatuses,omitempty",yaml:"containsStatuses,omitempty"`
    Hash *string `json:"hash,omitempty",yaml:"hash,omitempty",hash:"ignore"`
    ParticipatedByPeople *PeopleCollection `json:"participatedByPeople,omitempty",yaml:"participatedByPeople,omitempty"`
}