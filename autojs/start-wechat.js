
app.startActivity({
    packageName: "com.tencent.mm",
    className: "com.tencent.mm.ui.LauncherUI",
    extras: {

        // start voice
        // nofification_type: "talkroom_notification",
        // enter_chat_usrname: "rustgolang",

        // start chatting
        nofification_type: "new_msg_nofification",
        Main_User: "rustgolang",
      
        // nofification_type: "pushcontent_notification",
        // Main_FromUserName: "rustgolang",

        // "LauncherUI.Shortcut.LaunchType": "launch_type_voip",
        // "LauncherUI.Shortcut.LaunchType": "launch_type_voip_audio",
        // "LauncherUI.Shortcut.LaunchType": "launch_type_scan_qrcode",
        // "LauncherUI.Shortcut.LaunchType": "launch_type_offline_wallet",
        // "LauncherUI.Shortcut.LaunchType": "launch_type_my_qrcode",

        // "LauncherUI.From.Scaner.Shortcut": true,
        // "LauncherUI.switch.tab": "tab_find_friend",

        // "LauncherUI.From.Biz.Shortcut": true,
        // "LauncherUI.Shortcut.Username": "wellwellwork"

    }
});

toast("哈哈哈");