package databease

type TransactionDBTestSuite struct {
	suite.Suite
	DB * sql.DB
	client *entity.Client
	accountFrom *entity.Account
	client2 *entity.Client
	accountTo *entity.Account
	TransactionDB *TransactionDB
}


func (s *TransactionDBTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	s.Nil(err)
	s.DB = db
	db.Exec("CREATE TABLE clients (id varchar(255) PRIMARY KEY, name varchar(255), email varchar(255), created_at date)")
	db.Exec("CREATE TABLE accounts (id varchar(255) PRIMARY KEY, client_id varchar(255), balance decimal, created_at date)")
	db.Exec("CREATE TABLE transactions (id varchar(255) PRIMARY KEY, account_from_id varchar(255), account_to_id varchar(255), amount decimal, created_at date)")
	s.TransactionDB = NewTransactionDB(db)
	s.client, _ = entity.NewClient("John Doe", "j@j.com", 