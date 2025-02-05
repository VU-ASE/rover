use std::env;
use std::fs::File;
use std::io::Write;
use std::path::Path;

fn main() -> std::io::Result<()> {
    let out_dir = env::var_os("OUT_DIR").unwrap();
    let version = env!("CARGO_PKG_VERSION");

    let mut file = File::create(Path::new(&out_dir).join("version.rs"))?;

    writeln!(file, "#[allow(unused)]")?;
    writeln!(
        file,
        "pub const VERSION_MAJOR: u32 = {};",
        env!("CARGO_PKG_VERSION_MAJOR")
    )?;
    writeln!(file, "#[allow(unused)]")?;
    writeln!(
        file,
        "pub const VERSION_MINOR: u32 = {};",
        env!("CARGO_PKG_VERSION_MINOR")
    )?;
    writeln!(file, "#[allow(unused)]")?;
    writeln!(
        file,
        "pub const VERSION_PATCH: u32 = {};",
        env!("CARGO_PKG_VERSION_PATCH")
    )?;
    writeln!(file, "pub const VERSION: &str = \"{}\";", version)?;

    println!("cargo:rerun-if-changed=build.rs");

    Ok(())
}
