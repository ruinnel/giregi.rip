package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/asdine/storm/v3"
	"github.com/asdine/storm/v3/codec/protobuf"
	"github.com/go-redis/redis/v8"
	"github.com/patrickmn/go-cache"
	migrate "github.com/rubenv/sql-migrate"
	"github.com/ruinnel/giregi.rip-server/common"
	"github.com/ruinnel/giregi.rip-server/domain"
	"time"

	_archiveBoltRepository "github.com/ruinnel/giregi.rip-server/feature/archive/repository/bolt"
	_siteBoltRepository "github.com/ruinnel/giregi.rip-server/feature/site/repository/bolt"
	_tagBoltRepository "github.com/ruinnel/giregi.rip-server/feature/tag/repository/bolt"
	_tokenBoltRepository "github.com/ruinnel/giregi.rip-server/feature/token/repository/bolt"
	_userBoltRepository "github.com/ruinnel/giregi.rip-server/feature/user/repository/bolt"
	_webPageBoltRepository "github.com/ruinnel/giregi.rip-server/feature/webpage/repository/bolt"

	_archiveMysqlRepository "github.com/ruinnel/giregi.rip-server/feature/archive/repository/mysql"
	_siteMysqlRepository "github.com/ruinnel/giregi.rip-server/feature/site/repository/mysql"
	_tagMysqlRepository "github.com/ruinnel/giregi.rip-server/feature/tag/repository/mysql"
	_tokenMysqlRepository "github.com/ruinnel/giregi.rip-server/feature/token/repository/mysql"
	_userMysqlRepository "github.com/ruinnel/giregi.rip-server/feature/user/repository/mysql"
	_webPageMysqlRepository "github.com/ruinnel/giregi.rip-server/feature/webpage/repository/mysql"

	_archiveMemoryCache "github.com/ruinnel/giregi.rip-server/feature/archive/cache/memory"
	_archiveRedisCache "github.com/ruinnel/giregi.rip-server/feature/archive/cache/redis"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"sync"
)

var (
	driversMu   sync.RWMutex
	platform                  = common.PLATFORM_SERVER
	mysql       *sql.DB       = nil
	bolt        *storm.DB     = nil
	redisClient *redis.Client = nil
)

func Use(config *common.Config) error {
	logger := common.GetLogger()
	driversMu.Lock()
	defer driversMu.Unlock()

	err := checkConfig(config)
	if err != nil {
		return err
	}

	platform = config.Platform
	logger.Printf("use - %v", platform)
	switch platform {
	case "server":
		initMysql(config)
		redisClient = common.OpenRedis(config.Redis)
	case "desktop":
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
	mysql = common.OpenDatabase(config.Mysql)
	migrateMysql(config, mysql)
	boil.SetDB(mysql)
	// boil.DebugMode = true
}

func initBolt(config *common.Config) {
	logger := common.GetLogger()
	dbFile := config.Bolt.File
	db, err := storm.Open(dbFile, storm.Codec(protobuf.Codec))
	if err != nil {
		logger.Panicf("failed to connect database - %s", dbFile)
	}
	migrateBolt(config, db)
	bolt = db
}

func checkConfig(config *common.Config) error {
	if config.Platform == common.PLATFORM_SERVER {
		mysql := config.Mysql
		if len(mysql.Host) == 0 {
			return errors.New("require `Host`")
		}
		if mysql.Port == 0 {
			return errors.New("require `Port`")
		}
		if len(mysql.Name) == 0 {
			return errors.New("require `Name`")
		}
		if len(mysql.Username) == 0 {
			return errors.New("require `Username`")
		}
		if len(mysql.Password) == 0 {
			return errors.New("require `Password`")
		}
		if len(mysql.SQLMigrateSourcePath) == 0 {
			return errors.New("require `SQLMigrateSourcePath`")
		}
	} else if config.Platform == common.PLATFORM_DESKTOP {
		if len(config.Bolt.File) == 0 {
			return errors.New("require `File`")
		}
	}
	return nil
}

func migrateMysql(config *common.Config, db *sql.DB) {
	logger := common.GetLogger()
	srcPath := config.Mysql.SQLMigrateSourcePath
	source := migrate.FileMigrationSource{
		Dir: srcPath,
	}
	applyCount, err := migrate.Exec(db, "mysql", source, migrate.Up)
	if err != nil {
		panic(fmt.Sprintf("error: migration source(%s) not found. - %v", srcPath, err))
	}
	logger.Printf("migrate complete - %v", applyCount)
}

func migrateBolt(config *common.Config, db *storm.DB) {
	logger := common.GetLogger()

	err := db.Init(&domain.User{})
	if err != nil {
		panic("error: migrate bolt('user') fail.")
		return
	}
	err = db.Init(&domain.Token{})
	if err != nil {
		panic("error: migrate bolt('token') fail.")
		return
	}
	err = db.Init(&domain.Tag{})
	if err != nil {
		panic("error: migrate bolt('tag') fail.")
		return
	}
	err = db.Init(&domain.Site{})
	if err != nil {
		panic("error: migrate bolt('Site') fail.")
		return
	}
	err = db.Init(&domain.WebPage{})
	if err != nil {
		panic("error: migrate bolt('WebPage') fail.")
		return
	}
	err = db.Init(&domain.Archive{})
	if err != nil {
		panic("error: migrate bolt('Archive') fail.")
		return
	}
	err = db.Init(&domain.ArchiveTagMapping{})
	if err != nil {
		panic("error: migrate bolt('ArchiveTagMapping') fail.")
		return
	}
	logger.Printf("migrate bolt complete")
}

func User() domain.UserRepository {
	switch platform {
	case common.PLATFORM_SERVER:
		return _userMysqlRepository.NewUserRepository(mysql)
	case common.PLATFORM_DESKTOP:
		return _userBoltRepository.NewUserRepository(bolt)
	default:
		return nil
	}
}

func Archive() domain.ArchiveRepository {
	switch platform {
	case common.PLATFORM_SERVER:
		return _archiveMysqlRepository.NewArchiveRepository(mysql)
	case common.PLATFORM_DESKTOP:
		return _archiveBoltRepository.NewArchiveRepository(bolt)
	default:
		return nil
	}
}

func ArchiveCache() domain.ArchiveCache {
	switch platform {
	case common.PLATFORM_SERVER:
		return _archiveRedisCache.NewArchiveCache(redisClient)
	case common.PLATFORM_DESKTOP:
		return _archiveMemoryCache.NewArchiveCache(cache.New(5*time.Minute, 10*time.Minute))
	default:
		return nil
	}
}

func Token() domain.TokenRepository {
	switch platform {
	case common.PLATFORM_SERVER:
		return _tokenMysqlRepository.NewTokenRepository(mysql)
	case common.PLATFORM_DESKTOP:
		return _tokenBoltRepository.NewTokenRepository(bolt)
	default:
		return nil
	}
}

func Site() domain.SiteRepository {
	switch platform {
	case common.PLATFORM_SERVER:
		return _siteMysqlRepository.NewSiteRepository(mysql)
	case common.PLATFORM_DESKTOP:
		return _siteBoltRepository.NewSiteRepository(bolt)
	default:
		return nil
	}
}

func WebPage() domain.WebPageRepository {
	switch platform {
	case common.PLATFORM_SERVER:
		return _webPageMysqlRepository.NewWebPageRepository(mysql)
	case common.PLATFORM_DESKTOP:
		return _webPageBoltRepository.NewWebPageRepository(bolt)
	default:
		return nil
	}
}

func Tag() domain.TagRepository {
	switch platform {
	case common.PLATFORM_SERVER:
		return _tagMysqlRepository.NewTagRepository(mysql)
	case common.PLATFORM_DESKTOP:
		return _tagBoltRepository.NewTagRepository(bolt)
	default:
		return nil
	}
}
