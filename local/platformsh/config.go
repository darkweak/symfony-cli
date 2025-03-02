// Code generated by platformsh/generator/main.go
// DO NOT EDIT

/*
 * Copyright (c) 2021-present Fabien Potencier <fabien@symfony.com>
 *
 * This file is part of Symfony CLI project
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as
 * published by the Free Software Foundation, either version 3 of the
 * License, or (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program. If not, see <http://www.gnu.org/licenses/>.
 */

package platformsh

var availablePHPExts = map[string][]string{
	"amqp":            {"7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1", "8.2"},
	"apc":             {"5.4", "5.5"},
	"apcu":            {"5.4", "5.6", "7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1", "8.2"},
	"apcu_bc":         {"7.0", "7.1", "7.2", "7.3", "7.4"},
	"applepay":        {"7.0", "7.1", "7.4"},
	"bcmath":          {"7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1", "8.2"},
	"blackfire":       {"5.4", "5.5", "5.6", "7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1", "8.2"},
	"bz2":             {"7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1", "8.2"},
	"calendar":        {"7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1", "8.2"},
	"ctype":           {"7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1", "8.2"},
	"curl":            {"5.4", "5.5", "5.6", "7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1", "8.2"},
	"dba":             {"7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1", "8.2"},
	"dom":             {"7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1", "8.2"},
	"enchant":         {"5.4", "5.5", "5.6", "7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1", "8.2"},
	"event":           {"7.1", "7.2", "7.3", "7.4", "8.0", "8.1"},
	"exif":            {"7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1", "8.2"},
	"ffi":             {"7.4", "8.0", "8.1", "8.2"},
	"fileinfo":        {"7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1", "8.2"},
	"ftp":             {"7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1", "8.2"},
	"gd":              {"5.4", "5.5", "5.6", "7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1", "8.2"},
	"gearman":         {"5.4", "5.5", "5.6"},
	"geoip":           {"5.4", "5.5", "5.6", "7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1", "8.2"},
	"gettext":         {"7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1", "8.2"},
	"gmp":             {"5.4", "5.5", "5.6", "7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1", "8.2"},
	"http":            {"5.4", "5.5", "7.3", "7.4", "8.0", "8.1", "8.2"},
	"iconv":           {"7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1", "8.2"},
	"igbinary":        {"7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1", "8.2"},
	"imagick":         {"5.4", "5.5", "5.6", "7.0", "7.1", "7.2", "7.3", "7.4"},
	"imap":            {"5.4", "5.5", "5.6", "7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1", "8.2"},
	"interbase":       {"5.4", "5.5", "5.6", "7.0", "7.1", "7.2", "7.3", "7.4", "8.0"},
	"intl":            {"5.4", "5.5", "5.6", "7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1", "8.2"},
	"ioncube":         {"7.0", "7.1", "7.2", "7.3", "7.4", "8.0"},
	"json":            {"5.6", "7.0", "7.1", "7.2", "7.3", "7.4"},
	"ldap":            {"5.4", "5.5", "5.6", "7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1", "8.2"},
	"mailparse":       {"7.0", "7.1", "7.2", "7.4", "8.0", "8.1", "8.2"},
	"mbstring":        {"7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1", "8.2"},
	"mcrypt":          {"5.4", "5.5", "5.6", "7.0", "7.1"},
	"memcache":        {"5.4", "5.5", "5.6"},
	"memcached":       {"5.4", "5.5", "5.6", "7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1", "8.2"},
	"mongo":           {"5.4", "5.5", "5.6"},
	"mongodb":         {"7.1", "7.2", "7.3", "7.4", "8.0", "8.1"},
	"msgpack":         {"5.6", "7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1", "8.2"},
	"mssql":           {"5.4", "5.5", "5.6"},
	"mysql":           {"5.4", "5.5", "5.6"},
	"mysqli":          {"5.4", "5.5", "5.6", "7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1", "8.2"},
	"mysqlnd":         {"5.4", "5.5", "5.6", "7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1", "8.2"},
	"newrelic":        {"5.6", "7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1"},
	"oauth":           {"7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1"},
	"odbc":            {"5.4", "5.5", "5.6", "7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1", "8.2"},
	"opcache":         {"5.5", "5.6", "7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1", "8.2"},
	"pdo":             {"5.4", "5.5", "5.6", "7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1", "8.2"},
	"pdo_dblib":       {"5.4", "5.5", "5.6", "7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1", "8.2"},
	"pdo_firebird":    {"5.4", "5.5", "5.6", "7.0", "7.1", "7.2", "7.3", "7.4"},
	"pdo_mysql":       {"5.4", "5.5", "5.6", "7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1", "8.2"},
	"pdo_odbc":        {"5.4", "5.5", "5.6", "7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1", "8.2"},
	"pdo_pgsql":       {"5.4", "5.5", "5.6", "7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1", "8.2"},
	"pdo_sqlite":      {"5.4", "5.5", "5.6", "7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1", "8.2"},
	"pdo_sqlsrv":      {"7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1"},
	"pgsql":           {"5.4", "5.5", "5.6", "7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1", "8.2"},
	"phar":            {"7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1", "8.2"},
	"pinba":           {"5.4", "5.5", "5.6"},
	"posix":           {"7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1", "8.2"},
	"propro":          {"5.6"},
	"pspell":          {"5.4", "5.5", "5.6", "7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1", "8.2"},
	"pthreads":        {"7.1"},
	"raphf":           {"5.6", "7.4", "8.0", "8.1", "8.2"},
	"rdkafka":         {"8.1", "8.2"},
	"readline":        {"5.4", "5.5", "5.6", "7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1", "8.2"},
	"recode":          {"5.4", "5.5", "5.6", "7.0", "7.1", "7.2", "7.3"},
	"redis":           {"5.4", "5.5", "5.6", "7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1", "8.2"},
	"shmop":           {"7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1", "8.2"},
	"simplexml":       {"7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1", "8.2"},
	"snmp":            {"5.4", "5.5", "5.6", "7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1", "8.2"},
	"soap":            {"7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1", "8.2"},
	"sockets":         {"7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1", "8.2"},
	"sodium":          {"7.2", "7.3", "7.4", "8.0", "8.1", "8.2"},
	"sourceguardian":  {"7.0", "7.1"},
	"spplus":          {"5.4", "5.5"},
	"sqlite3":         {"5.4", "5.5", "5.6", "7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1", "8.2"},
	"sqlsrv":          {"7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1"},
	"ssh2":            {"5.4", "5.5", "5.6", "7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1", "8.2"},
	"sybase":          {"7.1", "7.2", "7.3", "7.4", "8.0", "8.1", "8.2"},
	"sysvmsg":         {"7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1", "8.2"},
	"sysvsem":         {"7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1", "8.2"},
	"sysvshm":         {"7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1", "8.2"},
	"tideways":        {"7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1", "8.2"},
	"tideways_xhprof": {"7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1"},
	"tidy":            {"5.4", "5.5", "5.6", "7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1", "8.2"},
	"tokenizer":       {"7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1", "8.2"},
	"uuid":            {"7.1", "7.2", "7.3", "7.4", "8.0", "8.1", "8.2"},
	"wddx":            {"7.0", "7.1", "7.2", "7.3", "7.4"},
	"xcache":          {"5.4", "5.5"},
	"xdebug":          {"7.1", "7.2", "7.3", "7.4", "8.0", "8.1", "8.2"},
	"xhprof":          {"5.4", "5.5", "5.6"},
	"xml":             {"7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1", "8.2"},
	"xmlreader":       {"7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1", "8.2"},
	"xmlrpc":          {"5.4", "5.5", "5.6", "7.0", "7.1", "7.2", "7.3", "7.4"},
	"xmlwriter":       {"7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1", "8.2"},
	"xsl":             {"5.4", "5.5", "5.6", "7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1", "8.2"},
	"yaml":            {"7.1", "7.2", "7.3", "7.4", "8.0", "8.1", "8.2"},
	"zbarcode":        {"7.0", "7.1", "7.2", "7.3"},
	"zendopcache":     {"5.4"},
	"zip":             {"7.0", "7.1", "7.2", "7.3", "7.4", "8.0", "8.1", "8.2"},
}

var availableServices = []*service{
	{
		Type: "chrome-headless",
		Versions: serviceVersions{
			Deprecated: []string{},
			Supported:  []string{"73", "80", "81", "83", "84", "86", "91", "95"},
		},
	},
	{
		Type: "elasticsearch",
		Versions: serviceVersions{
			Deprecated: []string{"0.9", "1.4", "1.7", "2.4", "5.2", "5.4", "6.5", "6.8", "7.2", "7.5", "7.7", "7.9", "7.10"},
			Supported: []string{},
		},
	},
	{
		Type: "influxdb",
		Versions: serviceVersions{
			Deprecated: []string{"1.2", "1.3", "1.7", "1.8", "2.2"},
			Supported:  []string{"2.3"},
		},
	},
	{
		Type: "kafka",
		Versions: serviceVersions{
			Deprecated: []string{"2.1", "2.2", "2.3", "2.4", "2.5", "2.6", "2.7"},
			Supported:  []string{"3.2"},
		},
	},
	{
		Type: "mariadb",
		Versions: serviceVersions{
			Deprecated: []string{"5.5", "10.0", "10.1", "10.2"},
			Supported:  []string{"10.3", "10.4", "10.5", "10.6", "10.11"},
		},
	},
	{
		Type: "memcached",
		Versions: serviceVersions{
			Deprecated: []string{},
			Supported:  []string{"1.4", "1.5", "1.6"},
		},
	},
	{
		Type: "mongodb",
		Versions: serviceVersions{
			Deprecated: []string{"3.0", "3.2", "3.4", "3.6"},
			Supported: []string{},
		},
	},
	{
		Type: "mongodb-enterprise",
		Versions: serviceVersions{
			Deprecated: []string{"4.0"},
			Supported:  []string{"4.2", "4.4", "5.0"},
		},
	},
	{
		Type: "mysql",
		Versions: serviceVersions{
			Deprecated: []string{"5.5", "10.0", "10.1", "10.2"},
			Supported:  []string{"10.3", "10.4", "10.5", "10.6", "10.11"},
		},
	},
	{
		Type: "network-storage",
		Versions: serviceVersions{
			Deprecated: []string{"1.0"},
			Supported:  []string{"2.0"},
		},
	},
	{
		Type: "opensearch",
		Versions: serviceVersions{
			Deprecated: []string{},
			Supported:  []string{"1.1", "1.2", "2"},
		},
	},
	{
		Type: "oracle-mysql",
		Versions: serviceVersions{
			Deprecated: []string{},
			Supported:  []string{"5.7", "8.0"},
		},
	},
	{
		Type: "postgresql",
		Versions: serviceVersions{
			Deprecated: []string{"9.3", "9.4", "9.5", "9.6", "10"},
			Supported:  []string{"11", "12", "13", "14", "15"},
		},
	},
	{
		Type: "rabbitmq",
		Versions: serviceVersions{
			Deprecated: []string{"3.5", "3.6", "3.7", "3.8"},
			Supported:  []string{"3.9", "3.10", "3.11"},
		},
	},
	{
		Type: "redis",
		Versions: serviceVersions{
			Deprecated: []string{"2.8", "3.0", "3.2", "4.0", "5.0", "6.0"},
			Supported:  []string{"6.2", "7.0"},
		},
	},
	{
		Type: "solr",
		Versions: serviceVersions{
			Deprecated: []string{"3.6", "4.10", "6.3", "6.6", "7.6", "7.7", "8.0", "8.4", "8.6"},
			Supported:  []string{"8.11", "9.1"},
		},
	},
	{
		Type: "varnish",
		Versions: serviceVersions{
			Deprecated: []string{"5.1", "5.2"},
			Supported:  []string{"6.0", "6.3", "7.1", "7.2"},
		},
	},
	{
		Type: "vault-kms",
		Versions: serviceVersions{
			Deprecated: []string{"1.6", "1.8"},
			Supported:  []string{"1.12"},
		},
	},
}
