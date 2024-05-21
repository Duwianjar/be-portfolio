package profileModal

import (
	"be-portfolio/config"
	"be-portfolio/entities"
	"time"
)

func GetAll() []entities.Profile {
	rows, err := config.DB.Query(`SELECT * FROM profile`)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var profiles []entities.Profile

	for rows.Next() {
		var profile entities.Profile
		if err := rows.Scan(&profile.Id, &profile.Name, &profile.Value, &profile.CreatedAt, &profile.UpdatedAt); err != nil {
			panic(err)
		}

		profiles = append(profiles, profile)
	}

	return profiles
}

func Create(profile entities.Profile) bool {
	result, err := config.DB.Exec(`
		INSERT INTO profile (name, value, created_at, updated_at)
		VALUE (? , ? , ? , ?)`,
		profile.Name, profile.Value, profile.CreatedAt, profile.UpdatedAt,
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


func Detail(id int) entities.Profile {
	row := config.DB.QueryRow(`SELECT id, name, value FROM profile WHERE id = ?`, id)

	var profile entities.Profile
	if err := row.Scan(&profile.Id, &profile.Name, &profile.Value); err != nil {
		panic(err)
	}

	return profile
}


func Update(id int, profile entities.Profile) bool {
	query, err := config.DB.Exec(`
		UPDATE profile SET name = ?, value = ?, updated_at = ? WHERE id = ?`,
		profile.Name, profile.Value, profile.UpdatedAt, id,
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
	_, err := config.DB.Exec(`DELETE FROM profile WHERE id = ?`, id)
	return err
}


func CV() entities.File {
	var name = "cv";
	row := config.DB.QueryRow(`SELECT id, name, address FROM files WHERE name = ?`, name)

	var file entities.File
	if err := row.Scan(&file.Id, &file.Name, &file.Address); err != nil {
		panic(err)
	}

	return file
}

func PP() entities.File {
	var name = "photoprofile";
	row := config.DB.QueryRow(`SELECT id, name, address FROM files WHERE name = ?`, name)

	var file entities.File
	if err := row.Scan(&file.Id, &file.Name, &file.Address); err != nil {
		panic(err)
	}

	return file
}


func UpdateCV(address string) bool {
	var name = "cv";
	query, err := config.DB.Exec(`
		UPDATE files SET address = ?, updated_at = ? WHERE name = ?`,
		address, time.Now(), name,
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

func UpdatePP(address string) bool {
	var name = "photoprofile";
	query, err := config.DB.Exec(`
		UPDATE files SET address = ?, updated_at = ? WHERE name = ?`,
		address, time.Now(), name,
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


