package treemodel

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

func GetTreePathByValueStringPath(
	path []string,
	m *gtk.TreeStore,
	str_col int,
) (*gtk.TreePath, error) {

	var ret *gtk.TreePath

	for _, j := range path {

		it, ok, err := FindTreePathByStringAndColOnSameLevel(
			m,
			ret,
			j,
			str_col,
		)
		if err != nil {
			return nil, err
		}

		if !ok {
			return nil, errors.New("not found")
		}

		ret = it
	}

	return ret, nil
}

func FindTreePathByStringAndColOnSameLevel(
	m *gtk.TreeStore,
	path *gtk.TreePath,
	str_val string,
	str_col int,
) (*gtk.TreePath, bool, error) {

	var ret *gtk.TreePath

	found := false

	err := TreeStoreForEach(
		m, path,
		func(
			m *gtk.TreeStore,
			path *gtk.TreePath,
			iter *gtk.TreeIter,
		) error {

			v, err := m.GetValue(iter, str_col)
			if err != nil {
				return nil
			}

			vv, err := v.GetString()
			if err != nil {
				return nil
			}

			if vv == str_val {
				ret, err = m.GetPath(iter)
				if err != nil {
					return err
				}

				found = true
				return errors.New("not error")
			}

			return nil
		},
	)

	if err != nil {
		if err.Error() == "not error" {
			err = nil
		} else {
			return nil, false, err
		}
	}

	if !found {
		return nil, false, nil
	}

	return ret, true, nil
}

func RenderTreePathString(
	m *gtk.TreeStore,
	path *gtk.TreePath,
	str_col int,
) ([]string, error) {
	ret := make([]string, 0)

	ind := path.GetIndices()
	for i := 0; i != len(ind); i++ {
		strs := make([]string, 0)
		for _, j := range ind[0 : i+1] {
			strs = append(strs, strconv.Itoa(j))
		}

		it, err := m.GetIterFromString(strings.Join(strs, ":"))
		if err != nil {
			return []string{}, err
		}

		val, err := m.GetValue(it, str_col)
		if err != nil {
			return []string{}, err
		}

		val_s, err := val.GetString()
		if err != nil {
			return []string{}, err
		}
		ret = append(ret, val_s)
	}

	return ret, nil
}

// TODO: NOTE: FIXME: at time of this writting (Mon Jan 15 03:52:17 MSK 2018)
//                    gotk3 doesn't support GtkTreeRowReference
//                    so I'm trying to do this with gtk.TreePath..
//                    but, probably, this will lead to errors..
// func CopyTreeItem(
// 	m *gtk.TreeStore,
// 	src, dst *gtk.TreePath,
// ) error {
// 	panic("this function isn't safe, do not use")
//
// 	{
// 		s := src.String()
// 		d := dst.String()
// 		if strings.HasPrefix(d, s+":") {
// 			// TODO: probably eventually this should be done possible, but not now
// 			return errors.New("dst can't be inside src")
// 		}
// 	}
//
// 	err := WalkTreeStore(
// 		m,
// 		src,
// 		func(
// 			m *gtk.TreeStore,
// 			path *gtk.TreePath,
// 			paths []*gtk.TreePath,
// 		) error {
// 			for _, i := range paths {
//
// 			}
// 			return nil
// 		},
// 	)
// 	if err != nil {
// 		return err
// 	}
//
// 	return nil
// }

func CopyTreeStoreType(src *gtk.TreeStore) (*gtk.TreeStore, error) {
	nmt := make([]glib.Type, 0)

	for i := 0; i != src.GetNColumns(); i++ {
		nmt = append(nmt, src.GetColumnType(i))
	}

	ret, err := gtk.TreeStoreNew(nmt...)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

func TreeStorePopulatePath(
	ts *gtk.TreeStore,
	ts_path *gtk.TreePath,
	values [][]*glib.Value,
) error {

	var it *gtk.TreeIter
	var err error

	if ts_path == nil {
		it = nil
	} else {
		it, err = ts.GetIter(ts_path)
		if err != nil {
			return err
		}
	}

	for _, i := range values {
		wit := ts.Append(it)

		for jj, j := range i {
			err = ts.SetValue(wit, jj, j)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func GetTreeStoreRowValues(src *gtk.TreeStore, path *gtk.TreePath) (
	[]*Value,
	error,
) {

	src_col_n := src.GetNColumns()

	ret := make([]*Value, 0)

	for i := 0; i != src_col_n; i++ {

		val, err := TreeStoreGetValue(src, path, i)
		if err != nil {
			return []*Value{}, err
		}

		ret = append(ret, val)
		if err != nil {
			return []*Value{}, err
		}
	}

	return ret, nil
}

func SetTreeStoreRowValues(
	dst *gtk.TreeStore,
	path *gtk.TreePath,
	values []*Value,
) error {

	var err error

	for ii, i := range values {
		err = i.TreeStoreSetValue(dst, path, ii)
		if err != nil {
			return err
		}
	}

	return nil
}

func PasteTreeStore(
	src *gtk.TreeStore,
	dst *gtk.TreeStore,
	path_to_use_as_parent_in_dst *gtk.TreePath,
) error {

	var err error
	var dst_lvl_track *gtk.TreeIter
	{

		p, err := gtk.TreePathNewFromString(
			path_to_use_as_parent_in_dst.String(),
		)
		if err != nil {
			return err
		}

		dst_lvl_track, err = dst.GetIter(p)
		if err != nil {
			return err
		}
	}

	var tfunc func(
		m *gtk.TreeStore,
		path *gtk.TreePath,
		iter *gtk.TreeIter,
	) error

	tfunc = func(
		m *gtk.TreeStore,
		path *gtk.TreePath,
		iter *gtk.TreeIter,
	) error {

		iter_path, err := m.GetPath(iter)
		if err != nil {
			return err
		}

		i := dst.Append(dst_lvl_track)
		i_path, err := dst.GetPath(i)
		if err != nil {
			return err
		}

		values, err := GetTreeStoreRowValues(m, iter_path)
		if err != nil {
			return err
		}

		err = SetTreeStoreRowValues(dst, i_path, values)
		if err != nil {
			return err
		}

		if m.IterHasChild(iter) {
			dst_lvl_track_copy, err := dst_lvl_track.Copy()
			if err != nil {
				return err
			}

			dst_lvl_track = i
			err = TreeStoreForEach(
				src,
				iter_path,
				tfunc,
			)
			if err != nil {
				return err
			}
			dst_lvl_track = dst_lvl_track_copy
		}

		return nil
	}

	err = TreeStoreForEach(
		src,
		nil,
		tfunc,
	)
	if err != nil {
		return err
	}

	return nil

}

func CopyTreeStore(
	src *gtk.TreeStore,
	start_path *gtk.TreePath,
) (
	*gtk.TreeStore,
	error,
) {

	ret, err := CopyTreeStoreType(src)
	if err != nil {
		return nil, err
	}

	var ret_lvl *gtk.TreePath

	{
		vals, err := GetTreeStoreRowValues(src, start_path)
		if err != nil {
			return nil, err
		}

		new_it := ret.Append(nil)

		new_path, err := ret.GetPath(new_it)
		if err != nil {
			return nil, err
		}

		err = SetTreeStoreRowValues(ret, new_path, vals)
		if err != nil {
			return nil, err
		}
	}

	var have_child bool
	{
		i, err := src.GetIter(start_path)
		if err != nil {
			return nil, err
		}

		have_child = src.IterHasChild(i)
	}

	if have_child {
		err = WalkTreeStore(
			src,
			start_path,
			func(
				m *gtk.TreeStore,
				path *gtk.TreePath,
				values [][]*Value,
			) error {

				var err error

				{
					start_path_str_list := strings.Split(start_path.String(), ":")
					path_str_list := strings.Split(path.String(), ":")

					path_str_list = path_str_list[len(start_path_str_list):]
					path_str_list = append(
						path_str_list[:0],
						append(
							[]string{"0"},
							path_str_list[0:]...,
						)...,
					)
					if len(path_str_list) == 0 {
						ret_lvl = nil
					} else {
						ret_lvl, err = gtk.TreePathNewFromString(
							strings.Join(path_str_list, ":"),
						)
						if err != nil {
							return err
						}
					}

				}

				var ret_lvl_itr *gtk.TreeIter
				if ret_lvl != nil {
					ret_lvl_itr, err = ret.GetIter(ret_lvl)
					if err != nil {
						return err
					}
				}

				for _, i := range values {
					it_new := ret.Append(ret_lvl_itr)

					it_new_path, err := ret.GetPath(it_new)
					if err != nil {
						return err
					}

					err = SetTreeStoreRowValues(ret, it_new_path, i)

				}

				return nil
			},
		)
		if err != nil {
			return nil, err
		}
	}

	return ret, nil
}

func ListTreeStore(
	m *gtk.TreeStore,
	path *gtk.TreePath,
) ([][]*Value, error) {

	ret := make([][]*Value, 0)

	err := TreeStoreForEach(
		m, path,
		func(
			m *gtk.TreeStore,
			path *gtk.TreePath,
			iter *gtk.TreeIter,
		) error {

			iter_path, err := m.GetPath(iter)
			if err != nil {
				return err
			}

			ret_line, err := GetTreeStoreRowValues(m, iter_path)
			if err != nil {
				return err
			}

			ret = append(ret, ret_line)
			return nil
		},
	)
	if err != nil {
		return [][]*Value{}, err
	}

	return ret, nil
}

func WalkTreeStore(
	m *gtk.TreeStore,
	path *gtk.TreePath,
	target func(
		m *gtk.TreeStore,
		path *gtk.TreePath,
		values [][]*Value,
	) error,
) error {

	values, err := ListTreeStore(m, path)
	if err != nil {
		return err
	}

	err = target(m, path, values)
	if err != nil {
		return err
	}

	err = TreeStoreForEach(
		m, path,
		func(
			m *gtk.TreeStore,
			path *gtk.TreePath,
			iter *gtk.TreeIter,
		) error {

			if m.IterHasChild(iter) {
				p, err := m.GetPath(iter)
				if err != nil {
					return err
				}
				err = WalkTreeStore(m, p, target)
				if err != nil {
					return err
				}
			}

			return nil
		},
	)
	if err != nil {
		return err
	}

	return nil
}

func TreeStoreForEach(
	m *gtk.TreeStore,
	path *gtk.TreePath,
	target func(
		m *gtk.TreeStore,
		path *gtk.TreePath,
		iter *gtk.TreeIter,
	) error,
) error {
	var err error
	var path_iter *gtk.TreeIter

	if path != nil {
		path_iter, err = m.GetIter(path)
		if err != nil {
			return err
		}
	}

	ok := true
	var it *gtk.TreeIter
	for {
		if it == nil {
			it = &gtk.TreeIter{}
			ok = m.IterChildren(path_iter, it)
		} else {
			ok = m.IterNext(it)
		}
		if !ok {
			break
		}

		err = target(
			m,
			path,
			it,
		)
		if err != nil {
			return err
		}

	}
	return nil
}

type Value struct {
	Type   glib.Type
	GValue *glib.Value
	//  Value  interface{}
}

func (self *Value) String() string {
	return fmt.Sprint(self.Interface())
}

func (self *Value) Interface() (interface{}, error) {
	return self.GValue.GoValue()
}

func TreeStoreGetValue(
	src *gtk.TreeStore,
	path *gtk.TreePath,
	colind int,
) (
	*Value,
	error,
) {

	iter, err := src.GetIter(path)
	if err != nil {
		return nil, err
	}

	val, err := src.GetValue(iter, colind)
	if err != nil {
		return nil, err
	}

	// goval, err := val.GoValue()
	// if err != nil {
	// 	return nil, err
	// }

	ret := new(Value)

	ret.GValue = val
	// ret.Value = goval
	ret.Type = src.GetColumnType(colind)

	return ret, nil
}

func (self *Value) TreeStoreSetValue(
	dst *gtk.TreeStore,
	path *gtk.TreePath,
	colind int,
) error {

	iter, err := dst.GetIter(path)
	if err != nil {
		return err
	}

	switch self.Type {
	default:
		panic("not implimented. Glib.Type == " + fmt.Sprint(self.Type))
	case glib.TYPE_STRING:
		v, err := self.GValue.GetString()
		if err != nil {
			return err
		}
		err = dst.SetValue(iter, colind, v)
		if err != nil {
			return err
		}
	case glib.TYPE_INT:
		v, err := self.GValue.GoValue()
		if err != nil {
			return err
		}
		vv, ok := v.(int)
		if !ok {
			return errors.New("can't treat GValue as int")
		}

		err = dst.SetValue(iter, colind, vv)
		if err != nil {
			return err
		}
	case glib.TYPE_BOOLEAN:
		v, err := self.GValue.GoValue()
		if err != nil {
			return err
		}
		vv, ok := v.(bool)
		if !ok {
			return errors.New("can't treat GValue as bool")
		}

		err = dst.SetValue(iter, colind, vv)
		if err != nil {
			return err
		}
	case gdk.PixbufGetType():
		v, err := self.GValue.GoValue()
		if err != nil {
			return err
		}
		vv, ok := v.(*gdk.Pixbuf)
		if !ok {
			return errors.New("can't treat GValue as *gdk.Pixbuf")
		}

		err = dst.SetValue(iter, colind, vv)
		if err != nil {
			return err
		}
	}

	return nil
}

func TreeStoreString(str *gtk.TreeStore, path *gtk.TreePath) (string, error) {
	ret := ""
	err := WalkTreeStore(
		str, path,
		func(
			m *gtk.TreeStore,
			path *gtk.TreePath,
			values [][]*Value,
		) error {
			if path != nil {
				ret += fmt.Sprintf("path: %s\n", path.String())
			} else {
				ret += fmt.Sprintf("path: nil\n")
			}
			for ii, i := range values {
				ret += fmt.Sprintf("  %d: %v\n", ii, i)
			}
			return nil
		},
	)
	if err != nil {
		return "", err
	}
	return ret, nil
}
