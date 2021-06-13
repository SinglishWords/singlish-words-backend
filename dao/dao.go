package dao

type notConnectedError struct{}

func (notConnectedError) Error() string {
	return "Not connect to mysql database."
}

type insertError struct{}

func (insertError) Error() string {
	return "Failed insert into database."
}
