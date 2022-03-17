
app.startActivity({
    packageName: "com.tencent.mm",
    className: "com.tencent.mm.ui.LauncherUI",
    extras: {

        // @start voice
        // nofification_type: "talkroom_notification",
        // enter_chat_usrname: "zoe_im",

        // @start chatting
        nofification_type: "new_msg_nofification",
        Main_User: "zoe_im", // 需要使用微信ID jsz1993 才不会错乱
        // Intro_Is_Muti_Talker: false,

        // MainUI_User_Last_Msg_Type: 34,
        
        // pushcontent_unread_count: 1,
        // Intro_Bottle_unread_count: 1,

        // From_fail_notify: true,
        // biz_chat_need_to_jump_to_chatting_ui: true,

      
        // nofification_type: "pushcontent_notification",
        // Main_FromUserName: "rustgolang",

        // @
        // "LauncherUI.Shortcut.LaunchType": "launch_type_voip",
        // "LauncherUI.Shortcut.LaunchType": "launch_type_voip_audio",
        // "LauncherUI.Shortcut.LaunchType": "launch_type_scan_qrcode",
        // "LauncherUI.Shortcut.LaunchType": "launch_type_offline_wallet",
        // "LauncherUI.Shortcut.LaunchType": "launch_type_my_qrcode",

        // @
        // "LauncherUI.From.Scaner.Shortcut": true,
        // "LauncherUI.From.Biz.Shortcut": true,
        // "LauncherUI.Shortcut.Username": "zoe_im"

        // @
        // "LauncherUI.switch.tab": "tab_find_friend",

        // "LauncherUI.From.Biz.Shortcut": true,
        // "LauncherUI.Shortcut.Username": "wellwellwork"

    }
});

toast("哈哈哈");