package db

func QueryAccessTokens() ([]string, error) {
	var tokens []string
	s := "SELECT * FROM accessTokens"
	if rows, err := sdb.Query(s); err != nil {
		return nil, err
	} else {
		for rows.Next() {
			var uid int
			var token string
			if err := rows.Scan(&uid, &token); err != nil {
				continue
			}
			tokens = append(tokens, token)
		}
	}
	return tokens, nil
}

func AddAccessToken(token string) error {
	s := "INSERT INTO accessTokens(token) values(?)"
	if stmt, err := sdb.Prepare(s); err != nil {
		return err
	} else {
		if _, err := stmt.Exec(token); err != nil {
			return err
		}
	}
	return nil
}

func DelAccessToken(token string) error {
	s := "DELETE FROM accessTokens where token=?"
	if stmt, err := sdb.Prepare(s); err != nil {
		return err
	} else {
		if _, err := stmt.Exec(token); err != nil {
			return err
		}
	}
	return nil
}

func UpdateAccessToken(oldToken, newToken string) error {
	s := "UPDATE accessTokens set token=? where token=?"
	if stmt, err := sdb.Prepare(s); err != nil {
		return err
	} else {
		if _, err := stmt.Exec(newToken, oldToken); err != nil {
			return err
		}
	}
	return nil
}
