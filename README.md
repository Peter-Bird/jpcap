# PcapAnalyzer

**PcapAnalyzer** is a command-line application for processing and analyzing pcapng files. It extracts and evaluates network packet information, offering insights into the data's structure and statistics.

## Features

- **File Analysis**: Reads pcapng files and verifies block types.
- **Statistics**: Provides detailed statistics after processing all packets.
- **Error Handling**: Logs issues with invalid files or unsupported block types.

## Getting Started

### Prerequisites

- **Go**: Version 1.21 or later.

### Installation

Clone the repository and build the project:

```bash
git clone <repository-url>
cd PcapAnalyzer
go build
