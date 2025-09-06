package task

func AddValidation(req AddReq) map[string][]string {
	validationErr := make(map[string][]string)

	if req.UserId == "" {
		validationErr["userId"] = append(validationErr["userId"], "cannot empty")
	}

	if req.Title == "" {
		validationErr["title"] = append(validationErr["title"], "cannot empty")
	}

	if req.Description == "" {
		validationErr["description"] = append(validationErr["description"], "cannot empty")
	}

	return validationErr
}
