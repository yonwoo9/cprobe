package mysql

import (
	"context"
	"fmt"

	"github.com/BurntSushi/toml"
	"github.com/cprobe/cprobe/exporter/mysql/collector"
	"github.com/cprobe/cprobe/lib/logger"
	"github.com/cprobe/cprobe/types"
	"github.com/prometheus/client_golang/prometheus"
)

type Global struct {
	User                  string   `toml:"user"`
	Password              string   `toml:"password"`
	SslCa                 string   `toml:"ssl_ca"`
	SslCert               string   `toml:"ssl_cert"`
	SslKey                string   `toml:"ssl_key"`
	TlsInsecureSkipVerify bool     `toml:"ssl_skip_verfication"`
	Tls                   string   `toml:"tls"`
	ScraperEnabled        []string `toml:"scraper_enabled"`
}

type Config struct {
	Global  *Global                 `toml:"global"`
	Queries []collector.CustomQuery `toml:"queries"`

	CollectGlobalStatus struct {
		Enabled bool `toml:"enabled"`
	} `toml:"collect_global_status"`
	CollectGlobalVariables struct {
		Enabled bool `toml:"enabled"`
	} `toml:"collect_global_variables"`
	CollectSlaveStatus struct {
		Enabled bool `toml:"enabled"`
	} `toml:"collect_slave_status"`
	CollectInfoSchemaInnodbCmp struct {
		Enabled bool `toml:"enabled"`
	} `toml:"collect_info_schema_innodb_cmp"`
	CollectInfoSchemaInnodbCmpmem struct {
		Enabled bool `toml:"enabled"`
	} `toml:"collect_info_schema_innodb_cmpmem"`
	CollectInfoSchemaQueryResponseTime struct {
		Enabled bool `toml:"enabled"`
	} `toml:"collect_info_schema_query_response_time"`
	CollectInfoSchemaProcesslist struct {
		Enabled         bool `toml:"enabled"`
		MinTime         int  `toml:"min_time"`
		ProcessesByUser bool `toml:"processes_by_user"`
		ProcessesByHost bool `toml:"processes_by_host"`
	} `toml:"collect_info_schema_processlist"`
	CollectInfoSchemaTables struct {
		Enabled   bool   `toml:"enabled"`
		Databases string `toml:"databases"`
	} `toml:"collect_info_schema_tables"`
	CollectInfoSchemaInnodbTablespaces struct {
		Enabled bool `toml:"enabled"`
	} `toml:"collect_info_schema_innodb_tablespaces"`
	CollectInfoSchemaInnodbMetrics struct {
		Enabled bool `toml:"enabled"`
	} `toml:"collect_info_schema_innodb_metrics"`
	CollectInfoSchemaUserstats struct {
		Enabled bool `toml:"enabled"`
	} `toml:"collect_info_schema_userstats"`
	CollectInfoSchemaClientstats struct {
		Enabled bool `toml:"enabled"`
	} `toml:"collect_info_schema_clientstats"`
	CollectInfoSchemaTablestats struct {
		Enabled bool `toml:"enabled"`
	} `toml:"collect_info_schema_tablestats"`
	CollectInfoSchemaSchemastats struct {
		Enabled bool `toml:"enabled"`
	} `toml:"collect_info_schema_schemastats"`
	CollectInfoSchemaReplicaHost struct {
		Enabled bool `toml:"enabled"`
	} `toml:"collect_info_schema_replica_host"`
	CollectMysqlUser struct {
		Enabled               bool `toml:"enabled"`
		CollectUserPrivileges bool `toml:"collect_user_privileges"`
	} `toml:"collect_mysql_user"`
	CollectAutoIncrementColumns struct {
		Enabled bool `toml:"enabled"`
	} `toml:"collect_auto_increment_columns"`
	CollectBinlogSize struct {
		Enabled bool `toml:"enabled"`
	} `toml:"collect_binlog_size"`
	CollectPerfSchemaTableiowaits struct {
		Enabled bool `toml:"enabled"`
	} `toml:"collect_perf_schema_tableiowaits"`
	CollectPerfSchemaIndexiowaits struct {
		Enabled bool `toml:"enabled"`
	} `toml:"collect_perf_schema_indexiowaits"`
	CollectPerfSchemaTablelocks struct {
		Enabled bool `toml:"enabled"`
	} `toml:"collect_perf_schema_tablelocks"`
	CollectPerfSchemaEventsstatements struct {
		Enabled         bool `toml:"enabled"`
		Limit           int  `toml:"limit"`
		Timelimit       int  `toml:"timelimit"`
		DigestTextLimit int  `toml:"digest_text_limit"`
	} `toml:"collect_perf_schema_eventsstatements"`
	CollectPerfSchemaEventsstatementssum struct {
		Enabled bool `toml:"enabled"`
	} `toml:"collect_perf_schema_eventsstatementssum"`
	CollectPerfSchemaEventswaits struct {
		Enabled bool `toml:"enabled"`
	} `toml:"collect_perf_schema_eventswaits"`
	CollectPerfSchemaFileEvents struct {
		Enabled bool `toml:"enabled"`
	} `toml:"collect_perf_schema_file_events"`
	CollectPerfSchemaFileInstances struct {
		Enabled      bool   `toml:"enabled"`
		Filter       string `toml:"filter"`
		RemovePrefix string `toml:"remove_prefix"`
	} `toml:"collect_perf_schema_file_instances"`
	CollectPerfSchemaMemoryEvents struct {
		Enabled      bool   `toml:"enabled"`
		RemovePrefix string `toml:"remove_prefix"`
	} `toml:"collect_perf_schema_memory_events"`
	CollectPerfSchemaReplicationGroupMembers struct {
		Enabled bool `toml:"enabled"`
	} `toml:"collect_perf_schema_replication_group_members"`
	CollectPerfSchemaReplicationGroupMemberStats struct {
		Enabled bool `toml:"enabled"`
	} `toml:"collect_perf_schema_replication_group_member_stats"`
	CollectPerfSchemaReplicationApplierStatusByWorker struct {
		Enabled bool `toml:"enabled"`
	} `toml:"collect_perf_schema_replication_applier_status_by_worker"`
	CollectSysUserSummary struct {
		Enabled bool `toml:"enabled"`
	} `toml:"collect_sys_user_summary"`
	CollectEngineTokudbStatus struct {
		Enabled bool `toml:"enabled"`
	} `toml:"collect_engine_tokudb_status"`
	CollectEngineInnodbStatus struct {
		Enabled bool `toml:"enabled"`
	} `toml:"collect_engine_innodb_status"`
	CollectHeartbeat struct {
		Enabled  bool   `toml:"enabled"`
		Database string `toml:"database"`
		Table    string `toml:"table"`
		UTC      bool   `toml:"utc"`
	} `toml:"collect_heartbeat"`
	CollectSlaveHosts struct {
		Enabled bool `toml:"enabled"`
	} `toml:"collect_slave_hosts"`
}

func (c *Config) EnabledScrapers() (ret []collector.Scraper) {
	if c.CollectGlobalStatus.Enabled {
		ret = append(ret, collector.ScrapeGlobalStatus{})
	}

	if c.CollectGlobalVariables.Enabled {
		ret = append(ret, collector.ScrapeGlobalVariables{})
	}

	if c.CollectSlaveStatus.Enabled {
		ret = append(ret, collector.ScrapeSlaveStatus{})
	}

	if c.CollectInfoSchemaInnodbCmp.Enabled {
		ret = append(ret, collector.ScrapeInnodbCmp{})
	}

	if c.CollectInfoSchemaInnodbCmpmem.Enabled {
		ret = append(ret, collector.ScrapeInnodbCmpMem{})
	}

	if c.CollectInfoSchemaQueryResponseTime.Enabled {
		ret = append(ret, collector.ScrapeQueryResponseTime{})
	}

	if c.CollectInfoSchemaProcesslist.Enabled {
		ret = append(ret, collector.ScrapeProcesslist{
			ProcesslistMinTime: c.CollectInfoSchemaProcesslist.MinTime,
			ProcessesByUser:    c.CollectInfoSchemaProcesslist.ProcessesByUser,
			ProcessesByHost:    c.CollectInfoSchemaProcesslist.ProcessesByHost,
		})
	}

	if c.CollectInfoSchemaTables.Enabled {
		ret = append(ret, collector.ScrapeTableSchema{
			TableSchemaDatabases: c.CollectInfoSchemaTables.Databases,
		})
	}

	if c.CollectInfoSchemaInnodbTablespaces.Enabled {
		ret = append(ret, collector.ScrapeInfoSchemaInnodbTablespaces{})
	}

	if c.CollectInfoSchemaInnodbMetrics.Enabled {
		ret = append(ret, collector.ScrapeInnodbMetrics{})
	}

	if c.CollectInfoSchemaUserstats.Enabled {
		ret = append(ret, collector.ScrapeUserStat{})
	}

	if c.CollectInfoSchemaClientstats.Enabled {
		ret = append(ret, collector.ScrapeClientStat{})
	}

	if c.CollectInfoSchemaTablestats.Enabled {
		ret = append(ret, collector.ScrapeTableStat{})
	}

	if c.CollectInfoSchemaSchemastats.Enabled {
		ret = append(ret, collector.ScrapeSchemaStat{})
	}

	if c.CollectInfoSchemaReplicaHost.Enabled {
		ret = append(ret, collector.ScrapeReplicaHost{})
	}

	if c.CollectMysqlUser.Enabled {
		ret = append(ret, collector.ScrapeUser{
			UserPrivilegesFlag: c.CollectMysqlUser.CollectUserPrivileges,
		})
	}

	if c.CollectAutoIncrementColumns.Enabled {
		ret = append(ret, collector.ScrapeAutoIncrementColumns{})
	}

	if c.CollectBinlogSize.Enabled {
		ret = append(ret, collector.ScrapeBinlogSize{})
	}

	if c.CollectPerfSchemaTableiowaits.Enabled {
		ret = append(ret, collector.ScrapePerfTableIOWaits{})
	}

	if c.CollectPerfSchemaIndexiowaits.Enabled {
		ret = append(ret, collector.ScrapePerfIndexIOWaits{})
	}

	if c.CollectPerfSchemaTablelocks.Enabled {
		ret = append(ret, collector.ScrapePerfTableLockWaits{})
	}

	if c.CollectPerfSchemaEventsstatements.Enabled {
		ret = append(ret, collector.ScrapePerfEventsStatements{
			PerfEventsStatementsLimit:           c.CollectPerfSchemaEventsstatements.Limit,
			PerfEventsStatementsTimeLimit:       c.CollectPerfSchemaEventsstatements.Timelimit,
			PerfEventsStatementsDigestTextLimit: c.CollectPerfSchemaEventsstatements.DigestTextLimit,
		})
	}

	if c.CollectPerfSchemaEventsstatementssum.Enabled {
		ret = append(ret, collector.ScrapePerfEventsStatementsSum{})
	}

	if c.CollectPerfSchemaEventswaits.Enabled {
		ret = append(ret, collector.ScrapePerfEventsWaits{})
	}

	if c.CollectPerfSchemaFileEvents.Enabled {
		ret = append(ret, collector.ScrapePerfFileEvents{})
	}

	if c.CollectPerfSchemaFileInstances.Enabled {
		ret = append(ret, collector.ScrapePerfFileInstances{
			PerformanceSchemaFileInstancesFilter:       c.CollectPerfSchemaFileInstances.Filter,
			PerformanceSchemaFileInstancesRemovePrefix: c.CollectPerfSchemaFileInstances.RemovePrefix,
		})
	}

	if c.CollectPerfSchemaMemoryEvents.Enabled {
		ret = append(ret, collector.ScrapePerfMemoryEvents{
			PerformanceSchemaMemoryEventsRemovePrefix: c.CollectPerfSchemaMemoryEvents.RemovePrefix,
		})
	}

	if c.CollectPerfSchemaReplicationGroupMembers.Enabled {
		ret = append(ret, collector.ScrapePerfReplicationGroupMembers{})
	}

	if c.CollectPerfSchemaReplicationGroupMemberStats.Enabled {
		ret = append(ret, collector.ScrapePerfReplicationGroupMemberStats{})
	}

	if c.CollectPerfSchemaReplicationApplierStatusByWorker.Enabled {
		ret = append(ret, collector.ScrapePerfReplicationApplierStatsByWorker{})
	}

	if c.CollectSysUserSummary.Enabled {
		ret = append(ret, collector.ScrapeSysUserSummary{})
	}

	if c.CollectEngineTokudbStatus.Enabled {
		ret = append(ret, collector.ScrapeEngineTokudbStatus{})
	}

	if c.CollectEngineInnodbStatus.Enabled {
		ret = append(ret, collector.ScrapeEngineInnodbStatus{})
	}

	if c.CollectHeartbeat.Enabled {
		ret = append(ret, collector.ScrapeHeartbeat{
			CollectHeartbeatDatabase: c.CollectHeartbeat.Database,
			CollectHeartbeatTable:    c.CollectHeartbeat.Table,
			CollectHeartbeatUtc:      c.CollectHeartbeat.UTC,
		})
	}

	if c.CollectSlaveHosts.Enabled {
		ret = append(ret, collector.ScrapeSlaveHosts{})
	}

	return
}

func ParseConfig(bs []byte) (*Config, error) {
	var c Config
	err := toml.Unmarshal(bs, &c)
	if err != nil {
		return nil, err
	}

	return &c, nil
}

// mysqld_exporter 原来的很多参数都是通过命令行传的，在 cprobe 的场景下，需要改造
// cprobe 是并发抓取很多个数据库实例的监控数据，不同的数据库实例其抓取参数可能不同
// 如果直接修改 collector pkg 下面的变量，就会有并发使用变量的问题
// 把这些自定义参数封装到一个一个的 collector.Scraper 对象中，每个 target 抓取时实例化这些 collector.Scraper 对象
func Scrape(ctx context.Context, address string, cfg *Config, ss *types.Samples) error {
	dsn, err := cfg.Global.FormDSN(address)
	if err != nil {
		return fmt.Errorf("failed to form dsn for %s: %s", address, err)
	}

	scrapers := cfg.EnabledScrapers()
	exporter := collector.New(ctx, dsn, scrapers, ss, cfg.Queries)

	ch := make(chan prometheus.Metric)
	go func() {
		exporter.Collect(ch)
		close(ch)
	}()

	for m := range ch {
		if err := ss.AddPromMetric(m); err != nil {
			logger.Warnf("failed to tranform prometheus metric: %s", err)
		}
	}

	return nil
}