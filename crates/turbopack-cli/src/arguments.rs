use std::{
    net::IpAddr,
    path::{Path, PathBuf},
};

use clap::{Args, Parser};
use turbopack_cli_utils::issue::IssueSeverityCliOption;

#[derive(Debug, Parser)]
#[clap(author, version, about, long_about = None)]
pub enum Arguments {
    Dev(DevArguments),
}

impl Arguments {
    /// The directory of the application. see [CommonArguments]::dir
    pub fn dir(&self) -> Option<&Path> {
        match self {
            Arguments::Dev(args) => args.common.dir.as_deref(),
        }
    }
}

#[derive(Debug, Args)]
pub struct CommonArguments {
    /// The directory of the application.
    /// If no directory is provided, the current directory will be used.
    #[clap(value_parser)]
    pub dir: Option<PathBuf>,

    /// The root directory of the project. Nothing outside of this directory can
    /// be accessed. e. g. the monorepo root.
    /// If no directory is provided, `dir` will be used.
    #[clap(long, value_parser)]
    pub root: Option<PathBuf>,

    /// Filter by issue severity.
    #[clap(short, long)]
    pub log_level: Option<IssueSeverityCliOption>,

    /// Show all log messages without limit.
    #[clap(long)]
    pub show_all: bool,

    /// Expand the log details.
    #[clap(long)]
    pub log_detail: bool,

    /// Whether to enable full task stats recording in Turbo Engine.
    #[clap(long)]
    pub full_stats: bool,

    /// Enable experimental garbage collection with the provided memory limit in
    /// MB.
    #[clap(long)]
    pub memory_limit: Option<usize>,
}

#[derive(Debug, Args)]
#[clap(author, version, about, long_about = None)]
pub struct DevArguments {
    #[clap(flatten)]
    pub common: CommonArguments,

    /// The port number on which to start the application
    /// Note: setting env PORT allows to configure port without explicit cli
    /// args. However, this is temporary measure to conform with existing
    /// next.js devserver and can be removed in the future.
    #[clap(short, long, value_parser, default_value_t = 3000, env = "PORT")]
    pub port: u16,

    /// Hostname on which to start the application
    #[clap(short = 'H', long, value_parser, default_value = "0.0.0.0")]
    pub hostname: IpAddr,

    /// Compile all, instead of only compiling referenced assets when their
    /// parent asset is requested
    #[clap(long)]
    pub eager_compile: bool,

    /// Don't open the browser automatically when the dev server has started.
    #[clap(long)]
    pub no_open: bool,

    // ==
    // = Inherited options from next-dev, need revisit later.
    // ==
    /// If port is not explicitly specified, use different port if it's already
    /// in use.
    #[clap(long)]
    pub allow_retry: bool,
}
