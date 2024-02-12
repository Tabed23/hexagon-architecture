package service

import (
	"fmt"
	"log"

	"github.com/gocql/gocql"
	"github.com/tabed23/hexagon-architecture/pkg/model"
)

type Service struct {
	db *gocql.Session
}

func NewService(db *gocql.Session) *Service {
	return &Service{db: db}
}

// Delete implements Repository.
func (s *Service) Delete(id string) (string, error) {
	if err := s.db.Query("DELETE FROM candy WHERE id = ?", id).Exec(); err != nil {
		return fmt.Sprintf("error deleting the %s", id), err
	}
	return fmt.Sprintf("id %s has been deleted", id), nil
}

// Get implements Repository.
func (s *Service) Get() ([]model.Candy, error) {
    var candys []model.Candy
	var candy model.Candy
    iter := s.db.Query("SELECT id, candyname, candy_price FROM candy").Iter()
    for iter.Scan(&candy.ID, &candy.Candyname, &candy.CandyPrice) {
        candys = append(candys, candy)
    }
    if err := iter.Close(); err != nil {
        log.Fatal(err)
        return nil, err
    }
    return candys, nil
}
// Post implements Repository.
func (s *Service) Post(c model.Candy) (string, error) {
	c.ID = gocql.TimeUUID()
	if err := s.db.Query("INSERT INTO candy(id, candyname, candy_price) VALUES(?, ?, ?)", c.ID, c.Candyname, c.CandyPrice).Exec(); err != nil {
		return fmt.Sprintf("Error inserting candyname: %v", c), err
	}

	return "New candy has been inserted", nil
}

// Put implements Repository.
func (s *Service) Put(id string, c model.Candy) (model.UpdatedCandy, error) {
	if err := s.db.Query("UPDATE candy SET candyname = ?, candy_price = ? WHERE id = ?",
		c.Candyname, c.CandyPrice, id).Exec(); err != nil {
		return model.UpdatedCandy{}, err
	}
	updatedCandy := model.UpdatedCandy{
        Candyname:  c.Candyname,
        CandyPrice: c.CandyPrice,
    }

	return updatedCandy, nil
}
