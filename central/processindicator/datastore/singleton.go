package datastore

import (
	"time"

	"github.com/pkg/errors"
	"github.com/stackrox/rox/central/globaldb"
	"github.com/stackrox/rox/central/globalindex"
	"github.com/stackrox/rox/central/processindicator/index"
	"github.com/stackrox/rox/central/processindicator/pruner"
	"github.com/stackrox/rox/central/processindicator/search"
	"github.com/stackrox/rox/central/processindicator/store"
	pgStore "github.com/stackrox/rox/central/processindicator/store/postgres"
	"github.com/stackrox/rox/central/processindicator/store/rocksdb"
	plopStore "github.com/stackrox/rox/central/processlisteningonport/store/postgres"
	"github.com/stackrox/rox/pkg/env"
	"github.com/stackrox/rox/pkg/logging"
	"github.com/stackrox/rox/pkg/sync"
	"github.com/stackrox/rox/pkg/utils"
)

const (
	pruneInterval     = 10 * time.Minute
	minArgsPerProcess = 5
)

var (
	once sync.Once

	ad DataStore

	log = logging.LoggerForModule()
)

func initialize() {
	var storage store.Store
	var plopStorage plopStore.Store
	var indexer index.Indexer
	if env.PostgresDatastoreEnabled.BooleanSetting() {
		storage = pgStore.New(globaldb.GetPostgres())
		plopStorage = plopStore.New(globaldb.GetPostgres())
		indexer = pgStore.NewIndexer(globaldb.GetPostgres())
	} else {
		storage = rocksdb.New(globaldb.GetRocksDB())
		// PLOP storage is only supported for PostgreSQL
		indexer = index.New(globalindex.GetProcessIndex())
	}
	searcher := search.New(storage, indexer)

	p := pruner.NewFactory(minArgsPerProcess, pruneInterval)

	var err error
	ad, err = New(storage, plopStorage, indexer, searcher, p)
	utils.CrashOnError(errors.Wrap(err, "unable to load datastore for process indicators"))
}

// Singleton provides the interface for non-service external interaction.
func Singleton() DataStore {
	once.Do(initialize)
	return ad
}
