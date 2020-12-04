package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/asdine/storm/v3"
	migrate "github.com/rubenv/sql-migrate"
	"github.com/ruinnel/giregi.rip-server/common"
	"github.com/ruinnel/giregi.rip-server/domain"
	_archiveRepository "github.com/ruinnel/giregi.rip-server/feature/archive/repository/mysql"
	_siteRepository "github.com/ruinnel/giregi.rip-server/feature/site/repository/mysql"
	_tagRepository "github.com/ruinnel/giregi.rip-server/feature/tag/repository/mysql"
	_tokenRepository "github.com/ruinnel/giregi.rip-server/feature/token/repository/mysql"
	_userRepository "github.com/ruinnel/giregi.rip-server/feature/user/repository/mysql"
	_webPageRepository "github.com/ruinnel/giregi.rip-server/feature/webpage/repository/mysql"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"sync"
)

type Database string

const (
	DATABASE_MYSQL = Database("mysql")
	DATABASE_BOLT  = Database("bolt")
)

var (
	driversMu sync.RWMutex
	dialect             = DATABASE_MYSQL
	mysql     *sql.DB   = nil
	bolt      *storm.DB = nil
)

func Use(config *common.Config) error {
	driversMu.Lock()
	defer driversMu.Unlock()

	dialectName := config.Database.Dialect
	if !(dialectName == string(DATABASE_MYSQL) || dialectName == string(DATABASE_BOLT)) {
		return errors.New("unknown dialect")
	}
	dialect = Database(dialectName)
	err := checkConfig(config)
	if err != nil {
		return err
	}
	switch dialect {
	case DATABASE_MYSQL:
		initMysql(config)
	case DATABASE_BOLT:
		initBolt(config)
	}
	return nil
}

func Disconnect() {
	logger := common.GetLogger()
	if mysql != nil {
		err := mysql.Close()
		if err != nil {
			logger.Fatal(err)
		}
	}
	if bolt != nil {
		err := bolt.Close()
		if err != nil {
			logger.Fatal(err)
		}
	}
}

func initMysql(config *common.Config) {
	mysql = common.OpenDatabase(config.Database)
	migrateDatabase(config, mysql)
	boil.SetDB(mysql)
	// boil.DebugMode = true
}

func initBolt(config *common.Config) {
	logger := common.GetLogger()
	dbFile := config.Database.File
	db, err := storm.Open(dbFile)
	if err != nil {
		logger.Panicf("failed to connect database - %s", dbFile)
	}
	bolt = db
}

func checkConfig(config *common.Config) error {
	cfg := config.Database
	if dialect == DATABASE_MYSQL {
		if len(cfg.Host) == 0 {
			return errors.New("require `Host`")
		}
		if cfg.Port == 0 {
			return errors.New("require `Port`")
		}
		if len(cfg.Name) == 0 {
			return errors.New("require `Name`")
		}
		if len(cfg.Username) == 0 {
			return errors.New("require `Username`")
		}
		if len(cfg.Password) == 0 {
			return errors.New("require `Password`")
		}
		if len(cfg.SQLMigrateSourcePath) == 0 {
			return errors.New("require `SQLMigrateSourcePath`")
		}
	} else if dialect == DATABASE_BOLT {
		if len(cfg.File) == 0 {
			return errors.New("require `File`")
		}
	}
	return nil
}

func migrateDatabase(config *common.Config, db *sql.DB) {
	logger := common.GetLogger()
	srcPath := config.Database.SQLMigrateSourcePath
	source := migrate.FileMigrationSource{
		Dir: srcPath,
	}
	applyCount, err := migrate.Exec(db, "mysql", source, migrate.Up)
	if err != nil {
		panic(fmt.Sprintf("error: migration source(%s) not found. - %v", srcPath, err))
	}
	logger.Printf("migrate complete - %v", applyCount)
}

func User() domain.UserRepository {
	switch dialect {
	case DATABASE_MYSQL:
		return _userRepository.NewUserRepository(mysql)
	case DATABASE_BOLT:
		return nil
	default:
		return nil
	}
}

func Archive() domain.ArchiveRepository {
	switch dialect {
	case DATABASE_MYSQL:
		return _archiveRepository.NewArchiveRepository(mysql)
	case DATABASE_BOLT:
		return nil
	default:
		return nil
	}
}

func Token() domain.TokenRepository {
	switch dialect {
	case DATABASE_MYSQL:
		return _tokenRepository.NewTokenRepository(mysql)
	case DATABASE_BOLT:
		return nil
	default:
		return nil
	}
}

func Site() domain.SiteRepository {
	switch dialect {
	case DATABASE_MYSQL:
		return _siteRepository.NewSiteRepository(mysql)
	case DATABASE_BOLT:
		return nil
	default:
		return nil
	}
}

func WebPage() domain.WebPageRepository {
	switch dialect {
	case DATABASE_MYSQL:
		return _webPageRepository.NewWebPageRepository(mysql)
	case DATABASE_BOLT:
		return nil
	default:
		return nil
	}
}

func Tag() domain.TagRepository {
	switch dialect {
	case DATABASE_MYSQL:
		return _tagRepository.NewTagRepository(mysql)
	case DATABASE_BOLT:
		return nil
	default:
		return nil
	}
}
