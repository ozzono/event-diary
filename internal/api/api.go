package api

import (
	"net/http"

	"event-diary/chart"
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

// @Summary Exibe relatório com contagem de convulsões registradas ao longo do dia
// @Produce  json
// @Success 200 "ok"
// @Router /relatorio [get]
func (h *handler) Report(c *gin.Context) {
	h.log.Info("show report")
	c.HTML(
		http.StatusOK,
		"bar.html",
		gin.H{
			"content": "Esse é um relatório de crises",
			"title":   "Relatório de crises",
			"url":     "/bar.html",
		},
	)
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
	newRecord.Description = c.Query("descricao")

	if err := newRecord.Valid(); err != nil {
		c.JSON(http.StatusBadRequest, web.APIError{ErrorMessage: err.Error(), ErrorCode: http.StatusBadRequest})
		return
	}

	if err := h.c.Create(newRecord); err != nil {
		h.log.Errorf("h.c.Create - %v", err)
		c.JSON(http.StatusInternalServerError, web.APIError{ErrorMessage: "contact system administrator", ErrorCode: http.StatusInternalServerError})
		return
	}
	go func() {
		records, err := h.c.AllRecords()
		if err != nil {
			h.log.Infof("h.c.AllRecords %v", err)
			return
		}
		reports, err := models.ReportData(records)
		if err != nil {
			h.log.Infof("models.ReportData %v", err)
			return
		}

		chart.BarChart(reports.ToBarData(), reports.HourArr())
	}()
	c.JSON(http.StatusOK, newRecord)
}

// @Summary Lista todos os registros
// @Produce  json
// @Success 200 {object} []models.Record	"ok"
// @Router /registros [get]
func (h *handler) AllRecords(c *gin.Context) {
	records, err := h.c.AllRecords()
	if err != nil {
		h.log.Errorf("h.c.Create - %v", err)
		c.JSON(http.StatusInternalServerError, web.APIError{ErrorMessage: "contact system administrator", ErrorCode: http.StatusInternalServerError})
		return
	}
	c.JSON(http.StatusOK, records)
}
