package skillModal

import (
	"be-portfolio/config"
	"be-portfolio/entities"
)

func GetAll() []entities.Skill {
	rows, err := config.DB.Query(`SELECT * FROM skill`)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var skills []entities.Skill

	for rows.Next() {
		var skill entities.Skill
		if err := rows.Scan(&skill.Id, &skill.Name, &skill.Level, &skill.Category, &skill.CreatedAt, &skill.UpdatedAt); err != nil {
			panic(err)
		}

		skills = append(skills, skill)
	}

	return skills
}

func Create(skill entities.Skill) bool {
	result, err := config.DB.Exec(`
		INSERT INTO skill (name, level, category, created_at, updated_at)
		VALUE (? , ? , ? , ?, ?)`,
		skill.Name, skill.Level, skill.Category, skill.CreatedAt, skill.UpdatedAt,
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

func Detail(id int) entities.Skill {
	row := config.DB.QueryRow(`SELECT id, name, level, category FROM skill WHERE id = ?`, id)

	var skill entities.Skill
	if err := row.Scan(&skill.Id, &skill.Name, &skill.Level, &skill.Category); err != nil {
		panic(err)
	}
	return skill
}

func Update(id int, skill entities.Skill) bool {
	query, err := config.DB.Exec(`
		UPDATE skill SET name = ?, level = ?, category = ?, updated_at = ? WHERE id = ?`,
		skill.Name, skill.Level, skill.Category, skill.UpdatedAt, id,
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
	_, err := config.DB.Exec(`DELETE FROM skill WHERE id = ?`, id)
	return err
}