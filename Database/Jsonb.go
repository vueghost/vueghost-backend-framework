package Database

import (
	"fmt"
	"sort"
	"strings"
)

//Jsonb
type Jsonb struct {
}

// NewJsonb.
func NewJsonb() *Jsonb {
	return &Jsonb{}
}

//Delete delete jsonb syntax generator.
func (d *Jsonb) Delete(from string, where map[string]interface{}) (PSQLString string) {
	var conditionArray []string
	for key, value := range where {
		conditionArray = append(conditionArray, fmt.Sprintf(`(%[1]s->i->'%[2]s' = '%#[3]v')`, from, key, value))
	}
	condition := strings.Join(conditionArray[:], "AND")
	return fmt.Sprintf(`%[1]s= %[1]s #-coalesce(('{' || ( SELECT i FROM generate_series(0, jsonb_array_length(%[1]s) - 1) AS i WHERE %[2]s) || '}')::text[], '{}')`, from, condition)
}

//Set set jsonb object syntax generator.
func (d *Jsonb) Set(field string, attributes map[string]interface{}, path interface{}) (PSQLString string) {
	attributesQuery := []string{}
	for attrKey, attrValue := range attributes {
		attributesQuery = append(attributesQuery, fmt.Sprintf(`'{"%s":%#v}'`, attrKey, attrValue))
	}
	if path == nil {
		return fmt.Sprintf(`%[1]s || %[3]s`, field, path, strings.Join(attributesQuery[:], "||"))
	}

	return fmt.Sprintf(`jsonb_set(%[1]s,'{%[2]s}', %[3]s)`, field, path, strings.Join(attributesQuery[:], "||"))
}

//Increase Increase jsonb field count.
func (d *Jsonb) Increase(field string, attributes map[string]int) (PSQLString string) {
	var attributesQuery []string
	i := 0

	for attrKey, attrValue := range attributes {
		query := ""
		if i <= 0 {
			query = fmt.Sprintf(`jsonb_set(%[1]s,'{%[2]s}', (COALESCE(%[1]s ->> '%[2]s', '0')::int + %[3]v)::text::jsonb)`, field, attrKey, attrValue)
		} else {
			query = fmt.Sprintf(`('{"%[2]s":' || (COALESCE(%[1]s->> '%[2]s', '0')::int + %[3]v)::text || '}')::text::jsonb`, field, attrKey, attrValue)

		}
		attributesQuery = append(attributesQuery, query)
		i++
	}
	q := fmt.Sprintf(`%[1]s=%[2]s`, field, strings.Join(attributesQuery[:], "||"))
	return q
}

//Decrease Decrease jsonb field count.
func (d *Jsonb) Decrease(field string, attributes map[string]int) (PSQLString string) {
	var attributesQuery []string
	i := 0
	for attrKey, attrValue := range attributes {
		query := ""
		if i <= 0 {
			query = fmt.Sprintf(`jsonb_set(%[1]s,'{%[2]s}', (COALESCE(%[1]s ->> '%[2]s', '0')::int - %[3]v)::text::jsonb)`, field, attrKey, attrValue)
		} else {
			query = fmt.Sprintf(`('{"%[2]s":' || (COALESCE(%[1]s->> '%[2]s', '0')::int - %[3]v)::text || '}')::text::jsonb`, field, attrKey, attrValue)

		}
		attributesQuery = append(attributesQuery, query)
		i++
	}
	sort.Strings(attributesQuery)
	q := fmt.Sprintf(`%[1]s=%[2]s`, field, strings.Join(attributesQuery[:], "||"))
	return q
}

//Get get jsonb syntax generator.
func (d *Jsonb) Get() {

}

//Find find jsonb syntax generator.
func (d *Jsonb) Find() {

}
