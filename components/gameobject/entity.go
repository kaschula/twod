package gameobject

type Entity interface {
	ID() string
}

func Remove[T Entity](entity T, entities []T) []T {
	if len(entities) == 0 {
		return entities
	}

	//if entity == nil {
	//	return entities
	//}

	var remaining []T
	for _, c := range entities {
		if c.ID() == entity.ID() {
			continue
		}
		remaining = append(remaining, c)
	}

	return remaining
}
