//ff:func feature=index type=rule control=iteration dimension=1
//ff:what checkSecs — every segment of some sections
package index

func checkSecs(ss []*Sec) error {
	for _, s := range ss {
		if err := checkSegs(s.Files); err != nil {
			return err
		}
	}
	return nil
}
