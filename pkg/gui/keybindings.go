package gui

import (
	"github.com/jesseduffield/gocui"
	"log"
	"strings"
)

// Binding - a keybinding mapping a key and modifier to a handler. The keypress
// is only handled if the given view has focus, or handled globally if the view
// is ""
type Binding struct {
	ViewName    string
	Contexts    []string
	Handler     func(*gocui.Gui, *gocui.View) error
	Key         interface{} // FIXME: find out how to get `gocui.Key | rune`
	Modifier    gocui.Modifier
	Description string
	Alternative string
}

// GetDisplayStrings returns the display string of a file
func (b *Binding) GetDisplayStrings(isFocused bool) []string {
	return []string{b.GetKey(), b.Description}
}

// GetKey is a function.
func (b *Binding) GetKey() string {
	key := 0

	switch b.Key.(type) {
	case rune:
		key = int(b.Key.(rune))
	case gocui.Key:
		if b.Key.(gocui.Key) == gocui.KeyCtrlA {
			return "ctrl+a"
		}
		if b.Key.(gocui.Key) == gocui.KeyCtrlB {
			return "ctrl+b"
		}
		if b.Key.(gocui.Key) == gocui.KeyCtrlC {
			return "ctrl+c"
		}
		if b.Key.(gocui.Key) == gocui.KeyCtrlD {
			return "ctrl+d"
		}
		if b.Key.(gocui.Key) == gocui.KeyCtrlE {
			return "ctrl+e"
		}
		if b.Key.(gocui.Key) == gocui.KeyCtrlF {
			return "ctrl+f"
		}
		if b.Key.(gocui.Key) == gocui.KeyCtrlG {
			return "ctrl+g"
		}
		if b.Key.(gocui.Key) == gocui.KeyCtrlH {
			return "ctrl+h"
		}
		if b.Key.(gocui.Key) == gocui.KeyCtrlI {
			return "ctrl+i"
		}
		if b.Key.(gocui.Key) == gocui.KeyCtrlJ {
			return "ctrl+j"
		}
		if b.Key.(gocui.Key) == gocui.KeyCtrlK {
			return "ctrl+k"
		}
		if b.Key.(gocui.Key) == gocui.KeyCtrlL {
			return "ctrl+l"
		}
		if b.Key.(gocui.Key) == gocui.KeyCtrlM {
			return "ctrl+m"
		}
		if b.Key.(gocui.Key) == gocui.KeyCtrlN {
			return "ctrl+n"
		}
		if b.Key.(gocui.Key) == gocui.KeyCtrlO {
			return "ctrl+o"
		}
		if b.Key.(gocui.Key) == gocui.KeyCtrlP {
			return "ctrl+p"
		}
		if b.Key.(gocui.Key) == gocui.KeyCtrlQ {
			return "ctrl+q"
		}
		if b.Key.(gocui.Key) == gocui.KeyCtrlR {
			return "ctrl+r"
		}
		if b.Key.(gocui.Key) == gocui.KeyCtrlS {
			return "ctrl+s"
		}
		if b.Key.(gocui.Key) == gocui.KeyCtrlT {
			return "ctrl+t"
		}
		if b.Key.(gocui.Key) == gocui.KeyCtrlU {
			return "ctrl+u"
		}
		if b.Key.(gocui.Key) == gocui.KeyCtrlV {
			return "ctrl+v"
		}
		if b.Key.(gocui.Key) == gocui.KeyCtrlW {
			return "ctrl+w"
		}
		if b.Key.(gocui.Key) == gocui.KeyCtrlX {
			return "ctrl+x"
		}
		if b.Key.(gocui.Key) == gocui.KeyCtrlY {
			return "ctrl+y"
		}
		if b.Key.(gocui.Key) == gocui.KeyCtrlZ {
			return "ctrl+z"
		}
		if b.Key.(gocui.Key) == gocui.KeyCtrlTilde {
			return "ctrl+~"
		}
		if b.Key.(gocui.Key) == gocui.KeyCtrl2 {
			return "ctrl+2"
		}
		if b.Key.(gocui.Key) == gocui.KeyCtrl3 {
			return "ctrl+3"
		}
		if b.Key.(gocui.Key) == gocui.KeyCtrl4 {
			return "ctrl+4"
		}
		if b.Key.(gocui.Key) == gocui.KeyCtrl5 {
			return "ctrl+5"
		}
		if b.Key.(gocui.Key) == gocui.KeyCtrl6 {
			return "ctrl+6"
		}
		if b.Key.(gocui.Key) == gocui.KeyCtrl7 {
			return "ctrl+7"
		}
		if b.Key.(gocui.Key) == gocui.KeyCtrl8 {
			return "ctrl+8"
		}
		if b.Key.(gocui.Key) == gocui.KeyCtrlSpace {
			return "ctrl+space"
		}
		if b.Key.(gocui.Key) == gocui.KeyCtrlBackslash {
			return "ctrl+\\"
		}
		if b.Key.(gocui.Key) == gocui.KeyCtrlLsqBracket {
			return "ctrl+["
		}
		if b.Key.(gocui.Key) == gocui.KeyCtrlRsqBracket {
			return "ctrl+]"
		}
		if b.Key.(gocui.Key) == gocui.KeyCtrlSlash {
			return "ctrl+/"
		}
		if b.Key.(gocui.Key) == gocui.KeyCtrlUnderscore {
			return "ctrl+_"
		}
		if b.Key.(gocui.Key) == gocui.KeyBackspace {
			return "backspace"
		}
		if b.Key.(gocui.Key) == gocui.KeyTab {
			return "tab"
		}
		if b.Key.(gocui.Key) == gocui.KeyEnter {
			return "enter"
		}
		if b.Key.(gocui.Key) == gocui.KeyEsc {
			return "esc"
		}
		if b.Key.(gocui.Key) == gocui.KeySpace {
			return "space"
		}
		if b.Key.(gocui.Key) == gocui.KeyF1 {
			return "f1"
		}
		if b.Key.(gocui.Key) == gocui.KeyF2 {
			return "f2"
		}
		if b.Key.(gocui.Key) == gocui.KeyF3 {
			return "f3"
		}
		if b.Key.(gocui.Key) == gocui.KeyF4 {
			return "f4"
		}
		if b.Key.(gocui.Key) == gocui.KeyF5 {
			return "f5"
		}
		if b.Key.(gocui.Key) == gocui.KeyF6 {
			return "f6"
		}
		if b.Key.(gocui.Key) == gocui.KeyF7 {
			return "f7"
		}
		if b.Key.(gocui.Key) == gocui.KeyF8 {
			return "f8"
		}
		if b.Key.(gocui.Key) == gocui.KeyF9 {
			return "f9"
		}
		if b.Key.(gocui.Key) == gocui.KeyF10 {
			return "f10"
		}
		if b.Key.(gocui.Key) == gocui.KeyF11 {
			return "f11"
		}
		if b.Key.(gocui.Key) == gocui.KeyF12 {
			return "f12"
		}
		if b.Key.(gocui.Key) == gocui.KeyInsert {
			return "insert"
		}
		if b.Key.(gocui.Key) == gocui.KeyDelete {
			return "delete"
		}
		if b.Key.(gocui.Key) == gocui.KeyHome {
			return "home"
		}
		if b.Key.(gocui.Key) == gocui.KeyEnd {
			return "end"
		}
		if b.Key.(gocui.Key) == gocui.KeyPgup {
			return "pgup"
		}
		if b.Key.(gocui.Key) == gocui.KeyPgdn {
			return "pgdown"
		}
		if b.Key.(gocui.Key) == gocui.KeyArrowUp {
			return "up"
		}
		if b.Key.(gocui.Key) == gocui.KeyArrowDown {
			return "down"
		}
		if b.Key.(gocui.Key) == gocui.KeyArrowLeft {
			return "left"
		}
		if b.Key.(gocui.Key) == gocui.KeyArrowRight {
			return "right"
		}
		key = int(b.Key.(gocui.Key))
	}

	// special keys
	switch key {
	case 27:
		return "esc"
	case 13:
		return "enter"
	case 32:
		return "space"
	case 65514:
		return "►"
	case 65515:
		return "◄"
	case 65517:
		return "▲"
	case 65516:
		return "▼"
	case 65508:
		return "PgUp"
	case 65507:
		return "PgDn"
	case 9:
		return "tab"
	}

	return string(key)
}

var keymap = map[string]interface{}{
	"<c-a>":       gocui.KeyCtrlA,
	"<c-b>":       gocui.KeyCtrlB,
	"<c-c>":       gocui.KeyCtrlC,
	"<c-d>":       gocui.KeyCtrlD,
	"<c-e>":       gocui.KeyCtrlE,
	"<c-f>":       gocui.KeyCtrlF,
	"<c-g>":       gocui.KeyCtrlG,
	"<c-h>":       gocui.KeyCtrlH,
	"<c-i>":       gocui.KeyCtrlI,
	"<c-j>":       gocui.KeyCtrlJ,
	"<c-k>":       gocui.KeyCtrlK,
	"<c-l>":       gocui.KeyCtrlL,
	"<c-m>":       gocui.KeyCtrlM,
	"<c-n>":       gocui.KeyCtrlN,
	"<c-o>":       gocui.KeyCtrlO,
	"<c-p>":       gocui.KeyCtrlP,
	"<c-q>":       gocui.KeyCtrlQ,
	"<c-r>":       gocui.KeyCtrlR,
	"<c-s>":       gocui.KeyCtrlS,
	"<c-t>":       gocui.KeyCtrlT,
	"<c-u>":       gocui.KeyCtrlU,
	"<c-v>":       gocui.KeyCtrlV,
	"<c-w>":       gocui.KeyCtrlW,
	"<c-x>":       gocui.KeyCtrlX,
	"<c-y>":       gocui.KeyCtrlY,
	"<c-z>":       gocui.KeyCtrlZ,
	"<c-~>":       gocui.KeyCtrlTilde,
	"<c-2>":       gocui.KeyCtrl2,
	"<c-3>":       gocui.KeyCtrl3,
	"<c-4>":       gocui.KeyCtrl4,
	"<c-5>":       gocui.KeyCtrl5,
	"<c-6>":       gocui.KeyCtrl6,
	"<c-7>":       gocui.KeyCtrl7,
	"<c-8>":       gocui.KeyCtrl8,
	"<c-space>":   gocui.KeyCtrlSpace,
	"<c-\\>":      gocui.KeyCtrlBackslash,
	"<c-[>":       gocui.KeyCtrlLsqBracket,
	"<c-]>":       gocui.KeyCtrlRsqBracket,
	"<c-/>":       gocui.KeyCtrlSlash,
	"<c-_>":       gocui.KeyCtrlUnderscore,
	"<backspace>": gocui.KeyBackspace,
	"<tab>":       gocui.KeyTab,
	"<enter>":     gocui.KeyEnter,
	"<esc>":       gocui.KeyEsc,
	"<space>":     gocui.KeySpace,
	"<f1>":        gocui.KeyF1,
	"<f2>":        gocui.KeyF2,
	"<f3>":        gocui.KeyF3,
	"<f4>":        gocui.KeyF4,
	"<f5>":        gocui.KeyF5,
	"<f6>":        gocui.KeyF6,
	"<f7>":        gocui.KeyF7,
	"<f8>":        gocui.KeyF8,
	"<f9>":        gocui.KeyF9,
	"<f10>":       gocui.KeyF10,
	"<f11>":       gocui.KeyF11,
	"<f12>":       gocui.KeyF12,
	"<insert>":    gocui.KeyInsert,
	"<delete>":    gocui.KeyDelete,
	"<home>":      gocui.KeyHome,
	"<end>":       gocui.KeyEnd,
	"<pgup>":      gocui.KeyPgup,
	"<pgdown>":    gocui.KeyPgdn,
	"<up>":        gocui.KeyArrowUp,
	"<down>":      gocui.KeyArrowDown,
	"<left>":      gocui.KeyArrowLeft,
	"<right>":     gocui.KeyArrowRight,
}

func (gui *Gui) getKey(name string) interface{} {
	key := gui.Config.GetUserConfig().GetString("keybinding." + name)
	if len(key) > 1 {
		binding := keymap[strings.ToLower(key)]
		if binding == nil {
			log.Fatal("Unrecognized binding: " + strings.ToLower(key))
		} else {
			return binding
		}
	} else if len(key) == 1 {
		return []rune(key)[0]
	}
	log.Fatal("Key empty for function: " + strings.ToLower(name))
	return nil
}

// GetInitialKeybindings is a function.
func (gui *Gui) GetInitialKeybindings() []*Binding {
	bindings := []*Binding{
		{
			ViewName: "",
			Key:      gui.getKey("universal.quit"),
			Modifier: gocui.ModNone,
			Handler:  gui.handleQuit,
		},
		{
			ViewName: "",
			Key:      gui.getKey("universal.quitWithoutChangingDirectory"),
			Modifier: gocui.ModNone,
			Handler:  gui.handleQuitWithoutChangingDirectory,
		},
		{
			ViewName: "",
			Key:      gui.getKey("universal.quit-alt1"),
			Modifier: gocui.ModNone,
			Handler:  gui.handleQuit,
		},
		{
			ViewName: "",
			Key:      gui.getKey("universal.return"),
			Modifier: gocui.ModNone,
			Handler:  gui.handleQuit,
		},
		{
			ViewName:    "",
			Key:         gui.getKey("universal.scrollUpMain"),
			Modifier:    gocui.ModNone,
			Handler:     gui.scrollUpMain,
			Alternative: "fn+up",
		},
		{
			ViewName:    "",
			Key:         gui.getKey("universal.scrollDownMain"),
			Modifier:    gocui.ModNone,
			Handler:     gui.scrollDownMain,
			Alternative: "fn+down",
		},
		{
			ViewName: "",
			Key:      gui.getKey("universal.scrollUpMain-alt1"),
			Modifier: gocui.ModNone,
			Handler:  gui.scrollUpMain,
		},
		{
			ViewName: "",
			Key:      gui.getKey("universal.scrollDownMain-alt1"),
			Modifier: gocui.ModNone,
			Handler:  gui.scrollDownMain,
		},
		{
			ViewName: "",
			Key:      gui.getKey("universal.scrollUpMain-alt2"),
			Modifier: gocui.ModNone,
			Handler:  gui.scrollUpMain,
		},
		{
			ViewName: "",
			Key:      gui.getKey("universal.scrollDownMain-alt2"),
			Modifier: gocui.ModNone,
			Handler:  gui.scrollDownMain,
		},
		{
			ViewName:    "",
			Key:         gui.getKey("universal.createRebaseOptionsMenu"),
			Modifier:    gocui.ModNone,
			Handler:     gui.handleCreateRebaseOptionsMenu,
			Description: gui.Tr.SLocalize("ViewMergeRebaseOptions"),
		},
		{
			ViewName:    "",
			Key:         gui.getKey("universal.pushFiles"),
			Modifier:    gocui.ModNone,
			Handler:     gui.pushFiles,
			Description: gui.Tr.SLocalize("push"),
		},
		{
			ViewName:    "",
			Key:         gui.getKey("universal.pullFiles"),
			Modifier:    gocui.ModNone,
			Handler:     gui.handlePullFiles,
			Description: gui.Tr.SLocalize("pull"),
		},
		{
			ViewName:    "",
			Key:         gui.getKey("universal.refresh"),
			Modifier:    gocui.ModNone,
			Handler:     gui.handleRefresh,
			Description: gui.Tr.SLocalize("refresh"),
		},
		{
			ViewName: "",
			Key:      gui.getKey("universal.optionMenu"),
			Modifier: gocui.ModNone,
			Handler:  gui.handleCreateOptionsMenu,
		},
		{
			ViewName: "",
			Key:      gui.getKey("universal.optionMenu-alt1"),
			Modifier: gocui.ModNone,
			Handler:  gui.handleCreateOptionsMenu,
		},
		{
			ViewName: "",
			Key:      gocui.MouseMiddle,
			Modifier: gocui.ModNone,
			Handler:  gui.handleCreateOptionsMenu,
		},
		{
			ViewName: "",
			Key:      gui.getKey("universal.createPatchOptionsMenu"),
			Modifier: gocui.ModNone,
			Handler:  gui.handleCreatePatchOptionsMenu,
		},
		{
			ViewName:    "status",
			Key:         gui.getKey("universal.edit"),
			Modifier:    gocui.ModNone,
			Handler:     gui.handleEditConfig,
			Description: gui.Tr.SLocalize("EditConfig"),
		},
		{
			ViewName:    "status",
			Key:         gui.getKey("universal.openFile"),
			Modifier:    gocui.ModNone,
			Handler:     gui.handleOpenConfig,
			Description: gui.Tr.SLocalize("OpenConfig"),
		},
		{
			ViewName:    "status",
			Key:         gui.getKey("status.checkForUpdate"),
			Modifier:    gocui.ModNone,
			Handler:     gui.handleCheckForUpdate,
			Description: gui.Tr.SLocalize("checkForUpdate"),
		},
		{
			ViewName:    "status",
			Key:         gui.getKey("status.recentRepos"),
			Modifier:    gocui.ModNone,
			Handler:     gui.handleCreateRecentReposMenu,
			Description: gui.Tr.SLocalize("SwitchRepo"),
		},
		{
			ViewName:    "files",
			Key:         gui.getKey("files.commitChanges"),
			Modifier:    gocui.ModNone,
			Handler:     gui.handleCommitPress,
			Description: gui.Tr.SLocalize("CommitChanges"),
		},
		{
			ViewName:    "files",
			Key:         gui.getKey("files.commitChangesWithoutHook"),
			Modifier:    gocui.ModNone,
			Handler:     gui.handleWIPCommitPress,
			Description: gui.Tr.SLocalize("commitChangesWithoutHook"),
		},
		{
			ViewName:    "files",
			Key:         gui.getKey("files.amendLastCommit"),
			Modifier:    gocui.ModNone,
			Handler:     gui.handleAmendCommitPress,
			Description: gui.Tr.SLocalize("AmendLastCommit"),
		},
		{
			ViewName:    "files",
			Key:         gui.getKey("files.commitChangesWithEditor"),
			Modifier:    gocui.ModNone,
			Handler:     gui.handleCommitEditorPress,
			Description: gui.Tr.SLocalize("CommitChangesWithEditor"),
		},
		{
			ViewName:    "files",
			Key:         gui.getKey("universal.select"),
			Modifier:    gocui.ModNone,
			Handler:     gui.handleFilePress,
			Description: gui.Tr.SLocalize("toggleStaged"),
		},
		{
			ViewName:    "files",
			Key:         gui.getKey("universal.remove"),
			Modifier:    gocui.ModNone,
			Handler:     gui.handleCreateDiscardMenu,
			Description: gui.Tr.SLocalize("viewDiscardOptions"),
		},
		{
			ViewName:    "files",
			Key:         gui.getKey("universal.edit"),
			Modifier:    gocui.ModNone,
			Handler:     gui.handleFileEdit,
			Description: gui.Tr.SLocalize("editFile"),
		},
		{
			ViewName:    "files",
			Key:         gui.getKey("universal.openFile"),
			Modifier:    gocui.ModNone,
			Handler:     gui.handleFileOpen,
			Description: gui.Tr.SLocalize("openFile"),
		},
		{
			ViewName:    "files",
			Key:         gui.getKey("files.ignoreFile"),
			Modifier:    gocui.ModNone,
			Handler:     gui.handleIgnoreFile,
			Description: gui.Tr.SLocalize("ignoreFile"),
		},
		{
			ViewName:    "files",
			Key:         gui.getKey("files.refreshFiles"),
			Modifier:    gocui.ModNone,
			Handler:     gui.handleRefreshFiles,
			Description: gui.Tr.SLocalize("refreshFiles"),
		},
		{
			ViewName:    "files",
			Key:         gui.getKey("files.stashAllChanges"),
			Modifier:    gocui.ModNone,
			Handler:     gui.handleStashChanges,
			Description: gui.Tr.SLocalize("stashAllChanges"),
		},
		{
			ViewName:    "files",
			Key:         gui.getKey("files.viewStashOptions"),
			Modifier:    gocui.ModNone,
			Handler:     gui.handleCreateStashMenu,
			Description: gui.Tr.SLocalize("viewStashOptions"),
		},
		{
			ViewName:    "files",
			Key:         gui.getKey("files.toggleStagedAll"),
			Modifier:    gocui.ModNone,
			Handler:     gui.handleStageAll,
			Description: gui.Tr.SLocalize("toggleStagedAll"),
		},
		{
			ViewName:    "files",
			Key:         gui.getKey("files.viewResetOptions"),
			Modifier:    gocui.ModNone,
			Handler:     gui.handleCreateResetMenu,
			Description: gui.Tr.SLocalize("viewResetOptions"),
		},
		{
			ViewName:    "files",
			Key:         gui.getKey("universal.goInto"),
			Modifier:    gocui.ModNone,
			Handler:     gui.handleEnterFile,
			Description: gui.Tr.SLocalize("StageLines"),
		},
		{
			ViewName:    "files",
			Key:         gui.getKey("files.fetch"),
			Modifier:    gocui.ModNone,
			Handler:     gui.handleGitFetch,
			Description: gui.Tr.SLocalize("fetch"),
		},
		{
			ViewName:    "files",
			Key:         gui.getKey("universal.executeCustomCommand"),
			Modifier:    gocui.ModNone,
			Handler:     gui.handleCustomCommand,
			Description: gui.Tr.SLocalize("executeCustomCommand"),
		},
		{
			ViewName:    "branches",
			Contexts:    []string{"local-branches"},
			Key:         gui.getKey("universal.select"),
			Modifier:    gocui.ModNone,
			Handler:     gui.handleBranchPress,
			Description: gui.Tr.SLocalize("checkout"),
		},
		{
			ViewName:    "branches",
			Contexts:    []string{"local-branches"},
			Key:         gui.getKey("branches.createPullRequest"),
			Modifier:    gocui.ModNone,
			Handler:     gui.handleCreatePullRequestPress,
			Description: gui.Tr.SLocalize("createPullRequest"),
		},
		{
			ViewName:    "branches",
			Contexts:    []string{"local-branches"},
			Key:         gui.getKey("branches.checkoutBranchByName"),
			Modifier:    gocui.ModNone,
			Handler:     gui.handleCheckoutByName,
			Description: gui.Tr.SLocalize("checkoutByName"),
		},
		{
			ViewName:    "branches",
			Contexts:    []string{"local-branches"},
			Key:         gui.getKey("branches.forceCheckoutBranch"),
			Modifier:    gocui.ModNone,
			Handler:     gui.handleForceCheckout,
			Description: gui.Tr.SLocalize("forceCheckout"),
		},
		{
			ViewName:    "branches",
			Contexts:    []string{"local-branches"},
			Key:         gui.getKey("universal.new"),
			Modifier:    gocui.ModNone,
			Handler:     gui.handleNewBranch,
			Description: gui.Tr.SLocalize("newBranch"),
		},
		{
			ViewName:    "branches",
			Contexts:    []string{"local-branches"},
			Key:         gui.getKey("universal.remove"),
			Modifier:    gocui.ModNone,
			Handler:     gui.handleDeleteBranch,
			Description: gui.Tr.SLocalize("deleteBranch"),
		},
		{
			ViewName:    "branches",
			Contexts:    []string{"local-branches"},
			Key:         gui.getKey("branches.rebaseBranch"),
			Modifier:    gocui.ModNone,
			Handler:     gui.handleRebaseOntoLocalBranch,
			Description: gui.Tr.SLocalize("rebaseBranch"),
		},
		{
			ViewName:    "branches",
			Contexts:    []string{"local-branches"},
			Key:         gui.getKey("branches.mergeIntoCurrentBranch"),
			Modifier:    gocui.ModNone,
			Handler:     gui.handleMerge,
			Description: gui.Tr.SLocalize("mergeIntoCurrentBranch"),
		},
		{
			ViewName:    "branches",
			Contexts:    []string{"local-branches"},
			Key:         gui.getKey("branches.FastForward"),
			Modifier:    gocui.ModNone,
			Handler:     gui.handleFastForward,
			Description: gui.Tr.SLocalize("FastForward"),
		},
		{
			ViewName:    "branches",
			Contexts:    []string{"tags"},
			Key:         gui.getKey("universal.select"),
			Modifier:    gocui.ModNone,
			Handler:     gui.handleCheckoutTag,
			Description: gui.Tr.SLocalize("checkout"),
		},
		{
			ViewName:    "branches",
			Contexts:    []string{"tags"},
			Key:         gui.getKey("universal.remove"),
			Modifier:    gocui.ModNone,
			Handler:     gui.handleDeleteTag,
			Description: gui.Tr.SLocalize("deleteTag"),
		},
		{
			ViewName:    "branches",
			Contexts:    []string{"tags"},
			Key:         gui.getKey("branches.pushTag"),
			Modifier:    gocui.ModNone,
			Handler:     gui.handlePushTag,
			Description: gui.Tr.SLocalize("pushTag"),
		},
		{
			ViewName:    "branches",
			Contexts:    []string{"tags"},
			Key:         gui.getKey("universal.new"),
			Modifier:    gocui.ModNone,
			Handler:     gui.handleCreateTag,
			Description: gui.Tr.SLocalize("createTag"),
		},
		{
			ViewName: "branches",
			Key:      gui.getKey("universal.nextBranchTab"),
			Modifier: gocui.ModNone,
			Handler:  gui.handleNextBranchesTab,
		},
		{
			ViewName: "branches",
			Key:      gui.getKey("universal.prevBranchTab"),
			Modifier: gocui.ModNone,
			Handler:  gui.handlePrevBranchesTab,
		},
		{
			ViewName:    "branches",
			Contexts:    []string{"remote-branches"},
			Key:         gui.getKey("universal.return"),
			Modifier:    gocui.ModNone,
			Handler:     gui.handleRemoteBranchesEscape,
			Description: gui.Tr.SLocalize("ReturnToRemotesList"),
		},
		{
			ViewName:    "commits",
			Key:         gui.getKey("commits.squashDown"),
			Modifier:    gocui.ModNone,
			Handler:     gui.handleCommitSquashDown,
			Description: gui.Tr.SLocalize("squashDown"),
		},
		{
			ViewName:    "commits",
			Key:         gui.getKey("commits.renameCommit"),
			Modifier:    gocui.ModNone,
			Handler:     gui.handleRenameCommit,
			Description: gui.Tr.SLocalize("renameCommit"),
		},
		{
			ViewName:    "commits",
			Key:         gui.getKey("commits.renameCommitWithEditor"),
			Modifier:    gocui.ModNone,
			Handler:     gui.handleRenameCommitEditor,
			Description: gui.Tr.SLocalize("renameCommitEditor"),
		},
		{
			ViewName:    "commits",
			Key:         gui.getKey("commits.viewResetOptions"),
			Modifier:    gocui.ModNone,
			Handler:     gui.handleCreateCommitResetMenu,
			Description: gui.Tr.SLocalize("resetToThisCommit"),
		},
		{
			ViewName:    "commits",
			Key:         gui.getKey("commits.markCommitAsFixup"),
			Modifier:    gocui.ModNone,
			Handler:     gui.handleCommitFixup,
			Description: gui.Tr.SLocalize("fixupCommit"),
		},
		{
			ViewName:    "commits",
			Key:         gui.getKey("commits.createFixupCommit"),
			Modifier:    gocui.ModNone,
			Handler:     gui.handleCreateFixupCommit,
			Description: gui.Tr.SLocalize("createFixupCommit"),
		},
		{
			ViewName:    "commits",
			Key:         gui.getKey("commits.squashAboveCommits"),
			Modifier:    gocui.ModNone,
			Handler:     gui.handleSquashAllAboveFixupCommits,
			Description: gui.Tr.SLocalize("squashAboveCommits"),
		},
		{
			ViewName:    "commits",
			Key:         gui.getKey("universal.remove"),
			Modifier:    gocui.ModNone,
			Handler:     gui.handleCommitDelete,
			Description: gui.Tr.SLocalize("deleteCommit"),
		},
		{
			ViewName:    "commits",
			Key:         gui.getKey("commits.moveDownCommit"),
			Modifier:    gocui.ModNone,
			Handler:     gui.handleCommitMoveDown,
			Description: gui.Tr.SLocalize("moveDownCommit"),
		},
		{
			ViewName:    "commits",
			Key:         gui.getKey("commits.moveUpCommit"),
			Modifier:    gocui.ModNone,
			Handler:     gui.handleCommitMoveUp,
			Description: gui.Tr.SLocalize("moveUpCommit"),
		},
		{
			ViewName:    "commits",
			Key:         gui.getKey("universal.edit"),
			Modifier:    gocui.ModNone,
			Handler:     gui.handleCommitEdit,
			Description: gui.Tr.SLocalize("editCommit"),
		},
		{
			ViewName:    "commits",
			Key:         gui.getKey("commits.amendToCommit"),
			Modifier:    gocui.ModNone,
			Handler:     gui.handleCommitAmendTo,
			Description: gui.Tr.SLocalize("amendToCommit"),
		},
		{
			ViewName:    "commits",
			Key:         gui.getKey("commits.pickCommit"),
			Modifier:    gocui.ModNone,
			Handler:     gui.handleCommitPick,
			Description: gui.Tr.SLocalize("pickCommit"),
		},
		{
			ViewName:    "commits",
			Key:         gui.getKey("commits.revertCommit"),
			Modifier:    gocui.ModNone,
			Handler:     gui.handleCommitRevert,
			Description: gui.Tr.SLocalize("revertCommit"),
		},
		{
			ViewName:    "commits",
			Key:         gui.getKey("commits.cherryPickCopy"),
			Modifier:    gocui.ModNone,
			Handler:     gui.handleCopyCommit,
			Description: gui.Tr.SLocalize("cherryPickCopy"),
		},
		{
			ViewName:    "commits",
			Key:         gui.getKey("commits.cherryPickCopyRange"),
			Modifier:    gocui.ModNone,
			Handler:     gui.handleCopyCommitRange,
			Description: gui.Tr.SLocalize("cherryPickCopyRange"),
		},
		{
			ViewName:    "commits",
			Key:         gui.getKey("commits.pasteCommits"),
			Modifier:    gocui.ModNone,
			Handler:     gui.HandlePasteCommits,
			Description: gui.Tr.SLocalize("pasteCommits"),
		},
		{
			ViewName:    "commits",
			Key:         gui.getKey("universal.goInto"),
			Modifier:    gocui.ModNone,
			Handler:     gui.handleSwitchToCommitFilesPanel,
			Description: gui.Tr.SLocalize("viewCommitFiles"),
		},
		{
			ViewName:    "commits",
			Key:         gui.getKey("universal.select"),
			Modifier:    gocui.ModNone,
			Handler:     gui.handleToggleDiffCommit,
			Description: gui.Tr.SLocalize("CommitsDiff"),
		},
		{
			ViewName:    "commits",
			Key:         gui.getKey("commits.tagCommit"),
			Modifier:    gocui.ModNone,
			Handler:     gui.handleTagCommit,
			Description: gui.Tr.SLocalize("tagCommit"),
		},
		{
			ViewName:    "stash",
			Key:         gui.getKey("universal.select"),
			Modifier:    gocui.ModNone,
			Handler:     gui.handleStashApply,
			Description: gui.Tr.SLocalize("apply"),
		},
		{
			ViewName:    "stash",
			Key:         gui.getKey("stash.popStash"),
			Modifier:    gocui.ModNone,
			Handler:     gui.handleStashPop,
			Description: gui.Tr.SLocalize("pop"),
		},
		{
			ViewName:    "stash",
			Key:         gui.getKey("universal.remove"),
			Modifier:    gocui.ModNone,
			Handler:     gui.handleStashDrop,
			Description: gui.Tr.SLocalize("drop"),
		},
		{
			ViewName: "commitMessage",
			Key:      gocui.KeyEnter,
			Modifier: gocui.ModNone,
			Handler:  gui.handleCommitConfirm,
		},
		{
			ViewName: "commitMessage",
			Key:      gocui.KeyEsc,
			Modifier: gocui.ModNone,
			Handler:  gui.handleCommitClose,
		},
		{
			ViewName: "credentials",
			Key:      gocui.KeyEnter,
			Modifier: gocui.ModNone,
			Handler:  gui.handleSubmitCredential,
		},
		{
			ViewName: "credentials",
			Key:      gocui.KeyEsc,
			Modifier: gocui.ModNone,
			Handler:  gui.handleCloseCredentialsView,
		},
		{
			ViewName: "menu",
			Key:      gui.getKey("universal.return"),
			Modifier: gocui.ModNone,
			Handler:  gui.handleMenuClose,
		},
		{
			ViewName: "menu",
			Key:      gui.getKey("universal.quit"),
			Modifier: gocui.ModNone,
			Handler:  gui.handleMenuClose,
		},
		{
			ViewName: "information",
			Key:      gocui.MouseLeft,
			Modifier: gocui.ModNone,
			Handler:  gui.handleDonate,
		},
		{
			ViewName:    "commitFiles",
			Key:         gui.getKey("universal.return"),
			Modifier:    gocui.ModNone,
			Handler:     gui.handleSwitchToCommitsPanel,
			Description: gui.Tr.SLocalize("goBack"),
		},
		{
			ViewName:    "commitFiles",
			Key:         gui.getKey("commitFiles.checkoutCommitFile"),
			Modifier:    gocui.ModNone,
			Handler:     gui.handleCheckoutCommitFile,
			Description: gui.Tr.SLocalize("checkoutCommitFile"),
		},
		{
			ViewName:    "commitFiles",
			Key:         gui.getKey("universal.remove"),
			Modifier:    gocui.ModNone,
			Handler:     gui.handleDiscardOldFileChange,
			Description: gui.Tr.SLocalize("discardOldFileChange"),
		},
		{
			ViewName:    "commitFiles",
			Key:         gui.getKey("universal.openFile"),
			Modifier:    gocui.ModNone,
			Handler:     gui.handleOpenOldCommitFile,
			Description: gui.Tr.SLocalize("openFile"),
		},
		{
			ViewName:    "commitFiles",
			Key:         gui.getKey("universal.select"),
			Modifier:    gocui.ModNone,
			Handler:     gui.handleToggleFileForPatch,
			Description: gui.Tr.SLocalize("toggleAddToPatch"),
		},
		{
			ViewName:    "commitFiles",
			Key:         gui.getKey("universal.goInto"),
			Modifier:    gocui.ModNone,
			Handler:     gui.handleEnterCommitFile,
			Description: gui.Tr.SLocalize("enterFile"),
		},
		{
			ViewName: "secondary",
			Key:      gocui.MouseWheelUp,
			Modifier: gocui.ModNone,
			Handler:  gui.scrollUpSecondary,
		},
		{
			ViewName: "secondary",
			Key:      gocui.MouseWheelDown,
			Modifier: gocui.ModNone,
			Handler:  gui.scrollDownSecondary,
		},
		{
			ViewName: "secondary",
			Contexts: []string{"normal"},
			Key:      gocui.MouseLeft,
			Modifier: gocui.ModNone,
			Handler:  gui.handleMouseDownSecondary,
		},
		{
			ViewName:    "main",
			Contexts:    []string{"normal"},
			Key:         gocui.MouseWheelDown,
			Modifier:    gocui.ModNone,
			Handler:     gui.scrollDownMain,
			Description: gui.Tr.SLocalize("ScrollDown"),
			Alternative: "fn+up",
		},
		{
			ViewName:    "main",
			Contexts:    []string{"normal"},
			Key:         gocui.MouseWheelUp,
			Modifier:    gocui.ModNone,
			Handler:     gui.scrollUpMain,
			Description: gui.Tr.SLocalize("ScrollUp"),
			Alternative: "fn+down",
		},
		{
			ViewName: "main",
			Contexts: []string{"normal"},
			Key:      gocui.MouseLeft,
			Modifier: gocui.ModNone,
			Handler:  gui.handleMouseDownMain,
		},
		{
			ViewName: "secondary",
			Contexts: []string{"staging"},
			Key:      gocui.MouseLeft,
			Modifier: gocui.ModNone,
			Handler:  gui.handleTogglePanelClick,
		},
		{
			ViewName:    "main",
			Contexts:    []string{"staging"},
			Key:         gui.getKey("universal.return"),
			Modifier:    gocui.ModNone,
			Handler:     gui.handleStagingEscape,
			Description: gui.Tr.SLocalize("ReturnToFilesPanel"),
		},
		{
			ViewName:    "main",
			Contexts:    []string{"staging"},
			Key:         gui.getKey("universal.select"),
			Modifier:    gocui.ModNone,
			Handler:     gui.handleStageSelection,
			Description: gui.Tr.SLocalize("StageSelection"),
		},
		{
			ViewName:    "main",
			Contexts:    []string{"staging"},
			Key:         gui.getKey("universal.remove"),
			Modifier:    gocui.ModNone,
			Handler:     gui.handleResetSelection,
			Description: gui.Tr.SLocalize("ResetSelection"),
		},
		{
			ViewName:    "main",
			Contexts:    []string{"staging"},
			Key:         gui.getKey("universal.togglePanel"),
			Modifier:    gocui.ModNone,
			Handler:     gui.handleTogglePanel,
			Description: gui.Tr.SLocalize("TogglePanel"),
		},
		{
			ViewName:    "main",
			Contexts:    []string{"patch-building"},
			Key:         gui.getKey("universal.return"),
			Modifier:    gocui.ModNone,
			Handler:     gui.handleEscapePatchBuildingPanel,
			Description: gui.Tr.SLocalize("ExitLineByLineMode"),
		},
		{
			ViewName:    "main",
			Contexts:    []string{"patch-building", "staging"},
			Key:         gui.getKey("universal.prevItem"),
			Modifier:    gocui.ModNone,
			Handler:     gui.handleSelectPrevLine,
			Description: gui.Tr.SLocalize("PrevLine"),
		},
		{
			ViewName:    "main",
			Contexts:    []string{"patch-building", "staging"},
			Key:         gui.getKey("universal.nextItem"),
			Modifier:    gocui.ModNone,
			Handler:     gui.handleSelectNextLine,
			Description: gui.Tr.SLocalize("NextLine"),
		},
		{
			ViewName: "main",
			Contexts: []string{"patch-building", "staging"},
			Key:      gui.getKey("universal.prevItem-alt"),
			Modifier: gocui.ModNone,
			Handler:  gui.handleSelectPrevLine,
		},
		{
			ViewName: "main",
			Contexts: []string{"patch-building", "staging"},
			Key:      gui.getKey("universal.nextItem-alt"),
			Modifier: gocui.ModNone,
			Handler:  gui.handleSelectNextLine,
		},
		{
			ViewName: "main",
			Contexts: []string{"patch-building", "staging"},
			Key:      gocui.MouseWheelUp,
			Modifier: gocui.ModNone,
			Handler:  gui.handleSelectPrevLine,
		},
		{
			ViewName: "main",
			Contexts: []string{"patch-building", "staging"},
			Key:      gocui.MouseWheelDown,
			Modifier: gocui.ModNone,
			Handler:  gui.handleSelectNextLine,
		},
		{
			ViewName:    "main",
			Contexts:    []string{"patch-building", "staging"},
			Key:         gui.getKey("universal.prevBlock"),
			Modifier:    gocui.ModNone,
			Handler:     gui.handleSelectPrevHunk,
			Description: gui.Tr.SLocalize("PrevHunk"),
		},
		{
			ViewName:    "main",
			Contexts:    []string{"patch-building", "staging"},
			Key:         gui.getKey("universal.nextBlock"),
			Modifier:    gocui.ModNone,
			Handler:     gui.handleSelectNextHunk,
			Description: gui.Tr.SLocalize("NextHunk"),
		},
		{
			ViewName: "main",
			Contexts: []string{"patch-building", "staging"},
			Key:      gui.getKey("universal.prevBlock-alt"),
			Modifier: gocui.ModNone,
			Handler:  gui.handleSelectPrevHunk,
		},
		{
			ViewName: "main",
			Contexts: []string{"patch-building", "staging"},
			Key:      gui.getKey("universal.nextBlock-alt"),
			Modifier: gocui.ModNone,
			Handler:  gui.handleSelectNextHunk,
		},
		{
			ViewName:    "main",
			Contexts:    []string{"staging"},
			Key:         gui.getKey("universal.edit"),
			Modifier:    gocui.ModNone,
			Handler:     gui.handleFileEdit,
			Description: gui.Tr.SLocalize("editFile"),
		},
		{
			ViewName:    "main",
			Contexts:    []string{"staging"},
			Key:         gui.getKey("universal.openFile"),
			Modifier:    gocui.ModNone,
			Handler:     gui.handleFileOpen,
			Description: gui.Tr.SLocalize("openFile"),
		},
		{
			ViewName:    "main",
			Contexts:    []string{"patch-building"},
			Key:         gui.getKey("universal.select"),
			Modifier:    gocui.ModNone,
			Handler:     gui.handleAddSelectionToPatch,
			Description: gui.Tr.SLocalize("StageSelection"),
		},
		{
			ViewName:    "main",
			Contexts:    []string{"patch-building"},
			Key:         gui.getKey("universal.remove"),
			Modifier:    gocui.ModNone,
			Handler:     gui.handleRemoveSelectionFromPatch,
			Description: gui.Tr.SLocalize("ResetSelection"),
		},
		{
			ViewName:    "main",
			Contexts:    []string{"patch-building", "staging"},
			Key:         gui.getKey("main.toggleDragSelect"),
			Modifier:    gocui.ModNone,
			Handler:     gui.handleToggleSelectRange,
			Description: gui.Tr.SLocalize("ToggleDragSelect"),
		},
		// Alias 'V' -> 'v'
		{
			ViewName:    "main",
			Contexts:    []string{"patch-building", "staging"},
			Key:         gui.getKey("main.toggleDragSelect-alt"),
			Modifier:    gocui.ModNone,
			Handler:     gui.handleToggleSelectRange,
			Description: gui.Tr.SLocalize("ToggleDragSelect"),
		},
		{
			ViewName:    "main",
			Contexts:    []string{"patch-building", "staging"},
			Key:         gui.getKey("main.toggleSelectHunk"),
			Modifier:    gocui.ModNone,
			Handler:     gui.handleToggleSelectHunk,
			Description: gui.Tr.SLocalize("ToggleSelectHunk"),
		},
		{
			ViewName: "main",
			Contexts: []string{"patch-building", "staging"},
			Key:      gocui.MouseLeft,
			Modifier: gocui.ModNone,
			Handler:  gui.handleMouseDown,
		},
		{
			ViewName: "main",
			Contexts: []string{"patch-building", "staging"},
			Key:      gocui.MouseLeft,
			Modifier: gocui.ModMotion,
			Handler:  gui.handleMouseDrag,
		},
		{
			ViewName: "main",
			Contexts: []string{"patch-building", "staging"},
			Key:      gocui.MouseWheelUp,
			Modifier: gocui.ModNone,
			Handler:  gui.handleMouseScrollUp,
		},
		{
			ViewName: "main",
			Contexts: []string{"patch-building", "staging"},
			Key:      gocui.MouseWheelDown,
			Modifier: gocui.ModNone,
			Handler:  gui.handleMouseScrollDown,
		},
		{
			ViewName:    "main",
			Contexts:    []string{"merging"},
			Key:         gui.getKey("universal.return"),
			Modifier:    gocui.ModNone,
			Handler:     gui.handleEscapeMerge,
			Description: gui.Tr.SLocalize("ReturnToFilesPanel"),
		},
		{
			ViewName:    "main",
			Contexts:    []string{"merging"},
			Key:         gui.getKey("universal.select"),
			Modifier:    gocui.ModNone,
			Handler:     gui.handlePickHunk,
			Description: gui.Tr.SLocalize("PickHunk"),
		},
		{
			ViewName:    "main",
			Contexts:    []string{"merging"},
			Key:         gui.getKey("main.PickBothHunks"),
			Modifier:    gocui.ModNone,
			Handler:     gui.handlePickBothHunks,
			Description: gui.Tr.SLocalize("PickBothHunks"),
		},
		{
			ViewName:    "main",
			Contexts:    []string{"merging"},
			Key:         gui.getKey("universal.prevBlock"),
			Modifier:    gocui.ModNone,
			Handler:     gui.handleSelectPrevConflict,
			Description: gui.Tr.SLocalize("PrevConflict"),
		},
		{
			ViewName:    "main",
			Contexts:    []string{"merging"},
			Key:         gui.getKey("universal.nextBlock"),
			Modifier:    gocui.ModNone,
			Handler:     gui.handleSelectNextConflict,
			Description: gui.Tr.SLocalize("NextConflict"),
		},
		{
			ViewName:    "main",
			Contexts:    []string{"merging"},
			Key:         gui.getKey("universal.prevItem"),
			Modifier:    gocui.ModNone,
			Handler:     gui.handleSelectTop,
			Description: gui.Tr.SLocalize("SelectTop"),
		},
		{
			ViewName:    "main",
			Contexts:    []string{"merging"},
			Key:         gui.getKey("universal.nextItem"),
			Modifier:    gocui.ModNone,
			Handler:     gui.handleSelectBottom,
			Description: gui.Tr.SLocalize("SelectBottom"),
		},
		{
			ViewName: "main",
			Contexts: []string{"mergin"},
			Key:      gocui.MouseWheelUp,
			Modifier: gocui.ModNone,
			Handler:  gui.handleSelectTop,
		},
		{
			ViewName: "main",
			Contexts: []string{"mergin"},
			Key:      gocui.MouseWheelDown,
			Modifier: gocui.ModNone,
			Handler:  gui.handleSelectBottom,
		},
		{
			ViewName: "main",
			Contexts: []string{"mergin"},
			Key:      gui.getKey("universal.prevBlock-alt"),
			Modifier: gocui.ModNone,
			Handler:  gui.handleSelectPrevConflict,
		},
		{
			ViewName: "main",
			Contexts: []string{"mergin"},
			Key:      gui.getKey("universal.nextBlock-alt"),
			Modifier: gocui.ModNone,
			Handler:  gui.handleSelectNextConflict,
		},
		{
			ViewName: "main",
			Contexts: []string{"mergin"},
			Key:      gui.getKey("universal.prevItem-alt"),
			Modifier: gocui.ModNone,
			Handler:  gui.handleSelectTop,
		},
		{
			ViewName: "main",
			Contexts: []string{"mergin"},
			Key:      gui.getKey("universal.nextItem-alt"),
			Modifier: gocui.ModNone,
			Handler:  gui.handleSelectBottom,
		},
		{
			ViewName:    "main",
			Contexts:    []string{"merging"},
			Key:         gui.getKey("main.undo"),
			Modifier:    gocui.ModNone,
			Handler:     gui.handlePopFileSnapshot,
			Description: gui.Tr.SLocalize("Undo"),
		},
		{
			ViewName: "branches",
			Contexts: []string{"remotes"},
			Key:      gui.getKey("universal.goInto"),
			Modifier: gocui.ModNone,
			Handler:  gui.handleRemoteEnter,
		},
		{
			ViewName:    "branches",
			Contexts:    []string{"remotes"},
			Key:         gui.getKey("universal.new"),
			Modifier:    gocui.ModNone,
			Handler:     gui.handleAddRemote,
			Description: gui.Tr.SLocalize("addNewRemote"),
		},
		{
			ViewName:    "branches",
			Contexts:    []string{"remotes"},
			Key:         gui.getKey("universal.remove"),
			Modifier:    gocui.ModNone,
			Handler:     gui.handleRemoveRemote,
			Description: gui.Tr.SLocalize("removeRemote"),
		},
		{
			ViewName:    "branches",
			Contexts:    []string{"remotes"},
			Key:         gui.getKey("universal.edit"),
			Modifier:    gocui.ModNone,
			Handler:     gui.handleEditRemote,
			Description: gui.Tr.SLocalize("editRemote"),
		},
		{
			ViewName:    "branches",
			Contexts:    []string{"remote-branches"},
			Key:         gui.getKey("universal.select"),
			Modifier:    gocui.ModNone,
			Handler:     gui.handleCheckoutRemoteBranch,
			Description: gui.Tr.SLocalize("checkout"),
		},
		{
			ViewName:    "branches",
			Contexts:    []string{"remote-branches"},
			Key:         gui.getKey("branches.mergeIntoCurrentBranch"),
			Modifier:    gocui.ModNone,
			Handler:     gui.handleMergeRemoteBranch,
			Description: gui.Tr.SLocalize("mergeIntoCurrentBranch"),
		},
		{
			ViewName:    "branches",
			Contexts:    []string{"remote-branches"},
			Key:         gui.getKey("universal.remove"),
			Modifier:    gocui.ModNone,
			Handler:     gui.handleDeleteRemoteBranch,
			Description: gui.Tr.SLocalize("deleteBranch"),
		},
		{
			ViewName:    "branches",
			Contexts:    []string{"remote-branches"},
			Key:         gui.getKey("branches.rebaseBranch"),
			Modifier:    gocui.ModNone,
			Handler:     gui.handleRebaseOntoRemoteBranch,
			Description: gui.Tr.SLocalize("rebaseBranch"),
		},
		{
			ViewName:    "branches",
			Contexts:    []string{"remote-branches"},
			Key:         gui.getKey("branches.setUpstream"),
			Modifier:    gocui.ModNone,
			Handler:     gui.handleSetBranchUpstream,
			Description: gui.Tr.SLocalize("setUpstream"),
		},
		{
			ViewName: "stash",
			Key:      gocui.MouseLeft,
			Modifier: gocui.ModNone,
			Handler:  gui.handleStashEntrySelect,
		},
		{
			ViewName: "status",
			Key:      gocui.MouseLeft,
			Modifier: gocui.ModNone,
			Handler:  gui.handleStatusClick,
		},
		{
			ViewName: "commitFiles",
			Key:      gocui.MouseLeft,
			Modifier: gocui.ModNone,
			Handler:  gui.handleCommitFilesClick,
		},
	}

	for _, viewName := range []string{"status", "branches", "files", "commits", "commitFiles", "stash", "menu"} {
		bindings = append(bindings, []*Binding{
			{ViewName: viewName, Key: gui.getKey("universal.togglePanel"), Modifier: gocui.ModNone, Handler: gui.nextView},
			{ViewName: viewName, Key: gui.getKey("universal.prevBlock"), Modifier: gocui.ModNone, Handler: gui.previousView},
			{ViewName: viewName, Key: gui.getKey("universal.nextBlock"), Modifier: gocui.ModNone, Handler: gui.nextView},
			{ViewName: viewName, Key: gui.getKey("universal.prevBlock-alt"), Modifier: gocui.ModNone, Handler: gui.previousView},
			{ViewName: viewName, Key: gui.getKey("universal.nextBlock-alt"), Modifier: gocui.ModNone, Handler: gui.nextView},
		}...)
	}

	// Appends keybindings to jump to a particular sideView using numbers
	for i, viewName := range []string{"status", "files", "branches", "commits", "stash"} {
		bindings = append(bindings, &Binding{ViewName: "", Key: rune(i+1) + '0', Modifier: gocui.ModNone, Handler: gui.goToSideView(viewName)})
	}

	for _, listView := range gui.getListViews() {
		bindings = append(bindings, []*Binding{
			{ViewName: listView.viewName, Contexts: []string{listView.context}, Key: gui.getKey("universal.prevItem-alt"), Modifier: gocui.ModNone, Handler: listView.handlePrevLine},
			{ViewName: listView.viewName, Contexts: []string{listView.context}, Key: gui.getKey("universal.prevItem"), Modifier: gocui.ModNone, Handler: listView.handlePrevLine},
			{ViewName: listView.viewName, Contexts: []string{listView.context}, Key: gocui.MouseWheelUp, Modifier: gocui.ModNone, Handler: listView.handlePrevLine},
			{ViewName: listView.viewName, Contexts: []string{listView.context}, Key: gui.getKey("universal.nextItem-alt"), Modifier: gocui.ModNone, Handler: listView.handleNextLine},
			{ViewName: listView.viewName, Contexts: []string{listView.context}, Key: gui.getKey("universal.nextItem"), Modifier: gocui.ModNone, Handler: listView.handleNextLine},
			{ViewName: listView.viewName, Contexts: []string{listView.context}, Key: gocui.MouseWheelDown, Modifier: gocui.ModNone, Handler: listView.handleNextLine},
			{ViewName: listView.viewName, Contexts: []string{listView.context}, Key: gocui.MouseLeft, Modifier: gocui.ModNone, Handler: listView.handleClick},
		}...)
	}

	return bindings
}

func (gui *Gui) keybindings(g *gocui.Gui) error {
	bindings := gui.GetInitialKeybindings()

	for _, binding := range bindings {
		if err := g.SetKeybinding(binding.ViewName, binding.Contexts, binding.Key, binding.Modifier, binding.Handler); err != nil {
			return err
		}
	}

	if err := g.SetTabClickBinding("branches", gui.onBranchesTabClick); err != nil {
		return err
	}

	return nil
}
