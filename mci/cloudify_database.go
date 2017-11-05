package mci

import (
	"errors"
	"fmt"
	"github.com/cloudfoundry-community/gautocloud/connectors/databases/dbtype"
	"github.com/cloudfoundry-community/gautocloud/connectors/databases/raw"
	"github.com/cloudfoundry-community/gautocloud/loader"
	"strconv"
)

func cloudifyDatabase(loader loader.Loader, mattermostConfig *MattermostConfig) error {
	loader.RegisterConnector(raw.NewMysqlRawConnector())
	loader.RegisterConnector(raw.NewPostgresqlRawConnector())
	var postSvc dbtype.PostgresqlDatabase
	err := loader.Inject(&postSvc)

	if err == nil {
		mattermostConfig.SqlSettings.DriverName = "postgres"
		mattermostConfig.SqlSettings.DataSource = fmt.Sprintf(
			"postgres://%s:%s@%s/%s?sslmode=disable&connect_timeout=10",
			postSvc.User,
			postSvc.Password,
			postSvc.Host+":"+strconv.Itoa(postSvc.Port),
			postSvc.Database)
		return nil
	}
	var mySvc dbtype.MysqlDatabase
	err = loader.Inject(&mySvc)
	if err != nil {
		return errors.New("Cannot find any mysql or postgres connected.")
	}
	mattermostConfig.SqlSettings.DriverName = "mysql"
	mattermostConfig.SqlSettings.DataSource = fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8mb4,utf8",
		mySvc.User,
		mySvc.Password,
		mySvc.Host+":"+strconv.Itoa(mySvc.Port),
		mySvc.Database)
	return nil
}
