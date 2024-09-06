# EGTS Packages Debugger

A Go-based command-line tool for parsing and describing EGTS protocol packets. This tool is designed to assist with debugging and understanding EGTS protocol messages by providing a detailed description of the packet’s content.

# Installation

To install the tool, make sure you have Go installed on your machine, then run the following command:
```sh
go install github.com/qed-tech/egts-pkg-debugger
```

# Prerequisites

```
Go 1.22 or later
```

# Usage

To use the tool, simply run the executable with an EGTS packet string as input. The following is an example of how to run the tool from the command line:

```sh
egts-pkg-debugger decode "0100020b002300020001871800010011e70300000202101500b6739d1b4fba3a9ed227bc350000000000000000003b07"
```

### Output: 
```json
{
  "PRV": 1,
  "SKID": 0,
  "PRF": "00",
  "RTE": "0",
  "ENA": "00",
  "CMP": "0",
  "PR": "11",
  "HL": 11,
  "HE": 0,
  "FDL": 16,
  "PID": 1,
  "PT": 0,
  "PRA": 0,
  "RCA": 0,
  "TTL": 0,
  "HCS": 245,
  "SFRD": {
    "RPID": 2,
    "PR": 0,
    "SDR": [
      {
        "RL": 6,
        "RN": 1,
        "SSOD": "0",
        "RSOD": "0",
        "GRP": "0",
        "RPP": "11",
        "TMFE": "0",
        "EVFE": "0",
        "OBFE": "0",
        "OID": 0,
        "EVID": 0,
        "TM": "0001-01-01T00:00:00Z",
        "SST": 2,
        "RST": 2,
        "RD": [
          {
            "SRT": 0,
            "SRL": 3,
            "SRD": {
              "CRN": 1,
              "RST": 0
            }
          }
        ]
      }
    ]
  },
  "SFRCS": 12775
}
```

# Contact

For questions or feedback, please contact  
 💬 Telegram [@VladislavBerezovskiy]((https://t.me/VladislavBerezovskiy)  )