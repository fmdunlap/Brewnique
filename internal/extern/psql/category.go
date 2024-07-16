package psql

import "brewnique.fdunlap.com/internal/data"

type CategoryDbRow struct {
	Id       int64
	Name     string
	ParentId *int64
}

func (c *CategoryDbRow) ToCategory() data.Category {
	return data.Category{
		Id:       c.Id,
		Name:     c.Name,
		ParentId: c.ParentId,
	}
}

func (p PostgresProvider) GetCategory(id int64) (*data.Category, error) {
	res := p.db.QueryRow("SELECT id, name, parent_id FROM categories WHERE id = $1", id)

	categoryRow := CategoryDbRow{}
	err := res.Scan(&categoryRow)
	if err != nil {
		return nil, err
	}

	category := categoryRow.ToCategory()

	return &category, nil
}

func (p PostgresProvider) ListCategories() ([]*data.Category, error) {
	rows, err := p.db.Query("SELECT id, name, parent_id FROM categories WHERE parent_id IS NULL")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []*data.Category
	for rows.Next() {
		var categoryRow CategoryDbRow
		err = rows.Scan(&categoryRow.Id, &categoryRow.Name, &categoryRow.ParentId)
		if err != nil {
			return nil, err
		}
		category := categoryRow.ToCategory()
		categories = append(categories, &category)
	}

	return categories, nil
}

func (p PostgresProvider) ListSubcategories(categoryId int64) ([]*data.Category, error) {
	rows, err := p.db.Query("SELECT id, name, parent_id FROM categories WHERE parent_id = $1", categoryId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []*data.Category
	for rows.Next() {
		var categoryRow CategoryDbRow
		err = rows.Scan(&categoryRow.Id, &categoryRow.Name, &categoryRow.ParentId)
		if err != nil {
			return nil, err
		}
		category := categoryRow.ToCategory()
		categories = append(categories, &category)
	}

	return categories, nil
}
