# Wisp Prover Server



The Prover is a fork of [Iden3's prover-server](https://github.com/iden3/prover-server). It is a REST API Wrapper
for [go-rapidsnark](https://github.com/iden3/go-rapidsnark)

## Installation

1. (Optional) Create / edit your config file. Defaults to `configs/dev.yaml`.
2. Prepare compiled circuits, zkey and verification key.
    1. Create a `circuits` directory if there isn't one.
          ```bash
          mkdir ./circuits
          ```
    2. Put your compiled circuit file and `final.zkey` and `vkey.json` into this directory.
3. Build the image
    ```bash
    docker build -t prover-server .
    ```
4. Run Prover
   ```
   docker run -it -p 8000:8000 prover-server
   ```
   If you want to use config, different from the default `dev` one you must pass it as an environmental
   variable `CONFIG={config}`

## API

### Generate proof

```
POST /api/v1/proof/generate
Content-Type: application/json
{
    "circuit": "multiplier", // name of the requested circuit as specified in the config
    "inputs": {...} // circuit specific inputs
}
```