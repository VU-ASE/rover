# Changelog

## [0.7.2](https://github.com/VU-ASE/rover/compare/v0.7.1...v0.7.2) (2025-03-04)


### Bug Fixes

* hard-coded version ([058573d](https://github.com/VU-ASE/rover/commit/058573d9cf67a814ef554712b4ad610eaabcb0b4))

## [0.7.1](https://github.com/VU-ASE/rover/compare/v0.7.0...v0.7.1) (2025-03-04)


### Bug Fixes

* roverd early check for root ([451f7eb](https://github.com/VU-ASE/rover/commit/451f7eb4ab031fe82dc0d7b19e2a7377122a2e5f))

## [0.7.0](https://github.com/VU-ASE/rover/compare/v0.6.0...v0.7.0) (2025-03-03)


### Features

* CLI approach to roverctl ([bb78304](https://github.com/VU-ASE/rover/commit/bb78304cf2c6fcda4b2ce95bba51fcd183ee6ce3))
* dangers page ([a296ad3](https://github.com/VU-ASE/rover/commit/a296ad3ccb2d682d7a881c057d5a60aec36a9888))
* info target ([0bc9cb0](https://github.com/VU-ASE/rover/commit/0bc9cb0ecfd0d6ffe2aac31c12717839d7ede73f))
* roverctl-web startup through docker ([e3ca9a5](https://github.com/VU-ASE/rover/commit/e3ca9a54be65233a5f7a8f8fa8b915c9d042b724))
* start roverctl-web container ([9329015](https://github.com/VU-ASE/rover/commit/9329015b9cb86a8db86bc756052327fdf170785e))
* update target ([0f3cc2a](https://github.com/VU-ASE/rover/commit/0f3cc2a340b5a78a3c04fc363941608fa9622613))

## [0.6.0](https://github.com/VU-ASE/rover/compare/v0.5.1...v0.6.0) (2025-03-03)


### Features

* add passthrough back ([e07dc70](https://github.com/VU-ASE/rover/commit/e07dc7073c6b63dd9ed1c9894a11df2127e7d23c))


### Bug Fixes

* rename passthrough to appropriate folder ([5c1d41e](https://github.com/VU-ASE/rover/commit/5c1d41e8f3a1bb3be425841ede4c9a80fe2c2861))
* roverd unclosed delimiter ([17d8e95](https://github.com/VU-ASE/rover/commit/17d8e95a9b322a205ffe46c5dd367459f59c9aef))

## [0.5.1](https://github.com/VU-ASE/rover/compare/v0.5.0...v0.5.1) (2025-03-03)


### Bug Fixes

* use lowercase for GHCR tag ([e62a583](https://github.com/VU-ASE/rover/commit/e62a58302e17c7ca9739349b60fc654cce7f7da3))

## [0.5.0](https://github.com/VU-ASE/rover/compare/v0.4.1...v0.5.0) (2025-03-03)


### Features

* added Dockerfile for roverctl-web ([ec9733b](https://github.com/VU-ASE/rover/commit/ec9733b2f507d03d0ec405e6bcc2d82519406e6b))
* added passthrough and debugging functionality ([90975ef](https://github.com/VU-ASE/rover/commit/90975ef34e1acf2ffc1f9ec740e18b600df7d561))
* app config and overview page ([1c76d8c](https://github.com/VU-ASE/rover/commit/1c76d8cd42abe86a77e5133806ac97921a4da5d5))
* enable and install debug logic from roverctl-web ([a760e30](https://github.com/VU-ASE/rover/commit/a760e30d6991e0322399aeb4d65beca78a06dcc6))
* github actions workflows to build and push docker images ([5c4b035](https://github.com/VU-ASE/rover/commit/5c4b035329c9998331228113913a28fcc72b7552))
* manage page (wip) and status overlay ([e9132ed](https://github.com/VU-ASE/rover/commit/e9132ed41535b49fbc10353c023d6935edf82640))
* service information ([c0119a5](https://github.com/VU-ASE/rover/commit/c0119a5c4715cb39d86fdecdf856ce3d48127595))
* service installation ([246a452](https://github.com/VU-ASE/rover/commit/246a452e0ae9b4d8489b1f852d1a425c7137e895))
* service selection toggles and ordening ([6ad02f4](https://github.com/VU-ASE/rover/commit/6ad02f49222b35caec422dc1e2c486276a50aa65))
* starting, stopping and modifying pipelines ([e54a176](https://github.com/VU-ASE/rover/commit/e54a1762fdc686ec0236874c1169ea9c20014bff))


### Bug Fixes

* consistent style for services ([a9f5fbe](https://github.com/VU-ASE/rover/commit/a9f5fbef311b50ce30ca3891a9c5591ec3c12d91))
* poll pipeline status ([a02bbf6](https://github.com/VU-ASE/rover/commit/a02bbf698dbd8888b4f8efcf82dfb2fdb7eca9a8))
* returning service_as optimistically for all endpoints ([a019fe1](https://github.com/VU-ASE/rover/commit/a019fe1bee5a0a5e591911b389ee599b647e0440))
* show correct panels for selected services ([5ba569e](https://github.com/VU-ASE/rover/commit/5ba569e06f205686fd5baa2dd7ff0ae765d06bb7))
* show debug toggle correctly ([3ea63f8](https://github.com/VU-ASE/rover/commit/3ea63f8a79ee4fe7ca0836407fb96d6f501fd94d))
* use dynamic env variables for configuration ([25d5c88](https://github.com/VU-ASE/rover/commit/25d5c88b217c9040aea12686a4b08b05443a8e03))

## [0.4.1](https://github.com/VU-ASE/rover/compare/v0.4.0...v0.4.1) (2025-02-12)


### Bug Fixes

* version utility and correct comparison ([9251512](https://github.com/VU-ASE/rover/commit/92515124a450c2edd2ad8aed10a20b7b42ba4b9c))

## [0.4.0](https://github.com/VU-ASE/rover/compare/v0.3.3...v0.4.0) (2025-02-11)


### Features

* service_as feature ([782ab78](https://github.com/VU-ASE/rover/commit/782ab7874c1eb6eec0e58ac73283f2f9ab940a5b))


### Bug Fixes

* hide updater screen when all services are installed ([f4dce16](https://github.com/VU-ASE/rover/commit/f4dce16e2ae91666336fa361d6c2d15db813c380))

## [0.3.3](https://github.com/VU-ASE/rover/compare/v0.3.2...v0.3.3) (2025-02-10)


### Bug Fixes

* update checking for default pipeline ([3f67220](https://github.com/VU-ASE/rover/commit/3f672204d7faaeae6ec4a3d2ec141ab52af4dc55))

## [0.3.2](https://github.com/VU-ASE/rover/compare/v0.3.1...v0.3.2) (2025-02-10)


### Bug Fixes

* version mismatch checking ([4ba63b9](https://github.com/VU-ASE/rover/commit/4ba63b901f88055e68db2a7234ca17659588798c))

## [0.3.1](https://github.com/VU-ASE/rover/compare/v0.3.0...v0.3.1) (2025-02-09)


### Bug Fixes

* roverctl replace FQN types ([f6a9c44](https://github.com/VU-ASE/rover/commit/f6a9c44dee576b4ba63f3e44ca5a07704ead245b))

## [0.3.0](https://github.com/VU-ASE/rover/compare/v0.2.0...v0.3.0) (2025-02-09)


### Features

* api adaptation ([d36262f](https://github.com/VU-ASE/rover/commit/d36262f0e0d336e08cc1a3a6e2484ad4e52c2431))
* default ASE pipeline updates and installs ([0abcdb0](https://github.com/VU-ASE/rover/commit/0abcdb0021d0c2c4158a52e479684615864518af))


### Bug Fixes

* make roverctl compatible with new API spec ([6eb675c](https://github.com/VU-ASE/rover/commit/6eb675c87be7b49d31ca3405906767df359405f3))

## [0.2.0](https://github.com/VU-ASE/rover/compare/v0.1.2...v0.2.0) (2025-02-07)


### Features

* api clean up and addition of alias error (BBC) ([1d685a6](https://github.com/VU-ASE/rover/commit/1d685a6578ddc0b9c3b48fa9ffeb0abecfefad22))
* install script that allows for installing a specific version ([f5bfcaf](https://github.com/VU-ASE/rover/commit/f5bfcaf2bada29e9e8732b61d79c93a53c3331d9))
* upgrade roverctl to go 1.22.12 ([6a19635](https://github.com/VU-ASE/rover/commit/6a1963502affb37a09328151d1c55622ffb14d0f))


### Bug Fixes

* error on roverctl install.sh if release does not exist or no binary ([1880064](https://github.com/VU-ASE/rover/commit/1880064e54514a502ae07bd21f690b8d953a7eea))
* roverctl install.sh, verify downloaded binary is valid ([381a5ed](https://github.com/VU-ASE/rover/commit/381a5ed2e809710a882a65e4affd3fb658ecb2e2))

## [0.1.2](https://github.com/VU-ASE/rover/compare/v0.1.1...v0.1.2) (2025-02-06)


### Bug Fixes

* version update script now accepts "v" prefix ([362c24c](https://github.com/VU-ASE/rover/commit/362c24c0b3264ce7d3aef7fdfabbfb6bea72d45a))

## [0.1.1](https://github.com/VU-ASE/rover/compare/v0.1.0...v0.1.1) (2025-02-06)


### Bug Fixes

* testing ci again ([573bb0f](https://github.com/VU-ASE/rover/commit/573bb0f5aafb6ca7b18208a4c71b1a14724de484))

## [0.1.0](https://github.com/VU-ASE/rover/compare/v0.0.6...v0.1.0) (2025-02-06)


### Features

* added manage script ([e1b1429](https://github.com/VU-ASE/rover/commit/e1b142989b1b0c0dfcaa55f0a496da56031288e4))


### Bug Fixes

* broken links ([7a4d2aa](https://github.com/VU-ASE/rover/commit/7a4d2aa8ee9bc87d5f06ea3c1f55a9f324cc57a7))
* ci testing ([269b7a3](https://github.com/VU-ASE/rover/commit/269b7a3b9ece651b79b22145db9eccb1cf018fa9))
* docs and install script for roverctl ([a069205](https://github.com/VU-ASE/rover/commit/a069205b726187a5ad7ad99c8f00a75f7569ce9f))
* public link for schematic in docs ([b26decf](https://github.com/VU-ASE/rover/commit/b26decf57eb08b15e0fd4c2c04e81d1f9b3a1fdd))

## [0.0.6](https://github.com/VU-ASE/rover/compare/v0.0.5...v0.0.6) (2025-02-06)


### Bug Fixes

* point install scripts to rover repository (was roverctl) ([b139324](https://github.com/VU-ASE/rover/commit/b139324d303d27c88f2f8d91d0aaf1b78a5f714a))

## [0.0.5](https://github.com/VU-ASE/rover/compare/v0.0.1...v0.0.5) (2025-02-06)


### Bug Fixes

* debug missing assets ([508f3dc](https://github.com/VU-ASE/rover/commit/508f3dcef4aab0035d5709d8f9a3922f5a1c26e6))
* exact match with roverd upload ([b8d38a6](https://github.com/VU-ASE/rover/commit/b8d38a6c927bc1539bece3833dff9ece119b07e6))
* remove whitespace from upload action ([0b89220](https://github.com/VU-ASE/rover/commit/0b892200cb4e12c3321ec080a40dc8286951ac3c))
* use correct file separator for asset upload ([699d913](https://github.com/VU-ASE/rover/commit/699d913477799a2e487821df38645650c3d5b7b5))


### Miscellaneous Chores

* release 0.0.1 ([0112285](https://github.com/VU-ASE/rover/commit/0112285b31119bb0d0262a3796a91415932d9c3e))
* release 0.0.5 ([aa2ed11](https://github.com/VU-ASE/rover/commit/aa2ed11e18453f8a6edc3f82cd503e38288e7060))

## [0.0.1](https://github.com/VU-ASE/rover/compare/v0.0.1...v0.0.1) (2025-02-06)


### Bug Fixes

* debug missing assets ([508f3dc](https://github.com/VU-ASE/rover/commit/508f3dcef4aab0035d5709d8f9a3922f5a1c26e6))
* exact match with roverd upload ([b8d38a6](https://github.com/VU-ASE/rover/commit/b8d38a6c927bc1539bece3833dff9ece119b07e6))
* remove whitespace from upload action ([0b89220](https://github.com/VU-ASE/rover/commit/0b892200cb4e12c3321ec080a40dc8286951ac3c))
* use correct file separator for asset upload ([699d913](https://github.com/VU-ASE/rover/commit/699d913477799a2e487821df38645650c3d5b7b5))


### Miscellaneous Chores

* release 0.0.1 ([0112285](https://github.com/VU-ASE/rover/commit/0112285b31119bb0d0262a3796a91415932d9c3e))

## [0.0.1](https://github.com/VU-ASE/rover/compare/v0.0.1...v0.0.1) (2025-02-06)


### Bug Fixes

* debug missing assets ([508f3dc](https://github.com/VU-ASE/rover/commit/508f3dcef4aab0035d5709d8f9a3922f5a1c26e6))
* exact match with roverd upload ([b8d38a6](https://github.com/VU-ASE/rover/commit/b8d38a6c927bc1539bece3833dff9ece119b07e6))
* use correct file separator for asset upload ([699d913](https://github.com/VU-ASE/rover/commit/699d913477799a2e487821df38645650c3d5b7b5))


### Miscellaneous Chores

* release 0.0.1 ([0112285](https://github.com/VU-ASE/rover/commit/0112285b31119bb0d0262a3796a91415932d9c3e))

## [0.0.1](https://github.com/VU-ASE/rover/compare/v0.0.3...v0.0.1) (2025-02-06)


### Bug Fixes

* exact match with roverd upload ([b8d38a6](https://github.com/VU-ASE/rover/commit/b8d38a6c927bc1539bece3833dff9ece119b07e6))


### Miscellaneous Chores

* release 0.0.1 ([0112285](https://github.com/VU-ASE/rover/commit/0112285b31119bb0d0262a3796a91415932d9c3e))

## [0.0.3](https://github.com/VU-ASE/rover/compare/v0.0.2...v0.0.3) (2025-02-06)


### Bug Fixes

* use correct file separator for asset upload ([699d913](https://github.com/VU-ASE/rover/commit/699d913477799a2e487821df38645650c3d5b7b5))

## [0.0.2](https://github.com/VU-ASE/rover/compare/v0.0.1...v0.0.2) (2025-02-06)


### Bug Fixes

* debug missing assets ([508f3dc](https://github.com/VU-ASE/rover/commit/508f3dcef4aab0035d5709d8f9a3922f5a1c26e6))

## [0.0.1](https://github.com/VU-ASE/rover/compare/v1.0.14...v0.0.1) (2025-02-06)


### Miscellaneous Chores

* release 0.0.1 ([abd806e](https://github.com/VU-ASE/rover/commit/abd806e59353ee5ea9da82a80612ff9ece653285))

## [1.0.14](https://github.com/VU-ASE/rover/compare/v1.0.13...v1.0.14) (2025-02-06)


### Bug Fixes

* extract from correct location ([d4106ba](https://github.com/VU-ASE/rover/commit/d4106ba17d7e8ec4d7fea3f3827aded9d4a25d0a))

## [1.0.13](https://github.com/VU-ASE/rover/compare/v1.0.12...v1.0.13) (2025-02-06)


### Bug Fixes

* extract artifact downloads first ([37762b7](https://github.com/VU-ASE/rover/commit/37762b77e75becc3a8fb960e83ddd72f47305729))

## [1.0.12](https://github.com/VU-ASE/rover/compare/v1.0.11...v1.0.12) (2025-02-06)


### Bug Fixes

* use correct version of action/artifact ([300beda](https://github.com/VU-ASE/rover/commit/300beda2031241877d4101279e5f9ec5cf8dcd07))

## [1.0.11](https://github.com/VU-ASE/rover/compare/v1.0.10...v1.0.11) (2025-02-06)


### Bug Fixes

* release assets on all success ([4448019](https://github.com/VU-ASE/rover/commit/44480199fcac3cc9c81e317ef7b34c270641ec6a))

## [1.0.10](https://github.com/VU-ASE/rover/compare/v1.0.9...v1.0.10) (2025-02-06)


### Bug Fixes

* remove unused view ([31f91cd](https://github.com/VU-ASE/rover/commit/31f91cd1bd482a2a1bb05af107ebb19b153d1de5))

## [1.0.9](https://github.com/VU-ASE/rover/compare/v1.0.8...v1.0.9) (2025-02-06)


### Bug Fixes

* permission issues for building ([8f82d02](https://github.com/VU-ASE/rover/commit/8f82d02b4b4bcd6a5930c5faa8fd504042a69f0b))

## [1.0.8](https://github.com/VU-ASE/rover/compare/v1.0.7...v1.0.8) (2025-02-06)


### Bug Fixes

* build to ./bin ([dba0f2b](https://github.com/VU-ASE/rover/commit/dba0f2b1cfd4deaa257beb7c0bbe241b02879f5d))

## [1.0.7](https://github.com/VU-ASE/rover/compare/v1.0.6...v1.0.7) (2025-02-06)


### Bug Fixes

* permissions on github actions ([0b7225a](https://github.com/VU-ASE/rover/commit/0b7225ae596dd3d1e1a29a46ad26b2f6e256d9dd))

## [1.0.6](https://github.com/VU-ASE/rover/compare/v1.0.5...v1.0.6) (2025-02-06)


### Bug Fixes

* create openapi directory if it does not exist ([1f881b8](https://github.com/VU-ASE/rover/commit/1f881b8846f38ceea597ab31be031ae7c60ecf96))

## [1.0.5](https://github.com/VU-ASE/rover/compare/v1.0.4...v1.0.5) (2025-02-06)


### Bug Fixes

* use docker buildx for builds ([aa4893b](https://github.com/VU-ASE/rover/commit/aa4893b9e97492b82198164711f7cca2e534f094))

## [1.0.4](https://github.com/VU-ASE/rover/compare/v1.0.3...v1.0.4) (2025-02-05)


### Bug Fixes

* run open-api generator as devuser ([e8e84c8](https://github.com/VU-ASE/rover/commit/e8e84c85578d578299798f469782f04e9acff0f0))

## [1.0.3](https://github.com/VU-ASE/rover/compare/v1.0.2...v1.0.3) (2025-02-05)


### Bug Fixes

* open-api-gen debugging ([013bce9](https://github.com/VU-ASE/rover/commit/013bce9f4a0838ccf91ddb53e429d27cb7132178))

## [1.0.2](https://github.com/VU-ASE/rover/compare/v1.0.1...v1.0.2) (2025-02-05)


### Bug Fixes

* do not build roverd on self-hosted ([af62651](https://github.com/VU-ASE/rover/commit/af626516335e3f1c2a5682aaa8fb54b5f72deb5c))
* do not publish openapi files ([f7f7ff0](https://github.com/VU-ASE/rover/commit/f7f7ff039253c5ebfc893fa0d53b5daad304968d))
* roverctl, build openapi files in Dockerfile ([af01eb6](https://github.com/VU-ASE/rover/commit/af01eb620b565518860cb84d5665595788b0de71))

## [1.0.1](https://github.com/VU-ASE/rover/compare/v1.0.0...v1.0.1) (2025-02-05)


### Bug Fixes

* move github.com/VU-ASE/roverctl to github.com/VU-ASE/rover/roverctl ([a35e0dc](https://github.com/VU-ASE/rover/commit/a35e0dc12895750545256432284054e560e9eb65))

## 1.0.0 (2025-02-05)


### Features

* unified roverd and roverctl under 'rover' repository ([f893365](https://github.com/VU-ASE/rover/commit/f893365bda94929d5cfa78a4a5d4407ba4d87d95))
