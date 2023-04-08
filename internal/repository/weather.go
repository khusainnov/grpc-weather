package repository

import "fmt"

const (
	cityTable = "city_table"
)

func (r *Repository) UploadWeather(city string) error {
	var exists bool

	querySelect := fmt.Sprintf("SELECT exists (SELECT 1 FROM %s WHERE city=$1);", cityTable)
	queryInsert := fmt.Sprintf("INSERT INTO %s (city, counter) VALUES ($1, 1);", cityTable)
	queryUpdate := fmt.Sprintf("UPDATE %s SET counter = counter + 1 WHERE city=$1;", cityTable)

	if err := r.db.QueryRow(querySelect, city).Scan(&exists); err != nil {
		return fmt.Errorf("error due select exists in %s, %w", cityTable, err)
	}

	if !exists {
		if _, err := r.db.Exec(queryInsert, city); err != nil {
			return fmt.Errorf("cannot insert city into %s, %w", cityTable, err)
		}

		return nil
	}

	if _, err := r.db.Exec(queryUpdate, city); err != nil {
		return fmt.Errorf("error due update %s, %w", cityTable, err)
	}

	return nil
}
