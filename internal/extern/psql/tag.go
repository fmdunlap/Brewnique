package psql

import "brewnique.fdunlap.com/internal/data"

type TagDbRow struct {
	Id   int64
	Name string
}

func (t *TagDbRow) ToTag() data.Tag {
	return data.Tag{
		Id:   t.Id,
		Name: t.Name,
	}
}

func (p PostgresProvider) GetTags() ([]*data.Tag, error) {
	rows, err := p.db.Query("SELECT id, name FROM tags")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tags []*data.Tag
	for rows.Next() {
		var tagRow TagDbRow
		err = rows.Scan(&tagRow.Id, &tagRow.Name)
		if err != nil {
			return nil, err
		}
		tag := tagRow.ToTag()
		tags = append(tags, &tag)
	}

	return tags, nil
}
