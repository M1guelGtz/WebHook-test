package application

import (
	"encoding/json"
	domain "github_wb/domain/value_objects"
	"log"
)

func ProcessPullRequest(payload []byte) int {
	var eventPayload domain.PullRequestEventPayload

	if err := json.Unmarshal(payload, &eventPayload); err != nil {
		return 500
	}
	if eventPayload.Action == "ready_for_review" {
		base := eventPayload.PullRequest.Base.Ref
		branch := eventPayload.PullRequest.Head.Ref
		user := eventPayload.PullRequest.User.Login
		pRID := eventPayload.PullRequest.ID
		body := eventPayload.PullRequest.Body

		log.Printf("Pull Request Listo para Revision:\nID:%d\nBase:%s\nHead:%s\nUser:%s\nBody:%s", pRID, base, branch, user,body)
	} else {
		log.Printf("Pull Request Action no es Closed: %s", eventPayload.Action)
	}

	return 200
}
