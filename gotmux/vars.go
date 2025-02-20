// Copyright (c) Gianluca Piccirillo
// This software is licensed under the MIT License.
// See the LICENSE file in the root directory for more information.

package gotmux

const (
	varActiveWindowIndex        = "active_window_index"
	varAlternateOn              = "alternate_on"
	varAlternateSavedX          = "alternate_saved_x"
	varAlternateSavedY          = "alternate_saved_y"
	varBufferCreated            = "buffer_created"
	varBufferName               = "buffer_name"
	varBufferSample             = "buffer_sample"
	varBufferSize               = "buffer_size"
	varClientActivity           = "client_activity"
	varClientCellHeight         = "client_cell_height"
	varClientCellWidth          = "client_cell_width"
	varClientControlMode        = "client_control_mode"
	varClientCreated            = "client_created"
	varClientDiscarded          = "client_discarded"
	varClientFlags              = "client_flags"
	varClientHeight             = "client_height"
	varClientKeyTable           = "client_key_table"
	varClientLastSession        = "client_last_session"
	varClientName               = "client_name"
	varClientPid                = "client_pid"
	varClientPrefix             = "client_prefix"
	varClientReadonly           = "client_readonly"
	varClientSession            = "client_session"
	varClientTermfeatures       = "client_termfeatures"
	varClientTermname           = "client_termname"
	varClientTermtype           = "client_termtype"
	varClientTty                = "client_tty"
	varClientUid                = "client_uid"
	varClientUser               = "client_user"
	varClientUtf8               = "client_utf8"
	varClientWidth              = "client_width"
	varClientWritten            = "client_written"
	varCommand                  = "command"
	varCommandListAlias         = "command_list_alias"
	varCommandListName          = "command_list_name"
	varCommandListUsage         = "command_list_usage"
	varConfigFiles              = "config_files"
	varCopyCursorLine           = "copy_cursor_line"
	varCopyCursorWord           = "copy_cursor_word"
	varCopyCursorX              = "copy_cursor_x"
	varCopyCursorY              = "copy_cursor_y"
	varCurrentFile              = "current_file"
	varCursorCharacter          = "cursor_character"
	varCursorFlag               = "cursor_flag"
	varCursorX                  = "cursor_x"
	varCursorY                  = "cursor_y"
	varHistoryBytes             = "history_bytes"
	varHistoryLimit             = "history_limit"
	varHistorySize              = "history_size"
	varHook                     = "hook"
	varHookClient               = "hook_client"
	varHookPane                 = "hook_pane"
	varHookSession              = "hook_session"
	varHookSessionName          = "hook_session_name"
	varHookWindow               = "hook_window"
	varHookWindowName           = "hook_window_name"
	varHost                     = "host"
	varHostShort                = "host_short"
	varInsertFlag               = "insert_flag"
	varKeypadCursorFlag         = "keypad_cursor_flag"
	varKeypadFlag               = "keypad_flag"
	varLastWindowIndex          = "last_window_index"
	varLine                     = "line"
	varMouseAllFlag             = "mouse_all_flag"
	varMouseAnyFlag             = "mouse_any_flag"
	varMouseButtonFlag          = "mouse_button_flag"
	varMouseHyperlink           = "mouse_hyperlink"
	varMouseLine                = "mouse_line"
	varMouseSgrFlag             = "mouse_sgr_flag"
	varMouseStandardFlag        = "mouse_standard_flag"
	varMouseStatusLine          = "mouse_status_line"
	varMouseStatusRange         = "mouse_status_range"
	varMouseUtf8Flag            = "mouse_utf8_flag"
	varMouseWord                = "mouse_word"
	varMouseX                   = "mouse_x"
	varMouseY                   = "mouse_y"
	varNextSessionId            = "next_session_id"
	varOriginFlag               = "origin_flag"
	varPaneActive               = "pane_active"
	varPaneAtBottom             = "pane_at_bottom"
	varPaneAtLeft               = "pane_at_left"
	varPaneAtRight              = "pane_at_right"
	varPaneAtTop                = "pane_at_top"
	varPaneBg                   = "pane_bg"
	varPaneBottom               = "pane_bottom"
	varPaneCurrentCommand       = "pane_current_command"
	varPaneCurrentPath          = "pane_current_path"
	varPaneDead                 = "pane_dead"
	varPaneDeadSignal           = "pane_dead_signal"
	varPaneDeadStatus           = "pane_dead_status"
	varPaneDeadTime             = "pane_dead_time"
	varPaneFg                   = "pane_fg"
	varPaneFormat               = "pane_format"
	varPaneHeight               = "pane_height"
	varPaneId                   = "pane_id"
	varPaneInMode               = "pane_in_mode"
	varPaneIndex                = "pane_index"
	varPaneInputOff             = "pane_input_off"
	varPaneLast                 = "pane_last"
	varPaneLeft                 = "pane_left"
	varPaneMarked               = "pane_marked"
	varPaneMarkedSet            = "pane_marked_set"
	varPaneMode                 = "pane_mode"
	varPanePath                 = "pane_path"
	varPanePid                  = "pane_pid"
	varPanePipe                 = "pane_pipe"
	varPaneRight                = "pane_right"
	varPaneSearchString         = "pane_search_string"
	varPaneStartCommand         = "pane_start_command"
	varPaneStartPath            = "pane_start_path"
	varPaneSynchronized         = "pane_synchronized"
	varPaneTabs                 = "pane_tabs"
	varPaneTitle                = "pane_title"
	varPaneTop                  = "pane_top"
	varPaneTty                  = "pane_tty"
	varPaneUnseenChanges        = "pane_unseen_changes"
	varPaneWidth                = "pane_width"
	varPid                      = "pid"
	varRectangleToggle          = "rectangle_toggle"
	varScrollPosition           = "scroll_position"
	varScrollRegionLower        = "scroll_region_lower"
	varScrollRegionUpper        = "scroll_region_upper"
	varSearchMatch              = "search_match"
	varSearchPresent            = "search_present"
	varSelectionActive          = "selection_active"
	varSelectionEndX            = "selection_end_x"
	varSelectionEndY            = "selection_end_y"
	varSelectionPresent         = "selection_present"
	varSelectionStartX          = "selection_start_x"
	varSelectionStartY          = "selection_start_y"
	varServerSessions           = "server_sessions"
	varSessionActivity          = "session_activity"
	varSessionAlerts            = "session_alerts"
	varSessionAttached          = "session_attached"
	varSessionAttachedList      = "session_attached_list"
	varSessionCreated           = "session_created"
	varSessionFormat            = "session_format"
	varSessionGroup             = "session_group"
	varSessionGroupAttached     = "session_group_attached"
	varSessionGroupAttachedList = "session_group_attached_list"
	varSessionGroupList         = "session_group_list"
	varSessionGroupManyAttached = "session_group_many_attached"
	varSessionGroupSize         = "session_group_size"
	varSessionGrouped           = "session_grouped"
	varSessionId                = "session_id"
	varSessionLastAttached      = "session_last_attached"
	varSessionManyAttached      = "session_many_attached"
	varSessionMarked            = "session_marked"
	varSessionName              = "session_name"
	varSessionPath              = "session_path"
	varSessionStack             = "session_stack"
	varSessionWindows           = "session_windows"
	varSocketPath               = "socket_path"
	varStartTime                = "start_time"
	varUid                      = "uid"
	varUser                     = "user"
	varVersion                  = "version"
	varWindowActive             = "window_active"
	varWindowActiveClients      = "window_active_clients"
	varWindowActiveClientsList  = "window_active_clients_list"
	varWindowActiveSessions     = "window_active_sessions"
	varWindowActiveSessionsList = "window_active_sessions_list"
	varWindowActivity           = "window_activity"
	varWindowActivityFlag       = "window_activity_flag"
	varWindowBellFlag           = "window_bell_flag"
	varWindowBigger             = "window_bigger"
	varWindowCellHeight         = "window_cell_height"
	varWindowCellWidth          = "window_cell_width"
	varWindowEndFlag            = "window_end_flag"
	varWindowFlags              = "window_flags"
	varWindowFormat             = "window_format"
	varWindowHeight             = "window_height"
	varWindowId                 = "window_id"
	varWindowIndex              = "window_index"
	varWindowLastFlag           = "window_last_flag"
	varWindowLayout             = "window_layout"
	varWindowLinked             = "window_linked"
	varWindowLinkedSessions     = "window_linked_sessions"
	varWindowLinkedSessionsList = "window_linked_sessions_list"
	varWindowMarkedFlag         = "window_marked_flag"
	varWindowName               = "window_name"
	varWindowOffsetX            = "window_offset_x"
	varWindowOffsetY            = "window_offset_y"
	varWindowPanes              = "window_panes"
	varWindowRawFlags           = "window_raw_flags"
	varWindowSilenceFlag        = "window_silence_flag"
	varWindowStackIndex         = "window_stack_index"
	varWindowStartFlag          = "window_start_flag"
	varWindowVisibleLayout      = "window_visible_layout"
	varWindowWidth              = "window_width"
	varWindowZoomedFlag         = "window_zoomed_flag"
	varWrapFlag                 = "wrap_flag"
)
