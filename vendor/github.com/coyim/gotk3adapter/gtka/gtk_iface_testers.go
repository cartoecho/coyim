package gtka

import "github.com/coyim/gotk3adapter/gtki"

func init() {
	gtki.AssertGtk(&RealGtk{})
	gtki.AssertAboutDialog(&aboutDialog{})
	gtki.AssertAccelGroup(&accelGroup{})
	gtki.AssertAdjustment(&adjustment{})
	gtki.AssertApplication(&application{})
	gtki.AssertApplicationWindow(&applicationWindow{})
	gtki.AssertAssistant(&assistant{})
	gtki.AssertBox(&box{})
	gtki.AssertBuilder(&builder{})
	gtki.AssertButton(&button{})
	gtki.AssertCellRenderer(&cellRenderer{})
	gtki.AssertCellRendererText(&cellRendererText{})
	gtki.AssertCellRendererToggle(&cellRendererToggle{})
	gtki.AssertCheckButton(&checkButton{})
	gtki.AssertCheckMenuItem(&checkMenuItem{})
	gtki.AssertComboBox(&comboBox{})
	gtki.AssertComboBoxText(&comboBoxText{})
	gtki.AssertCssProvider(&cssProvider{})
	gtki.AssertDialog(&dialog{})
	gtki.AssertEntry(&entry{})
	gtki.AssertEventBox(&eventBox{})
	gtki.AssertFileChooserDialog(&fileChooserDialog{})
	gtki.AssertGrid(&grid{})
	gtki.AssertHeaderBar(&headerBar{})
	gtki.AssertImage(&image{})
	gtki.AssertInfoBar(&infoBar{})
	gtki.AssertLabel(&label{})
	gtki.AssertListStore(&listStore{})
	gtki.AssertMenuBar(&menuBar{})
	gtki.AssertMenuItem(&menuItem{})
	gtki.AssertMenu(&menu{})
	gtki.AssertMessageDialog(&messageDialog{})
	gtki.AssertNotebook(&notebook{})
	gtki.AssertProgressBar(&progressBar{})
	gtki.AssertRevealer(&revealer{})
	gtki.AssertScrolledWindow(&scrolledWindow{})
	gtki.AssertSearchBar(&searchBar{})
	gtki.AssertSearchEntry(&searchEntry{})
	gtki.AssertSeparatorMenuItem(&separatorMenuItem{})
	gtki.AssertSettings(&settings{})
	gtki.AssertSpinner(&spinner{})
	gtki.AssertSpinButton(&spinButton{})
	gtki.AssertStyleContext(&styleContext{})
	gtki.AssertTextBuffer(&textBuffer{})
	gtki.AssertTextIter(&textIter{})
	gtki.AssertTextMark(&textMark{})
	gtki.AssertTextTagTable(&textTagTable{})
	gtki.AssertTextTag(&textTag{})
	gtki.AssertTextView(&textView{})
	gtki.AssertTreeIter(&treeIter{})
	gtki.AssertTreePath(&treePath{})
	gtki.AssertTreeSelection(&treeSelection{})
	gtki.AssertTreeStore(&treeStore{})
	gtki.AssertTreeView(&treeView{})
	gtki.AssertTreeViewColumn(&treeViewColumn{})
	gtki.AssertWidget(&widget{})
	gtki.AssertWindow(&window{})
}