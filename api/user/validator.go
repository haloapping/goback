package user

import "unicode"

func RegisterValidation(req UserRegisterReq) map[string][]string {
	validationErr := make(map[string][]string)

	if req.Username == "" {
		validationErr["username"] = append(validationErr["username"], "cannot empty")
	}

	nLowercase := 0
	nUppercase := 0
	nDigit := 0
	nPunctuation := 0
	for _, r := range req.Password {
		if unicode.IsLower(r) {
			nLowercase++
		}
		if unicode.IsUpper(r) {
			nUppercase++
		}
		if unicode.IsDigit(r) {
			nDigit++
		}
		if unicode.IsPunct(r) {
			nPunctuation++
		}
	}
	if len(req.Password) < 8 {
		validationErr["password"] = append(validationErr["password"], "min 8 character")
	}
	if nLowercase < 1 {
		validationErr["passoword"] = append(validationErr["passowrd"], "minimal 1 lowercase")
	}
	if nUppercase < 1 {
		validationErr["password"] = append(validationErr["password"], "minimal 1 uppercase")
	}
	if nDigit < 1 {
		validationErr["password"] = append(validationErr["passoword"], "minimal 1 digit")
	}
	if nPunctuation < 1 {
		validationErr["password"] = append(validationErr["password"], "minimal 1 punctuation")
	}

	if req.Password != req.ConfirmPassword {
		validationErr["confirmPassword"] = append(validationErr["confirmPassword"], "confirm password is not match with password")
	}

	return validationErr
}

func LoginValidation(req UserLoginReq) map[string][]string {
	validationErr := make(map[string][]string)

	if req.Username == "" {
		validationErr["username"] = append(validationErr["username"], "cannot empty")
	}

	if req.Password == "" {
		validationErr["password"] = append(validationErr["password"], "cannot empty")
	}

	return validationErr
}
