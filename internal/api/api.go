package api

import (
	"net/http"

	"event-diary/internal/models"
	"event-diary/internal/repo"
	"event-diary/utils"
	"event-diary/web"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type handler struct {
	c   *repo.Client
	log *zap.SugaredLogger
}

func NewHandler(log *zap.SugaredLogger) (*handler, error) {
	c, err := repo.NewClient(log)
	if err != nil {
		return nil, errors.Wrap(err, "repo.NewClient")
	}

	return &handler{
		c:   c,
		log: log,
	}, nil
}

// @Summary Adiciona um novo registro no diário
// @Produce  json
// @Param   ocorrido_agora    query    bool     true        "Evento ocorrido agora?"
// @Param   horario_evento    query    string   true        "Horário do evento"
// @Param   reporter          query    string   true        "Quem está registrando?"
// @Param   descricao         query    string   true        "Descrição do evento"
// @Success 200 {object} models.Record	"ok"
// @Failure 400 {object} web.APIError "insira dados faltantes"
// @Failure 500 {object} web.APIError "falha na inserção de dados"
// @Router /registro [get]
func (h *handler) AddRecord(c *gin.Context) {
	newRecord := &models.Record{}
	newRecord.EventTime = c.Query("horario_evento")
	newRecord.Reporter = c.Query("reporter")
	newRecord.RecordTime = utils.Now()
	if c.Query("ocorrido_agora") == "true" {
		newRecord.EventTime = newRecord.RecordTime
	} else {
		custom, err := utils.Custom(c.Query("horario_evento"))
		if err != nil {
			c.JSON(http.StatusBadRequest, web.APIError{ErrorMessage: err.Error(), ErrorCode: http.StatusBadRequest})
			return
		}
		newRecord.EventTime = custom
	}

	if err := newRecord.Valid(); err != nil {
		c.JSON(http.StatusBadRequest, web.APIError{ErrorMessage: err.Error(), ErrorCode: http.StatusBadRequest})
		return
	}

	if err := h.c.Create(newRecord); err != nil {
		h.log.Errorf("h.c.Create - %v", err)
		c.JSON(http.StatusInternalServerError, web.APIError{ErrorMessage: "contact system administrator", ErrorCode: http.StatusInternalServerError})
		return
	}
	c.JSON(http.StatusOK, newRecord)
}