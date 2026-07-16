package main

// performAction actually executes the subcommand's intended action, for example it posts
// to the endpoint that creates a new secret.
func performAction(c ClientConfig) (string, error) {
	switch c.Action {
	case ActionCreate:
		result, err := createSecret(c.URL, c.Data)
		if err != nil {
			return "", err
		}
		return result.Id, nil
	case ActionView:
		result, err := getSecret(c.URL, c.Id)
		if err != nil {
			return "", err
		}
		return result.Data, nil
	default:
		return "", ErrInvalidAction
	}
}
