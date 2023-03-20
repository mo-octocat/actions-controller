package webhooks

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"io"
	"log"
	"net/http"

	"github.tools.sap/actions-rollout-app/control/config"
)

func VerifySignature(payload []byte, signature string) bool {
	key := hmac.New(sha256.New, []byte(config.Config.GitHubWebhookSecret))
	key.Write([]byte(string(payload)))
	computedSignature := "sha256=" + hex.EncodeToString(key.Sum(nil))
	log.Printf("computed signature: %s", computedSignature)

	return computedSignature == signature
}

func WebhookHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	payload, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("Failed to read request body: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	eventType := r.Header.Get("X-GitHub-Event")
	switch eventType {
	case "workflow_dispatch":
		// handle workflow dispatch event
	case "workflow_run":
		// handle workflow run event
	case "workflow_job":
		// handle workflow job event
	default:
		log.Printf("Unsupported event type: %s", eventType)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	log.Printf("Received %s event", eventType)
	w.WriteHeader(http.StatusOK)
	w.Write(payload)
}
