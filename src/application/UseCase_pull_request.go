package application

import (
	domain "ejercicio/src/domain/value_objects"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)







func HandleGithubPingEvent(ctx *gin.Context) {
	log.Println("Evento Ping recibido de GitHub.")
	ctx.JSON(http.StatusOK, gin.H{"status": "Ping recibido"})
}

func HandleGithubPullRequestEvent(ctx *gin.Context, payload []byte) {
	var eventPayload domain.PullRequestEventPayload
	if err := json.Unmarshal(payload, &eventPayload); err != nil {
		log.Printf("Error al deserializar el payload del pull request: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Error al procesar el payload del pull request"})
		return
	}

	log.Printf(
		"Evento Pull Request recibido: Acción=%s, PR Título='%s', Rama Base='%s', Repositorio='%s'",
		eventPayload.Action, eventPayload.PullRequest.Title, eventPayload.PullRequest.Base.Ref, eventPayload.Repository.FullName)

	if eventPayload.Action == "closed" {
		log.Printf("Pull Request cerrado en el repositorio '%s': %s", eventPayload.Repository.FullName, eventPayload.PullRequest.Title)
		fmt.Printf("Pull Request cerrado: %s\n", eventPayload.PullRequest.Title)

		ctx.JSON(http.StatusOK, gin.H{
			"status":  "Pull Request cerrado detectado",
			"title":   eventPayload.PullRequest.Title,
			"base_ref": eventPayload.PullRequest.Base.Ref,
			"repo":    eventPayload.Repository.FullName,
			"url":     eventPayload.PullRequest.URL,
		})
		return
	}

	mainBranch := "develop"

	if eventPayload.PullRequest.Base.Ref == mainBranch {
		log.Printf("¡Pull Request a la rama '%s' detectado en el repositorio '%s'!", mainBranch, eventPayload.Repository.FullName)
		fmt.Printf("Pull Request detectado en la rama %s!\n", mainBranch)
	} else {
		log.Printf(
			"Pull Request detectado, pero no dirigido a la rama '%s'. Rama base: '%s'",
			mainBranch, eventPayload.PullRequest.Base.Ref)
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "Evento Pull Request recibido y procesado"})
}
