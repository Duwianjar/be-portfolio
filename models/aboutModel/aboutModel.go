package pboutModal

import (
	"be-portfolio/config"
	"be-portfolio/entities"
)

func GetAll() []entities.About {
	rows, err := config.DB.Query(`SELECT * FROM about`)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var abouts []entities.About

	for rows.Next() {
		var about entities.About
		if err := rows.Scan(&about.Id, &about.Name, &about.Value, &about.CreatedAt, &about.UpdatedAt); err != nil {
			panic(err)
		}

		abouts = append(abouts, about)
	}

	return abouts
}

func Create(about entities.About) bool {
	result, err := config.DB.Exec(`
		INSERT INTO about (name, value, created_at, updated_at)
		VALUE (? , ? , ? , ?)`,
		about.Name, about.Value, about.CreatedAt, about.UpdatedAt,
	)

	if err != nil {
		panic(err)
	}

	LastInsertID, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	return LastInsertID > 0
}

func Detail(id int) entities.About {
	row := config.DB.QueryRow(`SELECT id, name, value FROM about WHERE id = ?`, id)

	var about entities.About
	if err := row.Scan(&about.Id, &about.Name, &about.Value); err != nil {
		panic(err)
	}
	return about
}


func Update(id int, about entities.About) bool {
	query, err := config.DB.Exec(`
		UPDATE about SET name = ?, value = ?, updated_at = ? WHERE id = ?`,
		about.Name, about.Value, about.UpdatedAt, id,
	)
	if err != nil {
		panic(err)
	}

	result, err := query.RowsAffected()
	if err != nil {
		panic(err)
	}

	return result > 0
}

func Delete(id int) error {
	_, err := config.DB.Exec(`DELETE FROM about WHERE id = ?`, id)
	return err
}



func FOTO() entities.File {
	var name = "photoabout";
	row := config.DB.QueryRow(`SELECT id, name, address FROM files WHERE name = ?`, name)

	var file entities.File
	if err := row.Scan(&file.Id, &file.Name, &file.Address); err != nil {
		panic(err)
	}

	return file
}