package query_interface

type OdsDictQuerier interface {
	// select dict_type_code, count(dict_id) count from @@table
	// where is_valid = 1
	// group by dict_type_code
	CountGroupByType() ([]map[string]any, error)
}
