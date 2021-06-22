use clap::{App, Arg};
use std::thread;
use std::process;
use std::time::Duration;
use nix::unistd::{fork, ForkResult};


fn build_cli() -> App<'static> {
    App::new("rs-fork")
        .arg("-c, --count=[count] 'count to fork, default 5'")
        .arg("-d, --duration=[duration] 'duration to sleep, default 1000ms'")
        .arg(
            Arg::new("MODE")
                .about("What mode the program fork with")
                .index(1)
                .possible_values(&["thread", "process"])
                .required(true),
        )
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

fn fork_process(count: u32, duration: u32) {
    for i in 1..=count {
        println!("fork process id: {}", i);

        match unsafe{fork()} {
            Ok(ForkResult::Parent { .. }) => {
                thread::sleep(Duration::from_millis(u64::from(duration)));
            }
            Ok(ForkResult::Child) => thread::park(),
            Err(_) => println!("fork error"),
        }
    }
}

fn main() {
    let matches = build_cli().get_matches();

    let mode: &str = matches.value_of("MODE").unwrap();
    let count: u32 = matches.value_of_t("count").unwrap_or(5);
    let duration: u32 = matches.value_of_t("duration").unwrap_or(1000); 

    println!(
        "[{}] ===> {}: create {} every {}ms",
        process::id(),
        mode,
        count,
        duration,
    );

    match mode {
        "thread" => {
            fork_thread(count, duration);
        }
        "process" => {
            fork_process(count, duration);
        }
        _ => unreachable!(),
    }

    // wait here, or use while (!=0)
    thread::park();
}
