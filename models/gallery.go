package models

import (
	"database/sql"
	"errors"
	"fmt"
)

type Gallery struct {
	ID     int
	UserID int
	Title  string
}

type GalleryService struct {
	DB *sql.DB
}

func (service *GalleryService) Create(title string, userID int) (*Gallery, error) {
	gallery := Gallery{
		Title:  title,
		UserID: userID,
	}
	row := service.DB.QueryRow(`
	insert into galleries (title,user_id)
	values (&1,&2) returning id;`, gallery.Title, gallery.UserID)
	err := row.Scan(&gallery.ID)
	if err != nil {
		return nil, fmt.Errorf("create gallery: %w", err)
	}
	return &gallery, nil
}

func (service *GalleryService) ByID(id int) (*Gallery, error) {
	gallery := Gallery{
		ID: id,
	}
	row := service.DB.QueryRow(`
	select title,user_id
	from galleries
	where id = $1;`, gallery.ID)
	err := row.Scan(&gallery.Title, &gallery.UserID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, fmt.Errorf("query gallery by id : %w", err)

	}
	return &gallery, nil

}

func (service *GalleryService) ByUserID(userID int) ([]Gallery,error) {
	rows , err := service.DB.Query(`
	select id, title
	from galleries
	where user_id=$1;`,userID)
	if err != nil {
		return nil,fmt.Errorf("query galleries by user: %w",err)

	}
	var galleries []Gallery 
	for rows.Next() {
		gallery := Gallery {
			UserID: userID,

		}
		err = rows.Scan(&gallery.ID,gallery.Title)
		if err != nil {
			return nil , fmt.Errorf("query galleries by user: %w",err)
		}
		galleries = append(galleries, gallery)

	}
	if rows.Err() != nil {
		return nil , fmt.Errorf("query galleries by user: %w",err)
	}
	return galleries,nil
}

func (service *GalleryService) Update(galler *Gallery) error {
	_,err := service.DB.Exec(`
	UPDATE galleries
	set title = $2
	where id = $1;`,galler.ID,galler.Title)
	if err != nil{
		return fmt.Errorf("update gallery: %w",err)
	}
	return nil 
}

func (service *GalleryService) Delete(id int) error {
	_,err := service.DB.Exec(`
	delete from galleries
	where id = $1;`,id)
	if err != nil {
		return fmt.Errorf("delete gallery : %w",err)
	}
	return nil
}