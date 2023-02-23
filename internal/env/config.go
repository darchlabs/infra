package env

type Env struct {
	// onboarding service
	Port string `envconfig:"onboarding_port" required:"true"`

	// synchronizers service
	SynchronizersNodeURL          string `envconfig:"synchronizers_node_url" required:"true"`
	SynchronizersIntervalSeconds  string `envconfig:"synchronizers_interval_seconds" required:"true"`
	SynchronizersDatabaseFilepath string `envconfig:"synchronizers_database_filepath" required:"true"`
	SynchronizersPort             string `envconfig:"synchronizers_port" required:"true"`

	// jobs service
	JobsDatabaseFilepath string `envconfig:"jobs_database_filepath" required:"true"`
	JobsPort             string `envconfig:"jobs_port" required:"true"`
	JobsNodeURL          string `envconfig:"jobs_node_url" required:"true"`
	JobsPrivateKey       string `envconfig:"jobs_private_key" required:"true"`

	// nodes service
	NodesChain             string `envconfig:"nodes_chain" required:"true"`
	NodesApiServerPort     string `envconfig:"nodes_api_server_port" required:"true"`
	NodesRpcPort           string `envconfig:"nodes_rpc_port" required:"true"`
	NodesBaseChainDataPath string `envconfig:"nodes_base_chain_data_path" required:"true"`
	NodesNodeURL           string `envconfig:"nodes_node_url" required:"true"`
	NodesBlockNumber       string `envconfig:"nodes_block_number" required:"true"`

	// reporter service
	ReportSynchronizersDatabaseURL string `envconfig:"report_synchronizers_database_url" required:"true"`
	ReportSynchronizersServiceURL  string `envconfig:"report_synchronizers_service_url" required:"true"`
	ReportSynchronizersServiceType string `envconfig:"report_synchronizers_service_type" required:"true"`

	// webapp
	WebappPort             string `envconfig:"webapp_port" required:"true"`
	WebappRedisURL         string `envconfig:"webapp_redis_url" required:"true"`
	WebappSynchronizersURL string `envconfig:"webapp_synchronizers_url" required:"true"`
	WebappNodesURL         string `envconfig:"webapp_nodes_url" required:"true"`
	WebappJobsURL          string `envconfig:"webapp_jobs_url" required:"true"`
}
