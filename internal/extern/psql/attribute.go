package psql

import "brewnique.fdunlap.com/internal/data"

type AttributeDbRow struct {
	Id   int64
	Name string
	Type string
}

type AttributeValueDbRow struct {
	Id          int64
	AttributeId int64
	Value       string
}

type RecipeAttributeDbRow struct {
	Id          int64
	RecipeId    int64
	AttributeId int64
}

func (a *AttributeDbRow) ToAttribute() data.Attribute {
	return data.Attribute{
		Id:     a.Id,
		Name:   a.Name,
		Type:   a.Type,
		Values: make([]data.AttributeValue, 0),
	}
}

func (a *AttributeValueDbRow) ToAttributeValue() data.AttributeValue {
	return data.AttributeValue{
		Id:    a.Id,
		Value: a.Value,
	}
}

func (p *PostgresProvider) GetAttribute(id int64) (*data.Attribute, error) {
	res := p.db.QueryRow("SELECT id, name, type FROM attributes WHERE id = $1", id)

	attributeRow := AttributeDbRow{}
	err := res.Scan(&attributeRow)
	if err != nil {
		return nil, err
	}

	attribute := attributeRow.ToAttribute()

	return &attribute, nil
}

func (p *PostgresProvider) GetAttributeValues(attributeId int64) ([]*data.AttributeValue, error) {
	rows, err := p.db.Query("SELECT id, attribute_id, value FROM attribute_values WHERE attribute_id = $1", attributeId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var attributeValues []*data.AttributeValue
	for rows.Next() {
		var attributeValueRow AttributeValueDbRow
		err = rows.Scan(&attributeValueRow.Id, &attributeValueRow.AttributeId, &attributeValueRow.Value)
		if err != nil {
			return nil, err
		}
		attributeValue := attributeValueRow.ToAttributeValue()
		attributeValues = append(attributeValues, &attributeValue)
	}

	return attributeValues, nil
}

func (p *PostgresProvider) GetAttributes() ([]*data.Attribute, error) {
	rows, err := p.db.Query("SELECT id, name, type FROM attributes")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var attributes []*data.Attribute
	for rows.Next() {
		var attributeRow AttributeDbRow
		err = rows.Scan(&attributeRow.Id, &attributeRow.Name, &attributeRow.Type)
		if err != nil {
			return nil, err
		}
		attribute := attributeRow.ToAttribute()
		attributes = append(attributes, &attribute)
	}

	return attributes, nil
}
