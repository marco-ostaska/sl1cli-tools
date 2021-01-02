package sl1

import (
	"fmt"
	"path"
)

// LookupDescByURI looks up a description content for a provided URI
func (bInfo *BasicInfo) LookupDescByURI(uri string) (string, error) {

	for _, u := range *bInfo {
		if uri == u.URI {
			return u.Description, nil
		}
	}
	return uri, fmt.Errorf("%s: no such uri", uri)

}

// LookupIdxByDesc lookups up index position for provided BasicInfo.Description
func (bInfo *BasicInfo) LookupIdxByDesc(d string) (int, error) {

	for i, u := range *bInfo {
		if d == u.Description {
			return i, nil
		}
	}
	return 0, fmt.Errorf("%s: no such %v", d, path.Base(path.Dir((*bInfo)[0].URI)))

}

// LookupIDbyDesc looks up sl1id for a provided Description
func (bInfo *BasicInfo) LookupIDbyDesc(s string) (string, error) {
	i, err := bInfo.LookupIdxByDesc(s)
	if err != nil {
		return "", err
	}

	return path.Base((*bInfo)[i].URI), nil
}
