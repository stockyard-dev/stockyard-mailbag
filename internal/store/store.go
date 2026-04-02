package store
import ("database/sql";"fmt";"os";"path/filepath";"time";_ "modernc.org/sqlite")
type DB struct{db *sql.DB}
type Email struct {
	ID string `json:"id"`
	Subject string `json:"subject"`
	FromAddr string `json:"from_addr"`
	ToAddr string `json:"to_addr"`
	Body string `json:"body"`
	Status string `json:"status"`
	ThreadID string `json:"thread_id"`
	HasAttachment int `json:"has_attachment"`
	ReceivedAt string `json:"received_at"`
	CreatedAt string `json:"created_at"`
}
func Open(d string)(*DB,error){if err:=os.MkdirAll(d,0755);err!=nil{return nil,err};db,err:=sql.Open("sqlite",filepath.Join(d,"mailbag.db")+"?_journal_mode=WAL&_busy_timeout=5000");if err!=nil{return nil,err}
db.Exec(`CREATE TABLE IF NOT EXISTS emails(id TEXT PRIMARY KEY,subject TEXT NOT NULL,from_addr TEXT DEFAULT '',to_addr TEXT DEFAULT '',body TEXT DEFAULT '',status TEXT DEFAULT 'inbox',thread_id TEXT DEFAULT '',has_attachment INTEGER DEFAULT 0,received_at TEXT DEFAULT '',created_at TEXT DEFAULT(datetime('now')))`)
return &DB{db:db},nil}
func(d *DB)Close()error{return d.db.Close()}
func genID()string{return fmt.Sprintf("%d",time.Now().UnixNano())}
func now()string{return time.Now().UTC().Format(time.RFC3339)}
func(d *DB)Create(e *Email)error{e.ID=genID();e.CreatedAt=now();_,err:=d.db.Exec(`INSERT INTO emails(id,subject,from_addr,to_addr,body,status,thread_id,has_attachment,received_at,created_at)VALUES(?,?,?,?,?,?,?,?,?,?)`,e.ID,e.Subject,e.FromAddr,e.ToAddr,e.Body,e.Status,e.ThreadID,e.HasAttachment,e.ReceivedAt,e.CreatedAt);return err}
func(d *DB)Get(id string)*Email{var e Email;if d.db.QueryRow(`SELECT id,subject,from_addr,to_addr,body,status,thread_id,has_attachment,received_at,created_at FROM emails WHERE id=?`,id).Scan(&e.ID,&e.Subject,&e.FromAddr,&e.ToAddr,&e.Body,&e.Status,&e.ThreadID,&e.HasAttachment,&e.ReceivedAt,&e.CreatedAt)!=nil{return nil};return &e}
func(d *DB)List()[]Email{rows,_:=d.db.Query(`SELECT id,subject,from_addr,to_addr,body,status,thread_id,has_attachment,received_at,created_at FROM emails ORDER BY created_at DESC`);if rows==nil{return nil};defer rows.Close();var o []Email;for rows.Next(){var e Email;rows.Scan(&e.ID,&e.Subject,&e.FromAddr,&e.ToAddr,&e.Body,&e.Status,&e.ThreadID,&e.HasAttachment,&e.ReceivedAt,&e.CreatedAt);o=append(o,e)};return o}
func(d *DB)Update(e *Email)error{_,err:=d.db.Exec(`UPDATE emails SET subject=?,from_addr=?,to_addr=?,body=?,status=?,thread_id=?,has_attachment=?,received_at=? WHERE id=?`,e.Subject,e.FromAddr,e.ToAddr,e.Body,e.Status,e.ThreadID,e.HasAttachment,e.ReceivedAt,e.ID);return err}
func(d *DB)Delete(id string)error{_,err:=d.db.Exec(`DELETE FROM emails WHERE id=?`,id);return err}
func(d *DB)Count()int{var n int;d.db.QueryRow(`SELECT COUNT(*) FROM emails`).Scan(&n);return n}

func(d *DB)Search(q string, filters map[string]string)[]Email{
    where:="1=1"
    args:=[]any{}
    if q!=""{
        where+=" AND (subject LIKE ? OR body LIKE ?)"
        args=append(args,"%"+q+"%");args=append(args,"%"+q+"%");
    }
    if v,ok:=filters["status"];ok&&v!=""{where+=" AND status=?";args=append(args,v)}
    rows,_:=d.db.Query(`SELECT id,subject,from_addr,to_addr,body,status,thread_id,has_attachment,received_at,created_at FROM emails WHERE `+where+` ORDER BY created_at DESC`,args...)
    if rows==nil{return nil};defer rows.Close()
    var o []Email;for rows.Next(){var e Email;rows.Scan(&e.ID,&e.Subject,&e.FromAddr,&e.ToAddr,&e.Body,&e.Status,&e.ThreadID,&e.HasAttachment,&e.ReceivedAt,&e.CreatedAt);o=append(o,e)};return o
}

func(d *DB)Stats()map[string]any{
    m:=map[string]any{"total":d.Count()}
    rows,_:=d.db.Query(`SELECT status,COUNT(*) FROM emails GROUP BY status`)
    if rows!=nil{defer rows.Close();by:=map[string]int{};for rows.Next(){var s string;var c int;rows.Scan(&s,&c);by[s]=c};m["by_status"]=by}
    return m
}
