package port

type DB interface {
	CloseConnection() error
	Save(arg interface{}) error
	List(dest interface{}, conditions map[string]interface{}, limit int, offset int) error
	FindOne(dest interface{}, conditions map[string]interface{}) error
	DeleteOne(model interface{}, condition map[string]interface{}) error
	Raw(dest interface{}, query string, values ...interface{}) error
	FindById(dest interface{}, id int64) error
	Count(model interface{}, condition map[string]interface{}) (int64, error)
}
