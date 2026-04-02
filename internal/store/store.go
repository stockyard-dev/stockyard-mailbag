package store
import ("database/sql";"fmt";"os";"path/filepath";"time";_ "modernc.org/sqlite")
type DB struct{db *sql.DB}
type Email struct{
	ID string `json:"id"`
	Subject string `json:"subject"`
	From_ string `json:"from_addr"`
	To_ string `json:"to_addr"`
	Body string `json:"body"`
	Date string `json:"date"`
	Labels string `json:"labels"`
	Read_ string `json:"is_read"`
	CreatedAt string `json:"created_at"`
}
func Open(d string)(*DB,error){if err:=os.MkdirAll(d,0755);err!=nil{return nil,err};db,err:=sql.Open("sqlite",filepath.Join(d,"mailbag.db")+"?_journal_mode=WAL&_busy_timeout=5000");if err!=nil{return nil,err}
db.Exec(`CREATE TABLE IF NOT EXISTS emails(id TEXT PRIMARY KEY,subject TEXT NOT NULL,from_addr TEXT DEFAULT '',to_addr TEXT DEFAULT '',body TEXT DEFAULT '',date TEXT DEFAULT '',labels TEXT DEFAULT '',is_read TEXT DEFAULT 'false',created_at TEXT DEFAULT(datetime('now')))`)
return &DB{db:db},nil}
func(d *DB)Close()error{return d.db.Close()}
func genID()string{return fmt.Sprintf("%d",time.Now().UnixNano())}
func now()string{return time.Now().UTC().Format(time.RFC3339)}
func(d *DB)Create(e *Email)error{e.ID=genID();e.CreatedAt=now();_,err:=d.db.Exec(`INSERT INTO emails(id,subject,from_addr,to_addr,body,date,labels,is_read,created_at)VALUES(?,?,?,?,?,?,?,?,?)`,e.ID,e.Subject,e.From_,e.To_,e.Body,e.Date,e.Labels,e.Read_,e.CreatedAt);return err}
func(d *DB)Get(id string)*Email{var e Email;if d.db.QueryRow(`SELECT id,subject,from_addr,to_addr,body,date,labels,is_read,created_at FROM emails WHERE id=?`,id).Scan(&e.ID,&e.Subject,&e.From_,&e.To_,&e.Body,&e.Date,&e.Labels,&e.Read_,&e.CreatedAt)!=nil{return nil};return &e}
func(d *DB)List()[]Email{rows,_:=d.db.Query(`SELECT id,subject,from_addr,to_addr,body,date,labels,is_read,created_at FROM emails ORDER BY created_at DESC`);if rows==nil{return nil};defer rows.Close();var o []Email;for rows.Next(){var e Email;rows.Scan(&e.ID,&e.Subject,&e.From_,&e.To_,&e.Body,&e.Date,&e.Labels,&e.Read_,&e.CreatedAt);o=append(o,e)};return o}
func(d *DB)Delete(id string)error{_,err:=d.db.Exec(`DELETE FROM emails WHERE id=?`,id);return err}
func(d *DB)Count()int{var n int;d.db.QueryRow(`SELECT COUNT(*) FROM emails`).Scan(&n);return n}
