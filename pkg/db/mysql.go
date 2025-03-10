// Copyright (c) 2022 Silverton Data, Inc.
// You may use, distribute, and modify this code under the terms of the Apache-2.0 license, a copy of
// which may be found at https://github.com/silverton-io/buz/blob/main/LICENSE

package db

import "strconv"

// GenerateMysqlDsn generates a Mysql Dsn from the provided connection params
func GenerateMysqlDsn(params ConnectionParams) string {
	port := strconv.FormatUint(uint64(params.Port), 10)
	return params.User + ":" + params.Pass + "@tcp(" + params.Host + ":" + port + ")/" + params.Db + "?parseTime=true"
}
