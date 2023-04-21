package tools

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"

	_ "github.com/go-sql-driver/mysql"
)

func LoginDiscuzPassword(password string, salt string) (str string) {
	hash := md5.Sum([]byte(password))
	hashString := hex.EncodeToString(hash[:])

	finalHash := md5.Sum([]byte(hashString + salt))
	finalHashString := hex.EncodeToString(finalHash[:])

	return finalHashString
}

func LoginCheck(pass string, username string) bool {
	// 连接数据库
	db, err := sql.Open("mysql", "new_moebuss_com:WYi8ExaTEz@tcp(198.211.44.35:3306)/new_moebuss_com")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// 查询指定 username 的 password 和 salt
	var password string
	var salt string
	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM bbs_ucenter_members WHERE username=?", username).Scan(&count)
	if err != nil {
		panic(err.Error())
	}
	if count > 0 {
		err = db.QueryRow("SELECT password, salt FROM bbs_ucenter_members WHERE username=?", username).Scan(&password, &salt)
		if err != nil {
			panic(err.Error())
		}

		if password == LoginDiscuzPassword(pass, salt) {
			return true
		}
	} else {
		return false
	}
	return false
}
