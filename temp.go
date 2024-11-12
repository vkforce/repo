func IsMessageValid(message kafka.Message, ppcode string, filters map[string]string) bool {
	var collectionName string
	if filters != nil {
		collectionName = filters["CollectionName"]
	}

	var ppcodeMatch, collectionNameMatch bool

	for _, header := range message.Headers {
		if header.Key == "PPCODE" && string(header.Value) == ppcode {
			ppcodeMatch = true
		}
		if filters != nil && collectionName != "" && header.Key == "CollectionName" && string(header.Value) == collectionName {
			collectionNameMatch = true
		}
		if ppcodeMatch && (filters == nil || collectionNameMatch) {
			return true
		}
	}

	return false
}