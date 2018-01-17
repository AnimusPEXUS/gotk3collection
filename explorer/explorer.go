package explorer

import (
	"errors"
	"strings"

	"github.com/AnimusPEXUS/gotk3collection/explorer/ui"
	"github.com/AnimusPEXUS/gotk3collection/treemodel"
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
	"github.com/gotk3/gotk3/pango"
)

//go:generate go-bindata -pkg ui -o ./ui/explorer.go ./ui/explorer.glade

type Explorer struct {
	root    *gtk.Box
	tw, tw2 *gtk.TreeView
	// tw_column_icon *gtk.TreeViewColumn
	// tw_column_name *gtk.TreeViewColumn
	iw *gtk.IconView

	tb_cut           *gtk.ToolButton
	tb_copy          *gtk.ToolButton
	tb_paste         *gtk.ToolButton
	tb_delete        *gtk.ToolButton
	tb_s1            *gtk.SeparatorToolItem
	tb_new_directory *gtk.ToolButton
	tb_s2            *gtk.SeparatorToolItem
	tb_up            *gtk.ToolButton
	tb_down          *gtk.ToolButton

	entry_category *gtk.Entry

	lbl_active_view      *gtk.Label
	lbl_current_category *gtk.Label

	model        *gtk.TreeStore
	iconic_model *gtk.ListStore

	icon_column  int
	name_column  int
	isdir_column int

	active_view int

	cut_function      func(selected_items []*SelectedForOperationItem) error
	copy_function     func(selected_items []*SelectedForOperationItem) error
	paste_function    func(target *SelectedForOperationItem) error
	delete_function   func(selected_items []*SelectedForOperationItem) error
	mkdirall_function func(path string) error
}

func ExplorerNew() (*Explorer, error) {
	self := new(Explorer)

	builder, err := gtk.BuilderNew()
	if err != nil {
		panic(err.Error())
	}

	data := ui.MustAsset("ui/explorer.glade")

	err = builder.AddFromString(string(data))
	if err != nil {
		panic(err.Error())
	}

	if t, err := builder.GetObject("root"); err != nil {
		return nil, err
	} else {
		self.root = t.(*gtk.Box)
	}

	if t, err := builder.GetObject("entry_category"); err != nil {
		return nil, err
	} else {
		self.entry_category = t.(*gtk.Entry)
	}

	if t, err := builder.GetObject("tw"); err != nil {
		return nil, err
	} else {
		self.tw = t.(*gtk.TreeView)

		self.tw.Connect(
			"key-press-event",
			func() {
				self._SetActiveView(0)
			},
		)

		self.tw.Connect(
			"button-press-event",
			func() {
				self._SetActiveView(0)
			},
		)

	}

	if t, err := builder.GetObject("tw2"); err != nil {
		return nil, err
	} else {
		self.tw2 = t.(*gtk.TreeView)

		self.tw2.Connect(
			"key-press-event",
			func() {
				self._SetActiveView(1)
			},
		)

		self.tw2.Connect(
			"button-press-event",
			func() {
				self._SetActiveView(1)
			},
		)
	}

	if t, err := builder.GetObject("iw"); err != nil {
		return nil, err
	} else {
		self.iw = t.(*gtk.IconView)

		self.iw.Connect(
			"key-press-event",
			func() {
				self._SetActiveView(2)
			},
		)

		self.iw.Connect(
			"button-press-event",
			func() {
				self._SetActiveView(2)
			},
		)

	}

	if t, err := builder.GetObject("tb_cut"); err != nil {
		return nil, err
	} else {
		self.tb_cut = t.(*gtk.ToolButton)
	}

	if t, err := builder.GetObject("tb_copy"); err != nil {
		return nil, err
	} else {
		self.tb_copy = t.(*gtk.ToolButton)
	}

	if t, err := builder.GetObject("tb_paste"); err != nil {
		return nil, err
	} else {
		self.tb_paste = t.(*gtk.ToolButton)
	}

	if t, err := builder.GetObject("tb_delete"); err != nil {
		return nil, err
	} else {
		self.tb_delete = t.(*gtk.ToolButton)
	}

	if t, err := builder.GetObject("tb_new_directory"); err != nil {
		return nil, err
	} else {
		self.tb_new_directory = t.(*gtk.ToolButton)
	}

	if t, err := builder.GetObject("tb_up"); err != nil {
		return nil, err
	} else {
		self.tb_up = t.(*gtk.ToolButton)
	}

	if t, err := builder.GetObject("tb_down"); err != nil {
		return nil, err
	} else {
		self.tb_down = t.(*gtk.ToolButton)
	}

	if t, err := builder.GetObject("tb_s1"); err != nil {
		return nil, err
	} else {
		self.tb_s1 = t.(*gtk.SeparatorToolItem)
	}

	if t, err := builder.GetObject("tb_s2"); err != nil {
		return nil, err
	} else {
		self.tb_s2 = t.(*gtk.SeparatorToolItem)
	}

	if t, err := builder.GetObject("lbl_active_view"); err != nil {
		return nil, err
	} else {
		self.lbl_active_view = t.(*gtk.Label)
	}

	if t, err := builder.GetObject("lbl_current_category"); err != nil {
		return nil, err
	} else {
		self.lbl_current_category = t.(*gtk.Label)
	}

	self.tw.Connect(
		"row-activated",
		func(
			tree_view *gtk.TreeView,
			path *gtk.TreePath,
			column *gtk.TreeViewColumn,
		) {
			it, err := self.model.GetIter(path)
			if err != nil {
				return
			}
			isdir_val, err := self.model.GetValue(it, self.isdir_column)
			if err != nil {
				return
			}
			isdir_val_i, err := isdir_val.GoValue()
			if err != nil {
				return
			}
			var ok bool
			isdir_val_b, ok := isdir_val_i.(bool)
			if !ok {
				return
			}
			if isdir_val_b {
				t, err := treemodel.RenderTreePathString(self.model, path, 1)
				if err != nil {
					return
				}
				tt := strings.Join(t, "/")

				self.lbl_current_category.SetText(tt)
				self.entry_category.SetText(tt)

				self.iconic_model.Clear()
				ok := true
				var it2 *gtk.TreeIter
				for {
					if it2 == nil {
						it2 = &gtk.TreeIter{}
						ok = self.model.IterChildren(it, it2)
					} else {
						ok = self.model.IterNext(it2)
					}
					if !ok {
						break
					}
					nit := self.iconic_model.Append()
					t_name, err := self.model.GetValue(it2, self.name_column)
					if err != nil {
						return
					}
					t_name_v, err := t_name.GetString()
					if err != nil {
						return
					}

					t_icon, err := self.model.GetValue(it2, self.icon_column)
					if err != nil {
						return
					}
					t_icon_v, err := t_icon.GoValue()
					if err != nil {
						return
					}
					t_icon_vv, ok := t_icon_v.(*gdk.Pixbuf)
					if !ok {
						return
					}

					t_isdir, err := self.model.GetValue(it2, self.isdir_column)
					if err != nil {
						return
					}
					t_isdir_v, err := t_isdir.GoValue()
					if err != nil {
						return
					}
					t_isdir_vb, ok := t_isdir_v.(bool)
					if !ok {
						return
					}
					self.iconic_model.Set(
						nit,
						[]int{self.name_column, self.icon_column, self.isdir_column},
						[]interface{}{t_name_v, t_icon_vv, t_isdir_vb},
					)
				}
			}
		},
	)

	self.tb_cut.Connect(
		"clicked",
		func() {
			if self.cut_function != nil {
				s, err := self._GetSelectedForOperation()
				if err != nil {
					return
				}
				// TODO: error display
				self.cut_function(s)
			}
		},
	)

	self.tb_copy.Connect(
		"clicked",
		func() {
			if self.copy_function != nil {
				s, err := self._GetSelectedForOperation()
				if err != nil {
					return
				}
				// TODO: error display
				self.copy_function(s)
			}
		},
	)

	self.tb_paste.Connect(
		"clicked",
		func() {
			if self.paste_function != nil {
				s, err := self._GetSelectedForOperation()
				if err != nil {
					return
				}

				if len(s) != 1 {
					return
				}

				ss := s[0]
				// TODO: error display
				self.paste_function(ss)
			}
		},
	)

	self.tb_delete.Connect(
		"clicked",
		func() {
			if self.delete_function != nil {
				s, err := self._GetSelectedForOperation()
				if err != nil {
					return
				}
				// TODO: error display
				self.delete_function(s)
			}
		},
	)

	return self, nil
}

func (self *Explorer) _SetActiveView(index int) {
	if index >= 0 || index <= 2 {
		self.active_view = index
		lbl_active_view_text := "Active Widget: "
		switch index {
		case 0:
			lbl_active_view_text += "Left Tree View"
		case 1:
			lbl_active_view_text += "Middle Tree View"
		case 2:
			lbl_active_view_text += "Middle Icon View"
		}

		self.lbl_active_view.SetText(lbl_active_view_text)

	}
}

type SelectedForOperationItem struct {
	Path  string
	IsDir bool
}

func (self *Explorer) _GetSelectedForOperation() (
	[]*SelectedForOperationItem,
	error,
) {
	switch self.active_view {
	case 0:
		return self._GetSelectedForOperationFromView(self.tw)
	case 1:
		return self._GetSelectedForOperationFromView(self.tw2)
		// TODO
		// case 2:
		// 	return self._GetSelectedForOperationFromView(self.iw)
	}
	return []*SelectedForOperationItem{}, errors.New("programming error")
}

func (self *Explorer) _GetSelectedForOperationFromView(tw *gtk.TreeView) (
	[]*SelectedForOperationItem,
	error,
) {
	ret := make([]*SelectedForOperationItem, 0)

	m, err := tw.GetModel()
	if err != nil {
		return []*SelectedForOperationItem{}, err
	}

	s, err := tw.GetSelection()
	if err != nil {
		return []*SelectedForOperationItem{}, err
	}

	lst := s.GetSelectedRows(m)
	if err != nil {
		return []*SelectedForOperationItem{}, err
	}

	lst.Foreach(
		func(i interface{}) {
			// fmt.Println("lst.Foreach", reflect.TypeOf(i).String(), i)
			switch i.(type) {
			case *gtk.TreePath:
				// noop
			default:
				panic("programming error")
			}

			p := i.(*gtk.TreePath)

			str, err := treemodel.RenderTreePathString(self.model, p, 1)
			if err != nil {
				return
			}

			it, err := self.model.GetIter(p)
			if err != nil {
				return
			}

			it_v, err := self.model.GetValue(it, 2)
			if err != nil {
				return
			}

			it_vv, err := it_v.GoValue()
			if err != nil {
				return
			}

			it_vvb, ok := it_vv.(bool)
			if !ok {
				return
			}

			ret = append(
				ret,
				&SelectedForOperationItem{
					Path:  strings.Join(str, "/"),
					IsDir: it_vvb,
				},
			)

		},
	)

	return ret, nil
}

func (self *Explorer) SetControlls(
	can_new_file, can_new_directory,
	can_move, can_copy, can_delete bool,
) error {

	self.tb_cut.SetVisible(can_move)
	self.tb_copy.SetVisible(can_copy)
	self.tb_paste.SetVisible(can_copy || can_move)
	self.tb_delete.SetVisible(can_delete)

	self.tb_new_directory.SetVisible(can_new_directory)

	self.tb_s1.SetVisible(
		self.tb_cut.GetVisible() ||
			self.tb_copy.GetVisible() ||
			self.tb_paste.GetVisible() ||
			self.tb_delete.GetVisible(),
	)

	self.tb_s2.SetVisible(self.tb_new_directory.GetVisible())

	return nil
}

func (self *Explorer) SetColumns(icon_column, name_column, isdir_column int) error {

	self.icon_column = icon_column
	self.name_column = name_column
	self.isdir_column = isdir_column

	for _, i := range []*gtk.TreeView{
		self.tw,
		self.tw2,
	} {

		{
			{
				for i.GetNColumns() != 0 {
					i.RemoveColumn(i.GetColumn(0))
				}
			}

			{
				r, err := gtk.CellRendererPixbufNew()
				if err != nil {
					return err
				}

				c, err := gtk.TreeViewColumnNewWithAttribute(
					"Icon", r, "pixbuf", icon_column,
				)
				if err != nil {
					return err
				}

				// self.tw_column_name = c

				i.AppendColumn(c)
			}

			{
				r, err := gtk.CellRendererTextNew()
				if err != nil {
					return err
				}

				c, err := gtk.TreeViewColumnNewWithAttribute(
					"Name", r, "text", name_column,
				)
				if err != nil {
					return err
				}

				r.SetProperty("ellipsize", pango.ELLIPSIZE_END)

				c.SetClickable(true)
				c.SetSortColumnID(name_column)
				c.SetSortIndicator(true)

				// self.tw_column_icon = c

				i.AppendColumn(c)
			}
		}
	}
	{
		self.iw.SetProperty("text-column", self.name_column)
		self.iw.SetProperty("pixbuf-column", self.icon_column)
	}

	return nil
}

func (self *Explorer) GetWidget() *gtk.Box {
	return self.root
}

func (self *Explorer) GetModel() *gtk.TreeStore {
	return self.model
}

func (self *Explorer) SetModel(m *gtk.TreeStore) error {
	self.model = m
	self.tw.SetModel(m)

	nmt := make([]glib.Type, 0)

	for i := 0; i != m.GetNColumns(); i++ {
		nmt = append(nmt, m.GetColumnType(i))
	}

	if t, err := gtk.ListStoreNew(nmt...); err != nil {
		return err
	} else {
		self.iconic_model = t
		self.iw.SetModel(self.iconic_model)
		self.tw2.SetModel(self.iconic_model)
	}

	return nil
}

func (self *Explorer) SetCutFunction(
	fn func(selected_items []*SelectedForOperationItem) error,
) {
	self.cut_function = fn
}

func (self *Explorer) SetCopyFunction(
	fn func(selected_items []*SelectedForOperationItem) error,
) {
	self.copy_function = fn
}

func (self *Explorer) SetPasteFunction(
	fn func(target *SelectedForOperationItem) error,
) {
	self.paste_function = fn
}

func (self *Explorer) SetDeleteFunction(
	fn func(selected_items []*SelectedForOperationItem) error,
) {
	self.delete_function = fn
}

func (self *Explorer) SetMkDirAllFunction(
	fn func(path string) error,
) {
	self.mkdirall_function = fn
}
