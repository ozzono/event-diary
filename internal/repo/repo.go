package repo

import (
	"event-diary/internal/models"

	"github.com/pkg/errors"
	"go.uber.org/zap"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const recordDB = "record.db"

type Client struct {
	db  *gorm.DB
	log *zap.SugaredLogger
}

func NewClient(log *zap.SugaredLogger) (*Client, error) {
	db, err := gorm.Open(sqlite.Open(recordDB), &gorm.Config{})
	if err != nil {
		return nil, errors.Wrap(err, "gorm sqlite open")
	}
	db.AutoMigrate(&models.Record{})
	return &Client{
		db:  db,
		log: log,
	}, nil
}

func (c *Client) Create(r *models.Record) error {
	c.log.Info("inserting registry")
	result := c.db.Create(r)
	if result.Error != nil {
		c.log.Errorf("failed inserstion %v", result.Error)
		return errors.Wrap(result.Error, "c.db.Create")
	}
	c.log.Info("inserted registry")
	return nil
}

func (c *Client) Delete(r *models.Record) error {
	c.log.Info("deleting registry")
	result := c.db.Delete(&r)
	if result.Error != nil {
		c.log.Errorf("failed deletion %v", result.Error)
		return errors.Wrap(result.Error, "c.db.Delete")
	}
	c.log.Info("deleted registry")
	return nil
}

func (c *Client) AllRecords() ([]*models.Record, error) {
	c.log.Info("fetching all records")
	records := []*models.Record{}
	result := c.db.Find(&records)
	if result.Error != nil {
		return nil, errors.Wrap(result.Error, "c.db.Find")
	}
	c.log.Infof("records found %d", len(records))
	return records, nil
}
