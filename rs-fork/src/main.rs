use clap::{App};
use std::thread;
use std::process;
use std::time::Duration;


fn build_cli() -> App<'static> {
    App::new("rs-fork")
        .arg("-c, --count=[count] 'count to fork'")
        .arg("-d, --duration=[duration] 'duration to sleep'")
}

fn fork_thread(count: u32, duration: u32) {
    for i in 1..=count {
        println!("fork thread id: {}", i);

        thread::spawn(|| {
            thread::park();
        });
        
        thread::sleep(Duration::from_millis(u64::from(duration)));
    }
}

fn main() {
    let matches = build_cli().get_matches();

    let count: u32 = matches.value_of_t("count").unwrap_or(5);
    let duration: u32 = matches.value_of_t("duration").unwrap_or(1000); 

    println!("[{}] ===> create {} every {}ms", process::id(), count, duration);

    fork_thread(count, duration);

    // wait here, or use while (!=0)
    thread::park();
}
